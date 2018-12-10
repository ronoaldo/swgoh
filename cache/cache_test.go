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
	v, ok := c.Get(key)
	if ok {
		t.Errorf("Unexpected positive return from Get call to non-existing key!")
	}
	if v != nil {
		t.Errorf("Expected nil return for non-existing key, got '%v'", v)
	}

	c.Put(key, []byte(value))
	v, ok = c.Get(key)
	t.Logf("Result: %v, %v", string(v), ok)
	if !ok {
		t.Errorf("Expected to get cached value after saving, got %v, %v", v, ok)
	} else {
		if value != string(v) {
			t.Errorf("Excpected to have the same value from Get: %v != %v", string(value), v)
		}
	}

	time.Sleep(time.Second)
	v, ok = c.Get(key)
	if ok {
		t.Errorf("Expected expired item to not return true for Get")
	}
	if v != nil {
		t.Errorf("Unexpected result from %v: ", v)
	}
}
