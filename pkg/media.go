package pkg

type File interface {
	GetStorePath() string
	SetStorePath(string)
	GetName() string
	SetName(string)
	GetExtension() string
	SetExtension(string)
}

type Image interface {
	File
	GetWidth() int
	SetWidth(int)
	GetHeight() int
	SetHeight(int)
}
