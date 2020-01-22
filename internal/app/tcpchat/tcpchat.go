package tcpchat

import (
	"net"
)

type TCPChat struct {
	config *Config
}

func New(config *Config) *TCPChat {
	return &TCPChat{
		config: config,
	}
}

func (s *TCPChat) Start() error {

	listner, err := net.Listen("tcp", s.config.BindPort)
	if err != nil {
		return err
	}

	return nil
}
