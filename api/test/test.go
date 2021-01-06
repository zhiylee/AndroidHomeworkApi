package main

import (
	"fmt"
	"regexp"
)

func main()  {
	a:="http://android.refrainq.com/assets/avatar/169049883u8.jpg"

	reg:=regexp.MustCompile("http.//[^/]+/")
	rep:=reg.ReplaceAllLiteralString(a,"/")
	fmt.Println(rep)
}
