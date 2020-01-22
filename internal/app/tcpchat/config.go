package tcpchat

import (
	"encoding/json"
	"io"
	"os"
)

// ...
type Config struct {
	BindPort string `json:"bind_port"`
}

// ...
func NewConfig() *Config {
	return &Config{
		BindPort: ":8083",
	}
}

// ...
func (config *Config) DecodeJFile(configPath string) error {

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}

	defer file.Close()

	data := make([]byte, 64)
	for {
		_, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		return err
	}

	return nil
}
