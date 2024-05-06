package processor

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
)

func StartTurnBasedPromptSession(ctx context.Context, model *genai.GenerativeModel) {
	var responseContent genai.Part
	cs := model.StartChat()
	cs.History = []*genai.Content{
		// &genai.Content{
		// 	Parts: []genai.Part{
		// 		genai.Text(prompt),
		// 	},
		// 	Role: "user",
		// },
	}

	for {
		fmt.Printf("Prompt: ")
		in := bufio.NewReader((os.Stdin))
		prompt, err := in.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		res, err := cs.SendMessage(ctx, genai.Text(prompt))

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
		fmt.Printf("Prompt (No History): ")
		in := bufio.NewReader(os.Stdin)
		prompt, err := in.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

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
