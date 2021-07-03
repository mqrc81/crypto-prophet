package bot

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/xelaj/mtproto/telegram"
)

const (
	resourcesFilePath = "./resources"
)

func AuthorizeTelegram() (*telegram.Client, error) {
	sessionFile := filepath.Join(resourcesFilePath, "tg_session.json")
	publicKeysFile := filepath.Join(resourcesFilePath, "tg_public_keys.pem")

	appID, _ := strconv.Atoi(os.Getenv("TELEGRAM_ID"))

	return telegram.NewClient(telegram.ClientConfig{
		SessionFile:     sessionFile,                 // where to store session configuration
		ServerHost:      os.Getenv("MTPROTO_SERVER"), // host address of mtproto proxy server
		PublicKeysFile:  publicKeysFile,              // can be found at https://my.telegram.org
		AppID:           appID,                       // can be found at https://my.telegram.org
		AppHash:         os.Getenv("TELEGRAM_HASH"),  // can be found at https://my.telegram.org
		InitWarnChannel: true,                        // if errors should be received, otherwise client.Warnings will be set nil
	})
}
