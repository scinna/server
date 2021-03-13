package translations

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"io/fs"
	"net/http"
)

//go:embed langs
var fakeFS embed.FS
var bundle *i18n.Bundle

func Initialize() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_ = fs.WalkDir(fakeFS, "langs", func(path string, d fs.DirEntry, err error) error {
		if path != "langs" {
			_, err = loadMessageFileFS(bundle, fakeFS, path)
			if err != nil {
				fmt.Println("Could not load language "+d.Name()+": ", err)
			}
		}
		return nil
	})
}

func T(r *http.Request, id string) string {
	accept := ""
	if r != nil {
		accept = r.Header.Get("Accept-Language")
	}

	return TLang(accept, id)
}

func TLang(lang, id string) string {
	l := i18n.NewLocalizer(bundle, lang, "en")
	translated, err := l.Localize(&i18n.LocalizeConfig{
		MessageID: id,
		DefaultMessage: &i18n.Message{
			ID:    id,
			Zero:  id,
			One:   id,
			Two:   id,
			Few:   id,
			Many:  id,
			Other: id,
		},
	})

	if err != nil {
		return id
	}

	return translated
}

// Stolen from https://github.com/nicksnyder/go-i18n/pull/246
// Waiting for it to be merged in master before using it natively
func loadMessageFileFS(bundle *i18n.Bundle, fsys fs.FS, path string) (*i18n.MessageFile, error) {
	buf, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, err
	}

	return bundle.ParseMessageFileBytes(buf, path)
}
