/*
 * Copyright 2024 The Convērs Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package cvzlog provides a system logger for the service.
package cvzlog

import (
	"fmt"
	"os"
	"runtime"

	"github.com/convrz/convers/pkg/utils"
	"go.uber.org/zap"
)

func New() *zap.Logger {
	if os.Getenv("MODE") == "dev" {
		logger, _ := zap.NewDevelopment()
		return logger
	}

	logger, _ := zap.NewProduction()
	return logger
}

func Info(message string) {
	logger := New()
	defer utils.CallBack(logger.Sync)

	funcName := ""
	sugar := logger.Sugar()

	pc, file, line, ok := runtime.Caller(1)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	sugar.Infow(message,
		"function", funcName,
		"file", fmt.Sprintf("%s:%d", file, line),
	)
}

func Warn(message string) {
	logger := New()
	defer utils.CallBack(logger.Sync)

	funcName := ""
	sugar := logger.Sugar()

	pc, file, line, ok := runtime.Caller(1)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	sugar.Warnw(message,
		"function", funcName,
		"file", fmt.Sprintf("%s:%d", file, line),
	)
}

func Debug(message string) {
	logger := New()
	defer utils.CallBack(logger.Sync)

	funcName := ""
	sugar := logger.Sugar()

	pc, file, line, ok := runtime.Caller(1)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	sugar.Debugw(message,
		"function", funcName,
		"file", fmt.Sprintf("%s:%d", file, line),
	)
}

func Error(message string) {
	logger := New()
	defer utils.CallBack(logger.Sync)

	funcName := ""
	sugar := logger.Sugar()

	pc, file, line, ok := runtime.Caller(1)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	sugar.Errorw(message,
		"function", funcName,
		"file", fmt.Sprintf("%s:%d", file, line),
	)
}

func Fatal(message string) {
	logger := New()
	defer utils.CallBack(logger.Sync)

	funcName := ""
	sugar := logger.Sugar()

	pc, file, line, ok := runtime.Caller(1)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	sugar.Fatalw(message,
		"function", funcName,
		"file", fmt.Sprintf("%s:%d", file, line),
	)
}
