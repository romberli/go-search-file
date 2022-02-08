package excel

import (
	"strings"

	"github.com/pingcap/errors"
	"github.com/romberli/go-search-file/pkg/dependency"
	"github.com/romberli/go-search-file/pkg/util"
	"github.com/romberli/go-util/constant"
	"github.com/xuri/excelize/v2"
)

type Searcher struct {
	results []dependency.Result
}

func NewSearcher() dependency.Searcher {
	return newSearcher()
}

func newSearcher() *Searcher {
	return &Searcher{}
}

// GetResults returns the results
func (s *Searcher) GetResults() []dependency.Result {
	return s.results
}

func (s *Searcher) Search(path, keyword string) ([]dependency.Result, error) {
	var results []dependency.Result

	keywordList := strings.Split(keyword, constant.VerticalBarString)
	files, err := util.FindExcelFile(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		f, err := excelize.OpenFile(file)
		if err != nil {
			return nil, errors.Trace(err)
		}

		for _, sheet := range f.GetSheetList() {
			rows, err := f.GetRows(sheet)
			if err != nil {
				return nil, errors.Trace(err)
			}
			for i, row := range rows {
				for j, cell := range row {
					for _, word := range keywordList {
						word = strings.TrimSpace(word)
						if strings.Contains(cell, word) {
							results = append(results, NewResult(file, i+1, util.ConvertIntToAlphabet(j), word))
						}
					}
				}
			}
		}
	}

	return results, nil
}
