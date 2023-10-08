package media

type File struct {
	StorePath string `json:"store_path"`
	Url       string `json:"url"`
	Name      string `json:"name"`
	Ext       string `json:"ext"`
}

func (i *File) GetStorePath() string {
	return i.StorePath
}
func (i *File) SetStorePath(storePath string) {
	i.StorePath = storePath
}
func (i *File) GetUrl() string {
	return i.Url
}
func (i *File) SetUrl(url string) {
	i.Url = url
}
func (i *File) GetName() string {
	return i.Name
}
func (i *File) SetName(name string) {
	i.Name = name
}
func (i *File) GetExt() string {
	return i.Ext
}
func (i *File) SetExt(ext string) {
	i.Ext = ext
}
