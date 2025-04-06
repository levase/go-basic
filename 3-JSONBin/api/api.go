package api

type APIConfig struct {
	APIKey string
}

func NewAPIConfig(apiKey string) *APIConfig {
	return &APIConfig{APIKey: apiKey}
}
