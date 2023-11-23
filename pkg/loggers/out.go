package loggers

import "io"

type Out struct {
	io.Writer
	Name string
}
