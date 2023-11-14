package pkg

type Pipeline interface {
	Spider() Spider
	ProcessItem(Item) error
	Name() string
	SetName(string)
	Order() uint8
	SetOrder(uint8)
	FromSpider(Spider) error
}

type UnimplementedPipeline struct {
	name   string
	order  uint8
	spider Spider
}

func (p *UnimplementedPipeline) Spider() Spider {
	return p.spider
}
func (*UnimplementedPipeline) ProcessItem(Item) error {
	return nil
}
func (p *UnimplementedPipeline) Name() string {
	return p.name
}
func (p *UnimplementedPipeline) SetName(name string) {
	p.name = name
}
func (p *UnimplementedPipeline) Order() uint8 {
	return p.order
}
func (p *UnimplementedPipeline) SetOrder(order uint8) {
	p.order = order
}
func (p *UnimplementedPipeline) FromSpider(spider Spider) (err error) {
	if p == nil {
		return new(UnimplementedPipeline).FromSpider(spider)
	}

	p.spider = spider
	return
}
