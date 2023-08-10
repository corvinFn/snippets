package main

import "fmt"

type def struct {
}

func (s *def) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("DEF: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *def) saveOTPCache(otp string) {
	fmt.Printf("DEF: saving otp: %s to cache\n", otp)
}

func (s *def) getMessage(otp string) string {
	return "DEF OTP for login is " + otp
}

func (s *def) sendNotification(message string) error {
	fmt.Printf("DEF: sending def: %s\n", message)
	return nil
}

func (s *def) publishMetric() {
	fmt.Printf("DEF: publishing metrics\n")
}
