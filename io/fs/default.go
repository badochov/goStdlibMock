package fs

type Default struct{}

func New() Fs {
	return Default{}
}
