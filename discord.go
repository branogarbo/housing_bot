package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func initDiscord() (Discord, error) {
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		// fmt.Println("error creating Discord session:", err)
		return Discord{}, err
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection:", err)
		return Discord{}, err
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	return Discord{dg}, nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// fmt.Println(m.Content)
	cmd := "!house"

	if len(m.Content) < len(cmd) {
		return
	}

	cmdStr := m.Content[:len(cmd)]
	msg := m.Content[len(cmd)+1:]

	if m.Author.ID == s.State.User.ID || cmdStr != cmd {
		return
	}

	if msg == "ping" {
		s.ChannelMessageSend(channelID, "pong")
	}
}

func (d *Discord) notifyUser(message string) error {
	err := printToLog(message)
	if err != nil {
		log.Fatal(err)
	}

	_, err = d.Session.ChannelMessageSend(channelID, message)
	if err != nil {
		return err
	}

	return nil
}
