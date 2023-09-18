package conn

import (
	"net"
	"sync"
)

type Conn struct {
	conn     net.Conn
	key      string
	tag      sync.Map
	isClosed bool
	handle   func(*Conn)
}

func NewConn(conn net.Conn, key string) *Conn {
	return &Conn{
		conn:     conn,
		key:      key,
		tag:      sync.Map{},
		isClosed: false,
	}
}

func (c *Conn) Start() {
	if c.handle != nil {
		c.handle(c)
	}
}

func (c *Conn) Write(data []byte) (int, error) {
	return c.conn.Write(data)
}

func (c *Conn) Read(data []byte) (int, error) {
	return c.conn.Read(data)
}

func (c *Conn) IsClosed() bool {
	return c.isClosed
}

func (c *Conn) Close() error {
	c.isClosed = true
	return c.conn.Close()
}

func (c *Conn) SetHandle(handle func(*Conn)) {
	c.handle = handle
}

func (c *Conn) SetKey(key string) {
	c.key = key
}

func (c *Conn) GetKey() string {
	return c.key
}

func (c *Conn) GetTag(key string) any {
	value, _ := c.tag.Load(key)
	return value
}

func (c *Conn) SetTag(key string, value any) {
	c.tag.Store(key, value)
}

func (c *Conn) RemoveTag(key string) {
	c.tag.Delete(key)
}
