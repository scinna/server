package services

import (
	"os"
	"strconv"
	"crypto/rand"
	"encoding/base64"
	
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/matoous/go-nanoid"

    "golang.org/x/crypto/argon2"
)

type Provider struct {
	Db          *sqlx.DB
	ArgonParams *ArgonParams
	PicturePath string
}

func (prv *Provider) GenerateUID() (string, error) {
	alphabet, exists := os.LookupEnv("ID_ALPHABET")
	if !exists {
		alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ"
	}
	
	lengthStr, exists := os.LookupEnv("ID_LENGTH")
	length := 10
	if exists {
		length, _ = strconv.Atoi(lengthStr)
	}

	return gonanoid.Generate(alphabet, length)
}

func (prv *Provider) HashPassword(password string) (string, error) {
	salt, err := generateRandomBytes(prv.ArgonParams.SaltLength)
    if err != nil {
        return "", err
    }

    hash := argon2.IDKey([]byte(password), salt, prv.ArgonParams.Iterations, prv.ArgonParams.Memory, prv.ArgonParams.Parallelism, prv.ArgonParams.KeyLength)

    b64Salt := base64.RawStdEncoding.EncodeToString(salt)
    b64Hash := base64.RawStdEncoding.EncodeToString(hash)

    return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, prv.ArgonParams.Memory, prv.ArgonParams.Iterations, prv.ArgonParams.Parallelism, b64Salt, b64Hash), nil
}

// Temporary empty, will contain Postgres connection
func New(db *sqlx.DB, ap *ArgonParams, pictPath string) *Provider {
	return &Provider{
		Db:          db,
		ArgonParams: ap,
		PicturePath: pictPath,
	}
}


type ArgonParams struct {
    Memory      uint32
    Iterations  uint32
    Parallelism uint8
    SaltLength  uint32
    KeyLength   uint32
}

/** Is it really cryptographically secure as this website suggests ? **/
/** https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go **/
func generateRandomBytes(n uint32) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }

    return b, nil
}
