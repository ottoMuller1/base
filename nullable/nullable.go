package nullable



type Nullable[t any] struct{
	value t
	empty bool
}



func ToNullable[t any](value *t) Nullable[t] {

	if value == nil {
		return Nullable[t] {empty: true}
	}

	return Nullable[t] {
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



func (nullable Nullable[t]) FromNullable(defaultValue t) t {

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
