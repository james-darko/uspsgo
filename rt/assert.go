package rt

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func Assert(b bool) {
	if !b {
		log.Println("assertion failed")
		stack := debug.Stack()
		fmt.Println(string(stack))
		// if os.Getenv("ENV") == "prd" {
		// 	gotify.SendAndPray("assertion failed", string(stack), 1)
		// }
		os.Exit(1)
	}
}

func Assertf(b bool, format string, args ...any) {
	if !b {
		m := fmt.Sprintf(format, args...)
		log.Println(m)
		stack := debug.Stack()
		fmt.Println(string(stack))
		// if os.Getenv("ENV") == "prd" {
		// 	gotify.SendAndPray("assertion failed", m+"\n"+string(stack), 1)
		// }
		os.Exit(1)
	}
}

func AssertE(err error) {
	AssertNoErr(err)
}

func AssertNoErr(err error) {
	if err != nil {
		log.Printf("got asserted error: %v", err)
		stack := debug.Stack()
		fmt.Println(string(stack))
		// if os.Getenv("ENV") == "prd" {
		// 	gotify.SendAndPray("assertion failed", fmt.Sprintf("%v\n%s", err, string(stack)), 1)
		// }
		os.Exit(1)
	}
}

func AssertNoErrf(err error, format string, args ...any) {
	if err != nil {
		errStr := fmt.Sprintf(format+"\n", args...)
		log.Print(errStr)
		stack := debug.Stack()
		fmt.Println(string(stack))
		// if os.Getenv("ENV") == "prd" {
		// 	gotify.SendAndPray("assertion failed", fmt.Sprintf("%s%v\n%s", errStr, err, string(stack)), 1)
		// }
		os.Exit(1)
	}
}

func AssertErrf(err error, format string, args ...any) {
	if err != nil {
		errStr := fmt.Sprintf(format+"\n", args...)
		log.Print(errStr)
		stack := debug.Stack()
		fmt.Println(string(stack))
		// if os.Getenv("ENV") == "prd" {
		// 	gotify.SendAndPray("assertion failed", fmt.Sprintf("%s%v\n%s", errStr, err, string(stack)), 1)
		// }
		os.Exit(1)
	}
}
