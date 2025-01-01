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

package namespace

import (
	"strings"
	"testing"
)

func TestNamespaceManager(t *testing.T) {
	manager := New()

	// Add names to a multi-layered namespace using dots
	err := manager.Add("prod.backend.service-a")
	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}

	err = manager.Add("prod.backend.service-b")
	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}

	err = manager.Add("prod.frontend.service-c")
	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}

	// Try adding duplicate name
	err = manager.Add("prod.backend.service-a")
	if err == nil {
		t.Fatalf("Expected error when adding duplicate name")
	}

	// List services in the backend
	names, err := manager.List("prod.backend")
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}

	if len(names) != 2 {
		t.Errorf("Expected 2 services, got %d", len(names))
	}

	// Remove and check again
	err = manager.Remove("prod.backend.service-a")
	if err != nil {
		t.Fatalf("Remove failed: %v", err)
	}

	names, _ = manager.List("prod.backend")
	if len(names) != 1 {
		t.Errorf("Expected 1 service after remove, got %d", len(names))
	}

	t.Log(strings.Join(names, ""))
}
