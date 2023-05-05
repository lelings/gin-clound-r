package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

func GetUUid() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatal("[GetUUid error]:", err)
	}
	return uuid.String()
}

func GetPassword(password string) string {
	s := md5.New()
	s.Write([]byte(password))
	return hex.EncodeToString(s.Sum(nil))
}

func GetCode() string {
	rand.Seed(time.Now().Unix())
	code := ""
	for i := 0; i < 4; i++ {
		n := rand.Intn(10)
		temp := strconv.Itoa(n)
		code += temp
	}
	fmt.Println(code)
	return code
}
