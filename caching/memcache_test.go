package caching

import (
	"github.com/bradfitz/gomemcache/memcache"
	"testing"
	"time"
)

var client = NewMemCache("127.0.0.1:11211")

func TestMemCache_GetServers(t *testing.T) {
	if client.servers[0] != "127.0.0.1:11211" {
		t.Error("GetServers error")
		return
	}
}

func TestMemCache_SetValue(t *testing.T) {
	if err := client.SetValue("demo", []byte("1"), time.Second*5); err != nil {
		t.Error("GetValue", err)
		return
	}
}

func TestMemCache_GetValue(t *testing.T) {
	_, err := client.GetValue("demo")
	if err != nil {
		t.Error("GetValue", err)
		return
	}
	time.Sleep(time.Second * 5)
	if val, err := client.GetValue("demo"); err == nil && val != nil {
		t.Error("GetValue Expires failed")
		return
	}
}

func TestMemCache_DeleteValue(t *testing.T) {
	if err := client.DeleteValue("demo"); err != nil && err != memcache.ErrCacheMiss {
		t.Error("DeleteValue", err)
		return
	}
}
func TestMemCache_MultiGet(t *testing.T) {
	keys := []string{"demo1", "demo2", "demo3"}

	for _, key := range keys {
		client.SetValue(key, []byte(key), time.Minute)
	}

	keys = append(keys, "demo4")
	vals, err := client.MultiGet(keys)
	if err != nil {
		t.Error("MultiGet", err)
		return
	}
	if vals["demo4"] != nil {
		t.Error("MultiGet demo4 has value")
	}
}
