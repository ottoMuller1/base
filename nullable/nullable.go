package nullable

import (
	"encoding/json"
)

// nullable type
type Nullable[t any] struct {
	filled bool
	Error  error
	value  t
}

func (nl Nullable[T]) ToJSONByte() ([]byte, error) {
	if !nl.filled {
		return json.Marshal(nil)
	}

	return json.Marshal(
		struct {
			Value T `json:"value"`
		}{
			Value: nl.value,
		},
	)
}

func FromJSONByte[T any](data []byte) (*Nullable[T], error) {
	var temp *struct {
		Value T `json:"value"`
	}

	err := json.Unmarshal(data, &temp)

	if err != nil {
		return nil, err
	}

	if temp == nil {
		return nil, nil
	}

	return &Nullable[T]{true, nil, temp.Value}, nil
}

// null term
func Null[t any]() Nullable[t] {
	return Nullable[t]{filled: false}
}

// value to nullable
func ToNullable[t any](value t) Nullable[t] {

	return Nullable[t]{
		filled: true,
		value:  value,
	}

}

// value to nullable by pointer
func ToNullablePointer[t any](value *t) Nullable[t] {

	if value == nil {
		return Nullable[t]{filled: false}
	}

	return Nullable[t]{
		filled: true,
		value:  *value,
	}

}

// to pointer
func (nullable Nullable[t]) ToPointer() *t {

	if !nullable.filled {
		return nil
	}

	return &nullable.value

}

// pass an error use instead of .Err if inlinable Err reset is required
func (nullable Nullable[t]) PassError(err error) Nullable[t] {
	nullable.Error = err
	return nullable
}

// check if nullable is empty
func (nullable Nullable[t]) IsEmpty() bool {
	return !nullable.filled
}

// handle nullable
func Handle[t, k any](nullable Nullable[t], def k, handler func(t) k) k {

	isEmpty := nullable.IsEmpty()

	if isEmpty && nullable.Error != nil {
		panic(nullable.Error)
	}

	if isEmpty {
		return def
	}

	return handler(nullable.value)

}

// get value or default, invokes exception if not pure
func (nullable Nullable[t]) FromNullable(defaultValue t) t {

	idFunction := func(a t) t {
		return a
	}

	return Handle(nullable, defaultValue, idFunction)

}

// extra: get value if slice by index
func SliceIndex[t any](slice []t, index int) Nullable[t] {

	if index < 0 || index >= len(slice) {
		return Nullable[t]{filled: false}
	}

	return Nullable[t]{
		value:  slice[index],
		filled: true,
	}

}
