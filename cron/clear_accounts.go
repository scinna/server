package cron

import (
	"github.com/scinna/server/log"
	"github.com/scinna/server/services"
)

func ClearOldAccounts(prv *services.Provider) {
	log.Info("Cleaning up non-validated accounts")
	prv.DB.Exec(`DELETE FROM SCINNA_USER WHERE REGISTERED_AT < (NOW() - INTERVAL '1 HOUR')`)
}
