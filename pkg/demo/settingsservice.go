package demo

type ISettingService interface {
	GetSettings() *Settings
}

type Settings struct {
	Endpoint       string `json:"endpoint"`
	ApiKey         string `json:"apiKey"`
	Version        string `json:"version"`
	ChatModel      string `json:"chatModel"`
	EmbeddingModel string `json:"embeddingModel"`
}

var settings = &Settings{Endpoint: "https://api.openai.com/v1/engines/davinci-codex/completions", ApiKey: "sk-123"}

func (s *Settings) GetSettings() *Settings {
	return settings
}
