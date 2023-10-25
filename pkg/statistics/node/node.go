package node

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg/utils"
	"sync/atomic"
	"time"
)

type Status uint8

const (
	StatusUnknown = iota
	StatusOnline
	StatusOffline
)

func (s *Status) String() string {
	switch *s {
	case 1:
		return "online"
	case 2:
		return "offline"
	default:
		return "unknown"
	}
}

type Node struct {
	Status     `json:"status,omitempty"`
	Hostname   string          `json:"hostname,omitempty"`
	Ip         string          `json:"ip,omitempty"`
	Enable     bool            `json:"enable,omitempty"`
	Spider     uint32          `json:"spider,omitempty"`
	Schedule   uint32          `json:"schedule,omitempty"`
	Task       uint32          `json:"task,omitempty"`
	StartTime  utils.Timestamp `json:"start_time"`
	FinishTime utils.Timestamp `json:"finish_time"`
}

func (n *Node) WithStatus(status Status) *Node {
	n.Status = status
	return n
}
func (n *Node) WithEnable(enable bool) *Node {
	n.Enable = enable
	return n
}
func (n *Node) IncSpider() {
	atomic.AddUint32(&n.Spider, 1)
}
func (n *Node) IncTask() {
	atomic.AddUint32(&n.Task, 1)
}
func (n *Node) WithStartTime(t time.Time) *Node {
	n.StartTime = utils.Timestamp{
		Time: t,
	}
	return n
}
func (n *Node) WithFinishTime(t time.Time) *Node {
	n.FinishTime = utils.Timestamp{
		Time: t,
	}
	return n
}
func (n *Node) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(n)
	if err != nil {
		return
	}
	return
}
