package logger

import (
	"errors"
	"log"
	"os"
)

func DemoV1() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
	}
}

func DemoV2(logger *log.Logger) {
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
	}
}

// Call this w/ :
// No longer care about logger
//  logger := log.New(...)
//  DemoV3(log.Println)
func DemoV3(logFn func(...interface{})) {
	err := doTheThing()
	if err != nil {
		logFn("error in doTheThing():", err)
	}
}

// For DemoV4
type Logger interface {
	Println(...interface{})
	Printf(string, ...interface{})
}

// Call this w/
//  logger := log.New(...)
//  DemoV4(logger)
func DemoV4(logger Logger) {
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
		logger.Printf("error: %s\n", err)
	}
}

type Thing struct {
	Logger interface {
		Println(...interface{})
		Printf(string, ...interface{})
	}
}

func (t Thing) DemoV5() {
	err := doTheThing()
	if err != nil {
		t.Logger.Println("error in doTheThing():", err)
		t.Logger.Println("error: %s\n", err)
	}
}

func doTheThing() error {
	//return nil
	return errors.New("error opening file: abc.txt")
}
