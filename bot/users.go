package heimdallr

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/bwmarrin/discordgo"
)

//UserJoinHandler handles new users joining the server, and will welcome them.
// func UserJoinHandler(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
// 	_, _ = s.ChannelMessageSend(Config.WelcomeChannel, "Hello")
// 	_, _ = s.ChannelMessageSend(Config.AdminLogChannel, "Hello")

// 	welcomeMessage := Config.WelcomeMessage
// 	if strings.Count(welcomeMessage, "%s") > 0 {
// 		welcomeMessage = fmt.Sprintf(welcomeMessage, g.User.Mention(), Config.RulesChannel)
// 	}
// 	_, err := s.ChannelMessageSend(Config.WelcomeChannel, welcomeMessage)
// 	LogIfError(s, errors.Wrap(err, "sending message failed"))
// }

// //UserLeaveHandler wishes ex members goodbye
// func UserLeaveHandler(s *discordgo.Session, g *discordgo.GuildMemberRemove) {

// 	var name string
// 	if g.Nick != "" {
// 		name = g.Nick
// 	} else {
// 		name = g.User.Username
// 	}
// 	_, err := s.ChannelMessageSend(Config.WelcomeChannel, fmt.Sprintf("User `%s` (%s) has left the building.", name, g.User.Mention()))
// 	LogIfError(s, errors.Wrap(err, "sending message failed"))
// }

//UserJoinHandlerOriginal handles new users joining the server, and will welcome them.
func UserJoinHandlerOriginal(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
	_, _ = s.ChannelMessageSend(Config.WelcomeChannel, "Hello")

	welcomeMessage := Config.WelcomeMessage
	if strings.Count(welcomeMessage, "%s") > 0 {
		welcomeMessage = fmt.Sprintf(welcomeMessage, g.User.Mention())
	}
	_, err := s.ChannelMessageSend(Config.WelcomeChannel, welcomeMessage)
	LogIfError(s, errors.Wrap(err, "sending message failed"))
}

//UserLeaveHandlerOriginal wishes ex members goodbye
func UserLeaveHandlerOriginal(s *discordgo.Session, g *discordgo.GuildMemberRemove) {
	_, _ = s.ChannelMessageSend(Config.WelcomeChannel, "Goodbye")

	var name string
	if g.Nick != "" {
		name = g.Nick
	} else {
		name = g.User.Username
	}
	_, err := s.ChannelMessageSend(Config.WelcomeChannel, fmt.Sprintf("User `%s` (%s) has left the building.", name, g.User.Mention()))
	LogIfError(s, errors.Wrap(err, "sending message failed"))
}
