package filepath

type Default struct{}

func New() I {
	return Default{}
}
