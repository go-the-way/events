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
	BIHandler[E comparable, T, A any] struct{ *Handler[E, T] }
	BIEventHandler[T, A any]          func(sender T, args A)
)

func NewBIHandler[E comparable, T, A any]() *BIHandler[E, T, A] {
	return &BIHandler[E, T, A]{Handler: NewHandler[E, T]()}
}

func (h *BIHandler[E, T, A]) Bind(eventHandler BIEventHandler[T, A]) {
	var event E
	if value, loaded := h.m.Load(event); loaded {
		if handlers, ok := value.(*[]BIEventHandler[T, A]); ok {
			*handlers = append(*handlers, eventHandler)
		}
	} else {
		h.m.Store(event, &[]BIEventHandler[T, A]{eventHandler})
	}
}

func (h *BIHandler[E, T, A]) Fire(sender T, args A) {
	var event E
	if value, loaded := h.m.Load(event); loaded {
		if handlers, ok := value.(*[]BIEventHandler[T, A]); ok {
			wg := &sync.WaitGroup{}
			wg.Add(len(*handlers))
			for _, hr := range *handlers {
				go func(handler BIEventHandler[T, A]) { defer wg.Done(); handler(sender, args) }(hr)
			}
			wg.Wait()
		}
	}
}
