package handler

import (
	"github.com/vucongthanh92/courier-user-svc/config"
	"github.com/vucongthanh92/courier-user-svc/internal/domain/interfaces"
)

func Gracefully(cfg *config.AppConfig, cronService interfaces.CronJobServiceI) (err error) {
	return nil
}
