package main

import (
	"fmt"

	"github.com/phy0hk/go-utils/config"
	"github.com/phy0hk/go-utils/text"
)

func main() {
	config.SetVerbose(true)
	temp := text.ColorizeBoldUnderline("Hello World", text.UnderlineBoldGreen)
	fmt.Print(temp)
}
