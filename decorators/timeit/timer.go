package timer

import (
	"log"
	"time"

	funcs "github.com/SpaceDiverr/util"
)

/*
Decorator.
New logs the execution time of a given void no-argument function.

Difference from NewWithFuncNameInLog: no show of function name in log (makes sense, huh? XD).

New and NewWithFuncNameInLog share applicability options. Check out for tips of usage in 'Note' section of NewWithFuncNameInLog description.
*/
func New(f func()) func() {
	return func() {
		start := time.Now()
		f()
		elapsed := time.Since(start)

		log.Printf("-- func executed in %s --\n", elapsed)
	}
}

/*
Decorator.
NewWithFuncNameInLog logs the execution time of a given void no-argument function with its name.

isShowFuncName: A boolean flag to indicate
whether to show the name of the function in the log message.

methodOrFuncName: method or function itself (f) to be logged if isShowFuncName == true.
If isShowFuncName == false, methodOrFuncName must be nil (makes sense, huh? XD),
otherwise expect undefined behavior.

f: The function to be executed and timed.

Return: A function that wraps the original function and logs its execution time.

Note: if needed to measure exec time of either any struct method or func that takes arguments or/and returns something, consider the following:

What you could come up with straightforwardly:

	func main() {
		NewWithFuncNameInLog(true, isPalindrome, isPalindrome("Hello, world!"))()
	}

It won't work.

Solution: for 'f' use anonymous func (wrapper) that inside does logic with 'methodOrFunc'.

	func reverse(s string) string {
		runes := []rune(s)
		slices.Reverse(runes)
		return string(runes)
	}

	func isPalindrome(s string) bool {
		return s == reverse(s)
	}

	func main() {
		var res1 bool
		var res2 bool
		var s string = "I love Go!"

		timer1 := NewWithFuncNameInLog(true, isPalindrome, func() {
			res1 = isPalindrome(s)
		})
		timer2 := NewWithFuncNameInLog(false, nil, func() {
			res2 = isPalindrome(s)
		})

		timer1()
		timer2()

		fmt.Printf("isPalindrome(%s) == %v\n", s, res1 == res2)
	}

logs:

YYYY/MM/DD HH:MM:SS -- func executed in 583ns --

YYYY/MM/DD HH:MM:SS -- func isPalindrome executed in 458ns --

isPalindrome(I love Go!) == true

If your 'f' is void & no-argument, consider the following:

	func someFunc() {
	    n := 1
	    for n < 1024 {
	        n <<= 1
	    }
	    fmt.Printf("n == %d\n", n)
	}

	func main() {
		NewWithFuncNameInLog(false, nil, someFunc)()
		NewWithFuncNameInLog(true, someFunc, someFunc)()
	}

logs:

n == 1024

YYYY/MM/DD HH:MM:SS -- func executed in 19.722µs --

n == 1024

YYYY/MM/DD HH:MM:SS -- func someFunc executed in 29.625µs --
*/
func NewWithFuncNameInLog(isShowFuncName bool, methodOrFunc interface{}, f func()) func() {
	return func() {
		start := time.Now()
		f()
		elapsed := time.Since(start)

		var fName string
		if isShowFuncName && methodOrFunc != nil {
			fName = funcs.Name(methodOrFunc) + " "
		}

		log.Printf("-- func %sexecuted in %s --\n", fName, elapsed)
	}
}
