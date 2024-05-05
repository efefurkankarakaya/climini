package processor

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
)

func StartASinglePromptSession(ctx context.Context, model *genai.GenerativeModel) string {
	for {
		fmt.Printf("Prompt: ")
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
				fmt.Println(len(res.Candidates[i].Content.Parts))
				fmt.Println(res.Candidates[i].Content.Parts[j])
			}
		}
	}
}
