package app_types

type Operation int

const (
	Create Operation = iota
	Delete
	Update
)
