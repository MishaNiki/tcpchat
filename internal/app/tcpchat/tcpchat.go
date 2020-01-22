package tcpchat

import (
	"bufio"
	"fmt"
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

	fmt.Println(s.config.BindPort)
	list := NewListClients()
	listner, err := net.Listen("tcp", s.config.BindPort)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listner.Accept()
		if err != nil {
			return err
		}
		go handleConnection(list, conn)
	}
	return nil
}

func handleConnection(list *ListClients, conn net.Conn) {
	name := conn.RemoteAddr().String()
	fmt.Println(name + " connected")
	conn.Write([]byte("Hello, " + name + ", write your nickname\n"))
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	nickname := scanner.Text()
	client := list.Add(conn, nickname)

	list.mailing(client, "came in")
	defer list.Remove(client)
	defer list.mailing(client, "came out")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "->quit" {
			conn.Write([]byte("Bye\n\r"))
			fmt.Println(name, "disconnected")
			break
		} else if text != "" {
			list.mailing(client, text)
		}
	}
}

func (list *ListClients) mailing(client *Client, msg string) {
	for iter := list.first; iter != nil; iter = iter.next {
		if iter != client {
			iter.conn.Write([]byte(" " + client.nickname + ": " + msg + "\n\r"))
		}
	}
}
