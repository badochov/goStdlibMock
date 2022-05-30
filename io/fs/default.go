package fs

type Default struct{}

func New() I {
	return Default{}
}
