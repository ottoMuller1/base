package nullable



type Nullable[t any] struct{
	value t
	empty bool
}



func Null[t any]() Nullable[t] {
	return Nullable[t]{empty: true}
}



func ToNullable[t any](empty bool, value t) Nullable[t] {

	return Nullable[t] {
		empty: empty,
		value: value,
	}

}



func ToNullablePointer[t any](value *t) Nullable[t] {

	if value == nil {
		return Nullable[t]{empty: true}
	}

	return Nullable[t]{
		empty: false,
		value: *value,
	}

}



func (nullable Nullable[t]) IsEmpty() bool {
	return nullable.empty
}



func Handle[t, k any](nullable Nullable[t], def k, handler func(t) k) k {

	if nullable.IsEmpty() {
		return def
	}

	return handler(nullable.value)

}



func (nullable Nullable[t]) FromNullable(defaultValue t, clean bool) t {

	if nullable.empty && !clean {
		panic("Nullable is empty")
	}

	idFunction := func (a t) t {
		return a
	}

	return Handle(nullable, defaultValue, idFunction)

}



func SliceIndex[t any](slice []t, index int) Nullable[t] {

	if index < 0 || index >= len(slice) {
		return Nullable[t]{empty: true}
	}

	return Nullable[t]{
		value: slice[index],
		empty: false,
	}

}
