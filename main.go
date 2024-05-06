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
}
