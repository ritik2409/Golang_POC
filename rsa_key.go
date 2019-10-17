package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"

	"github.com/dgrijalva/jwt-go"
)

func main() {

	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	bytes, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	pem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: bytes,
	})
	var claim jwt.Claims

	claim("publickey") = string(pem)

	pem2 := []byte(claim("publickey").(string))
	log.Println(jwt.ParseRSAPublicKeyFromPEM(pem))
}
