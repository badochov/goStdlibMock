package io

type Default struct{}

func New() Io {
	return Default{}
}
