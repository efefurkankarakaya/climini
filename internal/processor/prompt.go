package processor

import (
	"climini/internal/utility"
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
)

func StartTurnBasedPromptSession(ctx context.Context, model *genai.GenerativeModel) {
	var responseContent genai.Part

	session := model.StartChat()
	session.History = []*genai.Content{
		// &genai.Content{
		// 	Parts: []genai.Part{
		// 		genai.Text(prompt),
		// 	},
		// 	Role: "user",
		// },
	}

	for {
		prompt := utility.ReadInput("Prompt: ")

		res, err := session.SendMessage(ctx, genai.Text(prompt))

		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(res.Candidates); i++ {
			for j := 0; j < len(res.Candidates[i].Content.Parts); j++ {
				responseContent = res.Candidates[i].Content.Parts[j]
				fmt.Println(responseContent)
			}
		}
	}
}

// No History
func StartASinglePromptSession(ctx context.Context, model *genai.GenerativeModel) {
	var responseContent genai.Part

	for {
		prompt := utility.ReadInput("Prompt (No History): ")

		res, err := model.GenerateContent(ctx, genai.Text(prompt))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(prompt)
		fmt.Println(len(res.Candidates))
		for i := 0; i < len(res.Candidates); i++ {
			for j := 0; j < len(res.Candidates[i].Content.Parts); j++ {
				// length := len(res.Candidates[i].Content.Parts)
				// fmt.Println(length)
				responseContent = res.Candidates[i].Content.Parts[j]
				fmt.Println(responseContent)
			}
		}
	}
}
