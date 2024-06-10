// store/store.go

package store

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

var mu sync.RWMutex
var inMemoryStore = make(map[string]Item)

type Item struct {
	URL    string
	Expiry time.Time
}

func Set(key, url string, duration time.Duration) {
	mu.Lock()
	defer mu.Unlock()
	inMemoryStore[key] = Item{
		URL:    url,
		Expiry: time.Now().Add(duration),
	}
}

func Get(key string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()
	item, exists := inMemoryStore[key]
	if !exists || item.Expiry.Before(time.Now()) {
		return "", false
	}
	return item.URL, true
}

func Delete(key string) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := inMemoryStore[key]
	if !exists {
		return errors.New("key not found")
	}

	delete(inMemoryStore, key)
	// log.Printf("Deleted entry with key: %s", key) // Add logging
	return nil
}

func GetAllDataFromInMemoryStore() map[string]string {
	mu.RLock()
	defer mu.RUnlock()

	data := make(map[string]string)
	for key, value := range inMemoryStore {
		data[key] = value.URL
	}

	return data
}

func Decr(key string) error {
	mu.Lock()
	defer mu.Unlock()
	item, exists := inMemoryStore[key]
	if !exists || item.Expiry.Before(time.Now()) {
		return errors.New("key not found")
	}

	val, err := strconv.Atoi(item.URL)
	if err != nil {
		return err
	}

	if val <= 0 {
		return errors.New("rate limit exceeded")
	}

	item.URL = strconv.Itoa(val - 1)
	inMemoryStore[key] = item
	return nil
}

func TTL(key string) (time.Duration, error) {
	mu.RLock()
	defer mu.RUnlock()
	item, exists := inMemoryStore[key]
	if !exists {
		return 0, errors.New("key not found")
	}
	return time.Until(item.Expiry), nil
}

func Keys(pattern string) ([]string, error) {
	mu.RLock()
	defer mu.RUnlock()
	keys := make([]string, 0, len(inMemoryStore))
	for k := range inMemoryStore {
		keys = append(keys, k)
	}
	return keys, nil
}