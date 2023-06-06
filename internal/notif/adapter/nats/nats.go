package nats

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"

	"github.com/tribefintech/microservices/internal/notif/domain"
)

type NATS struct {
	nc     *nats.Conn
	logger *logrus.Entry
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
	Course struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	} `json:"course"`
	ReceiverId string `json:"receiver_id"`
	CreatedAt  string `json:"created_at"`
	NoT        int8   `json:"number_of_times"` // 1 or 2
}

func NewNATS(url string, logger *logrus.Entry) *NATS {
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}
	return &NATS{
		nc: nc,
		logger: logger.WithFields(logrus.Fields{
			"adapter": "NATS",
		}),
	}
}

func (n *NATS) Run() <-chan domain.Event {
	source := make(chan domain.Event)
	n.subscribeSocialMessageCreated(source)
	n.subscribeEndorsementRequestCreated(source)
	n.subscribeEducationCourseProgressRemind(source)
	return source
}

func (n *NATS) subscribeSocialMessageCreated(source chan domain.Event) {
	_, err := n.nc.Subscribe(string(domain.SocialMessageCreatedCode), func(msg *nats.Msg) {
		logger := n.logger.WithFields(logrus.Fields{
			"event": msg.Subject,
		})

		logger.Info("received an event")

		var event SocialMessageCreated

		// todo: using when we have protobuf message
		// err := proto.Unmarshal(msg.Data, &event)
		// if err != nil {
		// 	logger.Errorf("error unmarshal event: %s", err.Error())
		// 	return
		// }

		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			logger.Errorf("error unmarshal event: %s", err.Error())
			return
		}
		senderId, err := uuid.Parse(event.Sender.Id)
		if err != nil {
			logger.Errorf("can not parse uuid: %s", err.Error())
			return
		}
		receiverIds := make([]uuid.UUID, 0)
		for _, v := range event.ReceiverIds {
			receiverId, err := uuid.Parse(v)
			if err != nil {
				logger.Errorf("can not parse uuid: %s", err.Error())
				return
			}
			receiverIds = append(receiverIds, receiverId)
		}
		source <- domain.NewSocialMessageCreated(domain.Sender{
			FirstName: event.Sender.FirstName,
			Id:        senderId,
			ImageURL:  event.Sender.ImageURL,
			LastName:  event.Sender.LastName,
		},
			event.Msg,
			&receiverIds,
			event.CreatedAt,
			event.ChannelId,
		)
	})
	if err != nil {
		panic(err)
	}

}

func (n *NATS) subscribeEndorsementRequestCreated(source chan domain.Event) {
	_, err := n.nc.Subscribe(string(domain.EndorsementRequestCreatedCode), func(msg *nats.Msg) {
		logger := n.logger.WithFields(logrus.Fields{
			"event": msg.Subject,
		})

		logger.Info("received an event")

		var event EndorsementRequestCreated

		// todo: using when we have protobuf message
		// err := proto.Unmarshal(msg.Data, &event)
		// if err != nil {
		// 	logger.Errorf("error unmarshal event: %s", err.Error())
		// 	return
		// }

		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			logger.Errorf("error unmarshal event: %s", err.Error())
			return
		}
		senderId, err := uuid.Parse(event.Sender.Id)
		if err != nil {
			logger.Errorf("can not parse uuid: %s", err.Error())
			return
		}

		receiverId, err := uuid.Parse(event.ReceiverId)
		if err != nil {
			logger.Errorf("can not parse uuid: %s", err.Error())
			return
		}

		source <- domain.NewEndorsementRequestCreated(&domain.Sender{
			FirstName: event.Sender.FirstName,
			Id:        senderId,
			ImageURL:  event.Sender.ImageURL,
			LastName:  event.Sender.LastName,
		},
			event.Endorsement.Id,
			receiverId,
			event.CreatedAt,
		)
	})
	if err != nil {
		panic(err)
	}
}

func (n *NATS) subscribeEducationCourseProgressRemind(source chan domain.Event) {
	_, err := n.nc.Subscribe(string(domain.EducationCourseProgressRemindCode), func(msg *nats.Msg) {
		logger := n.logger.WithFields(logrus.Fields{
			"event": msg.Subject,
		})

		logger.Info("received an event")

		var event EducationCourseProgressRemind

		// todo: using when we have protobuf message
		// err := proto.Unmarshal(msg.Data, &event)
		// if err != nil {
		// 	logger.Errorf("error unmarshal event: %s", err.Error())
		// 	return
		// }

		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			logger.Errorf("error unmarshal event: %s", err.Error())
			return
		}

		receiverId, err := uuid.Parse(event.ReceiverId)
		if err != nil {
			logger.Errorf("can not parse uuid: %s", err.Error())
			return
		}

		pe, err := domain.NewEducationCourseProgressRemind(domain.Course{
			Id:    event.Course.Id,
			Title: event.Course.Title,
		},
			receiverId,
			event.CreatedAt,
			event.NoT,
		)

		if err != nil {
			logger.Errorf("invalid event: %s", err.Error())
			return
		}

		source <- pe
	})

	if err != nil {
		panic(err)
	}
}
