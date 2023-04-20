package internal

type Image struct {
	Name      string
	Extension string
	Width     int
	Height    int
}

func (i *Image) GetName() string {
	return i.Name
}
func (i *Image) SetName(name string) {
	i.Name = name
}
func (i *Image) GetExtension() string {
	return i.Extension
}
func (i *Image) SetExtension(extension string) {
	i.Extension = extension
}
func (i *Image) GetWidth() int {
	return i.Width
}
func (i *Image) SetWidth(width int) {
	i.Width = width
}
func (i *Image) GetHeight() int {
	return i.Height
}
func (i *Image) SetHeight(height int) {
	i.Height = height
}

type OptionImage interface {
	GetName() string
	SetName(name string)
	GetExtension() string
	SetExtension(extension string)
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}
