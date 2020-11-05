/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"math/rand"
	"os"
	"time"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/briandowns/spinner"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/viper"
)

var (
	cfgFile    string
	status     string
	useCowsay  bool
	imOffended bool
)

var rootCmd = &cobra.Command{
	Use:   "okify",
	Short: "Feelings are more important than production.",
	Long: `Avoid getting non-positive feedback from anything.

Example:
	ls nonexistent-file | okify
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Spinner is pretty
		displaySpinner(3)

		// Table is also pretty
		statement := randomCompliment()
		if imOffended {
			statement = randomApology()
		}

		if useCowsay {
			output, _ := cowsay.Say(
				cowsay.Phrase(statement),
				cowsay.Type("default"),
				cowsay.BallonWidth(40),
			)
			fmt.Println(output)
			os.Exit(0)
		}

		printTable(statement)
		os.Exit(0)
	},
}

func printTable(s string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{s})
	t.Render()
}

func displaySpinner(seconds time.Duration) {
	s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
	s.Suffix = " Calculating non-offensive response"
	s.Start()
	time.Sleep(seconds * time.Second)
	s.Stop()
}

func randomCompliment() string {
	rand.Seed(time.Now().Unix())
	compliments := []string{
		"Everything is fine!\n",
		"Ignore the haters!\n",
		"Looks good to me!\n",
		"You are doing such a good job!\n",
		"No one is as good as you are!\n",
		"How are you still single!\n",
		"You are so handsome!\n",
	}
	complimentIndex := rand.Intn(len(compliments))
	return compliments[complimentIndex]
}

func randomApology() string {
	rand.Seed(time.Now().Unix())
	apologies := []string{
		"I'm so sorry!\n",
		"It's all my fault!\n",
		"How could I have been so stupid?!\n",
	}
	apologyIndex := rand.Intn(len(apologies))
	return apologies[apologyIndex]
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("You are doing great")
		os.Exit(0)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&status, "status", "", "how good the situation is")
	rootCmd.PersistentFlags().BoolVar(&imOffended, "im-offended", false, "you are owed an apology")
	rootCmd.PersistentFlags().BoolVar(&useCowsay, "cowsay", false, "use cowsay for printing messages")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println("You are a great person!")
			os.Exit(0)
		}

		// Search config in home directory with name ".okify" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".okify")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
