package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func generateKeyPair(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func encryptMessage(publicKey *rsa.PublicKey, message string) (string, error) {
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(message), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

func decryptMessage(privateKey *rsa.PrivateKey, encryptedMessage string) (string, error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", err
	}
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedBytes, nil)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}

func exportPrivateKeyAsPEM(privateKey *rsa.PrivateKey) ([]byte, error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	return pem.EncodeToMemory(block), nil
}

func exportPublicKeyAsPEM(publicKey *rsa.PublicKey) ([]byte, error) {
	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	return pem.EncodeToMemory(block), nil
}

func main() {

	privateKey, err := generateKeyPair(2048)
	if err != nil {
		fmt.Println("Erro ao gerar chave:", err)
		return
	}
	publicKey := &privateKey.PublicKey

	privatePEM, _ := exportPrivateKeyAsPEM(privateKey)
	publicPEM, _ := exportPublicKeyAsPEM(publicKey)

	fmt.Printf("Chave privada:\n%s\n", privatePEM)
	fmt.Printf("Chave pública:\n%s\n", publicPEM)

	message := "Esta é uma mensagem secreta!"

	encryptedMessage, err := encryptMessage(publicKey, message)
	if err != nil {
		fmt.Println("Erro ao cifrar a mensagem:", err)
		return
	}
	fmt.Printf("Mensagem cifrada: %s\n", encryptedMessage)

	decryptedMessage, err := decryptMessage(privateKey, encryptedMessage)
	if err != nil {
		fmt.Println("Erro ao decifrar a mensagem:", err)
		return
	}
	fmt.Printf("Mensagem decifrada: %s\n", decryptedMessage)
}
