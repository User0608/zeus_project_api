package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var (
	singKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

func LoadFiles(privateFilePath, publicFilePath string) error {
	var err error
	once.Do(func() {
		err = loadFiles(privateFilePath, publicFilePath)
	})
	return err
}

func loadFiles(privateFile, publicFile string) error {
	privateBytes, err := ioutil.ReadFile(privateFile)
	if err != nil {
		return err
	}
	publicBytes, err := ioutil.ReadFile(publicFile)
	if err != nil {
		return err
	}
	return parseRSA(privateBytes, publicBytes)
}

func parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	if singKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes); err != nil {
		return err
	}
	if verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes); err != nil {
		return err
	}
	return nil
}
