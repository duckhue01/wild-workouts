package notif

import (
	"encoding/json"
	"testing"

	"github.com/nats-io/nats.go"
)

type EndorsementRequestCreated struct {
	Endorsement struct {
		Id int `json:"id"`
	} `json:"endorsement"`

	Sender struct {
		Id        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		ImageURL  string `json:"image_url"`
	} `json:"sender"`
	ReceiverId string `json:"receiver_id"`
	CreatedAt  string `json:"created_at"`
}

type EducationCourseProgressRemind struct {
	Source struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	} `json:"course"`
	ReceiverId string `json:"receiver_id"`
	CreatedAt  string `json:"created_at"`
	NoT        int8   `json:"number_of_times"`
}

type SocialMessageCreated struct {
	ChannelId string `json:"chanel_id"`
	Msg       string `json:"msg"`
	Sender    struct {
		Id        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		ImageURL  string `json:"image_url"`
	} `json:"sender"`
	ReceiverIds []string `json:"receiver_ids"`
	CreatedAt   string   `json:"created_at"`
}

func TestSocialMessageCreatedEvent(t *testing.T) {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Error(err)
	}

	a := &SocialMessageCreated{
		ChannelId: "IKE",
		Msg:       "ike abc",
		Sender: struct {
			Id        string "json:\"id\""
			FirstName string "json:\"first_name\""
			LastName  string "json:\"last_name\""
			ImageURL  string "json:\"image_url\""
		}{
			Id:        "6f3d61e4-5fee-4b5a-995a-02f65765fb32",
			FirstName: "test",
			LastName:  "test",
			ImageURL:  "test",
		},
		ReceiverIds: []string{"6f3d61e4-5fee-4b5a-995a-02f65765fb32"},
		CreatedAt:   "2022-12-22T06:00:00.200Z",
	}

	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
	}

	err = nc.Publish("social.message.created", b)
	if err != nil {
		t.Error(err)
	}

	nc.Flush()
}

func TestEndorsementRequestCreatedEvent(t *testing.T) {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Error(err)
	}

	a := &EndorsementRequestCreated{
		Endorsement: struct {
			Id int "json:\"id\""
		}{
			Id: 1,
		},
		Sender: struct {
			Id        string "json:\"id\""
			FirstName string "json:\"first_name\""
			LastName  string "json:\"last_name\""
			ImageURL  string "json:\"image_url\""
		}{
			Id:        "test",
			FirstName: "test",
			LastName:  "test",
			ImageURL:  "test",
		},
		ReceiverId: "6f3d61e4-5fee-4b5a-995a-02f65765fb32",
		CreatedAt:  "2022-12-22T06:00:00.200Z",
	}

	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
	}

	err = nc.Publish("endorsement.request.created", b)
	if err != nil {
		t.Error(err)
	}

	nc.Flush()
}

func TestEducationCourseProgressRemindEvent(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Error(err)
	}

	a := &EducationCourseProgressRemind{
		Source: struct {
			Id    int    "json:\"id\""
			Title string "json:\"title\""
		}{
			Id:    10,
			Title: "introduction to duckhue1",
		},
		ReceiverId: "6f3d61e4-5fee-4b5a-995a-02f65765fb32",
		CreatedAt:  "2022-12-22T06:00:00.200Z",
		NoT:        int8(1),
	}

	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
	}

	err = nc.Publish("education.course.progress.created", b)
	if err != nil {
		t.Error(err)
	}

	nc.Flush()
}
