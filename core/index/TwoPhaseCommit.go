package index

type TwoPhaseCommit interface {
	PrepareCommit() int64
	Commit() int64
	Rollback()
}
