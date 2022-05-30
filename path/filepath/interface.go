package filepath

import "stdlibMock/io/fs"

type I interface {
	Separator() string
	ListSeparator() string
	ErrBadPattern() error
	SkipDir() error
	Abs(path string) (string, error)
	Base(path string) string
	Clean(path string) string
	Dir(path string) string
	EvalSymlinks(path string) (string, error)
	Ext(path string) string
	FromSlash(path string) string
	Glob(pattern string) (matches []string, err error)
	IsAbs(path string) bool
	Join(elem ...string) string
	Match(pattern, name string) (matched bool, err error)
	Rel(basepath, targpath string) (string, error)
	Split(path string) (dir, file string)
	SplitList(path string) []string
	ToSlash(path string) string
	VolumeName(path string) string
	Walk(root string, fn WalkFunc) error
	WalkDir(root string, fn fs.WalkDirFunc) error
}
