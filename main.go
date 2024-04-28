package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	var prompt string

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	for {
		fmt.Printf("Prompt: ")
		_, err = fmt.Scan(&prompt)

		if err != nil {
			log.Fatal(err)
		}

		res, err := model.GenerateContent(ctx, genai.Text(prompt))

		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(res.Candidates); i++ {
			for j := 0; j < len(res.Candidates[i].Content.Parts); j++ {
				fmt.Println(res.Candidates[i].Content.Parts[j])
			}
		}
	}
}
