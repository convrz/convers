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

package namespace

import (
	"context"
	"errors"
)

type Namespace struct {
	Name   string
	Labels map[string]string
}

func WithNamespace(ctx context.Context, ns string) context.Context {
	return context.WithValue(ctx, "namespace", ns)
}

func FromContext(ctx context.Context) (string, error) {
	ns := ctx.Value("namespace")
	if ns == nil {
		return "", errors.New("namespace not found in context")
	}
	return ns.(string), nil
}
