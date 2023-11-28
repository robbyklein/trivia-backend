package helpers

import (
	"fmt"

	"github.com/robbyklein/trivia-backend/models"
	"gorm.io/gorm"
)

func PopulateQuestions(db *gorm.DB) {
	questions := []models.Question{
		{
			Question: "What is the capital of Japan?",
			Answer:   "Tokyo",
		},
		{
			Question: "Who wrote the play 'Romeo and Juliet'?",
			Answer:   "William Shakespeare",
		},
		{
			Question: "What is the largest organ in the human body?",
			Answer:   "Skin",
		},
		{
			Question: "Which gas do plants absorb from the atmosphere?",
			Answer:   "Carbon dioxide (CO2)",
		},
		{
			Question: "Who painted the 'Mona Lisa'?",
			Answer:   "Leonardo da Vinci",
		},
		{
			Question: "What is the smallest planet in our solar system?",
			Answer:   "Mercury",
		},
		{
			Question: "In which year did Christopher Columbus discover America?",
			Answer:   "1492",
		},
		{
			Question: "What is the chemical symbol for gold?",
			Answer:   "Au",
		},
		{
			Question: "Which planet is known as the 'Morning Star'?",
			Answer:   "Venus",
		},
		{
			Question: "Who is the author of 'To Kill a Mockingbird'?",
			Answer:   "Harper Lee",
		},
		{
			Question: "What is the largest desert in the world?",
			Answer:   "Sahara Desert",
		},
		{
			Question: "Which element has the atomic number 1?",
			Answer:   "Hydrogen",
		},
		{
			Question: "What is the highest mountain in the world?",
			Answer:   "Mount Everest",
		},
		{
			Question: "Who painted the 'Starry Night'?",
			Answer:   "Vincent van Gogh",
		},
		{
			Question: "What is the chemical symbol for oxygen?",
			Answer:   "O",
		},
		{
			Question: "Which gas makes up the majority of Earth's atmosphere?",
			Answer:   "Nitrogen",
		},
		{
			Question: "What is the largest species of shark?",
			Answer:   "Whale Shark",
		},
		{
			Question: "Who wrote the 'Harry Potter' book series?",
			Answer:   "J.K. Rowling",
		},
		{
			Question: "What is the longest river in the world?",
			Answer:   "Nile River",
		},
		{
			Question: "In which year did the Titanic sink?",
			Answer:   "1912",
		},
		{
			Question: "What is the capital of Australia?",
			Answer:   "Canberra",
		},
		{
			Question: "What is the chemical symbol for water?",
			Answer:   "H2O",
		},
		{
			Question: "Which country is known as the Land of the Rising Sun?",
			Answer:   "Japan",
		},
		{
			Question: "Who was the first woman to fly solo across the Atlantic Ocean?",
			Answer:   "Amelia Earhart",
		},
		{
			Question: "What is the largest planet in our solar system?",
			Answer:   "Jupiter",
		},
		{
			Question: "Who painted the 'Sistine Chapel Ceiling'?",
			Answer:   "Michelangelo",
		},
		{
			Question: "What is the chemical symbol for carbon?",
			Answer:   "C",
		},
		{
			Question: "Which mammal can fly and is often associated with Halloween?",
			Answer:   "Bat",
		},
		{
			Question: "What is the largest continent in the world?",
			Answer:   "Asia",
		},
		{
			Question: "Who is the author of 'The Great Gatsby'?",
			Answer:   "F. Scott Fitzgerald",
		},
		{
			Question: "What is the chemical symbol for silver?",
			Answer:   "Ag",
		},
		{
			Question: "Which planet is known as the 'Red Planet'?",
			Answer:   "Mars",
		},
		{
			Question: "What is the national flower of Japan?",
			Answer:   "Cherry Blossom",
		},
		{
			Question: "Who is known as the 'Father of Modern Physics'?",
			Answer:   "Albert Einstein",
		},
		{
			Question: "What is the largest bird in the world?",
			Answer:   "Ostrich",
		},
	}

	// Loop through the questions and insert them into the database
	for _, question := range questions {
		result := db.Create(&question)
		if result.Error != nil {
			fmt.Printf("Error inserting question: %v\n", result.Error)
		} else {
			fmt.Printf("Inserted question: %s\n", question.Question)
		}
	}

	fmt.Println("Trivia questions populated successfully!")

}
