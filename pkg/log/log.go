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

// Package logger provides the logger for the service.
package log

import (
	"github.com/convrz/convers/pkg/utils"
	"go.uber.org/zap"
)

func New() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

func Info(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.Sugar()
	sugar.Info(message...)
}

func Infof(template string, message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.Sugar()
	sugar.Infof(template, message...)
}

func Debug(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.Sugar()
	sugar.Debug(message...)
}

func Error(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.Sugar()
	sugar.Error(message...)
}

func Fatal(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.Sugar()
	sugar.Fatal(message...)
}

func removeNil(input []interface{}) []interface{} {
	var result []interface{}
	for _, item := range input {
		if item != nil {
			result = append(result, item)
		}
	}
	return result
}
