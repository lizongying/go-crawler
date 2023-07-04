package media

type File struct {
	StorePath string
	Name      string
	Extension string
}

func (i *File) GetStorePath() string {
	return i.StorePath
}
func (i *File) SetStorePath(storePath string) {
	i.StorePath = storePath
}
func (i *File) GetName() string {
	return i.Name
}
func (i *File) SetName(name string) {
	i.Name = name
}
func (i *File) GetExtension() string {
	return i.Extension
}
func (i *File) SetExtension(extension string) {
	i.Extension = extension
}
