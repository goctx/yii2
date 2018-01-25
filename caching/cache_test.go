package caching

import "testing"

func TestCache_BuildKey(t *testing.T) {
	cache := new(Cache)
	cache.KeyPrefix = "d01"
	if "d01adc41d68149635b68228d9dfc7cc51bf" != cache.BuildKey("oHAf2t6xcRL0DS5-fPKhRWURVLl0") {
		t.Error("cache buildKey failed")
		return
	}
}
