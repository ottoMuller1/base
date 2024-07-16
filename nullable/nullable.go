package nullable


// nullable type
type Nullable[t any] struct{
	value t
	filled bool
}


// null term
func Null[t any]() Nullable[t] {
	return Nullable[t]{filled: false}
}


// value to nullable
func ToNullable[t any](value t) Nullable[t] {

	return Nullable[t] {
		filled: true,
		value: value,
	}

}


// value to nullable by pointer
func ToNullablePointer[t any](value *t) Nullable[t] {

	if value == nil {
		return Nullable[t]{filled: false}
	}

	return Nullable[t]{
		filled: true,
		value: *value,
	}

}


// to pointer
func (nullable Nullable[t]) ToPointer() *t {

	if !nullable.filled {
		return nil
	}

	return &nullable.value

}


// check if nullable is empty
func (nullable Nullable[t]) IsEmpty() bool {
	return !nullable.filled
}


// handle nullable
func Handle[t, k any](nullable Nullable[t], def k, handler func(t) k) k {

	if nullable.IsEmpty() {
		return def
	}

	return handler(nullable.value)

}


// get value or default, provide error if not clean
func (nullable Nullable[t]) FromNullable(defaultValue t, err error) t {

	if !nullable.filled && err != nil {
		panic(err)
	}

	idFunction := func (a t) t {
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
		value: slice[index],
		filled: true,
	}

}
