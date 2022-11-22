// Copyright 2022 events Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package events

import (
	"sync"
)

type (
	EventHandler[T any]          func(sender T)
	handler[E comparable, T any] struct{ m *sync.Map }
)

func Handler[E comparable, T any]() *handler[E, T] { return &handler[E, T]{&sync.Map{}} }

// Bind
//
// Bind the event handler on E
func (h *handler[E, T]) Bind(eventHandler EventHandler[T]) {
	var event E
	if value, loaded := h.m.Load(event); loaded {
		if handlers, ok := value.(*[]EventHandler[T]); ok {
			*handlers = append(*handlers, eventHandler)
		}
	} else {
		h.m.Store(event, &[]EventHandler[T]{eventHandler})
	}
}

// Fire
//
// Fire the event from sender
func (h *handler[E, T]) Fire(sender T) {
	var event E
	if value, loaded := h.m.Load(event); loaded {
		if handlers, ok := value.(*[]EventHandler[T]); ok {
			wg := &sync.WaitGroup{}
			wg.Add(len(*handlers))
			for _, hr := range *handlers {
				go func(handler EventHandler[T]) { defer wg.Done(); handler(sender) }(hr)
			}
			wg.Wait()
		}
	}
}
