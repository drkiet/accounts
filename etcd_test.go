package accounts

import "testing"

func TestPut(t *testing.T) {
	Config([]string{"localhost:2379", "localhost:2380"}, 5)

	err := Put("key1", "value1")
	if err != nil {
		t.Error("Unable to put key/value pair")
	}
}

func TestGet(t *testing.T) {
	Config([]string{"localhost:2379", "localhost:2380"}, 5)

	err := Put("key2", "value2")
	if err != nil {
		t.Error("Unable to put key/value pair")
	}
	val, err := Get("key2")
	if err != nil {
		t.Error("Unable to put key/value pair")
	}
	if val != "value2" {
		t.Error("Unable to get value by a key")
	}
}
