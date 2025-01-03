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

package init

import (
	"context"
	"fmt"
	"github.com/convrz/convers/core/cvzhook"
)

func init() {
	cvzhook.UseBefore(func(ctx context.Context) error {
		fmt.Println("before start")
		return nil
	})

	cvzhook.UseAfter(func(ctx context.Context) error {
		fmt.Println("after stop")
		return nil
	})
}
