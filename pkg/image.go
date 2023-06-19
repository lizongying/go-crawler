package pkg

type Image interface {
	GetName() string
	SetName(name string)
	GetExtension() string
	SetExtension(extension string)
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}
