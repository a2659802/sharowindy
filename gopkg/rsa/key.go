package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type PrivateKey struct {
	*rsa.PrivateKey
}
type PublicKey struct {
	*rsa.PublicKey
}

func NewKeys(b []byte) (*PrivateKey, *PublicKey, error) {
	block, _ := pem.Decode(b) //将密钥解析成私钥实例
	if block == nil {
		return nil, nil, errors.New("private key error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, nil, err
	}

	return &PrivateKey{PrivateKey: priv}, &PublicKey{PublicKey: &priv.PublicKey}, nil
}
func (k *PrivateKey) Encrypt(origData []byte) ([]byte, error) {
	return PrivateEncrypt(k.PrivateKey, origData)
}
func (k *PrivateKey) Decrypt(cipher []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, k.PrivateKey, cipher)
}

func (k *PublicKey) Encrypt(origData []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, k.PublicKey, origData)
}

func (k *PublicKey) Decrypt(cipher []byte) ([]byte, error) {
	return PublicDecrypt(k.PublicKey, cipher)
}
