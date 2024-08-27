package main

import (
	"fmt"

	"github.com/msalemor/gocontainer/pkg"
	"github.com/msalemor/gocontainer/pkg/demo"
)

var container = pkg.AppContainer

func initializeServices() {
	var settings demo.ISettingService = &demo.Settings{}
	container.Register("settings", settings)

	var chatService demo.IChatService = &demo.ChatService{Container: container}
	container.Register("chat", chatService)

	var webService demo.APIServer = &demo.Server{Container: container}
	container.Register("web", webService)
}

func init() {
	initializeServices()
}

func main() {
	var settingsService = container.Get("settings").(demo.ISettingService)
	fmt.Println(settingsService.GetSettings().Endpoint)

	var chatService = container.Get("chat").(demo.IChatService)
	fmt.Println(chatService.GetResponse("Hello, World!"))

	var webService = container.Get("web").(demo.APIServer)
	webService.Serve()
}
