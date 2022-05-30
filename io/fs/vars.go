package fs

import "io/fs"

func (Default) ErrInvalid() error {
	return fs.ErrInvalid
}
func (Default) ErrPermission() error {
	return fs.ErrPermission
}
func (Default) ErrExist() error {
	return fs.ErrExist
}
func (Default) ErrNotExist() error {
	return fs.ErrNotExist
}
func (Default) ErrClosed() error {
	return fs.ErrClosed
}
func (Default) SkipDir() error {
	return fs.SkipDir
}
