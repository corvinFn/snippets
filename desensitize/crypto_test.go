package desensitize

import (
	"testing"
)

func TestRsaEncode(t *testing.T) {
	// Init("conf")
	dt := "hello debugger!"
	encryptStr, err := rsaEncrypt(dt)
	if err != nil {
		t.Fatal(err)
	}

	decryptStr, err := rsaDecrypt(encryptStr)
	if err != nil {
		t.Fatal(err)
	}

	if dt != decryptStr {
		t.Fatalf("want:%s, got:%s", dt, decryptStr)
	}
}
