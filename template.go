package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
)

var t1 = template.Must(template.New("t1").Parse(`"{{23 }} < {{ 45}}"`))
var t2 = template.Must(template.New("t2").Parse(`"{{23 -}} < {{- 45}}"`))

func T1(wr io.Writer) {
	t1.Execute(wr, nil)
}

func T2(wr io.Writer) {
	t2.Execute(wr, nil)
}

func main() {
	T1(os.Stdout)
	fmt.Println()
	T2(os.Stdout)
	fmt.Println()
}
