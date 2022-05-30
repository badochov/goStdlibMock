package fs

import "io/fs"

func (Default) Glob(fsys FS, pattern string) (matches []string, err error) {
	return fs.Glob(fsys, pattern)
}
func (Default) ReadFile(fsys FS, name string) ([]byte, error) {
	return fs.ReadFile(fsys, name)
}
func (Default) ValidPath(name string) bool {
	return fs.ValidPath(name)
}
func (Default) WalkDir(fsys FS, root string, fn WalkDirFunc) error {
	return fs.WalkDir(fsys, root, fn)
}
func (Default) FileInfoToDirEntry(info FileInfo) DirEntry {
	return fs.FileInfoToDirEntry(info)
}
func (Default) ReadDir(fsys FS, name string) ([]DirEntry, error) {
	return fs.ReadDir(fsys, name)
}
func (Default) Sub(fsys FS, dir string) (FS, error) {
	return fs.Sub(fsys, dir)
}
func (Default) Stat(fsys FS, name string) (FileInfo, error) {
	return fs.Stat(fsys, name)
}
