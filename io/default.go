package io

type Default struct{}

func New() I {
	return Default{}
}
