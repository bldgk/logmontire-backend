package main

import (
	"logmontire/app"
	"logmontire/network/http"
	"logmontire/rmq"
	"logmontire/settings"
	"logmontire/storage"
)

func main() {
	settings.Load()
	http.Init(settings.Settings.HttpServer.Host, settings.Settings.HttpServer.Port)
	storage.Init(&settings.Settings.MongoDB)
	rmq.Init(settings.Settings.RabbitMQ.ConnectionString)
	app.Init()
	defer rmq.Close()
	http.Start()
}
