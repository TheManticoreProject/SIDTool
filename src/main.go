package main

import (
	"fmt"
	"strings"

	"github.com/TheManticoreProject/SIDTool/sid"
	"github.com/TheManticoreProject/SIDTool/utils"

	"github.com/TheManticoreProject/goopts/parser"
)

var (
	mode     string
	value    string
	toHex    bool
	toBytes  bool
	toString bool
)

func parseArgs() {
	ap := parser.ArgumentsParser{Banner: "SIDTool - by Remi GASCOU (Podalirius) @ TheManticoreProject - v1.2"}

	ap.NewStringPositionalArgument(&mode, "mode", "Mode.")

	// "convert"
	ap.NewStringArgument(&value, "-v", "--value", "", true, "Input value to be converted.")
	groupOutputFormats, err := ap.NewRequiredMutuallyExclusiveArgumentGroup("Conversion Options")
	if err != nil {
		fmt.Printf("[error] Error creating ArgumentGroup: %s\n", err)
	} else {
		groupOutputFormats.NewBoolArgument(&toHex, "-x", "--to-hex", false, "Convert the input value to hexadecimal.")
		groupOutputFormats.NewBoolArgument(&toBytes, "-b", "--to-bytes", false, "Convert the input value to bytes.")
		groupOutputFormats.NewBoolArgument(&toString, "-s", "--to-string", false, "Convert the input value to string.")
	}

	// "describe"

	// "lookup"

	ap.Parse()
}

func main() {
	parseArgs()

	if strings.EqualFold(mode, "convert") {
		value := strings.Trim(value, " ")

		sid := sid.SID{}

		// If input value is a string starting with S-
		if strings.EqualFold(value[:2], "S-") {
			sid.FromString(value)
		} else if utils.IsHexString(value) {
			sid.FromBytes(utils.BytesFromHexString(value))
		}

		// Export format
		if toHex {
			fmt.Printf("%s\n", utils.BytesToHexString(sid.ToBytes()))
		} else if toBytes {
			fmt.Printf("%s\n", utils.BytesToBytesString(sid.ToBytes()))
		} else if toString {
			fmt.Printf("%s\n", sid.ToString())
		}
	} else if strings.EqualFold(mode, "describe") {

	} else if strings.EqualFold(mode, "lookup") {

	}
}
