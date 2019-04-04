package stl

// Container base method
type Container interface {
	Size() int
	Clear()
	IsEmpty() bool
}
