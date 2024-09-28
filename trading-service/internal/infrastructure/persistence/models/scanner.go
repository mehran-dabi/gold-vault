package models

// Scanner interface is used to be able to scan both *sql.Row and *sql.Rows
type Scanner interface {
	Scan(dest ...interface{}) error
}
