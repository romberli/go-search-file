package excel

import (
	"github.com/romberli/go-search-file/pkg/message"
	"github.com/romberli/go-util/config"
)

const (
	ErrSearchExcel = 401001
)

func init() {
	initErrorMessage()
}

func initErrorMessage() {
	message.Messages[ErrSearchExcel] = config.NewErrMessage(message.DefaultMessageHeader, ErrSearchExcel, "search excel file failed")
}
