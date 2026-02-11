package config

import (
	"log"

	"github.com/phy0hk/go-utils/types"
)

var IsVerbose = new(bool)

func SetVerbose(verbose bool) {
	*IsVerbose = verbose
}
func VerboseLog(message string, color types.Color) {
	if *IsVerbose {
		log.Printf("%s%s%s\n", color, message, types.Reset)
	}
}
