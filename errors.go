package main

import (
	"fmt"
	"log"
)

func fatalErrorCheck(err error, customMessage string) {
	if err != nil {
		fmt.Println(customMessage)
		log.Fatal(err)
	}
}

func bytesWrittenCheck(bytesWritten, bytesExpected int) {
	if bytesWritten != bytesExpected {
		fmt.Println("Inncorrect number of bytes written to file")
	}
}
