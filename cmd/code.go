/*
Copyright © 2022 UZO

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
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Lucifergene/uzo/util"
	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:                   "code <zip_file_name>",
	Short:                 "Opens the project in your editor",
	Long:                  `Unzips the Project file and opens it in VS Code.`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Example:               `uzo code <project-name>.zip`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		var err error
		var argument string

		argument = args[0]

		FileExists, err := util.FileExists(argument)
		if err != nil {
			fmt.Println(err.Error())
		}

		if FileExists {
			fileName, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("File %s not found\n", argument)
			return
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		util.Unzip(fileName, wd)

		os.Chdir(util.FilenameWithoutExtension(fileName))

		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		commandCode := exec.Command("code", wd)
		err = commandCode.Run()
		if err != nil {
			fmt.Println("VS Code executable not found in %PATH%")
		}
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
