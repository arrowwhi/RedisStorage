package main

import (
	"fmt"
	"time"
)

func main() {
	store := NewStore()

	// Adding items to the store with different TTLs
	store.Set("key1", "value1", 5*time.Second)
	store.Set("key2", "value2", 10*time.Second)
	store.Set("key3", "value3", 0)  // No TTL
	store.Set("key4", "value4", -1) // TTL less than 0, ignored

	// Getting items from the store
	val1, found := store.Get("key1")
	if found {
		fmt.Println("Value of key1:", val1)
	} else {
		fmt.Println("Key1 not found")
	}

	val2, found := store.Get("key2")
	if found {
		fmt.Println("Value of key2:", val2)
	} else {
		fmt.Println("Key2 not found")
	}

	val3, found := store.Get("key3")
	if found {
		fmt.Println("Value of key3:", val3)
	} else {
		fmt.Println("Key3 not found")
	}

	// Waiting for some time to let the item with TTL expire
	time.Sleep(6 * time.Second)

	// Checking that the item with expired TTL was deleted
	val1AfterExpire, found := store.Get("key1")
	if found {
		fmt.Println("Value of key1 after expiration:", val1AfterExpire)
	} else {
		fmt.Println("Key1 not found after expiration")
	}

	// Deleting an item from the store
	store.Delete("key2")
	_, found = store.Get("key2")
	if !found {
		fmt.Println("Key2 deleted successfully")
	} else {
		fmt.Println("Key2 deletion failed")
	}
}
