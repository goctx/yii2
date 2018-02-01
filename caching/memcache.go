package caching

import (
	"github.com/bmizerany/mc"
	"time"
)

type MemCache struct {
	Cache
	client *mc.Conn
}

func NewMemCache(keyPrefix string, servers string) (*MemCache, error) {
	client, err := mc.Dial("tcp", servers)
	if err != nil {
		return nil, err
	}
	m := &MemCache{
		client: client,
	}
	m.KeyPrefix = keyPrefix
	return m, nil
}

func (m *MemCache) GetMemcache() *mc.Conn {
	return m.client
}

func (m *MemCache) Close() error {
	return m.client.Close()
}

func (m *MemCache) GetValue(key string) ([]byte, error) {
	if val, _, _, err := m.client.Get(m.BuildKey(key)); err != nil {
		return nil, err
	} else {
		return []byte(val), nil
	}
}

func (m *MemCache) SetValue(key string, val []byte, duration time.Duration) error {
	return m.client.Set(key, string(val), 0, 0, int(duration.Seconds()))
}

func (m *MemCache) DeleteValue(key string) error {
	return m.client.Del(m.BuildKey(key))
}

func (m *MemCache) Auth(username, password string) error {
	return m.client.Auth(username, password)
}
