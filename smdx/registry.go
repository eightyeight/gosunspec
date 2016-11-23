package smdx

import (
	"errors"
)

var (
	modelRegistry  = map[uint16]*ModelElement{}
	ErrNoSuchModel = errors.New("no such model")
)

// RegisterModel registers a new model element.
func RegisterModel(m *ModelElement) {
	modelRegistry[m.Id] = m
}

// GetModel answers the model element corresponding to the specified identifier
// or nil if no such element exists.
func GetModel(id uint16) *ModelElement {
	return modelRegistry[id]
}
