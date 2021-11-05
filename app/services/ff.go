package services

import (
	"fmt"
	channelpackage "logmontire/app/components/channel"
	"logmontire/app/components/environment"
	logpackage "logmontire/app/components/log"
	"logmontire/app/components/user"
	"logmontire/storage/mongo"
)

func Init() {
	listenNewEnvironments()
	listenNewChannels()
	listenNewLogs()
}

func listenNewEnvironments() {
	go func() {
		for {
			select {
			case env := <-environment.NewEnvironment:
				fmt.Println(env)
				for _, userId := range env.Users {
					fmt.Println(userId)
					u, err := user.FindById(userId)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(u)
				}
				break
			}
		}
	}()
}

func listenNewChannels() {
	go func() {
		for {
			select {
			case channel := <-channelpackage.NewChannel:
				go func() {
					fmt.Println("channel", channel)
				}()
				break
			}
		}
	}()
}

func listenNewLogs() {
	go func() {
		for {
			select {
			case log := <-logpackage.NewLog:
				go func() {
					fmt.Println("log", log)
					fmt.Println("logChannel", log.Channel)

					chss, err2 := channelpackage.Find(mongo.NewSelector(map[string]interface{}{
						"name": log.Channel,
					}))
					fmt.Println("chss", chss, len(chss), err2)

					chs, err := channelpackage.CreateChannel(map[string]interface{}{
						"name":           log.Channel,
						"environment_id": "2150",
					})

					fmt.Println("chs", chs)
					fmt.Println("err", err)
				}()
				break
			}
		}
	}()
}
