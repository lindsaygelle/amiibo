package main

var (
	_ rawAmiiboPresentedBy = (*RawAmiiboPresentedBy)(nil)
)

type rawAmiiboPresentedBy interface{}

type RawAmiiboPresentedBy string
