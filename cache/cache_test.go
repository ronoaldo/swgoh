package cache

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestBasicCacheOperations(t *testing.T) {
	file, err := ioutil.TempFile("", "cache")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	file.Close()
	t.Logf("Using cache file: %s", file.Name())

	key := "my-key"
	value := "my-value"

	c := NewCache(file.Name(), time.Second)
	var result string
	ok := c.Get(key, &result)
	if ok {
		t.Errorf("Unexpected positive return from Get call to non-existing key!")
	}
	if result != "" {
		t.Errorf("Expected nil return for non-existing key, got '%v'", result)
	}

	c.Put(key, value)
	ok = c.Get(key, &result)
	t.Logf("Result: %v, %v", result, ok)
	if !ok {
		t.Errorf("Expected to get cached value after saving, got %v, %v", result, ok)
	} else {
		if value != result {
			t.Errorf("Excpected to have the same value from Get: %v != %v", value, result)
		}
	}

	result = ""
	time.Sleep(time.Second)
	ok = c.Get(key, result)
	if ok {
		t.Errorf("Expected expired item to not return true for Get")
	}
	if result != "" {
		t.Errorf("Unexpected result from %v: ", result)
	}
}
