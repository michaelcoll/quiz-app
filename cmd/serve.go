/*
 * Copyright (c) 2022-2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/schoolbyhiit/quizz-app/internal/back"
	"github.com/schoolbyhiit/quizz-app/internal/back/domain/banner"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long: `
Starts the server`,
	Run: func(cmd *cobra.Command, args []string) {
		banner.Print(version, banner.Serve)

		module := back.New()
		module.GetPhotoController().Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
