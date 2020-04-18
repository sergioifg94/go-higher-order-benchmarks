package enumerable

type Enumerable interface {
	Map(func(interface{}) interface{}) Enumerable
	Filter(func(interface{}) bool) Enumerable
	Fold(interface{}, func(interface{}, interface{}) interface{}) interface{}
}