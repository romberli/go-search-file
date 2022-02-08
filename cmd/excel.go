/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/romberli/go-search-file/config"
	"github.com/romberli/go-search-file/pkg/excel"
	"github.com/romberli/go-search-file/pkg/message"
	mesgexcel "github.com/romberli/go-search-file/pkg/message/excel"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// excelCmd represents the Excel command
var excelCmd = &cobra.Command{
	Use:   "excel",
	Short: "excel command",
	Long:  `search excel file`,
	Run: func(cmd *cobra.Command, args []string) {
		// init config
		err := initConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf("%+v", message.NewMessage(message.ErrInitConfig, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		// search Excel file
		searcher := excel.NewSearcher()
		results, err := searcher.Search(viper.GetString(config.PathKey), viper.GetString(config.KeywordKey))
		if err != nil {
			fmt.Println(fmt.Sprintf("%+v", message.NewMessage(mesgexcel.ErrSearchExcel, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}

		for _, result := range results {
			fmt.Println(result.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(excelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// excelCmd.PersistentFlags().String("foo", "", "A help for foo")
	// excel
	excelCmd.PersistentFlags().StringVar(&path, "path", constant.DefaultRandomString, fmt.Sprintf("specify the path(default: %s)", constant.CurrentDir))
	excelCmd.PersistentFlags().StringVar(&keyword, "keyword", constant.DefaultRandomString, fmt.Sprintf("specify the keyword(default: %s)", constant.DefaultRandomString))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// excelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
