package bufio

type Default struct{}

func New() Bufio {
	return Default{}
}
