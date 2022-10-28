//Declare a main package. In Go, code executed as an application must be in a main package.
package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

//go应用程序入口，在程序开始执行并完成初始化后，第一个调用（程序的入口点）的函数是 main.main()（如：C 语言），该函数一旦返回就表示程序已成功执行并立即退出。
//该方法没有参数也没有返回值
//左大括号 { 必须与方法的声明放在同一行，这是编译器的强制规定，否则你在使用 gofmt 时就会出现错误提示：`build-error: syntax error: unexpected semicolon or newline before {`
//Go 语言虽然看起来不使用分号作为语句的结束，但实际上这一过程是由编译器自动完成，因此才会引发像上面这样的错误
func test() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Get a greeting message and print it.
	message, err := greetings.Hellos(names)

	if err != nil {
		// If an error was returned, print it to the console and
		// exit the program.
		log.Fatal(err)
	}

	fmt.Println(message)
}
