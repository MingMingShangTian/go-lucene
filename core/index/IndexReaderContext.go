package index

type IndexReaderContext struct {
	Parent          CompositeReaderContext
	IsTopLevel      bool
	DocBaseInParent int
	OrdInParent     int
	//todo
	//identity
}

func NewIndexReaderContext(parent CompositeReaderContext, ordInParent int, docBaseInParent int) *IndexReaderContext {
	return &IndexReaderContext{
		Parent:          parent,
		DocBaseInParent: docBaseInParent,
		OrdInParent:     ordInParent,
		IsTopLevel:      parent == nil,
	}
}

//todo
func (i *IndexReaderContext) Id() {

}

