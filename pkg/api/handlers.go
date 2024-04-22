// HANDLERS

package api

import (
	"Go-Quiz/pkg/model" // Importerar modellpaketet
	"encoding/json"     // Paket för att hantera JSON.
	"net/http"          // Paket för att hantera HTTP-protokollet.
)

// HandleQuestions skickar en lista med frågor till klienten.
func HandleQuestions(w http.ResponseWriter, r *http.Request) {
	questions := model.SampleQuestions // Hämtar frågor från modellen

	response, err := json.Marshal(questions) // Kodar frågorna till JSON-format
	if err != nil {
		http.Error(w, "Failed to encode questions", http.StatusInternalServerError) // Skickar ett felmeddelande om kodningen misslyckas
		return
	}

	w.Header().Set("Content-Type", "application/json") // Ställer in innehållstypen för svaret till JSON
	w.Write(response)                                  // Skriver det kodade svaret till klienten
}
