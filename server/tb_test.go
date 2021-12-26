// Copyright 2018-2020 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"testing"
)

func TestTBSystemAccount(t *testing.T) {
	s, _ := runTrustedServer(t)
	defer s.Shutdown()

	acc, _ := createAccount(s)
	s.setSystemAccount(acc)

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.sys == nil || s.sys.account == nil {
		t.Fatalf("Expected sys.account to be non-nil")
	}
	if s.sys.client == nil {
		t.Fatalf("Expected sys.client to be non-nil")
	}

	s.sys.client.mu.Lock()
	defer s.sys.client.mu.Unlock()
	if s.sys.client.echo {
		t.Fatalf("Internal clients should always have echo false")
	}
}
