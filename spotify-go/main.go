package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	state := uuid.New().String()
	if err := godotenv.Load(); err != nil {
		log.Panicln(err)
	}
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Panic("CLIENT_ID isn't provided")
	}
	secretID := os.Getenv("SECRET_ID")

	auth := spotifyauth.New(spotifyauth.WithClientID(clientID), spotifyauth.WithClientSecret(secretID))
	auth.Token(ctx, state, )
	spotify.New(auth.Client())
}
