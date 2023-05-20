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
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"github.com/michaelcoll/quiz-app/internal/back"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long: `
Starts the server`,
	Run: serve,
}

func serve(_ *cobra.Command, _ []string) {
	Print(version, Serve)

	module := back.New()

	err := module.GetService().Sync(context.Background())
	if err != nil {
		fmt.Printf("%s Can't sync quizzes (%v)\n", color.RedString("✗"), err)
		os.Exit(-1)
	}
	module.GetPhotoController().Serve()
}

func init() {
	serveCmd.Flags().StringP("repository-url", "r", "", "The url of the repository containing the quizzes")
	serveCmd.Flags().StringP("token", "t", "", "The P.A.T. used to access the repository")
	serveCmd.Flags().String("auth0-audience", "", "The Auth0 audience used in the clientId")
	serveCmd.Flags().String("restrict-email-domain", "", "New users will have to be in this domain to be created")

	_ = viper.BindPFlag("repository-url", serveCmd.Flags().Lookup("repository-url"))
	_ = viper.BindPFlag("token", serveCmd.Flags().Lookup("token"))
	_ = viper.BindPFlag("auth0-audience", serveCmd.Flags().Lookup("auth0-audience"))

	viper.SetDefault("repository-url", "https://github.com/michaelcoll/quiz-app.git")

	rootCmd.AddCommand(serveCmd)
}
