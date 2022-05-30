package os

type Default struct{}

func New() Os {
	return Default{}
}
