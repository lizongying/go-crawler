package pkg

type Filter interface {
	Exists(any) bool
	ExistsOrStore(any) bool
	Clean()
}
