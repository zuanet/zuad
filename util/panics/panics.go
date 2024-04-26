package panics

import (
	"fmt"
	"github.com/zuanet/zuad/infrastructure/logger"
	"os"
	"runtime/debug"
	"sync/atomic"
	"time"
)

const exitHandlerTimeout = 5 * time.Second

// HandlePanic recovers panics and then initiates a clean shutdown.
func HandlePanic(log *logger.Logger, zuaoutineName string, zuaoutineStackTrace []byte) {
	err := recover()
	if err == nil {
		return
	}

	reason := fmt.Sprintf("Fatal error in zuaoutine `%s`: %+v", zuaoutineName, err)
	exit(log, reason, debug.Stack(), zuaoutineStackTrace)
}

var zuaoutineLastID uint64

// ZuaoutineWrapperFunc returns a zuaoutine wrapper function that handles panics and writes them to the log.
func ZuaoutineWrapperFunc(log *logger.Logger) func(name string, spawnedFunction func()) {
	return func(name string, f func()) {
		stackTrace := debug.Stack()
		go func() {
			handleSpawnedFunction(log, stackTrace, name, f)
		}()
	}
}

// AfterFuncWrapperFunc returns a time.AfterFunc wrapper function that handles panics.
func AfterFuncWrapperFunc(log *logger.Logger) func(name string, d time.Duration, f func()) *time.Timer {
	return func(name string, d time.Duration, f func()) *time.Timer {
		stackTrace := debug.Stack()
		return time.AfterFunc(d, func() {
			handleSpawnedFunction(log, stackTrace, name, f)
		})
	}
}

// Exit prints the given reason to log and initiates a clean shutdown.
func Exit(log *logger.Logger, reason string) {
	exit(log, reason, nil, nil)
}

// Exit prints the given reason, prints either of the given stack traces (if not nil),
// waits for them to finish writing, and exits.
func exit(log *logger.Logger, reason string, currentThreadStackTrace []byte, zuaoutineStackTrace []byte) {
	exitHandlerDone := make(chan struct{})
	go func() {
		log.Criticalf("Exiting: %s", reason)
		if zuaoutineStackTrace != nil {
			log.Criticalf("Zuaoutine stack trace: %s", zuaoutineStackTrace)
		}
		if currentThreadStackTrace != nil {
			log.Criticalf("Stack trace: %s", currentThreadStackTrace)
		}
		log.Backend().Close()
		close(exitHandlerDone)
	}()

	select {
	case <-time.After(exitHandlerTimeout):
		fmt.Fprintln(os.Stderr, "Couldn't exit gracefully.")
	case <-exitHandlerDone:
	}
	fmt.Println("Exiting...")
	os.Exit(1)
}

func handleSpawnedFunction(log *logger.Logger, stackTrace []byte, spawnedFunctionName string, spawnedFunction func()) {
	zuaoutineID := atomic.AddUint64(&zuaoutineLastID, 1)
	zuaoutineName := fmt.Sprintf("%s %d", spawnedFunctionName, zuaoutineID)
	utilLog.Tracef("Started zuaoutine `%s`", zuaoutineName)
	defer utilLog.Tracef("Ended zuaoutine `%s`", zuaoutineName)
	defer HandlePanic(log, zuaoutineName, stackTrace)
	spawnedFunction()
}
