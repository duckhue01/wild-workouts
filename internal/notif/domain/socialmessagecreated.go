package domain

import (
	"github.com/google/uuid"
)

type Sender struct {
	FirstName string
	Id        uuid.UUID
	ImageURL  string
	LastName  string
}

type SocialMessageCreated struct {
	Sender      Sender
	Msg         string
	ReceiverIds *[]uuid.UUID
	CreateAt    string
	ChannelId   string
}

func NewSocialMessageCreated(sender Sender, msg string, receiverIds *[]uuid.UUID, createAt string, channelId string) *SocialMessageCreated {
	return &SocialMessageCreated{
		Sender:      sender,
		Msg:         msg,
		ReceiverIds: receiverIds,
		CreateAt:    createAt,
		ChannelId:   channelId,
	}
}

func (s *SocialMessageCreated) Event() EventCode {
	return SocialMessageCreatedCode
}
