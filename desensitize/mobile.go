package desensitize

import "errors"

func MobileEncrypt(plainMobile string) (string, string, error) {
	cipherMobile, err := rsaEncrypt(plainMobile)
	if err != nil {
		return "", "", err
	}

	desensitizedMobile, err := MobileDesensitize(plainMobile)
	if err != nil {
		return "", "", err
	}

	return cipherMobile, desensitizedMobile, nil

}

func MobileDecrypt(cipherMobile string) (string, error) {
	plainMobile, err := rsaDecrypt(cipherMobile)
	if err != nil {
		return "", err
	}

	return plainMobile, nil
}

func MobileDesensitize(plainMobile string) (string, error) {
	if len(plainMobile) < 11 {
		return "", errors.New("invalid mobile")
	}

	return plainMobile[:3] + "****" + plainMobile[7:], nil
}
