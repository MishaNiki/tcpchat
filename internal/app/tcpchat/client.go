package tcpchat

import (
	"net"
)

type Client struct {
	conn net.Conn
	//mesgOutCh chan []byte
	nickname string
	next     *Client
}

type ListClients struct {
	first *Client
	last  *Client
	size  int
}

// ...
func NewListClients() *ListClients {
	return &ListClients{
		first: nil,
		last:  nil,
		size:  0,
	}
}

// ...
func NewClient(conn net.Conn, nickname string) *Client {
	return &Client{
		conn: conn,
		//mesgOutCh: make(chan []byte, 1),
		nickname: nickname,
	}
}

// ...
func (list *ListClients) Add(connect net.Conn, name string) *Client {
	client := NewClient(connect, name)
	if list.size == 0 {
		list.first = client
		list.last = client
	} else {
		list.last.next = client
		list.last = client
	}
	list.size++
	return client
}

// ...
func (list *ListClients) Remove(client *Client) {
	if client == nil {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var beforeClient *Client
	for iter := list.first; iter != client; iter = iter.next {
		beforeClient = iter
	}

	if client == list.first {
		list.first = client.next
	}

	if client == list.last {
		list.last = beforeClient
	}
	if beforeClient != nil {
		beforeClient.next = client.next
	}
	client.Disconect()
	client = nil
	list.size--
}

// ...
func (list *ListClients) Clear() {
	for iter := list.first; iter != nil; iter = iter.next {
		iter.Disconect()
	}
	list.first = nil
	list.last = nil
	list.size = 0
}

// !!!
func (client *Client) Disconect() {
	//close(client.mesgOutCh)
	client.conn.Close() //!!!
}
