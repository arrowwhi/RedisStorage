package main

import (
	"testing"
	"time"
)

func TestSetAndGet(t *testing.T) {
	store := NewStore()

	store.Set("key1", "value1", time.Second*2)
	store.Set("key2", "value2", time.Second*2)

	val, ok := store.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("Expected 'value1' for key 'key1', got %v", val)
	}

	val, ok = store.Get("key2")
	if !ok || val != "value2" {
		t.Errorf("Expected 'value2' for key 'key2', got %v", val)
	}

	val, ok = store.Get("key3")
	if ok {
		t.Errorf("Expected 'false' for key 'key3', got %v", ok)
	}
}

func TestExpiration(t *testing.T) {
	store := NewStore()

	store.Set("key1", "value1", time.Second*1)

	time.Sleep(time.Second * 2)

	val, ok := store.Get("key1")
	if ok || val != nil {
		t.Errorf("Expected 'false' for key 'key1' after expiration, got %v", ok)
	}
}

func TestDelete(t *testing.T) {
	store := NewStore()

	store.Set("key1", "value1", time.Second*2)

	store.Delete("key1")

	val, ok := store.Get("key1")
	if ok || val != nil {
		t.Errorf("Expected 'false' for key 'key1' after deletion, got %v", ok)
	}
}

func TestConcurrentAccess(t *testing.T) {
	store := NewStore()

	// Concurrently set and get a key
	go func() {
		store.Set("key1", "value1", time.Second*2)
	}()

	go func() {
		time.Sleep(time.Millisecond * 500) // Wait for store.Set to finish
		val, ok := store.Get("key1")
		if !ok || val != "value1" {
			t.Errorf("Expected 'value1' for key 'key1', got %v", val)
		}
	}()

	// Concurrently delete a key
	go func() {
		store.Set("key2", "value2", time.Second*2)
	}()

	go func() {
		time.Sleep(time.Millisecond * 500) // Wait for store.Set to finish
		store.Delete("key2")
	}()

	// Concurrently set and delete a key
	go func() {
		store.Set("key3", "value3", time.Second*2)
	}()

	go func() {
		time.Sleep(time.Millisecond * 500) // Wait for store.Set to finish
		store.Delete("key3")
	}()

	time.Sleep(time.Second) // Wait for goroutines to finish
}

func TestZeroTTL(t *testing.T) {
	store := NewStore()

	store.Set("key1", "value1", 0)

	val, ok := store.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("Expected 'value1' for key 'key1', got %v", val)
	}
}

func TestNegativeTTL(t *testing.T) {
	store := NewStore()

	store.Set("key1", "value1", -time.Second)

	val, ok := store.Get("key1")
	if ok || val != nil {
		t.Errorf("Expected 'false' for key 'key1' with negative TTL, got %v", ok)
	}
}

func TestEmptyStore(t *testing.T) {
	store := NewStore()

	val, ok := store.Get("key1")
	if ok || val != nil {
		t.Errorf("Expected 'false' for key 'key1' in an empty store, got %v", ok)
	}

	store.Delete("key1") // Deleting non-existing key shouldn't cause an error
}
