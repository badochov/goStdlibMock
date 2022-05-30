package io

import "io"

func (Default) SeekStart() int {
	return io.SeekStart
}
func (Default) SeekCurrent() int {
	return io.SeekCurrent
}
func (Default) SeekEnd() int {
	return io.SeekEnd
}
func (Default) EOF() error {
	return io.EOF
}
func (Default) ErrClosedPipe() error {
	return io.ErrClosedPipe
}
func (Default) ErrNoProgress() error {
	return io.ErrNoProgress
}
func (Default) ErrShortBuffer() error {
	return io.ErrShortBuffer
}
func (Default) ErrShortWrite() error {
	return io.ErrShortWrite
}
func (Default) ErrUnexpectedEOF() error {
	return io.ErrUnexpectedEOF
}
func (Default) Discard() Writer {
	return io.Discard
}
