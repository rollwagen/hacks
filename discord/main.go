package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/samber/lo"
)

const SANSHolidayHackChallengeID = "783055461620514818"

func main() {
	token := os.Getenv("DISCORD_AUTHORIZATION_TOKEN")
	// fmt.Printf("token=%s\n", token)
	// session, err := discordgo.New("NTI3NzMyNDg2NzkzNDYxNzcw.G7xp_o.l6znB8dykvG643e6DO9B4LygFHoU2GEMOigD0Y")
	session, err := discordgo.New(token)
	// session, err := discordgo.New("NTI3NzMyNDg2NzkzNDYxNzcw.c9hPykVHRINYHLIrqx9CtxhukqE")
	// session, err = discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	// guilds, _ := session.UserGuilds(10, "aaa", "bbb")
	// for _, g := range guilds {
	// 	fmt.Println(g.Name)
	// }
	err = session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	// u, _ := session.User("@me")
	// fmt.Println(u)

	// for _, g := range session.State.Guilds {
	// 	fmt.Printf("name=%v id=%s\n", g.Name, g.ID)
	// }

	channels, err := session.GuildChannels(SANSHolidayHackChallengeID)
	_ = channels
	if err != nil {
		panic(err)
	}

	// for _, c := range channels {
	// 	fmt.Printf("channel name=%v id=%s\n", c.Name, c.ID)
	// }

	game := "1179764228362600529"
	// messages, _ := session.ChannelMessages(snowballFightID, 10, "", "", "")
	// for _, m := range messages {
	// 	fmt.Printf("%s: %v - %s\n", m.ID,m.Timestamp ,m.Content)
	// }
	messages := channelMessages(session, game, "")
	for _, m := range lo.Reverse(messages) {
		fmt.Println(m)
	}
}

func channelMessages(session *discordgo.Session, channelID string, beforeID string) []string {
	var allMessages []string

	messages, _ := session.ChannelMessages(channelID, 100, beforeID, "", "")
	for _, m := range messages {
		// fmt.Printf("%s: %v - %s\n", m.ID,m.Timestamp ,m.Content)
		allMessages = append(allMessages, m.Content)
	}

	if len(messages) == 100 {
		id := messages[100-1].ID
		allMessages = append(allMessages, channelMessages(session, channelID, id)...)
	}

	return allMessages
}

// 	dg.AddHandler(messageCreate)
//
// 	err = dg.Open()
// 	if err != nil {
// 		fmt.Println("error opening connection,", err)
// 		return
// 	}
//
// 	fmt.Println("Bot is now running. Press CTRL-C to exit.")
// 	sc := make(chan os.Signal, 1)
// 	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
// 	<-sc
//
// 	dg.Close()
// }
//
// func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	if m.Author.ID == s.State.User.ID {
// 		return
// 	}
//
// 	fmt.Printf("Message: %+v\n", m.Message)
// }
