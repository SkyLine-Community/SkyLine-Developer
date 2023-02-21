package SkyLine_Backend

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func User_Input(args ...Object) Object {
	// this function takes a few positional arguments
	// POS-1 -> The prompt for the user to enter data
	// POS-2 -> The expected return type float, string, int
	// POS-3 -> When to send the input or to cancle the input tag, this can be n for new line or t for tab etc
	if len(args) != 3 {
		return NewError("Function `input(...)` requires 3 positional arguments of type string. \n Arg1=prompt \n Arg2=Expected return type\nArg3=Cancle when... ")
	}
	var prompt, expected_type, drop_when Object
	prompt = args[0].(*String)
	expected_type = args[1].(*String)
	drop_when = args[2].(*String)
	retret := bufio.NewReader(os.Stdin)
	var out string
	fmt.Print(prompt.Inspect())
	et := expected_type.Inspect()
	for {
		switch drop_when.Inspect() {
		case "n":
			out, _ = retret.ReadString('\n')
			out = strings.Replace(out, "\n", "", -1)
		case "t":
			out, _ = retret.ReadString('\t')
			out = strings.Replace(out, "\t", "", -1)
		case "r":
			out, _ = retret.ReadString('\r')
			out = strings.Replace(out, "\r", "", -1)
		default:
			return NewError("Unsupported argument in placement `3` final argument in call to input(...) -> supported=(n,t,r)")
		}
		if out != "" {
			switch et {
			case "integer":
				c, x := strconv.ParseInt(out, 0, 64)
				if x != nil {
					return &Error{Message: "Could not return this value, it was not able to be parsed as a integer which means it was either a float or character but this function does not support that as input"}
				}
				return &Integer{Value: c}
			case "float":
				c, x := strconv.ParseFloat(out, 64)
				if x != nil {
					return &Error{Message: "Could not return this value, it was not able to parse as a float value which means it was either a character or an integer, this function input does not accept anything but a float value"}
				}
				return &Float{Value: c}
			case "string":
				return &String{Value: out}
			}
		}
	}
}

func IO_Clear() Object {
	// takes no positional arguments
	WIN := "\x1b[2J"
	LIN := "\x1b[H\x1b[2J\x1b[3J"
	if U.OperatingSystem == "windows" {
		fmt.Println(WIN)
	} else {
		fmt.Println(LIN)
	}
	return &String{Value: ""}
}

type Box struct {
	TL string
	TR string
	BL string
	BR string
	HZ string
	VT string
}

func IO_Box(args ...Object) Object {
	var BL Box
	var text string
	// Optionally takes 7 positonal arguments
	if len(args) == 7 {
		text = args[0].Inspect()  // Text for the box
		BL.TL = args[1].Inspect() // Top left
		BL.TR = args[2].Inspect() // Top right
		BL.BL = args[3].Inspect() // Bottom left
		BL.BR = args[4].Inspect() // Bottom right
		BL.HZ = args[5].Inspect() // Horizontal
		BL.VT = args[6].Inspect() // Verticle
	} else {
		if len(args) >= 1 {
			text = args[0].Inspect()
		} else {
			return NewError("Sorry this function takes 1 required positional argument and 6 other optional arguments")
		}
		BL = Box{
			TL: "┏",
			TR: "┓",
			BL: "┗",
			BR: "┛",
			HZ: "━",
			VT: "┃",
		}
	}
	l := strings.Split(text, "\n")
	var mlen int
	for _, lin := range l {
		if len(lin) > mlen {
			mlen = len(lin)
		}
	}
	var b string
	b += BL.TL + strings.Repeat(BL.HZ, mlen) + BL.TR + "\n"
	for _, line := range l {
		b += BL.VT + line + strings.Repeat(" ", mlen-len(line)) + BL.VT + "\n"
	}
	b += BL.BL + strings.Repeat(BL.HZ, mlen) + BL.BR + "\n"
	return &String{Value: b}
}

func IO_Listen(args ...Object) Object {
	// This function is a bit more complicated than the other IO functions
	// This will start a thread and simply return nothing but rather listen
	// for key based input such as CTRL+C
	if len(args) != 2 {
		return NewError("Sorry this function of call io does not support any other functions other than 2. ")
	}
	var sigtype string
	switch arg := args[0].(type) {
	case *String:
		sigtype = arg.Inspect()
	default:
		return NewError("Sorry first argument in call to io listen is not a string, this argument MUST be a string")
	}
	var sig os.Signal
	switch strings.ToLower(sigtype) {
	case "terminate":
		sig = syscall.SIGTERM
	case "kill":
		sig = syscall.SIGKILL
	case "hangup":
		sig = syscall.SIGHUP
	case "ctrl-c":
		sig = os.Interrupt
	case "user1":
		sig = syscall.SIGUSR1
	case "user2":
		sig = syscall.SIGUSR2
	default:
		return NewError("Sorry the first agrument in the list does not exist")
	}
	msg := args[1].Inspect()
	c := make(chan os.Signal)
	go func() {
		HandleListener(c, sig, msg, ExitGracefully)
	}()
	return &Nil{}
}

func ExitGracefully(msg string) {
	println(msg)
	os.Exit(0)
}

func HandleListener(c chan os.Signal, signalCHAN os.Signal, message string, run func(string)) {
	signal.Notify(c, signalCHAN)
	for s := <-c; ; s = <-c {
		switch {
		case signalCHAN == syscall.SIGUSR1 && s == syscall.SIGUSR1:
			run(message)
		case signalCHAN == os.Interrupt && s == os.Interrupt:
			run(message)
		case signalCHAN == syscall.SIGUSR2 && s == syscall.SIGUSR2:
			run(message)
		case signalCHAN == syscall.SIGHUP && s == syscall.SIGHUP:
			run(message)
		case signalCHAN == syscall.SIGTERM && s == syscall.SIGTERM:
			run(message)
		case signalCHAN == syscall.SIGKILL && s == syscall.SIGKILL:
			run(message)
		}
	}
}

// now create a function_listener | this is a concept not full idea but the idea was
// that we can listen for a key signal or os.signal that will allow us to execute a function
