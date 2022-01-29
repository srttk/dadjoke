/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random dadjoke",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {

	data := getJokeData("https://icanhazdadjoke.com")
	joke := Joke{}
	if err := json.Unmarshal(data, &joke); err != nil {
		log.Printf("Could not unmarshal data %v", err)
	}
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(2).
		PaddingBottom(2).
		PaddingRight(2).
		Width(50)

	fmt.Println(style.Render(joke.Joke))

}

func getJokeData(baseAPI string) []byte {

	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		fmt.Println("Could not request")
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "My Dadjoke app")

	respose, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("Couldnot make a request %v", err)
	}

	responseBytes, err := ioutil.ReadAll(respose.Body)

	if err != nil {
		log.Printf("Could not read response body %v", err)
	}

	return responseBytes

}
