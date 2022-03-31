package main

import (
	"Serve/tools/color_print"
)

func main() {
	color_print.ColorPrintln("This is test", color_print.Success("SUCCESS"))
	color_print.ColorPrint("This is test ", color_print.Warn("WARN\n"))
	color_print.ColorPrintf("This is test: %s\n", color_print.Fail("FAIL"))
}