package accounts

import (
	"log"
	"testing"
)

func TestCreateCredential(t *testing.T) {
	log.Println("test Create credential")
	Config([]string{"localhost:2379", "localhost:2380"}, 5)
	var user = "test_user1"
	var password = "test_password1234!@#$"
	if !CreateCredential(user, password) {
		t.Error("test create credential fails")
	}
	log.Println("test Create credential ends")
}

func TestCreateCredentialEmptyUser(t *testing.T) {
	log.Println("test Create credential empty user")
	Config([]string{"localhost:2379", "localhost:2380"}, 5)
	var user = ""
	var password = "test_password1234!@#$"
	if CreateCredential(user, password) {
		t.Error("test create credential empty user fails")
	}
	log.Println("test Create credential empty user ends")
}

func TestVerifyCredential(t *testing.T) {
	log.Println("test Verify credential")
	Config([]string{"localhost:2379", "localhost:2380"}, 5)
	var user = "test_user1"
	var password = "test_password1234!@#$"
	if !CreateCredential(user, password) {
		t.Error("test verify credential fails at Put")
	}
	if !VerifyCredential(user, password) {
		t.Error("test verify credential fails")
	}
	log.Println("test Verify credential ends")
}

func TestVerifyCredentialUserNotFound(t *testing.T) {
	log.Println("test Verify credential user not found")
	Config([]string{"localhost:2379", "localhost:2380"}, 5)
	var user = "test_user123456789"
	var password = "test_password1234!@#$"
	if VerifyCredential(user, password) {
		t.Error("test verify credential user not found fails")
	}
	log.Println("test Verify credential user not found ends")
}

func TestVerifyCredentialEmptyUser(t *testing.T) {
	log.Println("test Verify credential empty user")
	Config([]string{"localhost:2379", "localhost:2380"}, 5)
	var user = ""
	var password = "test_password1234!@#$"
	if VerifyCredential(user, password) {
		t.Error("test verify credential empty user fails")
	}
	log.Println("test Verify credential empty user ends")
}