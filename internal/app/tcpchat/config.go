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

	var lenBuf int
	for {
		len, e := file.Read(data)
		lenBuf += len
		if e == io.EOF {
			break
		}
	}
	err = json.Unmarshal(data[:lenBuf], config)
	if err != nil {
		return err
	}

	return nil
}
