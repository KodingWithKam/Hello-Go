/*
In Go, the main package is a special executable package that has to comply with two requirements:
1. the package name must be main,
2. and it must contain a function called main(), which will be the entry point for our program.
*/
package main // We have declared the package main, which means this will be an executable package.

import (
	"HelloGo/hello"
	"fmt"
	"github.com/fatih/color"
) // we have imported the fmt and the HelloGo/hello package that we created.

/*
we have imported the fmt and the example/hello package that we previously created.
*/
func main() {
	fmt.Println(hello.Welcome)

	fmt.Println(hello.HelloWorld())

	color.Blue("I'm blue daba dee daba di!")

}
