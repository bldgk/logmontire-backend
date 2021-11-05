package app

import (
	"logmontire/app/components"
	"logmontire/app/logging"
	"logmontire/app/services"
	"logmontire/log"
)

func Init() {
	components.Init()
	logging.Init()
	services.Init()
	lg("App started")
}

func lg(text string) {
	log.Log(log.LStr("App: "+text), log.ConsoleLog)
}
