package os

import "os"

func (Default) O_RDONLY() int {
	return os.O_RDONLY
}
func (Default) O_WRONLY() int {
	return os.O_WRONLY
}
func (Default) O_RDWR() int {
	return os.O_RDWR
}
func (Default) O_APPEND() int {
	return os.O_APPEND
}
func (Default) O_CREATE() int {
	return os.O_CREATE
}
func (Default) O_EXCL() int {
	return os.O_EXCL
}
func (Default) O_SYNC() int {
	return os.O_SYNC
}
func (Default) O_TRUNC() int {
	return os.O_TRUNC
}
func (Default) SEEK_SET() int {
	return os.SEEK_SET
}
func (Default) SEEK_CUR() int {
	return os.SEEK_CUR
}
func (Default) SEEK_END() int {
	return os.SEEK_END
}
func (Default) PathSeparator() string {
	return os.PathSeparator
}
func (Default) PathListSeparator() string {
	return os.PathListSeparator
}
func (Default) ModeDir() FileMode {
	return os.ModeDir
}
func (Default) ModeAppend() FileMode {
	return os.ModeAppend
}
func (Default) ModeExclusive() FileMode {
	return os.ModeExclusive
}
func (Default) ModeTemporary() FileMode {
	return os.ModeTemporary
}
func (Default) ModeSymlink() FileMode {
	return os.ModeSymlink
}
func (Default) ModeDevice() FileMode {
	return os.ModeDevice
}
func (Default) ModeNamedPipe() FileMode {
	return os.ModeNamedPipe
}
func (Default) ModeSocket() FileMode {
	return os.ModeSocket
}
func (Default) ModeSetuid() FileMode {
	return os.ModeSetuid
}
func (Default) ModeSetgid() FileMode {
	return os.ModeSetgid
}
func (Default) ModeCharDevice() FileMode {
	return os.ModeCharDevice
}
func (Default) ModeSticky() FileMode {
	return os.ModeSticky
}
func (Default) ModeIrregular() FileMode {
	return os.ModeIrregular
}
func (Default) ModeType() FileMode {
	return os.ModeType
}
func (Default) ModePerm() FileMode {
	return os.ModePerm
}
func (Default) DevNull() string {
	return os.DevNull
}
func (Default) ErrInvalid() error {
	return os.ErrInvalid
}
func (Default) ErrPermission() error {
	return os.ErrPermission
}
func (Default) ErrExist() error {
	return os.ErrExist
}
func (Default) ErrNotExist() error {
	return os.ErrNotExist
}
func (Default) ErrClosed() error {
	return os.ErrClosed
}
func (Default) ErrNoDeadline() error {
	return os.ErrNoDeadline
}
func (Default) ErrDeadlineExceeded() error {
	return os.ErrDeadlineExceeded
}
func (Default) Stdin() *os.File {
	return os.Stdin
}
func (Default) Stdout() *os.File {
	return os.Stdout
}
func (Default) Stderr() *os.File {
	return os.Stderr
}
func (Default) Args() []string {
	return os.Args
}
func (Default) ErrProcessDone() error {
	return os.ErrProcessDone
}
