package os

import (
	"io/fs"
	"os"
	"time"
)

type Os interface {
	Chdir(dir string) error
	Chmod(name string, mode FileMode) error
	Chown(name string, uid, gid int) error
	Chtimes(name string, atime time.Time, mtime time.Time) error
	Clearenv()
	DirFS(dir string) fs.FS
	Environ() []string
	Executable() (string, error)
	Exit(code int)
	Expand(s string, mapping func(string) string) string
	ExpandEnv(s string) string
	Getegid() int
	Getenv(key string) string
	Geteuid() int
	Getgid() int
	Getgroups() ([]int, error)
	Getpagesize() int
	Getpid() int
	Getppid() int
	Getuid() int
	Getwd() (dir string, err error)
	Hostname() (name string, err error)
	IsExist(err error) bool
	IsNotExist(err error) bool
	IsPathSeparator(c uint8) bool
	IsPermission(err error) bool
	IsTimeout(err error) bool
	Lchown(name string, uid, gid int) error
	Link(oldname, newname string) error
	LookupEnv(key string) (string, bool)
	Mkdir(name string, perm FileMode) error
	MkdirAll(path string, perm FileMode) error
	MkdirTemp(dir, pattern string) (string, error)
	NewSyscallError(syscall string, err error) error
	Pipe() (r File, w File, err error)
	ReadFile(name string) ([]byte, error)
	Readlink(name string) (string, error)
	Remove(name string) error
	RemoveAll(path string) error
	Rename(oldpath, newpath string) error
	SameFile(fi1, fi2 FileInfo) bool
	Setenv(key, value string) error
	Symlink(oldname, newname string) error
	TempDir() string
	Truncate(name string, size int64) error
	Unsetenv(key string) error
	UserCacheDir() (string, error)
	UserConfigDir() (string, error)
	UserHomeDir() (string, error)
	WriteFile(name string, data []byte, perm FileMode) error
	ReadDir(name string) ([]DirEntry, error)
	Create(name string) (File, error)
	CreateTemp(dir, pattern string) (File, error)
	NewFile(fd uintptr, name string) File
	Open(name string) (File, error)
	OpenFile(name string, flag int, perm FileMode) (File, error)
	Lstat(name string) (FileInfo, error)
	Stat(name string) (FileInfo, error)
	FindProcess(pid int) (Process, error)
	StartProcess(name string, argv []string, attr *ProcAttr) (Process, error)
	O_RDONLY() int
	O_WRONLY() int
	O_RDWR() int
	O_APPEND() int
	O_CREATE() int
	O_EXCL() int
	O_SYNC() int
	O_TRUNC() int
	SEEK_SET() int
	SEEK_CUR() int
	SEEK_END() int
	PathSeparator() string
	PathListSeparator() string
	ModeDir() FileMode
	ModeAppend() FileMode
	ModeExclusive() FileMode
	ModeTemporary() FileMode
	ModeSymlink() FileMode
	ModeDevice() FileMode
	ModeNamedPipe() FileMode
	ModeSocket() FileMode
	ModeSetuid() FileMode
	ModeSetgid() FileMode
	ModeCharDevice() FileMode
	ModeSticky() FileMode
	ModeIrregular() FileMode
	ModeType() FileMode
	ModePerm() FileMode
	DevNull() string
	ErrInvalid() error
	ErrPermission() error
	ErrExist() error
	ErrNotExist() error
	ErrClosed() error
	ErrNoDeadline() error
	ErrDeadlineExceeded() error
	Stdin() *os.File
	Stdout() *os.File
	Stderr() *os.File
	Args() []string
	ErrProcessDone() error
}
