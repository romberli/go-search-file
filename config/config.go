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

package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/pingcap/errors"
	"github.com/romberli/go-multierror"
	"github.com/romberli/go-search-file/pkg/message"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var (
	ValidLogLevels  = []string{"debug", "info", "warn", "warning", "error", "fatal"}
	ValidLogFormats = []string{"text", "json"}
)

// SetDefaultConfig set default configuration, it is the lowest priority
func SetDefaultConfig(baseDir string) {
	// daemon
	viper.SetDefault(DaemonKey, DefaultDaemon)
	// log
	defaultLogFile := filepath.Join(baseDir, DefaultLogDir, log.DefaultLogFileName)
	viper.SetDefault(LogFileNameKey, defaultLogFile)
	viper.SetDefault(LogLevelKey, log.DefaultLogLevel)
	viper.SetDefault(LogFormatKey, log.DefaultLogFormat)
	viper.SetDefault(LogMaxSizeKey, log.DefaultLogMaxSize)
	viper.SetDefault(LogMaxDaysKey, log.DefaultLogMaxDays)
	viper.SetDefault(LogMaxBackupsKey, log.DefaultLogMaxBackups)
	// excel
	viper.SetDefault(PathKey, constant.CurrentDir)
	viper.SetDefault(KeywordKey, constant.DefaultRandomString)
}

// ValidateConfig validates if the configuration is valid
func ValidateConfig() (err error) {
	merr := &multierror.Error{}

	// validate daemon section
	err = ValidateDaemon()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	// validate log section
	err = ValidateLog()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	// validate excel section
	err = ValidateExcel()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return errors.Trace(merr.ErrorOrNil())
}

// ValidateDaemon validates if daemon section is valid
func ValidateDaemon() error {
	_, err := cast.ToBoolE(viper.Get(DaemonKey))

	return errors.Trace(err)
}

// ValidateLog validates if log section is valid.
func ValidateLog() error {
	var valid bool

	merr := &multierror.Error{}

	// validate log.fileName
	logFileName, err := cast.ToStringE(viper.Get(LogFileNameKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	logFileName = strings.TrimSpace(logFileName)
	if logFileName == constant.EmptyString {
		merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyLogFileName))
	}
	isAbs := filepath.IsAbs(logFileName)
	if !isAbs {
		logFileName, err = filepath.Abs(logFileName)
		if err != nil {
			merr = multierror.Append(merr, errors.Trace(err))
		}
	}
	valid, _ = govalidator.IsFilePath(logFileName)
	if !valid {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogFileName, logFileName))
	}

	// validate log.level
	logLevel, err := cast.ToStringE(viper.Get(LogLevelKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	valid, err = common.ElementInSlice(ValidLogLevels, logLevel)
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if !valid {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogLevel, logLevel))
	}

	// validate log.format
	logFormat, err := cast.ToStringE(viper.Get(LogFormatKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	valid, err = common.ElementInSlice(ValidLogFormats, logFormat)
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if !valid {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogFormat, logFormat))
	}

	// validate log.maxSize
	logMaxSize, err := cast.ToIntE(viper.Get(LogMaxSizeKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if logMaxSize < MinLogMaxSize || logMaxSize > MaxLogMaxSize {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogMaxSize, MinLogMaxSize, MaxLogMaxSize, logMaxSize))
	}

	// validate log.maxDays
	logMaxDays, err := cast.ToIntE(viper.Get(LogMaxDaysKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if logMaxDays < MinLogMaxDays || logMaxDays > MaxLogMaxDays {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogMaxDays, MinLogMaxDays, MaxLogMaxDays, logMaxDays))
	}

	// validate log.maxBackups
	logMaxBackups, err := cast.ToIntE(viper.Get(LogMaxBackupsKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if logMaxBackups < MinLogMaxDays || logMaxBackups > MaxLogMaxDays {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogMaxBackups, MinLogMaxBackups, MaxLogMaxBackups, logMaxBackups))
	}

	return merr.ErrorOrNil()
}

// ValidateServer validates if server section is valid
func ValidateExcel() error {
	// validate path
	path, err := cast.ToStringE(viper.Get(PathKey))
	if err != nil {
		return errors.Trace(err)
	}
	if path == constant.EmptyString {
		return message.NewMessage(message.ErrEmptyPath)
	}
	isAbs := filepath.IsAbs(path)
	if !isAbs {
		path, err = filepath.Abs(path)
		if err != nil {
			return errors.Trace(err)
		}
	}
	ok, _ := govalidator.IsFilePath(path)
	if !ok {
		return errors.Trace(err)
	}

	keyword, err := cast.ToStringE(viper.Get(KeywordKey))
	if err != nil {
		return errors.Trace(err)
	}
	if keyword == constant.EmptyString {
		return message.NewMessage(message.ErrEmptyKeyword)
	}

	return nil
}

// TrimSpaceOfArg trims spaces of given argument
func TrimSpaceOfArg(arg string) string {
	args := strings.SplitN(arg, constant.EqualString, 2)

	switch len(args) {
	case 1:
		return strings.TrimSpace(args[0])
	case 2:
		argName := strings.TrimSpace(args[0])
		argValue := strings.TrimSpace(args[1])
		return fmt.Sprintf("%s=%s", argName, argValue)
	default:
		return arg
	}
}
