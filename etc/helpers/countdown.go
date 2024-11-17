package helpers

import (
	"backend/etc/Utime"
	"backend/models"
)

func CountDown(resp *models.GetAllLogisticsResp) {
	for i := range resp.Logistics {
		logistic := &resp.Logistics[i]
		if logistic.Status == "READY" || logistic.Status == "AT HOME" || logistic.Status == "ready" ||
			logistic.Status == "READY AT HOME" || logistic.Status == "LET US KNOW" {
			logistic.Countdown = Utime.Now().Sub(logistic.UpdateTime)
		} else if logistic.Status == "COVERED" {
			logistic.Countdown = logistic.UpdateTime.Sub(Utime.Now())
		} else if logistic.Status == "AT PU" || logistic.Status == "AT DEL" ||
			logistic.Status == "TRUCK ISSUES" {
			logistic.Countdown = 0
		} else if logistic.Status == "ETA" || logistic.Status == "ETA WILL BE LATE" {
			logistic.Countdown = logistic.UpdateTime.Sub(Utime.Now())
		}
	}
}
