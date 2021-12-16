package tests

import (
	"fmt"
	"main/utils/colors"
)

type Any interface {}

func PrintSuccess(message Any) {
	fmt.Println(colors.GREEN, "[SUCCESS]", message)
}

func PrintFail(message Any) {
	fmt.Println(colors.RED, "[FAIL]", message)
}
