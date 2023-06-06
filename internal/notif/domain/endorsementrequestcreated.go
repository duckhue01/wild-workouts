package domain

import "github.com/google/uuid"

type EndorsementRequestCreated struct {
	Sender        *Sender
	EndorsementId int
	ReceiverId    uuid.UUID
	CreatedAt     string
}

func NewEndorsementRequestCreated(sender *Sender, endorsementId int, receiverId uuid.UUID, createAt string) *EndorsementRequestCreated {
	return &EndorsementRequestCreated{
		Sender:        sender,
		EndorsementId: endorsementId,
		ReceiverId:    receiverId,
		CreatedAt:     createAt,
	}
}

func (s *EndorsementRequestCreated) Event() EventCode {
	return EndorsementRequestCreatedCode
}
