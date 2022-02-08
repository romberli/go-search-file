package excel

import (
	"fmt"

	"github.com/romberli/go-search-file/pkg/dependency"
)

var _ dependency.Result = (*Result)(nil)

type Result struct {
	file   string
	row    int
	column string
}

func NewResult(file string, row int, column string) *Result {
	return &Result{
		file:   file,
		row:    row,
		column: column,
	}
}

// GetFile returns the file
func (r *Result) GetFile() string {
	return r.file
}

// GetRow returns the row
func (r *Result) GetRow() int {
	return r.row
}

// GetColumn returns the column
func (r *Result) GetColumn() string {
	return r.column
}

// String returns the string value of the result
func (r *Result) String() string {
	return fmt.Sprintf("file: %s, row: %d, column: %s", r.GetFile(), r.GetRow(), r.GetColumn())
}
