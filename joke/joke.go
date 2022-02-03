package joke

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/charmbracelet/lipgloss"
)

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func GetRandomJoke() {

	data := GetJokeData("https://icanhazdadjoke.com")
	joke := Joke{}
	if err := json.Unmarshal(data, &joke); err != nil {
		log.Printf("Could not unmarshal data %v", err)
	}

	printStyle(joke.Joke)

}

func GetJokeData(baseAPI string) []byte {

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

func printStyle(str string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#03fc7b")).
		PaddingTop(2).
		PaddingLeft(2).
		PaddingBottom(2).
		PaddingRight(2)

	fmt.Println(style.Render(str))

}
