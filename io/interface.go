package io

type I interface {
	SeekStart() int
	SeekCurrent() int
	SeekEnd() int
	EOF() error
	ErrClosedPipe() error
	ErrNoProgress() error
	ErrShortBuffer() error
	ErrShortWrite() error
	ErrUnexpectedEOF() error
	Discard() Writer
	Copy(dst Writer, src Reader) (written int64, err error)
	CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
	CopyN(dst Writer, src Reader, n int64) (written int64, err error)
	Pipe() (PipeReader, PipeWriter)
	ReadAll(r Reader) ([]byte, error)
	ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
	ReadFull(r Reader, buf []byte) (n int, err error)
	WriteString(w Writer, s string) (n int, err error)
	NopCloser(r Reader) ReadCloser
	LimitReader(r Reader, n int64) Reader
	MultiReader(readers ...Reader) Reader
	TeeReader(r Reader, w Writer) Reader
	NewSectionReader(r ReaderAt, off int64, n int64) SectionReader
	MultiWriter(writers ...Writer) Writer
}
