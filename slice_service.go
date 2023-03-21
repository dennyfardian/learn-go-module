package learngomodule

import "fmt"

// tulis fungsi
func PrintSliceOfString(slc []string) {
	for _, item := range slc {
		fmt.Println(item)
	}
}
