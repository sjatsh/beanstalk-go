// Copyright 2020 The Gostalkd Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package network

import (
	"fmt"

	"github.com/edoger/zkits-runner"
)

func New(addr string, processor Processor) (Server, error) {
	return new(tcpServer), nil
}

type Server interface {
	runner.Task
	fmt.Stringer
}

type tcpServer struct {
}

func (s *tcpServer) Execute() error {
	return nil
}

func (s *tcpServer) Shutdown() error {
	return nil
}

func (s *tcpServer) String() string {
	return ""
}
