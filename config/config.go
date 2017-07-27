package config

import (
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"
	"crypto/rsa"
	"log"
)

var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)

const (
	// For simplicity these files are in the same folder as the app binary.
	// You shouldn't do this in production.
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

