package main

import (
	processor "climini/internal/processor"
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	processor.ReadArguments()

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	// processor.StartASinglePromptSession(ctx, model)
	processor.StartTurnBasedPromptSession(ctx, model)

	// for {
	// 	fmt.Printf("Prompt: ")
	// 	in := bufio.NewReader(os.Stdin)
	// 	prompt, err := in.ReadString('\n')

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	res, err := model.GenerateContent(ctx, genai.Text(prompt))

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(prompt)
	// 	fmt.Println(len(res.Candidates))
	// 	for i := 0; i < len(res.Candidates); i++ {
	// 		for j := 0; j < len(res.Candidates[i].Content.Parts); j++ {
	// 			fmt.Println(len(res.Candidates[i].Content.Parts))
	// 			fmt.Println(res.Candidates[i].Content.Parts[j])
	// 		}
	// 	}
	// }
}
