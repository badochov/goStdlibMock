package bufio

import "bufio"

func (Default) MaxScanTokenSize() int {
	return bufio.MaxScanTokenSize
}
func (Default) ErrInvalidUnreadByte() error {
	return bufio.ErrInvalidUnreadByte
}
func (Default) ErrInvalidUnreadRune() error {
	return bufio.ErrInvalidUnreadRune
}
func (Default) ErrBufferFull() error {
	return bufio.ErrBufferFull
}
func (Default) ErrNegativeCount() error {
	return bufio.ErrNegativeCount
}
func (Default) ErrTooLong() error {
	return bufio.ErrTooLong
}
func (Default) ErrNegativeAdvance() error {
	return bufio.ErrNegativeAdvance
}
func (Default) ErrAdvanceTooFar() error {
	return bufio.ErrAdvanceTooFar
}
func (Default) ErrBadReadCount() error {
	return bufio.ErrBadReadCount
}
func (Default) ErrFinalToken() error {
	return bufio.ErrFinalToken
}
