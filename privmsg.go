package main

import (
	"github.com/ryanwohara/reinze/runescape"
	irc "github.com/thoj/go-ircevent"
)

func addPrivmsg(irccon *irc.Connection) {
	irccon.AddCallback("PRIVMSG", func(event *irc.Event) {
		if event.Nick == "Dragon" {
			msg := event.Message()[1:]
			input := event.Message()[0]

			output := runescape.Matches(msg)

			if len(output) == 0 {
				return
			}
			if string(input) == "-" {
				irccon.Notice(event.Nick, output)
			} else if string(input) == "+" {
				irccon.Privmsgf(event.Arguments[0], output)
			}
		}
	})
}
