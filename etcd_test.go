package accounts

import (
	"log"
	"testing"
)

func TestPut(t *testing.T) {
	log.Println("testing Put")
	Config([]string{"localhost:2379", "localhost:2380"}, 5)

	err := Put("user1", HashAndSalt([]byte("Password1234!@#$")))
	if err != nil {
		t.Error("Unable to put key/value pair")
	}
	log.Println("testing Put ends")
}

func TestPutNoConnect(t *testing.T) {
	log.Println("testing Put no connect")
	Config([]string{}, 0)
	err := Put("user1", HashAndSalt([]byte("Password1234!@#$")))
	if err == nil {
		t.Error("Put no connect fails")
	}
	log.Println("testing Put no connect ends")
}

func TestPutEmptyKey(t *testing.T) {
	log.Println("testing Put empty key")
	Config([]string{"localhost:2379", "localhost:2380"}, 5)

	err := Put("", HashAndSalt([]byte("Password1234!@#$")))
	if err == nil {
		t.Error("Put empty key fails")
	}
	log.Println("testing Put empty key ends")
}

func TestGet(t *testing.T) {
	log.Println("testing Get")
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
	log.Println("testing Get ends")
}

func TestGetEmptyKey(t *testing.T) {
	log.Println("testing Get empty key")
	Config([]string{"localhost:2379", "localhost:2380"}, 5)

	_, err := Get("")
	if err == nil {
		t.Error("Get empty key fails")
	}
	log.Println("testing Get empty key ends")
}

func TestGetKeyNotFound(t *testing.T) {
	log.Println("testing Get key not found")
	Config([]string{"localhost:2379", "localhost:2380"}, 5)

	result, _ := Get("xyz123")
	log.Println(result)
	if "*** NO KEY ***" != result  {
		t.Error("Get key not found fails")
	}
	log.Println("testing Get key not found ends")
}

func TestGetNoConnect(t *testing.T) {
	log.Println("testing Get no connect")
	Config([]string{}, 0)

	_, err := Get("")
	if err == nil {
		t.Error("Get no connect fails")
	}
	log.Println("testing Get no connect ends")
}