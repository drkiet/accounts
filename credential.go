package accounts

import "log"

func CreateCredential(user, password string) bool {
	log.Println(user, password)
	err := Put(user, HashAndSalt([]byte(password)))
	if err != nil {
		return false
	}
	return true
}

func VerifyCredential(user, password string) bool {
	hash, err := Get(user)
	log.Println(user, hash)
	if err != nil {
		log.Println(err)
		return false
	}
	if !ComparePasswords(hash, []byte(password)) {
		return false
	}
	return true
}
