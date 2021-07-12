package bot

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/xelaj/mtproto/telegram"
)

const (
	resourcesFilePath           = "./resources"
	CrazyRussianTraderChannelID = 1234
)

func AuthorizeTelegram() (*telegram.Client, error) {
	sessionFile := filepath.Join(resourcesFilePath, "tg_session.json")
	publicKeysFile := filepath.Join(resourcesFilePath, "tg_public_keys.pem")

	appID, _ := strconv.Atoi(os.Getenv("TELEGRAM_ID"))

	return telegram.NewClient(telegram.ClientConfig{ // where to store session configuration
		SessionFile:     sessionFile,
		ServerHost:      os.Getenv("MTPROTO_SERVER"), // host address of mtproto server
		PublicKeysFile:  publicKeysFile,              // can be found at https://my.telegram.org
		AppID:           appID,                       // can be found at https://my.telegram.org
		AppHash:         os.Getenv("TELEGRAM_HASH"),  // can be found at https://my.telegram.org
		InitWarnChannel: true,                        // if errors should be received, otherwise client.Warnings will be set nil
	})
}

func getNewMessages(client *telegram.Client) {

}
