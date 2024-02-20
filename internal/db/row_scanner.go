package db

// This interface exists to make possible row.Scan and rows.Scan can be passed as argument to some function
type RowScanner interface {
	Scan(dest ...any) error
}
