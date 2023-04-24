package logger

import (
	"fmt"
	"log"
	"os"
)

/*
644 = owner of the file has read and write access,
while the group members and other users on the system only have read access.
This avoids the logs being easily tampered with.
*/
const filePermission = 0644

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
	storeFile     *os.File
)

func init() {
	var err error

	storeFile, err = os.OpenFile("gostore.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePermission)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(storeFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(storeFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(storeFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(storeFile, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Close() {
	var err error
	err = storeFile.Close()
	if err != nil {
		fmt.Printf("Problem closing log file %s", err)
		return
	}
}
