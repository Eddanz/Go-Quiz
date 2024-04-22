// CLI

package main

import (
	"bufio"         // Paket för buffrad inläsning från standardinmatningen.
	"bytes"         // Paket för hantering av byte-slices.
	"encoding/json" // Paket för JSON-serialisering och -deserialisering.
	"fmt"           // Grundläggande in- och utmatningsfunktioner.
	"io"            // Gränssnitt för in- och utmatning.
	"net/http"      // Paket för HTTP-klient och serverfunktioner.
	"os"            // Funktioner för operativsystem-interaktion.
	"strings"       // Funktioner för manipulation av strängar.

	"github.com/spf13/cobra" // Paket för att skapa kraftfulla CLI-program.
)

var questions []Question // Global slice för att lagra frågor.

type Question struct {
	ID      int
	Text    string
	Choices []string
}

func main() {
	// Startmeddelande
	fmt.Println("Welcome to the Quiz CLI. Type 'start' to load questions, or 'exit' to quit.")
	reader := bufio.NewReader(os.Stdin) // Skapar en ny buffrad läsare.

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n') // Läser inmatning från användaren.
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input) // Rensar whitespace från inmatningen.

		switch input {
		case "start":
			startCmd.Run(nil, nil) // Utför startkommandot
		case "exit":
			fmt.Println("Exiting quiz CLI...")
			return
		default:
			fmt.Println("Unknown command. Available commands are 'start' and 'exit'.")
		}
	}
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the quiz and submit answers",
	Run: func(cmd *cobra.Command, args []string) {
		// Hämta frågor från servern
		resp, err := http.Get("http://localhost:8080/questions")
		if err != nil {
			fmt.Println("Error fetching questions:", err)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		err = json.Unmarshal(body, &questions)
		if err != nil {
			fmt.Println("Error decoding questions:", err)
			return
		}

		// Kontrollera om frågor laddades framgångsrikt
		if len(questions) == 0 {
			fmt.Println("No questions available, please check the server.")
			return
		}

		// Fortsätt att samla in svar
		answers := make([]int, len(questions))
		for i := 0; i < len(questions); i++ {
			q := questions[i]
			fmt.Printf("\nQuestion %d: %s\n", i+1, q.Text)
			for idx, choice := range q.Choices {
				fmt.Printf("  %d. %s\n", idx+1, choice)
			}

			fmt.Print("Enter your answer (1-4): ")
			var response int
			_, err := fmt.Scan(&response) // Läs inmatning från användaren.
			if err != nil || response < 1 || response > len(q.Choices) {
				fmt.Println("Invalid input, please enter a number within the provided range.")
				i-- // Minska i för att upprepa frågan.
				continue
			}

			answers[i] = response - 1 // Lagrar det indexerade svaret.
		}

		// Skicka in de insamlade svaren.
		submitAnswers(answers)
	},
}

func submitAnswers(answers []int) {
	// Serialisera svaren till JSON-format.
	payload, err := json.Marshal(answers)
	if err != nil {
		fmt.Println("Error encoding answers:", err) // Skriv ut felmeddelande om serialisering misslyckas.
		return
	}

	// Skicka de serialiserade svaren till servern via en POST-förfrågan.
	resp, err := http.Post("http://localhost:8080/answers", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error submitting answers:", err) // Skriv ut felmeddelande om POST-förfrågan misslyckas.
		return
	}
	defer resp.Body.Close() // Se till att svarkroppen stängs oavsett vad som händer.

	// Läs svaret från servern.
	resultBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading result:", err) // Skriv ut felmeddelande om det inte går att läsa svaret.
		return
	}

	// Avkoda JSON-svaret från servern.
	var result map[string]interface{}
	err = json.Unmarshal(resultBody, &result)
	if err != nil {
		fmt.Println("Error decoding result:", err)
		return
	}

	// Skriv ut användarens poäng och percentil från serverns svar.
	fmt.Printf("Your score: %v, You were better than %.2f%% of all quizzers\n", result["score"], result["percentile"])
}
