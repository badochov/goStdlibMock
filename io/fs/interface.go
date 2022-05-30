package fs

type I interface {
	ErrInvalid() error
	ErrPermission() error
	ErrExist() error
	ErrNotExist() error
	ErrClosed() error
	SkipDir() error
	Glob(fsys FS, pattern string) (matches []string, err error)
	ReadFile(fsys FS, name string) ([]byte, error)
	ValidPath(name string) bool
	WalkDir(fsys FS, root string, fn WalkDirFunc) error
	FileInfoToDirEntry(info FileInfo) DirEntry
	ReadDir(fsys FS, name string) ([]DirEntry, error)
	Sub(fsys FS, dir string) (FS, error)
	Stat(fsys FS, name string) (FileInfo, error)
}
