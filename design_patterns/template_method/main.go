package main

import "fmt"

func main() {
	// otp := otp{}

	// smsOTP := &sms{
	//  otp: otp,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &email{
	//  otp: otp,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()
	o := otp{
		iOtp: &sms{},
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	o = otp{
		iOtp: &email{},
	}
	o.genAndSendOTP(4)

}
