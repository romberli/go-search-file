package excel

import (
	"fmt"

	"github.com/romberli/go-search-file/pkg/dependency"
)

var _ dependency.Result = (*Result)(nil)

type Result struct {
	file    string
	sheet   string
	row     int
	column  string
	keyword string
}

func NewResult(file, sheet string, row int, column string, keyword string) *Result {
	return &Result{
		file:    file,
		sheet:   sheet,
		row:     row,
		column:  column,
		keyword: keyword,
	}

}

// GetFile returns the file
func (r *Result) GetFile() string {
	return r.file
}

// GetSheet returns the sheet
func (r *Result) GetSheet() string {
	return r.sheet
}

// GetRow returns the row
func (r *Result) GetRow() int {
	return r.row
}

// GetColumn returns the column
func (r *Result) GetColumn() string {
	return r.column
}

// GetKeyword returns the keyword
func (r *Result) GetKeyword() string {
	return r.keyword
}

// String returns the string value of the result
func (r *Result) String() string {
	return fmt.Sprintf("文件名: %s, 工作表： %s 单元格: %s%d, 关键字: %s", r.GetFile(), r.GetSheet(), r.GetColumn(), r.GetRow(), r.GetKeyword())
}
