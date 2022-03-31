package color_print

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	Success = color.New(color.FgHiGreen).SprintFunc()
	SuccessBackground = color.New(color.BgGreen).SprintFunc()

	Fail = color.New(color.FgHiRed).SprintFunc()
	FailBackground = color.New(color.BgRed).SprintFunc()

	Warn = color.New(color.FgHiYellow).SprintFunc()
	WarnBackground = color.New(color.BgYellow).SprintFunc()

	Info = color.New(color.FgHiWhite).SprintFunc()
	InfoBackground = color.New(color.BgWhite).SprintFunc()
)

func ColorPrint(a ...interface{}) {
	fmt.Fprint(color.Output, a...)
}

func ColorPrintln(a ...interface{})  {
	fmt.Fprintln(color.Output, a...)
}

func ColorPrintf(format string, a ...interface{})  {
	fmt.Fprintf(color.Output, format, a...)
}
