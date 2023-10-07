package media

type Image struct {
	File
	Width  int `json:"width"`
	Height int `json:"height"`
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
