/*
Copyright 2024 The Convērs Authors.

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

package init

import (
	"github.com/convrz/convers/core/containers"
	"github.com/convrz/convers/core/services"
	"github.com/convrz/convers/x/greeter/v1/internal/biz"
	"github.com/convrz/convers/x/greeter/v1/internal/controllers"
	"github.com/convrz/convers/x/greeter/v1/internal/repos"
)

var (
	// delivery layer
	_ = containers.Inject(services.NewDefault)
	_ = containers.Inject(controllers.New)

	// domain layer
	_ = containers.Inject(biz.New)

	// repo layer
	_ = containers.InjectLifeCycle(repos.New, repos.OnStart, repos.OnStop)
)
