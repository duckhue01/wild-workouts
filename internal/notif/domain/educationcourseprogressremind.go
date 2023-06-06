package domain

import (
	"errors"

	"github.com/google/uuid"
)

type EducationCourseProgressRemind struct {
	ReceiverId uuid.UUID
	CreatedAt  string
	Course     Course
	NoT        int8
}

type Course struct {
	Id    int
	Title string
}

func NewEducationCourseProgressRemind(course Course,
	receiverId uuid.UUID,
	createAt string,
	not int8,
) (*EducationCourseProgressRemind, error) {

	if not < 1 && not > 2 {
		return nil, errors.New("invalid number of times")
	}

	return &EducationCourseProgressRemind{
		ReceiverId: receiverId,
		CreatedAt:  createAt,
		Course:     course,
		NoT:        not,
	}, nil
}

func (s *EducationCourseProgressRemind) Channel() ChannelCode {
	if s.NoT == 1 {
		return PushOrInAppChannel
	}
	return EmailChannel
}

func (s *EducationCourseProgressRemind) Event() EventCode {
	return EducationCourseProgressRemindCode
}
