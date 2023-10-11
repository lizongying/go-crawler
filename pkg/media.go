package pkg

type FileOptions struct {
	Url  bool `json:"url,omitempty"`
	Name bool `json:"name,omitempty"`
	Ext  bool `json:"ext,omitempty"`
}

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

type ImageOptions struct {
	FileOptions
	Width  bool `json:"width,omitempty"`
	Height bool `json:"height,omitempty"`
}

type Image interface {
	File
	GetWidth() int
	SetWidth(int)
	GetHeight() int
	SetHeight(int)
}
