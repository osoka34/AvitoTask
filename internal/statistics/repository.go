package statistics

type Repository interface {
	SelectByDates(params *SelectParams) ([]SelectOut, error)
	AddRows(params *InsertParams) error
}
