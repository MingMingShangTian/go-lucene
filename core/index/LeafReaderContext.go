package index

type LeafReaderContext struct {
	Ord                   int
	DocBase               int
	Reader                LeafReader
	LeavesContext                []LeafReaderContext
	IndexReaderContextPtr *IndexReaderContext
}

func NewLeafReaderContext(parent CompositeReaderContext, reader LeafReader, ord int, docBase int, leafOrd int, leafDocBase int) *LeafReaderContext {
	l := LeafReaderContext{
		Ord:                   leafOrd,
		DocBase:               leafDocBase,
		Reader:                reader,
		IndexReaderContextPtr: NewIndexReaderContext(parent, ord, docBase),
	}
	l.LeavesContext = []LeafReaderContext{l}
	return &l
}

func (l *LeafReaderContext) Leaves() []LeafReaderContext {
	if !l.IndexReaderContextPtr.IsTopLevel {
		panic("This is not a top-level context.")
	}

	return l.LeavesContext
}

func (l *LeafReaderContext) Children() []LeafReaderContext {
	return nil
}

func (l *LeafReaderContext) toString() {
	//todo
	//return "LeafReaderContext(" + reader + " docBase=" + docBase + " ord=" + ord + ")"
}

