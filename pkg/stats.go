package pkg

type Stats interface {
	RequestTotal() uint32
	IncRequestTotal() uint32
	RequestSuccess() uint32
	IncRequestSuccess() uint32
	RequestIgnore() uint32
	IncRequestIgnore() uint32
	RequestError() uint32
	IncRequestError() uint32
	ItemTotal() uint32
	IncItemTotal() uint32
	ItemSuccess() uint32
	IncItemSuccess() uint32
	ItemIgnore() uint32
	IncItemIgnore() uint32
	ItemError() uint32
	IncItemError() uint32
	StatusOk() uint32
	IncStatusOk() uint32
	StatusErr() uint32
	IncStatusErr() uint32
	GetMap() map[string]uint32
}

type StatsWithImage interface {
	Stats
	ImageTotal() uint32
	IncImageTotal() uint32
}
