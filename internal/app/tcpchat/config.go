package tcpchat

// ...
type Config struct {
	BindPort string `json:"bind_port"`
}

// ...
func NewConfig() *Config {
	return &Config{
		BindPort: ":8083"
	}
}

// ...
func (config *Config) DecodeJFile(configPath string) error {
	// Загрузка из файла
	
	return nil
}