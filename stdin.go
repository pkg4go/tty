package tty

import "runtime"
import "strings"
import "os/exec"
import "fmt"
import "os"

// Example:
//
// func main() {
//   fmt.Printf("1: InterceptChar - %s\n", tty.InterceptChar())
//   fmt.Printf("2: interceptLine - %s\n", tty.InterceptLine())

//   fmt.Printf("3: ReadChar - %s\n", tty.ReadChar())
//   fmt.Printf("4: ReadLine - %s\n", tty.ReadLine())
// }

var platform string
var sttyFlag string

func init() {
	platform = runtime.GOOS

	switch platform {
	case "darwin":
		sttyFlag = "-f"
	case "linux":
		sttyFlag = "-F"
	}
}

// intercept a char from terminal, no need input `Enter`, do not display on the screen
func InterceptChar() string {
	// disable input buffer
	e := exec.Command("stty", sttyFlag, "/dev/tty", "cbreak", "min", "1").Run()
	exitIfError(e)
	// do not display input characters on the screen
	e = exec.Command("stty", sttyFlag, "/dev/tty", "-echo").Run()
	exitIfError(e)

	defer exec.Command("stty", sttyFlag, "/dev/tty", "echo").Run()

	b := make([]byte, 1)
	os.Stdin.Read(b)

	return string(b[:])
}

// intercept a line of characters, no need input `Enter`, do not display on the screen
func InterceptLine() string {
	// disable input buffer
	e := exec.Command("stty", sttyFlag, "/dev/tty", "cbreak", "min", "1").Run()
	exitIfError(e)
	// do not display input characters on the screen
	e = exec.Command("stty", sttyFlag, "/dev/tty", "-echo").Run()
	exitIfError(e)

	defer exec.Command("stty", sttyFlag, "/dev/tty", "echo").Run()

	buffer := new([]byte)

	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)

		if int(b[0]) == 10 {
			break
		}

		*buffer = append(*buffer, b[0])
	}

	return string((*buffer)[:])
}

// read a char from terminal, no need input `Enter`
func ReadChar() string {
	// disable input buffer
	e := exec.Command("stty", sttyFlag, "/dev/tty", "cbreak", "min", "1").Run()
	exitIfError(e)

	b := make([]byte, 1)
	os.Stdin.Read(b)

	return string(b[:])
}

// read a line of characters, no need input `Enter`
func ReadLine() string {
	var s string
	fmt.Scanf("%s", &s)
	return strings.TrimSpace(s)
}

func exitIfError(e error) {
	if e != nil {
		fmt.Printf("error: %s", e.Error())
		os.Exit(1)
	}
}
