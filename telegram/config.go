package telegram


// Config contains some configuration variables for Telegram Adapter.
type Config struct {
	// Token declares the API token to integrate with Gitter.
	Token string `json:"token" yaml:"token"`	
}

// NewConfig creates and returns a new Config instance with default settings.
// Token is empty at this point as there can not be a default value.
// Use json.Unmarshal, yaml.Unmarshal, or manual manipulation to populate the blank value or override those default values.
func NewConfig() *Config {
	return &Config{
		Token: "",		
	}
}
