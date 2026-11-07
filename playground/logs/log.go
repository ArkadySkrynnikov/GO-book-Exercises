package logs

import (
	"bytes"
	"fmt"
	"log"
)

func LogsTry() {
	var (
		buf      bytes.Buffer
		myLogger = log.New(&buf, "logger: ", log.Lshortfile)
	)
	myLogger.SetPrefix("pidoras/")
	myLogger.Print("myLogger")

	fmt.Print(&buf)
}
