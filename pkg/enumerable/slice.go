package enumerable

type SliceEnumerable []interface{}

func (input SliceEnumerable) Map(f func(interface{}) interface{}) Enumerable {
	output := make(SliceEnumerable, len(input))

	for i, elem := range input {
		output[i] = f(elem)
	}

	return output
}

func (input SliceEnumerable) Filter(f func(interface{}) bool) Enumerable {
	output := SliceEnumerable{}

	for _, elem := range input {
		if f(elem) {
			output = append(output, elem)
		}
	}

	return output
}

func (input SliceEnumerable) Fold(initialValue interface{}, f func(interface{}, interface{}) interface{}) interface{} {
	accumulated := initialValue

	for _, elem := range input {
		accumulated = f(accumulated, elem)
	}

	return accumulated
}
