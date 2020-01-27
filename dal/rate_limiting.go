package dal

import (
	"net/http"

	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"github.com/scinna/server/utils"
)

// APIRateLimiting limits the rate of the requests for the API
func APIRateLimiting(prv *services.Provider, r *http.Request) error {
	ip := utils.ReadUserIP(prv.Config.HeaderIPField, r)

	rq := `
		INSERT INTO RATE_LIMITER (IP, STARTING_TIME, AMT_REQ)
		VALUES ($1, CURRENT_TIMESTAMP, 1)
		ON CONFLICT (IP) DO 
		UPDATE SET AMT_REQ = CASE 
			WHEN rate_limiter.STARTING_TIME > (CURRENT_TIMESTAMP - INTERVAL '1 min') THEN
				rate_limiter.AMT_REQ + 1
			ELSE
				1
		end, STARTING_TIME = CASE
			WHEN rate_limiter.STARTING_TIME > (CURRENT_TIMESTAMP - INTERVAL '1 min') THEN
				rate_limiter.STARTING_TIME
			ELSE
				CURRENT_TIMESTAMP
		end
		RETURNING (AMT_REQ)`

	var amtRq int
	err := prv.Db.QueryRow(rq, ip).Scan(&amtRq)
	if err != nil {
		return err
	}

	if amtRq > prv.Config.RateLimiting {
		return serrors.ErrorRateLimited
	}

	return nil
}
