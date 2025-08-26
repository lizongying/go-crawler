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

var ErrQueueTimeout = errors.New("queue timeout")
var ErrTimeout = errors.New("timeout")

var ErrSpiderNotFound = errors.New("spider not found")

var ErrYieldRequestFailed = errors.New("yield request failed")
var ErrYieldExtraFailed = errors.New("yield extra failed")
var ErrYieldItemFailed = errors.New("yield item failed")
var ErrNilProxy = errors.New("proxy enabled but proxy is nil")
