package components

import (
	"logmontire/app/components/channel"
	"logmontire/app/components/environment"
	"logmontire/app/components/home"
	"logmontire/app/components/log"
	"logmontire/app/components/user"
)

func Init() {
	home.Init()
	user.Init()
	environment.Init()
	channel.Init()
	log.Init()
}
