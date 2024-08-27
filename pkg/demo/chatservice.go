package demo

import "github.com/msalemor/gocontainer/pkg"

type IChatService interface {
	GetResponse(input string) string
}

type ChatService struct {
	Container *pkg.Container
}

func (c *ChatService) GetResponse(input string) string {
	settings = c.Container.Get("settings").(ISettingService).GetSettings()
	return input + " -> " + settings.Endpoint
}
