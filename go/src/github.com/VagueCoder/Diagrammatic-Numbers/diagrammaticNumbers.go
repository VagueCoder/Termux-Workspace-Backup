package main

import (
	"fmt"
	"os"
	"strings"
)

var numbers map[string][7]string = map[string][7]string{
	"0": [7]string{
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},

	"1": [7]string{
		"  1  ",
		" 11  ",
		"1 1  ",
		"  1  ",
		"  1  ",
		"  1  ",
		"11111"},

	"2": [7]string{
		" 222 ",
		"2   2",
		"    2",
		"   2 ",
		"  2  ",
		" 2   ",
		"22222"},

	"3": [7]string{
		" 333 ",
		"3   3",
		"    3",
		"  3  ",
		"    3",
		"3   3",
		" 333 "},

	"4": [7]string{
		"    4 ",
		"   44 ",
		"  4 4 ",
		" 4  4 ",
		"444444",
		"    4 ",
		"    4 "},

	"5": [7]string{
		"55555",
		"5    ",
		"5    ",
		"5555 ",
		"    5",
		"    5",
		"5555 "},

	"6": [7]string{
		"  6666",
		" 6    ",
		"6     ",
		"66666 ",
		"6    6",
		"6    6",
		" 6666 "},

	"7": [7]string{
		"7777777",
		"      7",
		"     7 ",
		" 77777 ",
		"   7   ",
		"  7    ",
		" 7     "},

	"8": [7]string{
		" 888 ",
		"8   8",
		"8   8",
		"  8  ",
		"8   8",
		"8   8",
		" 888 "},

	"9": [7]string{
		" 9999 ",
		"9    9",
		"9    9",
		" 99999",
		"     9",
		"    9 ",
		"9999  "},
}

func concat(num string) [7]string {
	var result [7]string
	for _, i := range strings.Split(num, "") {
		for j := 0; j < 7; j++ {
			result[j] += numbers[i][j] + "   "
		}
	}
	return result
}

func main() {
	for _, v := range concat(os.Args[1]) {
		fmt.Println(v)
	}
}
