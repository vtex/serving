/*
Copyright 2018 The Knative Authors

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

package gc

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDeepCopy(t *testing.T) {
	src := Config{
		RetainSinceCreateTime:     5 * time.Minute,
		RetainSinceLastActiveTime: 5 * time.Minute,
		MinNonActiveRevisions:     1,
		MaxNonActiveRevisions:     5,
	}

	if diff := cmp.Diff(src, *src.DeepCopy()); diff != "" {
		t.Error("Unexpected DeepCopy (-want, +got):", diff)
	}
}

func TestDeepCopyInto(t *testing.T) {
	var dest Config
	src := Config{
		RetainSinceCreateTime:     5 * time.Minute,
		RetainSinceLastActiveTime: 5 * time.Minute,
		MinNonActiveRevisions:     1,
		MaxNonActiveRevisions:     5,
	}

	src.DeepCopyInto(&dest)
	if diff := cmp.Diff(src, dest); diff != "" {
		t.Error("Unexpected DeepCopyInto (-want, +got):", diff)
	}
}
