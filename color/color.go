package color

import (
	"fmt"
	"io"
)

// type  int

const Color (
    black
)

func Reset(w io.Writer) {
    fmt.Fprint(w, "\x1b[0m")
}

func Set(w io.Writer, color) {

}
