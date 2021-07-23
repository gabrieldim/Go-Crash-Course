package main

import (
	"fmt"
	"math"
	"github.com\gabrieldim\golang_crash_course\03_packages\strutil"
)

func main() {

	fmt.Println(math.Floor(3.4))
	fmt.Println(math.Ceil(5.7))
	fmt.Println(math.Sqrt(64))
	fmt.Println(strutil.Reverse("Hello backwards."))
}
