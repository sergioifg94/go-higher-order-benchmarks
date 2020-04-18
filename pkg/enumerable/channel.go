package enumerable

type ChanEnumerable <-chan interface{}

func (input ChanEnumerable) Map(f func(interface{}) interface{}) Enumerable {
	output := make(chan interface{})

	go func() {
		for elem := range input {
			mapped := f(elem)
			output <- mapped
		}

		close(output)
	}()

	return ChanEnumerable(output)
}

func (input ChanEnumerable) Filter(f func(interface{}) bool) Enumerable {
	output := make(chan interface{})

	go func() {
		for elem := range input {
			if f(elem) {
				output <- elem
			}
		}

		close(output)
	}()

	return ChanEnumerable(output)
}

func (input ChanEnumerable) Fold(initialValue interface{}, f func(interface{}, interface{}) interface{}) interface{} {
	accumulated := initialValue

	for elem := range input {
		accumulated = f(accumulated, elem)
	}

	return accumulated
}
