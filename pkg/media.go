package pkg

type File interface {
	GetStorePath() string
	SetStorePath(string)
	GetUrl() string
	SetUrl(string)
	GetName() string
	SetName(string)
	GetExt() string
	SetExt(string)
}

type Image interface {
	File
	GetWidth() int
	SetWidth(int)
	GetHeight() int
	SetHeight(int)
}
