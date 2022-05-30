package os

import (
	"os"
	"stdlibMock/io"
	"stdlibMock/io/fs"
	"syscall"
	"time"
)

type FileMode = fs.FileMode
type FileInfo = fs.FileInfo
type File interface {
	Fd() uintptr
	Readdir(n int) ([]FileInfo, error)
	Readdirnames(n int) (names []string, err error)
	ReadDir(n int) ([]DirEntry, error)
	Name() string
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	ReadFrom(r io.Reader) (n int64, err error)
	Write(b []byte) (n int, err error)
	WriteAt(b []byte, off int64) (n int, err error)
	Seek(offset int64, whence int) (ret int64, err error)
	WriteString(s string) (n int, err error)
	Chmod(mode FileMode) error
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
	Stat() (FileInfo, error)
	Close() error
	Truncate(size int64) error
	Sync() error
	Chown(uid, gid int) error
	Chdir() error
}
type Process interface {
	Release() error
	Kill() error
	Wait() (ProcessState, error)
	Signal(sig os.Signal) error
	Pid() int
}

type ProcAttr = os.ProcAttr

type ProcessState interface {
	UserTime() time.Duration
	SystemTime() time.Duration
	Exited() bool
	Success() bool
	Sys() any
	SysUsage() any
	Pid() int
	String() string
	ExitCode() int
}

type DirEntry = fs.DirEntry
type PathError = fs.PathError
type Signal = os.Signal
type SyscallError = os.SyscallError

type processWrapper struct{ *os.Process }

func (p *processWrapper) Pid() int {
	return p.Process.Pid
}

func (p *processWrapper) Wait() (ProcessState, error) {
	return p.Wait()
}

func toProcess(p *os.Process) Process {
	return &processWrapper{p}
}
