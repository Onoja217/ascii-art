package main

import (
	"fmt"
	"os"
	"strings"
)
func getBannerFile(args []string) string {
	if len(args) == 3 {
		switch args[2] {
		case "standard", "shadow", "thinkertoy":
			return args[2] + ".txt"
			fmt.Println("Invalid banner name")
			return ""
		}
	}
	return "standard.txt"

}

func laodBanner(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file:", err)
		os.Exit(1)
	}
	return strings.Split(string(data), "\n")
}