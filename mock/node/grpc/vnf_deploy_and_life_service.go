// Copyright 2019 Smart-Edge.com, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/smartedgemec/controller-ce/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type vnfDeployLifeService struct {
	vnfs []*pb.VNF
}

func (s *vnfDeployLifeService) Deploy(
	ctx context.Context,
	vnf *pb.VNF,
) (*empty.Empty, error) {
	vnf.Status = pb.LifecycleStatus_READY
	s.vnfs = append(s.vnfs, vnf)

	return &empty.Empty{}, nil
}

func (s *vnfDeployLifeService) GetStatus(
	ctx context.Context,
	id *pb.VNFID,
) (*pb.LifecycleStatus, error) {
	vnf := s.find(id.Id)

	if vnf != nil {
		return &pb.LifecycleStatus{
			Status: vnf.Status,
		}, nil
	}

	return nil, status.Errorf(codes.NotFound, "VNF %s not found", id.Id)
}

func (s *vnfDeployLifeService) Redeploy(
	ctx context.Context,
	vnf *pb.VNF,
) (*empty.Empty, error) {
	i := s.findIndex(vnf.Id)

	if i < len(s.vnfs) {
		vnf.Status = s.vnfs[i].Status
		s.vnfs[i] = vnf
		return &empty.Empty{}, nil
	}

	return nil, status.Errorf(
		codes.NotFound, "VNF %s not found", vnf.Id)
}

func (s *vnfDeployLifeService) Undeploy(
	ctx context.Context,
	id *pb.VNFID,
) (*empty.Empty, error) {
	i := s.findIndex(id.Id)

	if i < len(s.vnfs) {
		s.delete(i)
		return &empty.Empty{}, nil
	}

	return nil, status.Errorf(codes.NotFound, "VNF %s not found", id.Id)
}

func (s *vnfDeployLifeService) Start(
	ctx context.Context,
	cmd *pb.LifecycleCommand,
) (*empty.Empty, error) {
	vnf := s.find(cmd.Id)

	if vnf != nil {
		switch vnf.Status {
		case pb.LifecycleStatus_READY:
		case pb.LifecycleStatus_STOPPED:
		default:
			return nil, status.Errorf(
				codes.FailedPrecondition, "VNF %s not stopped or ready",
				cmd.Id)
		}

		vnf.Status = pb.LifecycleStatus_RUNNING
		return &empty.Empty{}, nil
	}

	return nil, status.Errorf(
		codes.NotFound, "VNF %s not found", cmd.Id)
}

func (s *vnfDeployLifeService) Stop(
	ctx context.Context,
	cmd *pb.LifecycleCommand,
) (*empty.Empty, error) {
	vnf := s.find(cmd.Id)

	if vnf != nil {
		if vnf.Status != pb.LifecycleStatus_RUNNING {
			return nil, status.Errorf(
				codes.FailedPrecondition, "VNF %s not running", cmd.Id)
		}

		vnf.Status = pb.LifecycleStatus_STOPPED
		return &empty.Empty{}, nil
	}

	return nil, status.Errorf(
		codes.NotFound, "VNF %s not found", cmd.Id)
}

func (s *vnfDeployLifeService) Restart(
	ctx context.Context,
	cmd *pb.LifecycleCommand,
) (*empty.Empty, error) {
	vnf := s.find(cmd.Id)

	if vnf != nil {
		if vnf.Status != pb.LifecycleStatus_RUNNING {
			return nil, status.Errorf(
				codes.FailedPrecondition, "VNF %s not running", cmd.Id)
		}

		return &empty.Empty{}, nil
	}

	return nil, status.Errorf(
		codes.NotFound, "VNF %s not found", cmd.Id)
}

func (s *vnfDeployLifeService) find(id string) *pb.VNF {
	for _, vnf := range s.vnfs {
		if vnf.Id == id {
			return vnf
		}
	}

	return nil
}

func (s *vnfDeployLifeService) findIndex(id string) int {
	for i, vnf := range s.vnfs {
		if vnf.Id == id {
			return i
		}
	}

	return len(s.vnfs)
}

func (s *vnfDeployLifeService) delete(i int) {
	copy(s.vnfs[i:], s.vnfs[i+1:])
	s.vnfs[len(s.vnfs)-1] = nil
	s.vnfs = s.vnfs[:len(s.vnfs)-1]
}