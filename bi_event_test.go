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

import "testing"

type (
	MyBIEvent  struct{}
	MyBISender struct{ ID uint }
	Extra      struct{}
)

func TestBIEvent(t *testing.T) {
	myHr := NewBIHandler[MyBIEvent, MyBISender, Extra]()
	myHr.Bind(func(sender MyBISender, extra Extra) { t.Log(sender.ID) })
	myHr.Bind(func(sender MyBISender, extra Extra) { t.Log(sender.ID) })
	myHr.Bind(func(sender MyBISender, extra Extra) { t.Log(sender.ID) })
	myHr.Bind(func(sender MyBISender, extra Extra) { t.Log(sender.ID) })
	myHr.Fire(MyBISender{1000}, Extra{})
	myHr.Fire(MyBISender{2000}, Extra{})
	myHr.Fire(MyBISender{3000}, Extra{})
}
