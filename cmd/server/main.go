// SERVER

package main

import (
	"Go-Quiz/pkg/model" // Importera paketet som innehåller modellen för frågor.
	"encoding/json"     // Paket för att hantera JSON, används för att serialisera och deserialisera data.
	"net/http"          // Paket för att hantera HTTP-protokollet, används för att skapa webbservrar.
	"sort"              // Paket för att sortera slice-datastrukturer.
	"sync"              // Paket för att hantera synkronisering mellan go-rutiner.
)

// Question definierar strukturen för en fråga.
type Question struct {
	ID       int      `json:"id"`        // Unikt ID för frågan.
	Text     string   `json:"text"`      // Texten för frågan.
	Choices  []string `json:"choices"`   // Lista över valmöjligheter.
	AnswerID int      `json:"answer_id"` // ID för det korrekta svaret.
}

// Global variabel för att lagra frågor och poäng.
var quiz = &model.Quiz{
	Questions: model.SampleQuestions, // Antas initiera korrekt med exempelfrågor.
	Scores:    []int{},               // Slice för att spåra poängen för alla genomförda quiz.
	Mutex:     sync.Mutex{},          // Mutex för att säkerställa trådsäkerhet.
}

func main() {
	http.HandleFunc("/questions", handleQuestions) // Hantera GET-förfrågningar till /questions.
	http.HandleFunc("/answers", handleAnswers)     // Hantera POST-förfrågningar till /answers.
	http.ListenAndServe(":8080", nil)              // Starta servern på port 8080.
}

// handleQuestions skickar alla frågor som JSON via HTTP.
func handleQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Sätt Content-Type header till application/json.
	if err := json.NewEncoder(w).Encode(model.SampleQuestions); err != nil {
		http.Error(w, "Failed to encode questions", http.StatusInternalServerError) // Skicka felmeddelande om encoding misslyckas.
	}
}

// UserAnswers håller en lista över användarens svar.
type UserAnswers struct {
	Answers []int `json:"answers"` // Användarinskickade svars-IDn.
}

// handleAnswers behandlar användarsvar och beräknar poäng.
func handleAnswers(w http.ResponseWriter, r *http.Request) {
	var submittedAnswers []int
	if err := json.NewDecoder(r.Body).Decode(&submittedAnswers); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest) // Skicka felmeddelande om kroppen är ogiltig.
		return
	}

	score := 0
	for i, answer := range submittedAnswers {
		if i < len(model.SampleQuestions) && model.SampleQuestions[i].AnswerID == answer {
			score++ // Öka poängen om svaret är rätt.
		}
	}

	// Lås mutexen innan du åtkommer den globala poänglistan.
	quiz.Mutex.Lock()
	// Lägg till poängen i den globala listan och beräkna percentilen inom det låsta området.
	quiz.Scores = append(quiz.Scores, score)
	percentile := computePercentile(quiz.Scores, score)
	quiz.Mutex.Unlock() // Lås upp mutexen efter att ha modifierat den globala listan.

	response := map[string]interface{}{
		"score":      score,
		"percentile": percentile,
	}
	json.NewEncoder(w).Encode(response) // Kodar och skickar svaret som JSON.
}

// computePercentile beräknar procentandelen av poäng som är mindre än eller lika med den aktuella poängen.
func computePercentile(scores []int, userScore int) float64 {
	sort.Ints(scores) // Sorterar poängen.
	var rank int
	for _, score := range scores {
		if score <= userScore {
			rank++ // Räkna hur många poäng som är mindre än eller lika med användarens poäng.
		}
	}
	return float64(rank) / float64(len(scores)) * 100 // Beräkna percentilen.
}
