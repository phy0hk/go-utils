package types

type Color string

const (
	Red           Color = "\033[31m"
	Green         Color = "\033[32m"
	Yellow        Color = "\033[33m"
	Blue          Color = "\033[34m"
	Magenta       Color = "\033[35m"
	Cyan          Color = "\033[36m"
	White         Color = "\033[37m"
	BrightRed     Color = "\033[91m"
	BrightGreen   Color = "\033[92m"
	BrightYellow  Color = "\033[93m"
	BrightBlue    Color = "\033[94m"
	BrightMagenta Color = "\033[95m"
	BrightCyan    Color = "\033[96m"
	BrightWhite   Color = "\033[97m"
	Reset         Color = "\033[0m"
)

type UnderLine string

const (
	UnderlineRed           UnderLine = "\033[4;31m"
	UnderlineGreen         UnderLine = "\033[4;32m"
	UnderlineYellow        UnderLine = "\033[4;33m"
	UnderlineBlue          UnderLine = "\033[4;34m"
	UnderlineMagenta       UnderLine = "\033[4;35m"
	UnderlineCyan          UnderLine = "\033[4;36m"
	UnderlineWhite         UnderLine = "\033[4;37m"
	UnderlineBrightRed     UnderLine = "\033[4;91m"
	UnderlineBrightGreen   UnderLine = "\033[4;92m"
	UnderlineBrightYellow  UnderLine = "\033[4;93m"
	UnderlineBrightBlue    UnderLine = "\033[4;94m"
	UnderlineBrightMagenta UnderLine = "\033[4;95m"
	UnderlineBrightCyan    UnderLine = "\033[4;96m"
	UnderlineBrightWhite   UnderLine = "\033[4;97m"
)

type BackgroundColor string

const (
	BackgroundBlack         BackgroundColor = "\033[40m"
	BackgroundRed           BackgroundColor = "\033[41m"
	BackgroundGreen         BackgroundColor = "\033[42m"
	BackgroundYellow        BackgroundColor = "\033[43m"
	BackgroundBlue          BackgroundColor = "\033[44m"
	BackgroundMagenta       BackgroundColor = "\033[45m"
	BackgroundCyan          BackgroundColor = "\033[46m"
	BackgroundWhite         BackgroundColor = "\033[47m"
	BackgroundBrightRed     BackgroundColor = "\033[48;5;196m"
	BackgroundBrightGreen   BackgroundColor = "\033[48;5;118m"
	BackgroundBrightYellow  BackgroundColor = "\033[48;5;226m"
	BackgroundBrightBlue    BackgroundColor = "\033[48;5;75m"
	BackgroundBrightMagenta BackgroundColor = "\033[48;5;201m"
	BackgroundBrightCyan    BackgroundColor = "\033[48;5;117m"
	BackgroundBrightWhite   BackgroundColor = "\033[48;5;231m"
)

type ColorBold string

const (
	BoldRed           ColorBold = "\033[1;31m"
	BoldGreen         ColorBold = "\033[1;32m"
	BoldYellow        ColorBold = "\033[1;33m"
	BoldBlue          ColorBold = "\033[1;34m"
	BoldMagenta       ColorBold = "\033[1;35m"
	BoldCyan          ColorBold = "\033[1;36m"
	BoldWhite         ColorBold = "\033[1;37m"
	BoldBrightRed     ColorBold = "\033[1;91m"
	BoldBrightGreen   ColorBold = "\033[1;92m"
	BoldBrightYellow  ColorBold = "\033[1;93m"
	BoldBrightBlue    ColorBold = "\033[1;94m"
	BoldBrightMagenta ColorBold = "\033[1;95m"
	BoldBrightCyan    ColorBold = "\033[1;96m"
	BoldBrightWhite   ColorBold = "\033[1;97m"
)

type UnderLineBold string

const (
	UnderlineBoldRed           UnderLineBold = "\033[4;1;31m"
	UnderlineBoldGreen         UnderLineBold = "\033[4;1;32m"
	UnderlineBoldYellow        UnderLineBold = "\033[4;1;33m"
	UnderlineBoldBlue          UnderLineBold = "\033[4;1;34m"
	UnderlineBoldMagenta       UnderLineBold = "\033[4;1;35m"
	UnderlineBoldCyan          UnderLineBold = "\033[4;1;36m"
	UnderlineBoldWhite         UnderLineBold = "\033[4;1;37m"
	UnderlineBoldBrightRed     UnderLineBold = "\033[4;1;91m"
	UnderlineBoldBrightGreen   UnderLineBold = "\033[4;1;92m"
	UnderlineBoldBrightYellow  UnderLineBold = "\033[4;1;93m"
	UnderlineBoldBrightBlue    UnderLineBold = "\033[4;1;94m"
	UnderlineBoldBrightMagenta UnderLineBold = "\033[4;1;95m"
	UnderlineBoldBrightCyan    UnderLineBold = "\033[4;1;96m"
	UnderlineBoldBrightWhite   UnderLineBold = "\033[4;1;97m"
)
