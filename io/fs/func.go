package fs

import "io/fs"

func (Default) Glob(fsys FS, pattern string) (matches []string, err error) {
	matches, err = fs.Glob(fsys, pattern)
	return matches, castError(err)
}
func (Default) ReadFile(fsys FS, name string) ([]byte, error) {
	c, err := fs.ReadFile(fsys, name)
	return c, castError(err)
}
func (Default) ValidPath(name string) bool {
	return fs.ValidPath(name)
}
func (Default) WalkDir(fsys FS, root string, fn WalkDirFunc) error {
	return castError(fs.WalkDir(fsys, root, fn))
}
func (Default) FileInfoToDirEntry(info FileInfo) DirEntry {
	return fs.FileInfoToDirEntry(info)
}
func (Default) ReadDir(fsys FS, name string) ([]DirEntry, error) {
	e, err := fs.ReadDir(fsys, name)
	return e, castError(err)
}
func (Default) Sub(fsys FS, dir string) (FS, error) {
	f, err := fs.Sub(fsys, dir)
	return f, castError(err)
}
func (Default) Stat(fsys FS, name string) (FileInfo, error) {
	i, err := fs.Stat(fsys, name)
	return i, castError(err)
}
