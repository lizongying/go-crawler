package crawler

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/queue"
	"github.com/lizongying/go-crawler/pkg/utils"
	"sync/atomic"
	"time"
)

type Crawler struct {
	pkg.CrawlerStatus `json:"status,omitempty"`
	Id                string               `json:"id,omitempty"`
	Hostname          string               `json:"hostname,omitempty"`
	Ip                string               `json:"ip,omitempty"`
	Enable            bool                 `json:"enable,omitempty"`
	Spider            uint32               `json:"spider,omitempty"`
	Job               uint32               `json:"job,omitempty"`
	Task              uint32               `json:"task,omitempty"`
	Request           uint32               `json:"request,omitempty"`
	Item              uint32               `json:"item,omitempty"`
	StartTime         utils.Timestamp      `json:"start_time"`
	FinishTime        utils.Timestamp      `json:"finish_time"`
	UpdateTime        utils.Timestamp      `json:"update_time,omitempty"`
	StatusList        *queue.PriorityQueue `json:"status_list,omitempty"`
}

func (n *Crawler) WithStatusAndTime(status pkg.CrawlerStatus, t time.Time) pkg.StatisticsCrawler {
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
func (n *Crawler) WithId(id string) pkg.StatisticsCrawler {
	n.Id = id
	return n
}
func (n *Crawler) withStatus(status pkg.CrawlerStatus) *Crawler {
	n.CrawlerStatus = status
	return n
}
func (n *Crawler) WithEnable(enable bool) *Crawler {
	n.Enable = enable
	return n
}
func (n *Crawler) IncSpider() {
	atomic.AddUint32(&n.Spider, 1)
}
func (n *Crawler) DecSpider() {
	atomic.AddUint32(&n.Spider, ^uint32(0))
}
func (n *Crawler) IncJob() {
	atomic.AddUint32(&n.Job, 1)
}
func (n *Crawler) DecJob() {
	atomic.AddUint32(&n.Job, ^uint32(0))
}
func (n *Crawler) IncTask() {
	atomic.AddUint32(&n.Task, 1)
}
func (n *Crawler) DecTask() {
	atomic.AddUint32(&n.Task, ^uint32(0))
}
func (n *Crawler) IncRequest() {
	atomic.AddUint32(&n.Request, 1)
}
func (n *Crawler) DecRequest() {
	atomic.AddUint32(&n.Request, ^uint32(0))
}
func (n *Crawler) IncItem() {
	atomic.AddUint32(&n.Item, 1)
}
func (n *Crawler) DecItem() {
	atomic.AddUint32(&n.Item, ^uint32(0))
}
func (n *Crawler) withStartTime(t time.Time) *Crawler {
	n.StartTime = utils.Timestamp{
		Time: t,
	}
	return n
}
func (n *Crawler) withFinishTime(t time.Time) *Crawler {
	n.FinishTime = utils.Timestamp{
		Time: t,
	}
	return n
}
func (n *Crawler) withUpdateTime(t time.Time) *Crawler {
	n.UpdateTime = utils.Timestamp{
		Time: t,
	}
	return n
}
func (n *Crawler) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(n)
	if err != nil {
		return
	}
	return
}
