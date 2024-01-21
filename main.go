package main

import (
	"lab.sda1.net/nexryai/altcore/internal/core/boot"
)

func main() {
	boot.Init()

	go func() {
		boot.StartWebServer()
	}()

	boot.QueueProcessDaemon()
}
