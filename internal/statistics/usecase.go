package statistics

type Usecase interface {
	SelectForCsv(params *SelectParams) (*StatResponse, error)
	AddRows(params *InsertParams) error
}
