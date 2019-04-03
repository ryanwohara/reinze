package main

import (
	"fmt"

	"crypto/tls"

	"../goreinze/runescape"
	irc "github.com/thoj/go-ircevent"
)

const channel = "#reinze"
const serverssl = "irc.swiftirc.net:6697"

func main() {
	runescape.Hello()
	fmt.Println(runescape.GetUsersOnline())
	ircnick1 := "PiKick"
	irccon := irc.IRC(ircnick1, "IRCTestSSL")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
	irccon.AddCallback("366", func(e *irc.Event) {})
	export(irccon)
	err := irccon.Connect(serverssl)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}

func addInvite(irccon *irc.Connection) {
	irccon.AddCallback("INVITE", func(event *irc.Event) {
		if event.Nick == "Dragon" {
			irccon.Join(event.Arguments[1])
		}
	})
}

func addPrivmsg(irccon *irc.Connection) {
	irccon.AddCallback("PRIVMSG", func(event *irc.Event) {
		if event.Nick == "Dragon" {
			if event.Message() == "TEST" {
				irccon.Privmsgf(event.Arguments[0], "Test over SSL successful\n")
			} else if event.Message() == "-players" {
				irccon.Notice(event.Nick, "There are currently "+runescape.GetUsersOnline()+" players online.")
			} else if event.Message() == "+players" {
				irccon.Privmsgf(event.Arguments[0], "There are currently %s players online.", runescape.GetUsersOnline())
			}
		}
	})
}

type binFunc func(irccon *irc.Connection)

func export(irccon *irc.Connection) {
	available := []binFunc{addInvite, addPrivmsg}
	for a := 0; a < len(available); a++ {
		available[a](irccon)
	}
}