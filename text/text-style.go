package text

import (
	"fmt"

	"github.com/phy0hk/go-utils/types"
)

func Colorize(text string, color types.Color) string {
	return fmt.Sprintf("%s%s%s", color, text, types.Reset)
}

func ColorizeBackground(text string, background types.BackgroundColor, color types.Color) string {
	return fmt.Sprintf("%s%s%s%s", background, color, text, types.Reset)
}

func ColorizeBold(text string, color types.ColorBold) string {
	return fmt.Sprintf("%s%s%s", color, text, types.Reset)
}

func ColorizeUnderline(text string, colorUnderline types.UnderLine) string {
	return fmt.Sprintf("%s%s%s", colorUnderline, text, types.Reset)
}

func ColorizeBoldUnderline(text string, color types.UnderLineBold) string {
	return fmt.Sprintf("%s%s%s", color, text, types.Reset)
}
