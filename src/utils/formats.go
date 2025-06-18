package utils

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
)

func IsHexString(s string) bool {
	match, _ := regexp.MatchString("^[0-9a-fA-F]+$", s)
	return match
}

func IsBase64String(s string) ([]byte, bool){
	decoded, err := base64.StdEncoding.DecodeString(s)
	return decoded, err == nil
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

func BytesToBase64String(data []byte) string {
	result := base64.StdEncoding.EncodeToString(data)
	return result
}

func HexFromBytesString(data []byte) string {
	var result string
	for _, b := range data {
		result += fmt.Sprintf("\\x%02x", b)
	}
	return result
}
