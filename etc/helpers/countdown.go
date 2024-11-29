package helpers

import (
	"backend/etc/Utime"
	"backend/models"
)

func CountDown(resp *models.GetAllLogisticsResp) {
	now := Utime.Now()

	for i := range resp.Companies {
		for j := range resp.Companies[i].Logistics {
			logistic := &resp.Companies[i].Logistics[j]
			switch logistic.Status {
			case "READY", "AT HOME", "READY AT HOME", "LET US KNOW":
				logistic.Countdown = now.Sub(logistic.UpdateTime).String()
			case "COVERED", "ETA", "ETA WILL BE LATE":
				logistic.Countdown = logistic.UpdateTime.Sub(now).String()
			case "AT PU", "AT DEL", "TRUCK ISSUES":
				logistic.Countdown = ""
			default:
				logistic.Countdown = ""
			}
		}
	}
}
