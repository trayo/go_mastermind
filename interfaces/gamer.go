package interfaces

import (
	"bufio"

	"github.com/trayo/go_mastermind/print"
)

type Gamer interface {
	Play(stdin *bufio.Reader, printer print.Printer)
}
