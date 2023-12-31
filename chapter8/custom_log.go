package main

import (
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log files:", err)
	}

	logFlags := log.Ldate | log.Ltime | log.Lshortfile
	Trace = log.New(io.Discard, "TRACE: ", logFlags)
	Info = log.New(os.Stdout, "INFO: ", logFlags)
	Warning = log.New(os.Stdout, "TRACE: ", logFlags)
	Error = log.New(io.MultiWriter(file, os.Stderr), "TRACE: ", logFlags)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
