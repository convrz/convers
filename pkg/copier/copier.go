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

package copier

import (
	"encoding/json"

	"github.com/convrz/convers/pkg/pb/proto"
	google "google.golang.org/protobuf/proto"
)

func CopyMsg(src, dst google.Message) error {
	bytes, err := proto.Marshal(src)
	if err != nil {
		return err
	}

	return proto.Unmarshal(bytes, dst)
}

func Copy(src, dst interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, dst)
}
