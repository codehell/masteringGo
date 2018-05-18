package main

import (
	"path/filepath"
	"os"
	"log/syslog"
	"log"
	"fmt"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO, programName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO: Loggin in Go!")
	fmt.Println("Wuff!")
	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_MAIL: Loggin in GO")
	fmt.Println("Wuff2!!")

}
