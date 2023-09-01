package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/gob"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type JwkKey struct {
	D   string `json:"d"`
	Dp  string `json:"dp"`
	Dq  string `json:"dq"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
	P   string `json:"p"`
	Q   string `json:"q"`
	Qi  string `json:"qi"`
}

func GetJwk() (*JwkKey, error) {
	privPemStr := os.Getenv("PRIVATE_KEY_PEM")
	privPemBytes := []byte(privPemStr)

	block, _ := pem.Decode([]byte(string(privPemBytes)))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}

	rsaKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	key, err := jwk.FromRaw(rsaKey)
	if err != nil {
		return nil, err
	}

	if _, ok := key.(jwk.RSAPrivateKey); !ok {
		return nil, err
	}
	jwk.AssignKeyID(key)

	buf, err := json.MarshalIndent(key, "", "  ")
	if err != nil {
		return nil, err
	}

	jwkKey := JwkKey{}
	json.Unmarshal(buf, &jwkKey)

	return &jwkKey, nil
}

func GenerateJwk() error {
	reader := rand.Reader
	bitSize := 2048
	rsaKey, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return fmt.Errorf("failed to generate rsa key: %s", err.Error())
	}

	publicKey := rsaKey.PublicKey

	err = saveGobKey("private.key", rsaKey)
	if err != nil {
		return err
	}

	err = savePEMKey("private.pem", rsaKey)
	if err != nil {
		return err
	}

	err = saveGobKey("public.key", publicKey)
	if err != nil {
		return err
	}

	err = savePublicPEMKey("public.pem", publicKey)
	if err != nil {
		return err
	}

	return nil
}

func savePEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	privateKey := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	if err != nil {
		return err
	}

	return nil
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) error {
	asn1Bytes, err := asn1.Marshal(pubkey)
	if err != nil {
		return err
	}

	pemkey := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		return err
	}

	return nil
}

func saveGobKey(fileName string, key interface{}) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	if err != nil {
		return err
	}

	return nil
}
