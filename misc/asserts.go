package misc

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Ensure err is nil othervise exit
func Ensure(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}

// Assert fails the test if the condition is false
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// Ok fails the test if an err is not nil
func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// Equals fails the test if exp is not equal to act
func Equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\nexp: %#v\ngot: %#v\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

// NotEquals fails the test if exp is equal to act
func NotEquals(tb testing.TB, exp, act interface{}) {
	if reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\nnot exp: %#v\ngot: %#v\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
