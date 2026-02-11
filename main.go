package main

import (
	"fmt"

	"github.com/phy0hk/go-utils/config"
	"github.com/phy0hk/go-utils/text"
	"github.com/phy0hk/go-utils/types"
)

func main() {
	config.SetVerbose(true)
	temp := text.ColorizeBoldUnderline("Hello World", types.UnderlineBoldGreen)
	fmt.Print(temp)
}
