package utils

import (
	"bufio"
	"fmt"
	"os"
)

var MyStdin = os.Stdin

// GO自带的fmt.Scanln将空格也当作结束符，若需要读取含有空格的句子请使用该方法
func Scanln(a *string) {
	reader := bufio.NewReader(MyStdin)
	data, _, err := reader.ReadLine()
	if err != nil {
		fmt.Print(err)
	}
	*a = string(data)
}

func ScanLine(a *string) {
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}

	*a = string(b)
}
