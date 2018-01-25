package caching

import (
	"github.com/bradfitz/gomemcache/memcache"
	"testing"
	"time"
)

var client, _ = NewMemCache("", "127.0.0.1:11211")

func TestMemCache_SetValue(t *testing.T) {
	if err := client.SetValue("demo", []byte("1"), time.Second*5); err != nil {
		t.Error("GetValue", err)
		return
	}
}

func TestMemCache_GetValue(t *testing.T) {
	val, err := client.GetValue("demo")
	if err != nil || string(val) != "1" {
		t.Error("GetValue", err)
		return
	}
}

func TestMemCache_DeleteValue(t *testing.T) {
	if err := client.DeleteValue("demo"); err != nil && err != memcache.ErrCacheMiss {
		t.Error("DeleteValue", err)
		return
	}
}
