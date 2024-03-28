package main

import "strings"

func pascal2Snake(s string) string {
	var result string
	for i, r := range s {
		if i == 0 {
			result += strings.ToLower(string(r))
			continue
		}
		if i > 0 && 'A' <= r && r <= 'Z' {
			result += "_"
		}
		result += string(r)
	}
	return result
}
