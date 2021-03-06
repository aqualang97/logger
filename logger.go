package logger

import (
	"fmt"
	"os"
	"time"
)

var formatTime = time.Now().Format(time.RFC1123)

//var fileHere *os.File

type Logger struct {
	FileToWrite *os.File
}

func (thisLogger Logger) Info(message string, a ...interface{}) {
	message = "---INFO---\n" + formatTime + "\n" + message + "\n---INFO---\n\n"
	fmt.Fprintf(thisLogger.FileToWrite, message, a...)

}
func (thisLogger Logger) InfoConsole(message string, a ...interface{}) {
	message = "---INFO---\n" + formatTime + "\n" + message + "\n---INFO---\n\n"
	fmt.Printf(message, a...)

}
func (thisLogger Logger) Warning(message string, a ...interface{}) {
	message = "---WARN---\n" + formatTime + "\n" + message + "\n---WARN---\n\n"
	fmt.Fprintf(thisLogger.FileToWrite, message, a...)
}
func (thisLogger Logger) WarningConsole(message string, a ...interface{}) {
	message = "---WARN---\n" + formatTime + "\n" + message + "\n---WARN---\n\n"
	fmt.Printf(message, a...)
}

func (thisLogger Logger) Debug(message string, a ...interface{}) {
	message = "---DEBUG---\n" + formatTime + "\n" + message + "\n---DEBUG---\n\n"
	fmt.Fprintf(thisLogger.FileToWrite, message, a...)
}
func (thisLogger Logger) DebugConsole(message string, a ...interface{}) {
	message = "---DEBUG---\n" + formatTime + "\n" + message + "\n---DEBUG---\n\n"
	fmt.Printf(message, a...)
}

func (thisLogger Logger) Error(message string, a ...interface{}) {
	message = "---ERROR---\n" + formatTime + "\n" + message + "\n---ERROR---\n\n"
	fmt.Fprintf(thisLogger.FileToWrite, message, a...)
}
func (thisLogger Logger) ErrorConsole(message string, a ...interface{}) {
	message = "---ERROR---\n" + formatTime + "\n" + message + "\n---ERROR---\n\n"
	fmt.Printf(message, a...)
}

/*
func TakeFile(file *os.File) {
	fileHere = file
}
*/

func NewLogger(File *os.File, a ...interface{}) Logger {
	return Logger{FileToWrite: File}
} /*

func main() {
	file, err := os.Create("logfile")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	msgInfo := "User %s, payment %.2f\n%s"
	Info(file, msgInfo, "email", 123.123, "SUCCESSFUL")
	msgErr := "User %s has reg err with \n%s"
	Error(file, msgErr, "email", "ERROR")
	msgDebug := "User %s reg success but \n%s"
	Debug(file, msgDebug, "email", "DEBUG MSG")
	msgWarn := "User %s reg success but \n%s"
	Warning(file, msgWarn, "email", "WARNING MSG")

}
*/
