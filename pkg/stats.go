package pkg

type Stats interface {
	GetRequestTotal() uint32
	IncRequestTotal() uint32
	GetRequestSuccess() uint32
	IncRequestSuccess() uint32
	GetRequestIgnore() uint32
	IncRequestIgnore() uint32
	GetRequestError() uint32
	IncRequestError() uint32
	GetItemTotal() uint32
	IncItemTotal() uint32
	GetItemSuccess() uint32
	IncItemSuccess() uint32
	GetItemIgnore() uint32
	IncItemIgnore() uint32
	GetItemError() uint32
	IncItemError() uint32
	GetStatusOk() uint32
	IncStatusOk() uint32
	GetStatusErr() uint32
	IncStatusErr() uint32
}

type StatsWithImage interface {
	Stats
	GetImageTotal() uint32
	IncImageTotal() uint32
}
