package prefix

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const split = " | "

func StartCommand(name string, args ...string) error {
	c := exec.Command(name, args...)
	c.Stdout = &Writer{Prefix: name + split, Writer: os.Stdout}
	c.Stderr = &Writer{Prefix: name + split, Writer: os.Stderr}
	if err := c.Run(); err != nil {
		return fmt.Errorf("%s%s%w", name, split, err)
	}

	return nil
}

type Writer struct {
	Prefix     string
	Writer     io.Writer
	skipPrefix bool
}

func (w *Writer) Write(p []byte) (n int, err error) {
	n = len(p)

	if !w.skipPrefix {
		// Line should be prefixed (it is the first line, or the last Write ended in a newline).
		p = append([]byte(w.Prefix), p...)
		w.skipPrefix = true
	}

	p = bytes.ReplaceAll(p, []byte("\n"), []byte("\n"+w.Prefix))
	if bytes.HasSuffix(p, []byte(w.Prefix)) {
		p = bytes.TrimSuffix(p, []byte(w.Prefix))
		w.skipPrefix = false
	}

	_, err = w.Writer.Write(p)
	return
}
