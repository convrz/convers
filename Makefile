# Copyright 2024 The Convērs Authors.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

test:
	@echo "Running tests..."

lint:
	@golangci-lint run

build.greeter:
	docker buildx build -f ./x/greeter/v1/Dockerfile -t cvz.x.greeter.v1:latest .

build.gateway:
	docker buildx build -f ./cmd/gateway/Dockerfile -t cvz.gateway:latest .

build.srx:
	docker buildx build -f ./cmd/srx/Dockerfile -t cvz.srx:latest .
