// Copyright 2015 - António Meireles  <antonio.meireles@reformi.st>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
)

var (
	utilsCmd = &cobra.Command{
		Use:    "utils",
		Short:  "Some developer focused utilies",
		Run:    utilsCommand,
		Hidden: true,
	}
	manCmd = &cobra.Command{
		Use:     "mkMan",
		Short:   "Generates man pages",
		PreRunE: defaultPreRunE,
		Run:     manCommand,
	}
	mkdownCmd = &cobra.Command{
		Use:     "mkMkdown",
		Short:   "Generates Markdown documentation",
		PreRunE: defaultPreRunE,
		Run:     mkdownCommand,
	}
)

func utilsCommand(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func manCommand(cmd *cobra.Command, args []string) {
	header := &doc.GenManHeader{
		Title: "corectl", Source: " ",
	}
	manDir := "documentation/man/"
	// check if target dir exists.if not create it
	if _, err := os.Stat(manDir); os.IsNotExist(err) {
		if err = os.MkdirAll(manDir, 0644); err != nil {
			fmt.Printf("failed to create %s. "+
				"No manual pages generated\n", manDir)
		}
	}
	doc.GenManTree(RootCmd, header, manDir)
}

func mkdownCommand(cmd *cobra.Command, args []string) {
	markdownDir := "documentation/markdown/"
	// check if target dir exists.if not create it
	if _, err := os.Stat(markdownDir); os.IsNotExist(err) {
		if err = os.MkdirAll(markdownDir, 0644); err != nil {
			fmt.Printf("failed to create %s. "+
				"No markdown documentation generated\n", markdownDir)
		}
	}
	doc.GenMarkdownTree(RootCmd,
		engine.pwd+"/documentation/markdown/")
}

func init() {
	utilsCmd.AddCommand(manCmd)
	utilsCmd.AddCommand(mkdownCmd)
	RootCmd.AddCommand(utilsCmd)
}
