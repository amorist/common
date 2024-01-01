package password

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword 加密用户密码
func EncryptPassword(password, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("EncryptPassword ERR", err)
	}
	return string(hash), err
}

// CheckPassword 检查密码是否正确
func CheckPassword(hashedPassword, newPassword, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(newPassword+salt))
	if err != nil {
		fmt.Println("CheckPassword ERR", err)
		return false
	}
	return true
}

// RandomString 随机生成字符串
// RandomString(8, "A")
// RandomString(8, "a0")
// RandomString(20, "Aa0")
func RandomString(randLength int, randType string) (result string) {
	var num = "0123456789"
	var lower = "abcdefghijklmnopqrstuvwxyz"
	var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := bytes.Buffer{}
	if strings.Contains(randType, "0") {
		b.WriteString(num)
	}
	if strings.Contains(randType, "a") {
		b.WriteString(lower)
	}
	if strings.Contains(randType, "A") {
		b.WriteString(upper)
	}
	var str = b.String()
	var strLen = len(str)
	if strLen == 0 {
		result = ""
		return
	}

	rand.Seed(time.Now().UnixNano())
	b = bytes.Buffer{}
	for i := 0; i < randLength; i++ {
		b.WriteByte(str[rand.Intn(strLen)])
	}
	result = b.String()
	return
}
