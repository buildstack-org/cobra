// Copyright 2013-2023 The Cobra Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cobra

// Demonstration tests for the Avrea Test Intelligence fork. They are not part
// of the upstream project: they exist so the fork's runs show a failing test,
// a flaky test, a slow test and a skipped test.

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestAvreaDemoAlwaysFails(t *testing.T) {
	got, want := len("cobra"), 6
	if got != want {
		t.Fatalf("demo failure: command name length = %d, want %d", got, want)
	}
}

func TestAvreaDemoOneSubtestFails(t *testing.T) {
	for _, name := range []string{"alpha", "beta", "gamma"} {
		t.Run(name, func(t *testing.T) {
			if name == "beta" {
				t.Errorf("demo failure: flag %q was not registered", name)
			}
		})
	}
}

// Fails on every second workflow run, so its history alternates.
func TestAvreaDemoFlaky(t *testing.T) {
	run, err := strconv.Atoi(os.Getenv("GITHUB_RUN_NUMBER"))
	if err != nil {
		t.Skip("not running in a workflow")
	}
	if run%2 == 1 {
		t.Fatalf("demo failure: timed out waiting for completion on run %d", run)
	}
}

func TestAvreaDemoSlow(t *testing.T) {
	time.Sleep(4 * time.Second)
}

func TestAvreaDemoSkipped(t *testing.T) {
	t.Skip("demo skip: needs a shell that is not installed")
}
