package main

import "C"
import (
	log "github.com/sirupsen/logrus"
	"go_client/outline"
	"sync"
	"os"
)

var outlineClient *outline.OutlineClient
var outlineMu sync.Mutex
var logFile *os.File

//export StartOutline
func StartOutline(key *C.char) {
    logFile, err := os.OpenFile("outline.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
    	log.Fatalf("Failed to open log file: %v", err)
    }
    log.SetOutput(logFile)

	str_key := C.GoString(key)
	keyPtr := &str_key

	outlineMu.Lock()
	defer outlineMu.Unlock()

	if outlineClient != nil {
		log.Infof("Disconnect existing outline client")
		err = outlineClient.Disconnect()
		if err != nil {
			log.Errorf("Failed to disconnect existing outline client: %v", err)
			return
		}
	}

	outlineClient = outline.NewClient(keyPtr)
	log.Infof("Connect outline client")
	err = outlineClient.Connect()
	if err != nil {
		log.Errorf("Failed to connect outline client: %v", err)
	}
}

//export StopOutline
func StopOutline() {
	outlineMu.Lock()
	defer outlineMu.Unlock()
	if outlineClient != nil {
		log.Infof("Disconnect outline client")
		err := outlineClient.Disconnect()
		if err != nil {
			log.Errorf("Failed to disconnect outline client: %v", err)
		}
		outlineClient = nil
	}
}
