package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；
//标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）
/**
	因此符合规范的函数一般写成如下的形式：
	func functionName(parameter_list) (return_value_list) {
   		…
	}
	parameter_list 的形式为 (param1 type1, param2 type2, …)
	return_value_list 的形式为 (ret1 type1, ret2 type2, …)

*/
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}

	return messages, nil
}

// init sets initial values for variables used in the function.
// Add an init function to seed the rand package with the current time.
// Go executes init functions automatically at program startup, after global variables have been initialized.
// For more about init functions, see Effective Go.
func init() {
	rand.Seed(time.Now().UnixNano())
}

//lowcase better小写字母命名的函数,只会暴露在本包内(local package)
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!"}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
