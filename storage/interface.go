package storage

import "context"

type (
	// ID MessageKey in underlying KV Engine
	// TODO is 64 too large?
	ID    [64]byte
	Value []byte
)

type KVStorage interface {
	Init(context.Context) error
	Destroy(context.Context) error

	Get(context.Context, ID) (Value, error)
	Put(context.Context, ID, Value) error
	Delete(context.Context, ID) error

	BatchGet(context.Context, ...ID) ([]Value, error)
	BatchDelete(context.Context, ...ID) error

	RangeGet(context.Context, ID, ID) ([]Value, error)
	RangeDelete(context.Context, ID, ID) error
}

