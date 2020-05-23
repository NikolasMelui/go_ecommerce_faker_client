package efclient

// Config ...
type Config struct {
	BaseURL string `toml:"base_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BaseURL: "http://localhost:1337",
	}
}
