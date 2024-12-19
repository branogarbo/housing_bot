package main

import "github.com/bwmarrin/discordgo"

type Bot struct {
	Session *discordgo.Session
}

type HTTPoptions struct {
  URL string
  Method string
  Body string
  Headers map[string]string
}

