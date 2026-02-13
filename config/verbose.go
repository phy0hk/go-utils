package config

import (
	"log"

	"github.com/phy0hk/go-utils/text"
)

var IsVerbose = new(bool)

func SetVerbose(verbose bool) {
	*IsVerbose = verbose
}
func VerboseLog(message string, color text.Color) {
	if *IsVerbose {
		log.Printf("%s%s%s\n", color, message, text.Reset)
	}
}
