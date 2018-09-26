package cryptoutils

import "golang.org/x/crypto/bcrypt"

// NewPwd creates a new pwd with salt
func NewPwd(plainPWD string, cost int) (pwd string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPWD), cost)
	return string(bytes), err
}

// IsPwdOK checks plainPWD and plainPWD are matching
func IsPwdOK(plainPWD, encryptedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPwd), []byte(plainPWD))
	return err == nil
}
