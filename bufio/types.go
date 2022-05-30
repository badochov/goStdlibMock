package bufio

import (
	"bufio"
	"github.com/badochov/goStdlibMock/io"
)

type ReadWriter struct {
	Reader
	Writer
}

type Reader interface {
	Size() int
	Reset(r io.Reader)
	Peek(n int) ([]byte, error)
	Discard(n int) (discarded int, err error)
	Read(p []byte) (n int, err error)
	ReadByte() (byte, error)
	UnreadByte() error
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
	Buffered() int
	ReadSlice(delim byte) (line []byte, err error)
	ReadLine() (line []byte, isPrefix bool, err error)
	ReadBytes(delim byte) ([]byte, error)
	ReadString(delim byte) (string, error)
	WriteTo(w io.Writer) (n int64, err error)
}

type Writer interface {
	Size() int
	Reset(w io.Writer)
	Flush() error
	Available() int
	AvailableBuffer() []byte
	Buffered() int
	Write(p []byte) (nn int, err error)
	WriteByte(c byte) error
	WriteRune(r rune) (size int, err error)
	WriteString(s string) (int, error)
	ReadFrom(r io.Reader) (n int64, err error)
}

type SplitFunc = bufio.SplitFunc

type Scanner interface {
	MaxTokenSize(n int)
	ErrOrEOF() error
	Err() error
	Bytes() []byte
	Text() string
	Scan() bool
	Buffer(buf []byte, max int)
	Split(split SplitFunc)
}
