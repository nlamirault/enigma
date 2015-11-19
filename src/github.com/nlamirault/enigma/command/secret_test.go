// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

// import (
// 	"testing"

// 	"github.com/mitchellh/cli"
// )

// func TestSecretWithoutAction(t *testing.T) {
// 	ui := new(cli.MockUi)
// 	c := &SecretCommand{
// 		UI: ui,
// 	}

// 	args := []string{
// 		"--bucket", "foo",
// 		"--region", "region1",
// 	}

// 	if code := c.Run(args); code != 1 {
// 		t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
// 	}
// }

// func TestSecretWithoutBucket(t *testing.T) {
// 	ui := new(cli.MockUi)
// 	c := &SecretCommand{
// 		UI: ui,
// 	}

// 	args := []string{
// 		"--bucket", "foo",
// 		"--region", "region1",
// 		"--action", "get-text",
// 	}

// 	if code := c.Run(args); code != 1 {
// 		t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
// 	}
// }

// func TestSecretWithoutRegion(t *testing.T) {
// 	ui := new(cli.MockUi)
// 	c := &SecretCommand{
// 		UI: ui,
// 	}

// 	args := []string{
// 		"--region", "region1",
// 		"--action", "get-text",
// 	}

// 	if code := c.Run(args); code != 1 {
// 		t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
// 	}
// }

// func TestSecretWithInvalidAction(t *testing.T) {
// 	ui := new(cli.MockUi)
// 	c := &SecretCommand{
// 		UI: ui,
// 	}

// 	args := []string{
// 		"--bucket", "foo",
// 		"--region", "region1",
// 		"--action", "list",
// 	}

// 	if code := c.Run(args); code != 1 {
// 		t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
// 	}
// }
