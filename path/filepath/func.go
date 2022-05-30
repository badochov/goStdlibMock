package filepath

import (
	"github.com/badochov/goStdlibMock/io/fs"
	"path/filepath"
)

func (Default) Abs(path string) (string, error) {
	return filepath.Abs(path)
}
func (Default) Base(path string) string {
	return filepath.Base(path)
}
func (Default) Clean(path string) string {
	return filepath.Clean(path)
}
func (Default) Dir(path string) string {
	return filepath.Dir(path)
}
func (Default) EvalSymlinks(path string) (string, error) {
	return filepath.EvalSymlinks(path)
}
func (Default) Ext(path string) string {
	return filepath.Ext(path)
}
func (Default) FromSlash(path string) string {
	return filepath.FromSlash(path)
}
func (Default) Glob(pattern string) (matches []string, err error) {
	return filepath.Glob(pattern)
}
func (Default) IsAbs(path string) bool {
	return filepath.IsAbs(path)
}
func (Default) Join(elem ...string) string {
	return filepath.Join(elem...)
}
func (Default) Match(pattern, name string) (matched bool, err error) {
	return filepath.Match(pattern, name)
}
func (Default) Rel(basepath, targpath string) (string, error) {
	return filepath.Rel(basepath, targpath)
}
func (Default) Split(path string) (dir, file string) {
	return filepath.Split(path)
}
func (Default) SplitList(path string) []string {
	return filepath.SplitList(path)
}
func (Default) ToSlash(path string) string {
	return filepath.ToSlash(path)
}
func (Default) VolumeName(path string) string {
	return filepath.VolumeName(path)
}
func (Default) Walk(root string, fn WalkFunc) error {
	return filepath.Walk(root, fn)
}
func (Default) WalkDir(root string, fn fs.WalkDirFunc) error {
	return filepath.WalkDir(root, fn)
}
