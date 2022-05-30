package bufio

import (
	"bufio"
	"github.com/badochov/goStdlibMock/io"
)

func (Default) ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return bufio.ScanBytes(data, atEOF)
}
func (Default) ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return bufio.ScanLines(data, atEOF)
}
func (Default) ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return bufio.ScanRunes(data, atEOF)
}
func (Default) ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return bufio.ScanWords(data, atEOF)
}
func (Default) NewReadWriter(r Reader, w Writer) *ReadWriter {
	return &ReadWriter{r, w}
}
func (Default) NewReader(rd io.Reader) Reader {
	return bufio.NewReader(rd)
}
func (Default) NewReaderSize(rd io.Reader, size int) Reader {
	return bufio.NewReaderSize(rd, size)
}
func (Default) NewScanner(r io.Reader) Scanner {
	return bufio.NewScanner(r)
}
func (Default) NewWriter(w io.Writer) Writer {
	return bufio.NewWriter(w)
}
func (Default) NewWriterSize(w io.Writer, size int) Writer {
	return bufio.NewWriterSize(w, size)
}
