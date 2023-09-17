package conn

import "sync"

type ConnManager struct {
	conns map[string]*Conn
	mux   *sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		conns: make(map[string]*Conn),
		mux:   new(sync.RWMutex),
	}
}

func (cm *ConnManager) AddConn(conn *Conn) {
	cm.mux.Lock()
	defer cm.mux.Unlock()
	cm.conns[conn.GetKey()] = conn
}

func (cm *ConnManager) RemoveConn(conn *Conn) {
	cm.mux.Lock()
	defer cm.mux.Unlock()
	delete(cm.conns, conn.GetKey())
}

func (cm *ConnManager) GetConn(key string) *Conn {
	cm.mux.RLock()
	defer cm.mux.RUnlock()
	return cm.conns[key]
}

func (cm *ConnManager) ReSetKey(oldKey, newKey string) {
	cm.mux.Lock()
	defer cm.mux.Unlock()
	conn := cm.conns[oldKey]
	delete(cm.conns, oldKey)
	conn.SetKey(newKey)
	cm.conns[newKey] = conn
}
