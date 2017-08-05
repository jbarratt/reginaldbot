package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kyokomi/emoji"
	"github.com/pelletier/go-toml"
	"github.com/tucnak/telebot"
)

func main() {
	conf, err := toml.LoadFile("conf.toml")
	if err != nil {
		log.Fatalf("Unable to parse conf file: %v", err)
	}
	jrnlPath := conf.Get("jrnl.path").(string)
	err = jrnlEntry("test entry", jrnlPath)
	if err != nil {
		log.Fatalf("Error storing journal: %v\n", err)
	}

	botToken := conf.Get("telegram.token").(string)
	myID := int(conf.Get("telegram.myID").(int64))
	myUser := telebot.User{ID: myID}
	if err != nil {
		log.Fatalln(err)
	}

	bot, err := telebot.NewBot(botToken)
	if err != nil {
		log.Fatalln(err)
	}

	messages := make(chan telebot.Message, 100)
	bot.Listen(messages, 1*time.Second)

	bot.SendMessage(myUser, emoji.Sprint("ReginaldBot at your service :tophat:"), nil)

	for message := range messages {
		if message.Sender.ID != myID {
			bot.SendMessage(message.Chat, "I'm sorry, I'm a private bot.", nil)
			continue
		}
		parts := strings.SplitN(message.Text, " ", 2)
		cmd := parts[0]
		body := parts[1]

		if cmd == "/jrnl" || cmd == "/j" {
			err := jrnlEntry(body, jrnlPath)
			if err != nil {
				log.Printf("Error saving journal: %v", err)
				bot.SendMessage(message.Chat, "unable to save, sorry.", nil)
			} else {
				bot.SendMessage(message.Chat, emoji.Sprint("Saved :+1:"), nil)
			}
			continue
		}
		bot.SendMessage(message.Chat, "I'm sorry, unknown command.", nil)
	}
}

func jrnlEntry(content, jrnlPath string) error {
	t := time.Now()
	year := fmt.Sprintf("%d", t.Year())
	fname := fmt.Sprintf("%02d.md", t.Month())
	heading := fmt.Sprintf("### %s\n\n", t.Format("Mon Jan _2, 3:04PM"))

	// make the base directory
	storeDir := filepath.Join(jrnlPath, year)
	os.MkdirAll(storeDir, os.ModePerm)

	jrnlFile := filepath.Join(storeDir, fname)

	f, err := os.OpenFile(jrnlFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(heading + content + "\n\n")
	if err != nil {
		return err
	}
	return nil
}
