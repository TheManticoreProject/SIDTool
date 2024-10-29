package utils

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func IsHexString(s string) bool {
	match, _ := regexp.MatchString("^[0-9a-fA-F]+$", s)
	return match
}

func BytesToBytesString(data []byte) string {
	result := ""
	for _, b := range data {
		result += fmt.Sprintf("\\x%02x", b)
	}
	return result
}

func BytesFromBytesString(hexStr string) ([]byte, error) {
	// Remove all occurrences of "\x" from the string
	cleaned := strings.ReplaceAll(hexStr, "\\x", "")

	// Decode the hexadecimal string into bytes
	bytes, err := hex.DecodeString(cleaned)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func BytesFromHexString(hexStr string) []byte {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil
	}
	return bytes
}

func BytesToHexString(data []byte) string {
	result := ""
	for _, b := range data {
		result += fmt.Sprintf("%02x", b)
	}
	return result
}

func HexFromBytesString(data []byte) string {
	var result string
	for _, b := range data {
		result += fmt.Sprintf("\\x%02x", b)
	}
	return result
}
