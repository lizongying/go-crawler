package stats

import "sync/atomic"

type Stats struct {
	RequestTotal   uint32
	RequestSuccess uint32
	RequestIgnore  uint32
	RequestError   uint32
	ItemTotal      uint32
	ItemSuccess    uint32
	ItemIgnore     uint32
	ItemError      uint32
	StatusOk       uint32
	StatusErr      uint32
}

func (s *Stats) GetRequestTotal() uint32 {
	return atomic.LoadUint32(&s.RequestTotal)
}
func (s *Stats) IncRequestTotal() uint32 {
	return atomic.AddUint32(&s.RequestTotal, 1)
}
func (s *Stats) GetRequestSuccess() uint32 {
	return atomic.LoadUint32(&s.RequestSuccess)
}
func (s *Stats) IncRequestSuccess() uint32 {
	return atomic.AddUint32(&s.RequestSuccess, 1)
}
func (s *Stats) GetRequestIgnore() uint32 {
	return atomic.LoadUint32(&s.RequestIgnore)
}
func (s *Stats) IncRequestIgnore() uint32 {
	return atomic.AddUint32(&s.RequestIgnore, 1)
}
func (s *Stats) GetRequestError() uint32 {
	return atomic.LoadUint32(&s.RequestError)
}
func (s *Stats) IncRequestError() uint32 {
	return atomic.AddUint32(&s.RequestError, 1)
}
func (s *Stats) GetItemTotal() uint32 {
	return atomic.LoadUint32(&s.ItemTotal)
}
func (s *Stats) IncItemTotal() uint32 {
	return atomic.AddUint32(&s.ItemTotal, 1)
}
func (s *Stats) GetItemSuccess() uint32 {
	return atomic.LoadUint32(&s.ItemSuccess)
}
func (s *Stats) IncItemSuccess() uint32 {
	return atomic.AddUint32(&s.ItemSuccess, 1)
}
func (s *Stats) GetItemIgnore() uint32 {
	return atomic.LoadUint32(&s.ItemIgnore)
}
func (s *Stats) IncItemIgnore() uint32 {
	return atomic.AddUint32(&s.ItemIgnore, 1)
}
func (s *Stats) GetItemError() uint32 {
	return atomic.LoadUint32(&s.ItemError)
}
func (s *Stats) IncItemError() uint32 {
	return atomic.AddUint32(&s.ItemError, 1)
}
func (s *Stats) GetStatusOk() uint32 {
	return atomic.LoadUint32(&s.StatusOk)
}
func (s *Stats) IncStatusOk() uint32 {
	return atomic.AddUint32(&s.StatusOk, 1)
}
func (s *Stats) GetStatusErr() uint32 {
	return atomic.LoadUint32(&s.StatusErr)
}
func (s *Stats) IncStatusErr() uint32 {
	return atomic.AddUint32(&s.StatusOk, 1)
}

type ImageStats struct {
	Stats
	ImageTotal uint32
}

func (s *ImageStats) GetImageTotal() uint32 {
	return atomic.LoadUint32(&s.ImageTotal)
}
func (s *ImageStats) IncImageTotal() uint32 {
	return atomic.AddUint32(&s.ImageTotal, 1)
}
