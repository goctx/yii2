package caching

import (
	"strings"
	"time"
	"xialeistudio/yii2/helpers/crypto"
	"xialeistudio/yii2/helpers/ctype"
)

type Cache struct {
	KeyPrefix string
}

func (m Cache) BuildKey(key string) string {
	if ctype.AlNum(key) && len(key) <= 32 {
		key = m.KeyPrefix + key
	} else {
		key = m.KeyPrefix + crypto.Md5(strings.NewReader(key))
	}
	return key
}

type CacheInterface interface {
	GetValue(key string) ([]byte, error)

	SetValue(key string, val []byte, duration time.Duration) (error)

	DeleteValue(key string) (error)

	MultiGet(keys []string) (map[string][]byte, error)
}
