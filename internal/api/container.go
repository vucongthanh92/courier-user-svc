package api

import (
	"github.com/vucongthanh92/courier-user-svc/internal/api/cron"
	"github.com/vucongthanh92/courier-user-svc/internal/api/grpc"
	"github.com/vucongthanh92/courier-user-svc/internal/api/http"
	"github.com/vucongthanh92/courier-user-svc/internal/domain/interfaces"
	"github.com/vucongthanh92/courier-user-svc/internal/worker"
)

type ApiContainer struct {
	HttpServer     *http.Server
	GrpcServer     *grpc.Server
	CronServer     *cron.Server
	OutboxWorker   *worker.OutboxWorker
	EventPublisher interfaces.IntegrationEventPublisherI
}

func NewApiContainer(
	http *http.Server,
	grpc *grpc.Server,
	cron *cron.Server,
	outboxWorker *worker.OutboxWorker,
	eventPublisher interfaces.IntegrationEventPublisherI,
) *ApiContainer {
	return &ApiContainer{
		HttpServer:     http,
		GrpcServer:     grpc,
		CronServer:     cron,
		OutboxWorker:   outboxWorker,
		EventPublisher: eventPublisher,
	}
}
