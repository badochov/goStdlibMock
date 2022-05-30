package io

import (
	"io"
)

type Writer = io.Writer
type Closer = io.Closer
type Reader = io.Reader
type ReadCloser = io.ReadCloser
type ByteReader = io.ByteReader
type ByteScanner = io.ByteScanner
type ByteWriter = io.ByteWriter
type LimitedReader = io.LimitedReader
type PipeReader interface {
	Close() error
	CloseWithError(err error) error
	Read(data []byte) (n int, err error)
}
type PipeWriter interface {
	Close() error
	CloseWithError(err error) error
	Write(data []byte) (n int, err error)
}
type ReadSeekCloser = io.ReadSeekCloser
type ReadSeeker = io.ReadSeeker
type ReadWriteCloser = io.ReadWriteCloser
type ReadWriteSeeker = io.ReadWriteSeeker
type ReadWriter = io.ReadWriter
type ReaderAt = io.ReaderAt
type ReaderFrom = io.ReaderFrom
type RuneReader = io.RuneReader
type RuneScanner = io.RuneScanner
type SectionReader interface {
	Reader
	Seeker
	ReaderAt
}
type Seeker = io.Seeker
type StringWriter = io.StringWriter
type WriteCloser = io.WriteCloser
type WriteSeeker = io.WriteSeeker
type WriterAt = io.WriterAt
type WriterTo = io.WriterTo
