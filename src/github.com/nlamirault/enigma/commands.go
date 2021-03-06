// Copyright (C) 2015, 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/mitchellh/cli"

	"github.com/nlamirault/enigma/command"
)

// Commands is the mapping of all the available Terraform commands.
var (
	Commands map[string]cli.CommandFactory
	UI       cli.Ui
)

// type Meta struct {
// 	UI cli.Ui
// }

func init() {
	UI = &cli.ColoredUi{
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			Reader:      os.Stdin,
			ErrorWriter: os.Stderr,
		},
		OutputColor: cli.UiColorNone,
		InfoColor:   cli.UiColorGreen,
		ErrorColor:  cli.UiColorRed,
	}

	Commands = map[string]cli.CommandFactory{
		"bucket": func() (cli.Command, error) {
			return &command.BucketCommand{
				UI: UI,
			}, nil
		},
		"secret": func() (cli.Command, error) {
			return &command.SecretCommand{
				UI: UI,
			}, nil
		},
	}
}
