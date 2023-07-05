package pkg

import (
	"errors"
)

var DontStopErr = errors.New("don't stop")
var BreakErr = errors.New("break")
var ErrIgnoreRequest = errors.New("IgnoreRequest")
var ErrNotAllowRequest = errors.New("NotAllowRequest")
var ErrIgnoreResponse = errors.New("IgnoreResponse")
var ErrNeedRetry = errors.New("need retry")
var ErrUrlLengthLimit = errors.New("UrlLengthLimit")
var ErrDropItem = errors.New("DropItem")
