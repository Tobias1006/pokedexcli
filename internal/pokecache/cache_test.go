package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGetLocation(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add("loc", c.key, c.val)
			val, ok := cache.Get("loc", c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}
func TestAddGetPokemon(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add("pok", c.key, c.val)
			val, ok := cache.Get("pok", c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoopLocation(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("loc", "https://example.com", []byte("testdata"))

	_, ok := cache.Get("loc", "https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("loc", "https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestReapLoopLongWaitLocation(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = (baseTime * 100) + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("loc", "https://example.com", []byte("testdata"))

	_, ok := cache.Get("loc", "https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("loc", "https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
func TestReapLoopPokemon(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("pok", "https://example.com", []byte("testdata"))

	_, ok := cache.Get("pok", "https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("pok", "https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestReapLoopLongWaitPokemon(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = (baseTime * 100) + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("pok", "https://example.com", []byte("testdata"))

	_, ok := cache.Get("pok", "https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("pok", "https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
