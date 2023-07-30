package desensitize

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path"
	"sync"
)

var (
	once       sync.Once
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func Init(filepath string) {
	once.Do(func() {
		genPrivateKey(path.Join(filepath, "private.pem"))
		genPublicKey(path.Join(filepath, "public.pem"))
	})
}

func genPrivateKey(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// pem解码
	block, _ := pem.Decode(buf)
	// X509解码
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
}

func genPublicKey(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// pem解码
	block, _ := pem.Decode(buf)
	// x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	publicKey = publicKeyInterface.(*rsa.PublicKey)
}

func rsaEncrypt(plainText string) (string, error) {
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
	if err != nil {
		return "", err
	}

	return string(cipherText), nil
}

func rsaDecrypt(cipherText string) (string, error) {
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, []byte(cipherText))
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
