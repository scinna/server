package utils

import (
	"fmt"
	"github.com/scinna/server/models"
	"github.com/scinna/server/services"
	"os"
	"os/exec"
)

func Thumbnailize(prv *services.Provider, user *models.User, media *models.Media) error {
	parentFolder := prv.Config.MediaPath + user.UserID + "/"
	_ = os.MkdirAll(parentFolder, os.ModePerm)

	source := prv.Config.MediaPath + media.Path
	mediaPath := source + "_thumb"

	// [0] is to take the first frame when it's a gif
	// -quality to have a lesser quality (It's a thumbnail so no need for something extra
	// -thumbnail remove all extra infos and scale it down to 256x256
	imagickCommand := fmt.Sprintf("convert '%v[0]' -quality 50 -thumbnail 256x256\\> %v", source, mediaPath)

	cmd := exec.Command("sh", "-c", imagickCommand)
	return cmd.Run()
}
