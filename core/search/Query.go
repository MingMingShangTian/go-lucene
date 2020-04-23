package search

type Query struct {

}

func (*Query) ToString(filed string) {
	return filed
}