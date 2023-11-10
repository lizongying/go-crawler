package node

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/queue"
	"github.com/lizongying/go-crawler/pkg/utils"
	"sync/atomic"
	"time"
)

type Node struct {
	pkg.CrawlerStatus `json:"status,omitempty"`
	Id                string               `json:"id,omitempty"`
	Hostname          string               `json:"hostname,omitempty"`
	Ip                string               `json:"ip,omitempty"`
	Enable            bool                 `json:"enable,omitempty"`
	Spider            uint32               `json:"spider,omitempty"`
	Job               uint32               `json:"job,omitempty"`
	Task              uint32               `json:"task,omitempty"`
	Record            uint32               `json:"record,omitempty"`
	StartTime         utils.Timestamp      `json:"start_time"`
	FinishTime        utils.Timestamp      `json:"finish_time"`
	UpdateTime        utils.Timestamp      `json:"update_time,omitempty"`
	StatusList        *queue.PriorityQueue `json:"status_list,omitempty"`
}

func (n *Node) WithStatusAndTime(status pkg.CrawlerStatus, t time.Time) pkg.StatisticsNode {
	n.withStatus(status)
	n.withUpdateTime(t)
	switch status {
	case pkg.CrawlerStatusRunning:
		n.withStartTime(t)
	case pkg.CrawlerStatusStopped:
		n.withFinishTime(t)
	}

	if n.StatusList == nil {
		n.StatusList = queue.NewPriorityQueue(10)
	}
	n.StatusList.Push(queue.NewItem(status, t.UnixNano()))
	return n
}
func (n *Node) WithId(id string) pkg.StatisticsNode {
	n.Id = id
	return n
}
func (n *Node) withStatus(status pkg.CrawlerStatus) *Node {
	n.CrawlerStatus = status
	return n
}
func (n *Node) WithEnable(enable bool) *Node {
	n.Enable = enable
	return n
}
func (n *Node) IncSpider() {
	atomic.AddUint32(&n.Spider, 1)
}
func (n *Node) DecSpider() {
	atomic.AddUint32(&n.Spider, ^uint32(0))
}
func (n *Node) IncJob() {
	atomic.AddUint32(&n.Job, 1)
}
func (n *Node) DecJob() {
	atomic.AddUint32(&n.Job, ^uint32(0))
}
func (n *Node) IncTask() {
	atomic.AddUint32(&n.Task, 1)
}
func (n *Node) DecTask() {
	atomic.AddUint32(&n.Task, ^uint32(0))
}
func (n *Node) IncRecord() {
	atomic.AddUint32(&n.Record, 1)
}
func (n *Node) DecRecord() {
	atomic.AddUint32(&n.Record, ^uint32(0))
}
func (n *Node) withStartTime(t time.Time) *Node {
	n.StartTime = utils.Timestamp{
		Time: t,
	}
	return n
}
func (n *Node) withFinishTime(t time.Time) *Node {
	n.FinishTime = utils.Timestamp{
		Time: t,
	}
	return n
}
func (n *Node) withUpdateTime(t time.Time) *Node {
	n.UpdateTime = utils.Timestamp{
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
