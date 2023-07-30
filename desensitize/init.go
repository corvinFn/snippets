package desensitize

import (
	"crypto/rsa"
	"crypto/x509"
	"embed"
	"encoding/pem"
)

//go:embed conf/*.pem
var pemFiles embed.FS

func init() {
	once.Do(func() {
		genPrivateKey1(pemFiles, "conf/private.pem")
		genPublicKey1(pemFiles, "conf/public.pem")
	})
}

func genPrivateKey1(f embed.FS, name string) {
	buf, err := f.ReadFile(name)
	if err != nil {
		panic(err)
	}
	// pem解码
	block, _ := pem.Decode(buf)
	// X509解码
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
}

func genPublicKey1(f embed.FS, name string) {
	buf, err := f.ReadFile(name)
	if err != nil {
		panic(err)
	}
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
