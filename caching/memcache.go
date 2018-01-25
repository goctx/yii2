package caching

import (
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

type MemCache struct {
	Cache
	client  *memcache.Client
	servers []string
}

func NewMemCache(keyPrefix string, maxConn int, servers ...string) *MemCache {
	client := memcache.New(servers...)
	m := &MemCache{
		client:  client,
		servers: servers,
	}
	m.KeyPrefix = keyPrefix
	m.client.MaxIdleConns = maxConn
	return m
}

func (m *MemCache) GetMemcache() *memcache.Client {
	return m.client
}

func (m *MemCache) GetServers() []string {
	return m.servers
}

func (m *MemCache) GetValue(key string) ([]byte, error) {
	if item, err := m.client.Get(m.BuildKey(key)); err != nil {
		return nil, err
	} else {
		return item.Value, nil
	}
}

func (m *MemCache) SetValue(key string, val []byte, duration time.Duration) error {
	item := &memcache.Item{
		Key:        m.BuildKey(key),
		Value:      val,
		Expiration: int32(duration.Seconds()),
	}
	return m.client.Set(item)
}

func (m *MemCache) DeleteValue(key string) error {
	return m.client.Delete(m.BuildKey(key))
}

func (m *MemCache) MultiGet(keys []string) (map[string][]byte, error) {
	data := make(map[string][]byte)
	queryKeys := make([]string, 0)
	for i := range keys {
		queryKeys = append(queryKeys, m.BuildKey(keys[i]))
	}
	items, err := m.client.GetMulti(queryKeys)
	if err != nil {
		return nil, err
	}
	for key, item := range items {
		for i, queryKey := range queryKeys {
			if key == queryKey {
				data[keys[i]] = item.Value
				break
			}
		}
	}
	return data, nil
}
