package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

var usage = make(map[string]int64)

func byteInUse(username string) int64 {
	return usage[username]
}

const sender = "notifications@example.com"
const password = "example"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d butes of storage, %d%% of your quota.`

func checkQuota(username string) {
	used := byteInUse(username)
	const quota = 1000 * 1000 * 1000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}
