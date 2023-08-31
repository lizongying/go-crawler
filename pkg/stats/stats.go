package stats

import "sync/atomic"

type Stats struct {
	requestTotal   uint32
	requestSuccess uint32
	requestIgnore  uint32
	requestError   uint32
	itemTotal      uint32
	itemSuccess    uint32
	itemIgnore     uint32
	itemError      uint32
	statusOk       uint32
	statusErr      uint32
}

func (s *Stats) RequestTotal() uint32 {
	return atomic.LoadUint32(&s.requestTotal)
}
func (s *Stats) IncRequestTotal() uint32 {
	return atomic.AddUint32(&s.requestTotal, 1)
}
func (s *Stats) RequestSuccess() uint32 {
	return atomic.LoadUint32(&s.requestSuccess)
}
func (s *Stats) IncRequestSuccess() uint32 {
	s.IncRequestTotal()
	return atomic.AddUint32(&s.requestSuccess, 1)
}
func (s *Stats) RequestIgnore() uint32 {
	return atomic.LoadUint32(&s.requestIgnore)
}
func (s *Stats) IncRequestIgnore() uint32 {
	s.IncRequestTotal()
	return atomic.AddUint32(&s.requestIgnore, 1)
}
func (s *Stats) RequestError() uint32 {
	return atomic.LoadUint32(&s.requestError)
}
func (s *Stats) IncRequestError() uint32 {
	s.IncRequestTotal()
	return atomic.AddUint32(&s.requestError, 1)
}
func (s *Stats) ItemTotal() uint32 {
	return atomic.LoadUint32(&s.itemTotal)
}
func (s *Stats) IncItemTotal() uint32 {
	return atomic.AddUint32(&s.itemTotal, 1)
}
func (s *Stats) ItemSuccess() uint32 {
	return atomic.LoadUint32(&s.itemSuccess)
}
func (s *Stats) IncItemSuccess() uint32 {
	s.IncItemTotal()
	return atomic.AddUint32(&s.itemSuccess, 1)
}
func (s *Stats) ItemIgnore() uint32 {
	return atomic.LoadUint32(&s.itemIgnore)
}
func (s *Stats) IncItemIgnore() uint32 {
	s.IncItemTotal()
	return atomic.AddUint32(&s.itemIgnore, 1)
}
func (s *Stats) ItemError() uint32 {
	return atomic.LoadUint32(&s.itemError)
}
func (s *Stats) IncItemError() uint32 {
	s.IncItemTotal()
	return atomic.AddUint32(&s.itemError, 1)
}
func (s *Stats) StatusOk() uint32 {
	return atomic.LoadUint32(&s.statusOk)
}
func (s *Stats) IncStatusOk() uint32 {
	return atomic.AddUint32(&s.statusOk, 1)
}
func (s *Stats) StatusErr() uint32 {
	return atomic.LoadUint32(&s.statusErr)
}
func (s *Stats) IncStatusErr() uint32 {
	return atomic.AddUint32(&s.statusOk, 1)
}
func (s *Stats) GetMap() map[string]uint32 {
	return map[string]uint32{
		"requestTotal":   s.RequestTotal(),
		"requestSuccess": s.RequestSuccess(),
		"requestIgnore":  s.RequestIgnore(),
		"requestError":   s.RequestError(),
		"itemTotal":      s.ItemTotal(),
		"itemSuccess":    s.ItemSuccess(),
		"itemIgnore":     s.ItemIgnore(),
		"itemError":      s.ItemError(),
		"statusOk":       s.StatusOk(),
		"statusErr":      s.StatusErr(),
	}
}

type ImageStats struct {
	Stats
	imageTotal uint32
}

func (s *ImageStats) ImageTotal() uint32 {
	return atomic.LoadUint32(&s.imageTotal)
}
func (s *ImageStats) IncImageTotal() uint32 {
	return atomic.AddUint32(&s.imageTotal, 1)
}
func (s *ImageStats) GetMap() map[string]uint32 {
	return map[string]uint32{
		"requestTotal":   s.RequestTotal(),
		"requestSuccess": s.RequestSuccess(),
		"requestIgnore":  s.RequestIgnore(),
		"requestError":   s.RequestError(),
		"itemTotal":      s.ItemTotal(),
		"itemSuccess":    s.ItemSuccess(),
		"itemIgnore":     s.ItemIgnore(),
		"itemError":      s.ItemError(),
		"statusOk":       s.StatusOk(),
		"statusErr":      s.StatusErr(),
		"imageTotal":     s.ImageTotal(),
	}
}

type FileStats struct {
	Stats
	fileTotal uint32
}

func (s *FileStats) FileTotal() uint32 {
	return atomic.LoadUint32(&s.fileTotal)
}
func (s *FileStats) IncFileTotal() uint32 {
	return atomic.AddUint32(&s.fileTotal, 1)
}
func (s *FileStats) GetMap() map[string]uint32 {
	return map[string]uint32{
		"requestTotal":   s.RequestTotal(),
		"requestSuccess": s.RequestSuccess(),
		"requestIgnore":  s.RequestIgnore(),
		"requestError":   s.RequestError(),
		"itemTotal":      s.ItemTotal(),
		"itemSuccess":    s.ItemSuccess(),
		"itemIgnore":     s.ItemIgnore(),
		"itemError":      s.ItemError(),
		"statusOk":       s.StatusOk(),
		"statusErr":      s.StatusErr(),
		"fileTotal":      s.FileTotal(),
	}
}

type MediaStats struct {
	Stats
	imageTotal uint32
	fileTotal  uint32
}

func (s *MediaStats) ImageTotal() uint32 {
	return atomic.LoadUint32(&s.imageTotal)
}
func (s *MediaStats) IncImageTotal() uint32 {
	return atomic.AddUint32(&s.imageTotal, 1)
}
func (s *MediaStats) FileTotal() uint32 {
	return atomic.LoadUint32(&s.fileTotal)
}
func (s *MediaStats) IncFileTotal() uint32 {
	return atomic.AddUint32(&s.fileTotal, 1)
}
func (s *MediaStats) GetMap() map[string]uint32 {
	return map[string]uint32{
		"requestTotal":   s.RequestTotal(),
		"requestSuccess": s.RequestSuccess(),
		"requestIgnore":  s.RequestIgnore(),
		"requestError":   s.RequestError(),
		"itemTotal":      s.ItemTotal(),
		"itemSuccess":    s.ItemSuccess(),
		"itemIgnore":     s.ItemIgnore(),
		"itemError":      s.ItemError(),
		"statusOk":       s.StatusOk(),
		"statusErr":      s.StatusErr(),
		"imageTotal":     s.ImageTotal(),
		"fileTotal":      s.FileTotal(),
	}
}
