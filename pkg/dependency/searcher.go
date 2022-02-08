package dependency

type Result interface {
	// GetFile returns the file
	GetFile() string
	// GetRow returns the row
	GetRow() int
	// GetColumn returns the column
	GetColumn() string
	// String returns the string value of the result
	String() string
}

type Searcher interface {
	Search(path, keyword string) ([]Result, error)
}
