package main

import (
	"encoding/base32"
	"fmt"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func main() {
	secret := []byte("solution.itlab@gmail.comHENNGECHALLENGE003")
	code, err := totp.GenerateCodeCustom(base32.StdEncoding.EncodeToString(secret), time.Now(), totp.ValidateOpts{
		Algorithm: otp.AlgorithmSHA512,
		Digits:    10,
	})
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Your key, sir.\n%v", code)
}
