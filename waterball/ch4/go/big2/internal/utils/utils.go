// Package utils 提供與領域無關的輔助函式（輸入解析、scanner 封裝等）。
package utils

import (
	"bufio"
	"strconv"
	"strings"
)

// ParseIndices 解析一行空白分隔的非負整數，例如 "0 1 2" -> []int{0,1,2}。
func ParseIndices(line string) ([]int, error) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return nil, strconv.ErrSyntax
	}
	out := make([]int, 0, len(parts))
	for _, s := range parts {
		i, err := strconv.Atoi(s)
		if err != nil || i < 0 {
			return nil, err
		}
		out = append(out, i)
	}
	return out, nil
}

// ReadLineFromScanner 回傳一個函式，每次呼叫從 scanner 讀取下一行。供需要「讀一行」的呼叫端使用。
func ReadLineFromScanner(scanner *bufio.Scanner) func() string {
	return func() string {
		if scanner.Scan() {
			return scanner.Text()
		}
		return ""
	}
}
