package io

import "io"

func (Default) Copy(dst Writer, src Reader) (written int64, err error) {
	return io.Copy(dst, src)
}
func (Default) CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	return io.CopyBuffer(dst, src, buf)
}
func (Default) CopyN(dst Writer, src Reader, n int64) (written int64, err error) {
	return io.CopyN(dst, src, n)
}
func (Default) Pipe() (PipeReader, PipeWriter) {
	return io.Pipe()
}
func (Default) ReadAll(r Reader) ([]byte, error) {
	return io.ReadAll(r)
}
func (Default) ReadAtLeast(r Reader, buf []byte, min int) (n int, err error) {
	return io.ReadAtLeast(r, buf, min)
}
func (Default) ReadFull(r Reader, buf []byte) (n int, err error) {
	return io.ReadFull(r, buf)
}
func (Default) WriteString(w Writer, s string) (n int, err error) {
	return io.WriteString(w, s)
}
func (Default) NopCloser(r Reader) ReadCloser {
	return io.NopCloser(r)
}
func (Default) LimitReader(r Reader, n int64) Reader {
	return io.LimitReader(r, n)
}
func (Default) MultiReader(readers ...Reader) Reader {
	return io.MultiReader(readers...)
}
func (Default) TeeReader(r Reader, w Writer) Reader {
	return io.TeeReader(r, w)
}
func (Default) NewSectionReader(r ReaderAt, off int64, n int64) SectionReader {
	return io.NewSectionReader(r, off, n)
}
func (Default) MultiWriter(writers ...Writer) Writer {
	return io.MultiWriter(writers...)
}
