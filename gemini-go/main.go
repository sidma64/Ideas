package main

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type model struct {
	text    string
	gemini  genai.GenerativeModel
	cursor  int
	width   int
	height  int
	lastKey string
	textField textinput.Model
}

func initialModel(ctx context.Context) model {
	key := os.Getenv("GEMINI_API_KEY")
	if key == "" {
		panic("GEMINI_API_KEY is required")
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(key))
	if err != nil {
		panic(err)
	}
	return model{
		gemini: *client.GenerativeModel("gemini-1.0-pro"),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		key := msg.String()
		switch {
		case len(key) == 1:
			m.text = m.text[m.cursor:] + key + m.text[:m.cursor]
		case key == "ctrl+c":
			return m, tea.Quit
		}
		m.lastKey = key
	}
	return m, nil
}

func (m model) View() (view string) {
	view += fmt.Sprintf("Width: %v\n", m.width)
	view += fmt.Sprintf("Height: %v\n", m.height)
	view += fmt.Sprintf("Key: %v\n", m.lastKey)
	return view
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	ctx := context.Background()
	p := tea.NewProgram(initialModel(ctx))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
