package notifs

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/tribefintech/microservices/internal/notif/domain"
)

type (
	EndorsementRequestCreatedNotifData struct {
		Sender        *Sender
		EndorsementId int
		CreateAt      string
	}
	SocialMessageCreatedNotifData struct {
		Sender    *Sender
		Msg       string
		ChannelId string
		CreateAt  string
	}

	EducationCourseProgressRemindData struct {
		Course   *domain.Course
		CreateAt string
	}

	Sender struct {
		Id        uuid.UUID
		FirstName string
		LastName  string
		ImageURL  string
	}

	Receiver struct {
		Id uuid.UUID
	}
)

// todo: make getter setter function and make all properties private
type Notif struct {
	Id               uuid.UUID
	Seen             bool
	Event            domain.EventCode
	ReceiverId       uuid.UUID
	SenderId         uuid.NullUUID
	Data             interface{}
	IsReceiverActive bool
	CreateAt         time.Time
	channel          domain.ChannelCode
}

func NewNotif(receiverId uuid.UUID,
	data interface{},
	isReceiverActive bool,
	senderId uuid.NullUUID,
	event domain.EventCode,
) *Notif {
	return &Notif{
		Id:               uuid.New(),
		Seen:             false,
		Event:            event,
		ReceiverId:       receiverId,
		Data:             data,
		IsReceiverActive: isReceiverActive,
		SenderId:         senderId,
		CreateAt:         time.Now(),
	}
}

func (s *Notif) SetChannel(c domain.ChannelCode) {
	s.channel = c
}

func (s *Notif) Channel() domain.ChannelCode {
	switch s.Event {

	case domain.EndorsementRequestCreatedCode:
		if s.IsReceiverActive {
			return domain.InAppChannel
		}
		return domain.PushChannel

	case domain.SocialMessageCreatedCode:
		if s.IsReceiverActive {
			return domain.InAppChannel
		}
		return domain.PushChannel

	case domain.EducationCourseProgressRemindCode:
		if s.channel == domain.PushOrInAppChannel {
			if s.IsReceiverActive {
				return domain.InAppChannel
			}
			return domain.PushChannel
		}

		return s.channel

	default:
		return s.channel
	}
}

func (s *Notif) ByteData() ([]byte, error) {
	b, err := json.Marshal(s.Event)
	return b, err
}
