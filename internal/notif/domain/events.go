package domain

type EventCode string

type ChannelCode string

const (
	SocialMessageCreatedCode          EventCode = "social.message.created"
	EndorsementRequestCreatedCode     EventCode = "endorsement.request.created"
	EducationCourseProgressRemindCode EventCode = "education.course.progress.created"
)

const (
	EmailChannel       ChannelCode = "email"
	PushChannel        ChannelCode = "push"
	InAppChannel       ChannelCode = "in-app"
	PushOrInAppChannel ChannelCode = "push-or-in-app"
	UndefinedChannel   ChannelCode = "undefined"
)

type Event interface {
	Event() EventCode
}
