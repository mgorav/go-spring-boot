package main

import (
	"github.com/pineda89/golang-springboot/actuator"
	"github.com/pineda89/golang-springboot/config"
	"github.com/pineda89/golang-springboot/eureka"
	"github.com/pineda89/golang-springboot-archetype/handler"
	"os"
)

func main() {
	config.LoadConfig()

	go handler.StartWebServer(config.Configuration["server.port"].(int))

	go actuator.InitializeActuator()
	go eureka.Register(config.Configuration)
	eureka.CaptureInterruptSignal()
	eureka.Deregister()
	os.Exit(0)
}

