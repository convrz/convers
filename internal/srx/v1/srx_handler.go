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

package srx

import (
	"context"
	"github.com/convrz/convers/api/srx/v1"
)

type SRX struct {
	srx.UnimplementedSrxServiceServer
}

func (srx *SRX) RegisterService(ctx context.Context, info *srx.ServiceInfo) (*srx.RegisterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (srx *SRX) DiscoverService(ctx context.Context, request *srx.ServiceRequest) (*srx.ServiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (srx *SRX) Heartbeat(ctx context.Context, info *srx.ServiceInfo) (*srx.HeartbeatResponse, error) {
	//TODO implement me
	panic("implement me")
}
