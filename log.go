package goofy

import (
	"io"
	"log"
	"os"
)

var InfoLog = log.New(os.Stdout, "INFO: ", log.Lshortfile)
var DebugLog = log.New(io.Discard, "DEBUG: ", log.Lshortfile)
var ErrLog = log.New(os.Stderr, "ERROR: ", log.Lshortfile)
