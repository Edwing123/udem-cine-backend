package main

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "[INFO]", log.LstdFlags)
	errorLogger = log.New(os.Stderr, "[INFO]", log.LstdFlags)
)
