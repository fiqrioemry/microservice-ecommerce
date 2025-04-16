package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomUserAvatar() string {
	avatars := []string{
		"Shopia", "Jameson", "Emery", "Sawyer", "Maria",
		"Chase", "Jocelyn", "Liliana", "Robert", "Christian", "Nolan",
	}

	rand.Seed(time.Now().UnixNano())
	avatar := avatars[rand.Intn(len(avatars))]

	return fmt.Sprintf("https://api.dicebear.com/9.x/fun-emoji/svg?seed=%s", avatar)
}

func GenerateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	otp := ""
	for i := 0; i < length; i++ {
		otp += string(digits[rand.Intn(len(digits))])
	}
	return otp
}
