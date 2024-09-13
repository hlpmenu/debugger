package debugger

import (
	"os"

	"gopkg.hlmpn.dev/pkg/funcutils"
	errutil "gopkg.hlmpn.dev/pkg/go-logger/errutil"
)

// Runs a a function if debug mode is enabled
// Pass in a function and any arguments
// Example:
//
//	func DoSomething(s string) {
//		fmt.Println(s)
//	}
//
// func main() {
// debug.RunFunc(DoSomething, "Hello, World!") // prints "Hello, World!"
// }
// For functions with no arguments, use RunSimple
// For functions with multiple arguments, use RunFuncAny
func (d *Debug) RunFunc(fn func(), args ...interface{}) {
	if d.IsDebug() {
		fn()
	}
}

// RunFuncAny takes any function as an argument and runs it
func (d *Debug) RunFuncAny(fn func(any ...interface{}), arg1 interface{}, args ...interface{}) {
	if d.IsDebug() {
		var f = &funcutils.Function{}
		f = funcutils.CreateFunc(fn, arg1, args)
		f.Run()
	}

}

// GoRun takes any function as an argument and runs it in a goroutine
// Example:
//
//	func DoSomething(s string) {
//		fmt.Println(s)
//	}
//
// func main() {
// debug.GoRun(DoSomething, "Hello, World!") // prints "Hello, World!"

func (d *Debug) GoRun(fn func(any ...interface{}), arg1 interface{}, args ...interface{}) func() {
	if d.IsDebug() {
		var f = &funcutils.Function{}
		f = funcutils.CreateFunc(fn, arg1, args)
		return func() {
			go func() {
				f.Run()
			}()
		}
	}
	return func() {}

}

func (d *Debug) RunFuncUtilFunc(fn funcutils.Function) {
	if d.IsDebug() {
		fn.Run()
	}
}

func (d *Debug) SetEnv() {
	os.Setenv("DEBUG", "true")
	os.Setenv("DEBUG_LEVEL", "debug")
	os.Setenv("debug", "true")
	os.Setenv("debug_level", "debug")
}

// RunFunc takes any function as an argument and runs it
func RunFuncAlways(fn func(a interface{}, b ...interface{}), i interface{}, args ...interface{}) {
	fn(i, args)
}

func Try(fn func(), catch func(interface{})) error {
	defer func() {
		if r := recover(); r != nil {
			catch(r)
			errutil.New(r.(string)).Unwrap()
		}
	}()
	fn()
	return nil
}
