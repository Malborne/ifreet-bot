package commands

import (
	"fmt"

	heimdallr "github.com/Malborne/ifreet/tree/master/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/docopt/docopt-go"
	"github.com/pkg/errors"
)

var mystudentsCommand = command{
	"mystudents",
	commandMystudents,
	"gives a list of students registered in a circle",
	[]string{
		"",
	},
	[]string{
		"",
	},
}

//commandMystudents gives a list of students registered in a circle
func commandMystudents(s *discordgo.Session, m *discordgo.MessageCreate, args docopt.Opts) error {

	guildID := m.GuildID
	guild, err := heimdallr.GetGuild(s, guildID)
	if err != nil {
		return err
	}

	author, err := heimdallr.GetMember(s, guildID, m.Author.ID)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Message Author with ID %s was not found.", m.Author.ID))
		return errors.Wrap(err, "getting author failed")
	}

	if !hasRole(author, heimdallr.Config.CricleLeaderRole) || !heimdallr.IsCircleMember(author, guild) {
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You are not a leader of any circle and do not have students."))
		return errors.Wrap(err, "sending message failed")
	}

	circleName := ""

	if hasRole(author, heimdallr.Config.OmerIbnAlKhattabRole) {
		circleName = "Omer Ibn Al Khattab's Circle"
	} else if hasRole(author, heimdallr.Config.AbuBakrAlSiddeeqRole) {
		circleName = "Abu Bakr Al Siddeeq's Circle"
	} else if hasRole(author, heimdallr.Config.AliBinAbiTaalibRole) {
		circleName = "Ali Ibn Abi Talib's Circle"
	} else if hasRole(author, heimdallr.Config.SistersCircleRole) {
		circleName = "Sisters Circle"
	}
	students, err := heimdallr.GetStudents(circleName)
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Could not find the circle members in the database."))
		return errors.Wrap(err, "getting the sheetLink failed")
	}
	if students[0].ID != "" {
		var fields []*discordgo.MessageEmbedField
		for _, student := range students {
			member, _ := heimdallr.GetMember(s, guildID, student.ID)
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  member.Mention(),
				Value: student.SheetLink,
			})
		}
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    author.User.Username,
				IconURL: author.User.AvatarURL(""),
			},
			Title:  circleName,
			Fields: fields,
		})
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The student is not registered in the database. Make sure you add the student first."))
		return errors.Wrap(err, "getting the sheetLink failed")
	}
	return errors.Wrap(err, "sending message failed")

}
