package accounts

import (
	"log"
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	log.Println("test Salted Hashing")
	pwd := "GreenDog1234!@#$"
	hash := HashAndSalt([]byte(pwd))
	if len(hash) <= 0 {
		t.Error("Invalid hashing returned")
	}
	log.Println("test Salted Hashing ends", hash)
}

func TestComparePasswords(t *testing.T) {
	log.Println("Salted Hashing Verification")
	pwd := "GreenDog1234!@#$"
	hash := HashAndSalt([]byte(pwd))
	if len(hash) <= 0 {
		t.Error("Invalid hashing returned")
	}
	log.Println("Salted Hash", hash)
	if !ComparePasswords(hash, []byte(pwd)) {
		t.Error("Hash comparison failed.")
	}
	log.Println("Successful Hash compare.")
}

func TestCompareFailure(t *testing.T) {
	log.Println("Negative Salted Hashing Verification")
	pwd := "GreenDog1234!@#$"
	hash := HashAndSalt([]byte(pwd))
	if len(hash) <= 0 {
		t.Error("Invalid hashing returned")
	}
	log.Println("Salted Hash", hash)
	if ComparePasswords(hash, []byte(pwd+"x")) {
		t.Error("Negative Hash comparison failed.")
	}
	log.Println("Successful Negative Hash compare.")

}