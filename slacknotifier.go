package checkupd

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"fmt"

	"github.com/nlopes/slack"
)

type EndpointState struct {
	Name       string
	URL        string
	LastChange int64
	LastStatus StatusText
}

// SlackNotifier is the main struct consist of all the sub component including slack api, real-time messaing api and face detector
type SlackNotifier struct {
	ID             string                   `json:"-"`
	RTM            *slack.RTM               `json:"-"`
	SlackApi       *slack.Client            `json:"-"`
	EndpointStates map[string]EndpointState `json:"-"`
	ChannelID      string                   `json:"channel"`
	SlackToken     string                   `json:"token"`
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
		SlackToken:     slackToken,
		SlackApi:       slack.New(slackToken),
		ChannelID:      channelID,
		EndpointStates: make(map[string]EndpointState),
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
	t.RTM = t.SlackApi.NewRTM()
	go t.RTM.ManageConnection()

	for msg := range t.RTM.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			t.ID = ev.Info.User.ID
			log.Println(ev.Info.User.ID, "Connected")
		case *slack.MessageEvent:
			t.messageHandler(ev)
		case *slack.RTMError:
			log.Println(ev.Error())
		case *slack.InvalidAuthEvent:
			log.Println("Failed to Authenticate")
			return
		default:
		}
	}
}

func (s *SlackNotifier) GetType() string {
	return "slack"
}

func (s *SlackNotifier) Notify(results []Result) error {
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{}
	// decide whether to send announcements or not
	boolSend := false

	tempState := make(map[string]EndpointState)

	for _, result := range results {
		state, ok := s.EndpointStates[result.Title]
		if !ok {
			state = EndpointState{
				Name:       result.Title,
				URL:        result.Endpoint,
				LastChange: 0,
				LastStatus: result.Status(),
			}
		}
		// If the endpoint is Down
		switch {
		case result.Down:

			if state.LastStatus == "healthy" || state.LastChange == 0 {
				boolSend = true
				attachment := slack.Attachment{
					Title: result.Title + " is currently down",
					Text:  result.Endpoint,
					Color: "danger",
					Fields: []slack.AttachmentField{
						slack.AttachmentField{
							Title: "Last Checked",
							Value: time.Unix(0, result.Timestamp).Format("2006-01-02-15:04:05"),
							Short: true,
						},
						slack.AttachmentField{
							Title: "Last Up",
							Value: time.Unix(0, state.LastChange).Format("2006-01-02-15:04:05"),
							Short: true,
						},
					},
				}
				params.Attachments = append(params.Attachments, attachment)
				state.LastChange = result.Timestamp
				state.LastStatus = result.Status()
			}
			lastResultTime := time.Unix(0, result.Timestamp)
			lastChangeTime := time.Unix(0, state.LastChange)
			diffMinutes := lastResultTime.Sub(lastChangeTime).Minutes()
			if state.LastStatus == "down" && diffMinutes > 5.0 {
				boolSend = true
				attachment := slack.Attachment{
					Title: result.Title + " still down",
					Text:  result.Endpoint,
					Color: "danger",
					Fields: []slack.AttachmentField{
						slack.AttachmentField{
							Title: "Last Checked",
							Value: time.Unix(0, result.Timestamp).Format("2006-01-02-15:04:05"),
							Short: true,
						},
						slack.AttachmentField{
							Title: "Last Up",
							Value: time.Unix(0, state.LastChange).Format("2006-01-02-15:04:05"),
							Short: true,
						},
					},
				}
				params.Attachments = append(params.Attachments, attachment)
				state.LastChange = result.Timestamp
				state.LastStatus = result.Status()
			}
		case result.Healthy:
			if state.LastStatus == "down" {
				boolSend = true
				attachment := slack.Attachment{
					Title: result.Title + " just got resurrected",
					Text:  result.Endpoint,
					Color: "good",
					Fields: []slack.AttachmentField{
						slack.AttachmentField{
							Title: "Last Checked",
							Value: time.Unix(0, result.Timestamp).Format("2006-01-02-15:04:05"),
							Short: true,
						},
						slack.AttachmentField{
							Title: "Last Down",
							Value: time.Unix(0, state.LastChange).Format("2006-01-02-15:04:05"),
							Short: true,
						},
					},
				}
				params.Attachments = append(params.Attachments, attachment)
				state.LastChange = result.Timestamp
				state.LastStatus = result.Status()
			}
		}
		tempState[result.Title] = state
	}

	s.EndpointStates = tempState

	var err error
	if boolSend {
		headerMessage := fmt.Sprintf("<!channel> Service Announcements")
		_, _, err = s.SlackApi.PostMessage(s.ChannelID, headerMessage, params)
	}
	return err
}
