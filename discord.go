package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func initDiscord() (Bot, error) {
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		return Bot{}, err
	}

	bot := Bot{dg}

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Intents |= discordgo.IntentsMessageContent

	dg.AddHandler(bot.onMessageCreate)

	err = dg.Open()
	if err != nil {
		return Bot{}, err
	}

	return bot, nil
}

func (b Bot) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := "!"

	if len(m.Content) <= len(prefix) {
		return
	}

	cmdStr := m.Content[:len(prefix)]
	msg := m.Content[len(prefix):]

	if m.Author.ID == s.State.User.ID || cmdStr != prefix {
		return
	}

	switch msg {
	case "check":
		err := b.checkHousingPage(true)
		if err != nil {
			log.Fatal(err)
		}
	case "help":
		err := b.printPossibleCmds()
		if err != nil {
			log.Fatal(err)
		}
	case "stop":
		if !isChecking {
			b.notifyUser("Housing check is aleady stopped!")
		} else {
			b.notifyUser("Stopping housing check...")
			isChecking = false
		}

	case "start":
		if isChecking {
			b.notifyUser("Housing check is aleady running!")
		} else {
			b.notifyUser("Starting housing check...")
			isChecking = true
		}
	default:
		b.notifyUser("Please enter a valid command.")

		err := b.printPossibleCmds()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (b Bot) notifyUser(message string) error {
	err := printToLog(message)
	if err != nil {
		return err
	}

	_, err = b.Session.ChannelMessageSend(channelID, message)
	if err != nil {
		return err
	}

	return nil
}

func (b Bot) printPossibleCmds() error {
	msg := "Possible commands:\n"
	for _, cmd := range possibleCmds {
		msg += "\t- " + cmd + "\n"
	}

	_, err := b.Session.ChannelMessageSend(channelID, msg)
	if err != nil {
		return err
	}

	return nil
}
