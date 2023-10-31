package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "1234567"
	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}
	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not login")
	}
	fmt.Println("Signed in")
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("bcrypt error from pw: %w", err)
	}
	return bs, nil

}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil

}
