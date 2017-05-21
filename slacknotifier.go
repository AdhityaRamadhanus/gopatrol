package gopatrol

import (
	"math/rand"
	"strings"
	"time"

	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/nlopes/slack"
)

// SlackNotifier is the main struct consist of all the sub component including slack api, real-time messaing api and face detector
type SlackNotifier struct {
	ID         string        `json:"-"`
	RTM        *slack.RTM    `json:"-"`
	SlackAPI   *slack.Client `json:"-"`
	ChannelID  string        `json:"channel" validate:"required"`
	SlackToken string        `json:"token" validate:"required"`
	Type       string        `json:"type" validate:"required"`
}

const (
	helpText = "How can I help you?"
)

var (
	greetingPattern  = "hi bot|hello bot"
	greetingPrefixes = []string{"Hi", "Hello", "Howdy", "Wazzzup", "Hey"}
)

// NewSlackNotifier create new Thug bot
func NewSlackNotifier(slackToken string, channelID string) *SlackNotifier {
	slackNotifier := &SlackNotifier{
		SlackToken: slackToken,
		SlackAPI:   slack.New(slackToken),
		ChannelID:  channelID,
	}
	go slackNotifier.run()
	return slackNotifier
}

func (t *SlackNotifier) messageHandler(ev *slack.MessageEvent) {
	if ev.Type == "message" &&
		(strings.HasPrefix(strings.ToLower(ev.Text), "hi bot") ||
			strings.HasPrefix(strings.ToLower(ev.Text), "hello bot")) {
		go t.helloWorld(ev)
		return
	}

	if ev.Type == "message" && strings.HasPrefix(strings.ToLower(ev.Text), "bot help") {
		go t.help(ev)
		return
	}
}

func (t *SlackNotifier) helloWorld(ev *slack.MessageEvent) (err error) {
	rand.Seed(time.Now().UnixNano())
	msg := greetingPrefixes[rand.Intn(len(greetingPrefixes))] + " <@" + ev.User + ">!"
	t.RTM.SendMessage(t.RTM.NewTypingMessage(ev.Channel))
	t.RTM.SendMessage(t.RTM.NewOutgoingMessage(msg, ev.Channel))
	return nil
}

func (t *SlackNotifier) help(ev *slack.MessageEvent) (err error) {
	t.RTM.SendMessage(t.RTM.NewOutgoingMessage(helpText, ev.Channel))
	return nil
}

func (t *SlackNotifier) run() {
	t.RTM = t.SlackAPI.NewRTM()
	go t.RTM.ManageConnection()

	for msg := range t.RTM.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			t.ID = ev.Info.User.ID
			log.WithField("bot_id", ev.Info.User.ID).Info("Connected to Slack Channel")
		case *slack.MessageEvent:
			t.messageHandler(ev)
		case *slack.RTMError:
			log.Error("Slack Error", ev.Error())
		case *slack.InvalidAuthEvent:
			log.Error("Failed to get authenticated in slack")
			return
		default:
		}
	}
}

//GetType return what type of notifier is thiss
func (t SlackNotifier) GetType() string {
	return "slack"
}

//Notify implements notifier interface, send slack message about an event
func (t SlackNotifier) Notify(result Result) error {
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{}
	// decide whether to send announcements or not
	// If the endpoint is Down
	switch {
	case result.Down:
		attachment := slack.Attachment{
			Title: result.Name + " is currently down",
			Text:  result.URL,
			Color: "danger",
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "Event Timestamp",
					Value: result.Timestamp.Format("2006-01-02-15:04:05"),
					Short: true,
				},
			},
		}
		params.Attachments = append(params.Attachments, attachment)
	case result.Healthy:
		attachment := slack.Attachment{
			Title: result.Name + " is up",
			Text:  result.URL,
			Color: "good",
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "Event Timestamp",
					Value: result.Timestamp.Format("2006-01-02-15:04:05"),
					Short: true,
				},
			},
		}
		params.Attachments = append(params.Attachments, attachment)
	}

	var err error
	headerMessage := fmt.Sprintf("<!channel> Service Announcements")
	_, _, err = t.SlackAPI.PostMessage(t.ChannelID, headerMessage, params)
	return err
}
