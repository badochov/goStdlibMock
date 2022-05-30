package os

import (
	"github.com/badochov/stdlibMock/io/fs"
	"os"
	"time"
)

func (Default) Chdir(dir string) error {
	return os.Chdir(dir)
}
func (Default) Chmod(name string, mode FileMode) error {
	return os.Chmod(name, mode)
}
func (Default) Chown(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}
func (Default) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}
func (Default) Clearenv() {
	os.Clearenv()
}
func (Default) DirFS(dir string) fs.FS {
	return os.DirFS(dir)
}
func (Default) Environ() []string {
	return os.Environ()
}
func (Default) Executable() (string, error) {
	return os.Executable()
}
func (Default) Exit(code int) {
	os.Exit(code)
}
func (Default) Expand(s string, mapping func(string) string) string {
	return os.Expand(s, mapping)
}
func (Default) ExpandEnv(s string) string {
	return os.ExpandEnv(s)
}
func (Default) Getegid() int {
	return os.Getegid()
}
func (Default) Getenv(key string) string {
	return os.Getenv(key)
}
func (Default) Geteuid() int {
	return os.Geteuid()
}
func (Default) Getgid() int {
	return os.Getgid()
}
func (Default) Getgroups() ([]int, error) {
	return os.Getgroups()
}
func (Default) Getpagesize() int {
	return os.Getpagesize()
}
func (Default) Getpid() int {
	return os.Getpid()
}
func (Default) Getppid() int {
	return os.Getppid()
}
func (Default) Getuid() int {
	return os.Getuid()
}
func (Default) Getwd() (dir string, err error) {
	return os.Getwd()
}
func (Default) Hostname() (name string, err error) {
	return os.Hostname()
}
func (Default) IsExist(err error) bool {
	return os.IsExist(err)
}
func (Default) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}
func (Default) IsPathSeparator(c uint8) bool {
	return os.IsPathSeparator(c)
}
func (Default) IsPermission(err error) bool {
	return os.IsPermission(err)
}
func (Default) IsTimeout(err error) bool {
	return os.IsTimeout(err)
}
func (Default) Lchown(name string, uid, gid int) error {
	return os.Lchown(name, uid, gid)
}
func (Default) Link(oldname, newname string) error {
	return os.Link(oldname, newname)
}
func (Default) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}
func (Default) Mkdir(name string, perm FileMode) error {
	return os.Mkdir(name, perm)
}
func (Default) MkdirAll(path string, perm FileMode) error {
	return os.MkdirAll(path, perm)
}
func (Default) MkdirTemp(dir, pattern string) (string, error) {
	return os.MkdirTemp(dir, pattern)
}
func (Default) NewSyscallError(syscall string, err error) error {
	return os.NewSyscallError(syscall, err)
}
func (Default) Pipe() (r File, w File, err error) {
	return os.Pipe()
}
func (Default) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}
func (Default) Readlink(name string) (string, error) {
	return os.Readlink(name)
}
func (Default) Remove(name string) error {
	return os.Remove(name)
}
func (Default) RemoveAll(path string) error {
	return os.RemoveAll(path)
}
func (Default) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}
func (Default) SameFile(fi1, fi2 FileInfo) bool {
	return os.SameFile(fi1, fi2)
}
func (Default) Setenv(key, value string) error {
	return os.Setenv(key, value)
}
func (Default) Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}
func (Default) TempDir() string {
	return os.TempDir()
}
func (Default) Truncate(name string, size int64) error {
	return os.Truncate(name, size)
}
func (Default) Unsetenv(key string) error {
	return os.Unsetenv(key)
}
func (Default) UserCacheDir() (string, error) {
	return os.UserCacheDir()
}
func (Default) UserConfigDir() (string, error) {
	return os.UserConfigDir()
}
func (Default) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}
func (Default) WriteFile(name string, data []byte, perm FileMode) error {
	return os.WriteFile(name, data, perm)
}
func (Default) Create(name string) (File, error) {
	return os.Create(name)
}
func (Default) CreateTemp(dir, pattern string) (File, error) {
	return os.CreateTemp(dir, pattern)
}
func (Default) NewFile(fd uintptr, name string) File {
	return os.NewFile(fd, name)
}
func (Default) Open(name string) (File, error) {
	return os.Open(name)
}
func (Default) OpenFile(name string, flag int, perm FileMode) (File, error) {
	return os.OpenFile(name, flag, perm)
}
func (Default) Lstat(name string) (FileInfo, error) {
	return os.Lstat(name)
}
func (Default) Stat(name string) (FileInfo, error) {
	return os.Stat(name)
}
func (Default) FindProcess(pid int) (Process, error) {
	p, err := os.FindProcess(pid)
	return toProcess(p), err
}

func (Default) StartProcess(name string, argv []string, attr *ProcAttr) (Process, error) {
	p, err := os.StartProcess(name, argv, attr)
	return toProcess(p), err
}

func (d Default) ReadDir(name string) ([]DirEntry, error) {
	return os.ReadDir(name)
}
