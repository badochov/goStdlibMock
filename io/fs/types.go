package fs

import "io/fs"

type DirEntry = fs.DirEntry
type FS = fs.FS
type File = fs.File
type FileInfo = fs.FileInfo
type FileMode interface {
	String() string
	IsDir() bool
	IsRegular() bool
	Perm() FileMode
	Type() FileMode
	ToFsFileMode() fs.FileMode
}
type GlobFS = fs.GlobFS
type PathError struct{ *fs.PathError }
type ReadDirFs = fs.ReadDirFS
type ReadDirFile = fs.ReadDirFile
type ReadFileFS = fs.ReadFileFS
type StatFS = fs.StatFS
type SubFS = fs.SubFS
type WalkDirFunc = fs.WalkDirFunc

type modeWrapper struct{ fs.FileMode }

func (m modeWrapper) Perm() FileMode {
	return ToFileMode(m.FileMode)
}

func (m modeWrapper) Type() FileMode {
	return ToFileMode(m.FileMode)
}

func (m modeWrapper) ToFsFileMode() fs.FileMode {
	return m.FileMode
}

func ToFileMode(mode fs.FileMode) FileMode {
	return modeWrapper{mode}
}

func castError(err error) error {
	if e, ok := err.(*fs.PathError); ok {
		return &PathError{e}
	}
	return err
}
