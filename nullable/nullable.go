package nullable


// nullable type
type Nullable[t any] struct{
	filled    bool
	errPassed bool
	err       error
	value     t
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



// put an error into nullable
func (nullable Nullable[t]) PassError(err error) Nullable[t] {
	
	if !nullable.errPassed {
		nullable.err = err
		nullable.errPassed = true
		return nullable
	}

	return nullable

}





// get an error from nullable
func (nullable Nullable[t]) GetError() error {
	return nullable.err
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

	isEmpty := nullable.IsEmpty()

	if isEmpty && nullable.err != nil {
		panic(nullable.err)
	}

	if isEmpty {
		return def
	}

	return handler(nullable.value)

}


// get value or default, invokes exception if not pure
func (nullable Nullable[t]) FromNullable(defaultValue t) t {

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
