package main

import (
	"app/app"
	"app/domain"
	"app/infra/repository"
	"app/infra/transaction"
	"app/view/web"
	"os"
	"time"
)

func main() {

	var logger = app.NewZeroLogger()

	var application = app.NewApplication(
		logger,
		transaction.NewDbTransactionManager(app.NewDb()),
		domain.NewUsersService(repository.NewDbUsersRepository()),
	)

	var server = web.NewWebServer(8888, time.Second, application)

	logger.Fatal().Err(server.ListenAndServe()).Send()

	os.Exit(0)
}
