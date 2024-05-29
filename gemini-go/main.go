package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"github.com/charmbracelet/glamour"

)

var chat *genai.ChatSession

func main() {
	print(">")
	ctx := context.Background()
	godotenv.Load()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with most use cases
	model := client.GenerativeModel("gemini-1.0-pro")
	chat = model.StartChat()

  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
		text := scanner.Text() // Get the current line of text
		if text == "" {
			break // Exit loop if an empty line is entered
		}
		input(ctx, text)
		print(">")
	}
}

func input(ctx context.Context, in string) {
	res, err := chat.SendMessage(ctx, genai.Text(in))
	if err != nil {
		log.Println(err)
		return
	}
	parts := res.Candidates[0].Content.Parts
	for _, part := range parts {
		out, _ := glamour.Render(part, "dark")
		fmt.Println(out)
	}
}
