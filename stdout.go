package tty

import "fmt"
import "io"
import "os"

func write(s string) {
	io.WriteString(os.Stdout, s)
}

func Clear() {
	write("\033[2J")
}

func ClearLine() {
	write("\033[2K")
}

func Move(x, y int) {
	write(fmt.Sprintf("\033[%d;%dH", x, y))
}

func Up(x int) {
	write(fmt.Sprintf("\033[%dA", x))
}

func Down(x int) {
	write(fmt.Sprintf("\033[%dB", x))
}

func Right(x int) {
	write(fmt.Sprintf("\033[%dC", x))
}

func Left(x int) {
	write(fmt.Sprintf("\033[%dD", x))
}
