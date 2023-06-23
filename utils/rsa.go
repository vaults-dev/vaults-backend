package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func GetRsaPrivateKey() (*rsa.PrivateKey, error) {
	privPemBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(string(privPemBytes)))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	rsaKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("ssh.ParsePKCS1PrivateKey error: %s", err.Error())
	}

	return rsaKey, nil
}

func GetRsaPublicKey() (*rsa.PublicKey, error) {
	pubPemBytes, err := ioutil.ReadFile("public.pem")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(string(pubPemBytes)))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	rsaKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("ssh.ParsePKCS1PublicKey error: %s", err.Error())
	}

	return rsaKey, nil
}
