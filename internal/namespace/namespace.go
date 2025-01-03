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
	"errors"
	"strings"
	"sync"
)

type INamespace interface {
	Add(fullPath string) error
	Remove(fullPath string) error
	List(ns string) ([]string, error)
}

// New initializes a new Namespace
func New() *Namespace {
	return &Namespace{
		namespaces: make(map[string]map[string]bool),
	}
}

// Namespace manages nested namespaces using dots .
type Namespace struct {
	mu         sync.RWMutex
	namespaces map[string]map[string]bool
}

// Add adds a name to a multi-layered namespace using dots .
func (n *Namespace) Add(fullPath string) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	layers := strings.Split(fullPath, ".")
	if len(layers) < 2 {
		return errors.New("invalid namespace format")
	}

	ns := strings.Join(layers[:len(layers)-1], ".")
	name := layers[len(layers)-1]

	if n.namespaces[ns] == nil {
		n.namespaces[ns] = make(map[string]bool)
	}

	if n.namespaces[ns][name] {
		return errors.New("name already exists in namespace")
	}

	n.namespaces[ns][name] = true
	return nil
}

// Remove removes a name from the namespace
func (n *Namespace) Remove(fullPath string) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	layers := strings.Split(fullPath, ".")
	if len(layers) < 2 {
		return errors.New("invalid namespace format")
	}

	ns := strings.Join(layers[:len(layers)-1], ".")
	name := layers[len(layers)-1]

	if _, ok := n.namespaces[ns]; !ok {
		return errors.New("namespace does not exist")
	}

	if !n.namespaces[ns][name] {
		return errors.New("name not found in namespace")
	}

	delete(n.namespaces[ns], name)

	if len(n.namespaces[ns]) == 0 {
		delete(n.namespaces, ns)
	}

	return nil
}

// List lists all names in a multi-layered namespace
func (n *Namespace) List(ns string) ([]string, error) {
	n.mu.RLock()
	defer n.mu.RUnlock()

	if _, ok := n.namespaces[ns]; !ok {
		return nil, errors.New("namespace does not exist")
	}

	names := make([]string, 0, len(n.namespaces[ns]))
	for name := range n.namespaces[ns] {
		names = append(names, name)
	}

	return names, nil
}
