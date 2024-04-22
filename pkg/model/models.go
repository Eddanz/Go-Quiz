// MODELS

package model

import "sync"

// Question representerar en fråga i ett quiz med flera svarsalternativ och index för det korrekta svaret.
type Question struct {
	ID       int      // Unikt identifierare för frågan
	Text     string   // Texten för frågan
	Choices  []string // En skärva med möjliga svar
	AnswerID int      // Indexet för det korrekta svaret i Choices
}

// Quiz innehåller en skärva av Questions och håller reda på poängen för alla avslutade quiz.
type Quiz struct {
	Questions []Question // En skärva av Question-objekt
	Scores    []int      // En skärva för att spåra poängen för alla genomförda quiz
	Mutex     sync.Mutex // Mutex för att säkerställa trådsäkerhet.
}

// Quizfrågor att lägga till
var SampleQuestions = []Question{
	{
		ID:       1,
		Text:     "What is the capital of France?",
		Choices:  []string{"Paris", "London", "Berlin", "Madrid"},
		AnswerID: 0,
	},
	{
		ID:       2,
		Text:     "Which element has the chemical symbol 'O'?",
		Choices:  []string{"Gold", "Oxygen", "Silver", "Iron"},
		AnswerID: 1,
	},
	{
		ID:       3,
		Text:     "Who wrote 'Macbeth'?",
		Choices:  []string{"Charles Dickens", "William Shakespeare", "Jane Austen", "Mark Twain"},
		AnswerID: 1,
	},
	{
		ID:       4,
		Text:     "What is the largest planet in our solar system?",
		Choices:  []string{"Earth", "Mars", "Jupiter", "Saturn"},
		AnswerID: 2,
	},
	{
		ID:       5,
		Text:     "What year did the Titanic sink in the Atlantic Ocean on April 15, during its maiden voyage from Southampton?",
		Choices:  []string{"1912", "1898", "1905", "1923"},
		AnswerID: 0,
	},
	{
		ID:       6,
		Text:     "What is the speed of light in vacuum, approximately?",
		Choices:  []string{"150,000 km/s", "300,000 km/s", "120,000 km/s", "500,000 km/s"},
		AnswerID: 1,
	},
	{
		ID:       7,
		Text:     "What programming language is known for its simplicity and readability?",
		Choices:  []string{"C++", "Python", "Java", "JavaScript"},
		AnswerID: 1,
	},
	{
		ID:       8,
		Text:     "Which company is known as the 'Big Blue'?",
		Choices:  []string{"Microsoft", "Intel", "IBM", "Facebook"},
		AnswerID: 2,
	},
	{
		ID:       9,
		Text:     "What is the main ingredient in traditional Japanese miso soup?",
		Choices:  []string{"Tofu", "Rice", "Miso Paste", "Seaweed"},
		AnswerID: 2,
	},
	{
		ID:       10,
		Text:     "What is the process called in which plants make their food using sunlight?",
		Choices:  []string{"Respiration", "Photosynthesis", "Digestion", "Fermentation"},
		AnswerID: 1,
	},
}
