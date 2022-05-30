package filepath

import "path/filepath"

func (Default) Separator() string {
	return filepath.Separator
}
func (Default) ListSeparator() string {
	return filepath.ListSeparator
}
func (Default) ErrBadPattern() error {
	return filepath.ErrBadPattern
}
func (Default) SkipDir() error {
	return filepath.SkipDir
}
