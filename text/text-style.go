package text

import (
	"fmt"
)

func Colorize(text string, color Color) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}

func ColorizeBackground(text string, background BackgroundColor, color Color) string {
	return fmt.Sprintf("%s%s%s%s", background, color, text, Reset)
}

func ColorizeBold(text string, color ColorBold) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}

func ColorizeUnderline(text string, colorUnderline UnderLine) string {
	return fmt.Sprintf("%s%s%s", colorUnderline, text, Reset)
}

func ColorizeBoldUnderline(text string, color UnderLineBold) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}
