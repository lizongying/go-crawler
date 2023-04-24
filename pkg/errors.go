package pkg

import (
	"errors"
)

var DontStopErr = errors.New("don't stop")
var BreakErr = errors.New("break")
var ErrIgnoreRequest = errors.New("IgnoreRequest")
var ErrDropItem = errors.New("DropItem")
