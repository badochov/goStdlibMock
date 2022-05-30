package bufio

import "stdlibMock/io"

type I interface {
	MaxScanTokenSize() int
	ErrInvalidUnreadByte() error
	ErrInvalidUnreadRune() error
	ErrBufferFull() error
	ErrNegativeCount() error
	ErrTooLong() error
	ErrNegativeAdvance() error
	ErrAdvanceTooFar() error
	ErrBadReadCount() error
	ErrFinalToken() error
	ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
	ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
	ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
	ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
	NewReadWriter(r Reader, w Writer) *ReadWriter
	NewReader(rd io.Reader) Reader
	NewReaderSize(rd io.Reader, size int) Reader
	NewScanner(r io.Reader) Scanner
	NewWriter(w io.Writer) Writer
	NewWriterSize(w io.Writer, size int) Writer
}
