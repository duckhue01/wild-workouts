// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package ports

const (
	JWTScopes = "JWT.Scopes"
)

// Defines values for Event.
const (
	EventEducationCourseProgressRemind Event = "education.course.progress.remind"
	EventEndorsementRequestCreated     Event = "endorsement.request.created"
	EventSocialMessageCreated          Event = "social.message.created"
)

// EducationCourseProgressRemind defines model for EducationCourseProgressRemind.
type EducationCourseProgressRemind struct {
	// Course information about the course
	Course *struct {
		// Id id of the course
		Id int `json:"id"`

		// Title title of the course
		Title string `json:"title"`
	} `json:"course,omitempty"`

	// CreateAt time when user send message
	CreateAt string `json:"create_at"`
	Event    Event  `json:"event"`
}

// EndorsementRequestCreated defines model for EndorsementRequestCreated.
type EndorsementRequestCreated struct {
	// CreateAt time when user send message
	CreateAt string `json:"create_at"`

	// EndorsementId id of endorsement
	EndorsementId int   `json:"endorsement_id"`
	Event         Event `json:"event"`
	Sender        struct {
		// FirstName user first name
		FirstName string `json:"first_name"`

		// Id id of endorsement sender
		Id string `json:"id"`

		// ImageUrl avatar of sender
		ImageUrl string `json:"image_url"`

		// LastName user last name
		LastName string `json:"last_name"`
	} `json:"sender"`
}

// Error defines model for Error.
type Error struct {
	Slug string `json:"slug"`
}

// Event defines model for Event.
type Event string

// Notification defines model for Notification.
type Notification struct {
	// Data the data of notification. It depends on the event type
	Data *map[string]interface{} `json:"data,omitempty"`

	// Event the event id
	Event *string `json:"event,omitempty"`

	// Id the id of notification
	Id *string `json:"id,omitempty"`

	// IsSeen describe use is seen the notification or not
	IsSeen *bool `json:"is_seen,omitempty"`

	// SenderId the id of the sender
	SenderId *string `json:"sender_id,omitempty"`
}

// Ping defines model for Ping.
type Ping = string

// SocialMessageCreated defines model for SocialMessageCreated.
type SocialMessageCreated struct {
	// ChannelId the channel id that have message
	ChannelId string `json:"channel_id"`

	// CreateAt time when user send message
	CreateAt string `json:"create_at"`
	Event    Event  `json:"event"`

	// Msg message content
	Msg    string `json:"msg"`
	Sender struct {
		// FirstName user first name
		FirstName string `json:"first_name"`

		// Id id of message sender
		Id string `json:"id"`

		// ImageUrl avatar of sender
		ImageUrl string `json:"image_url"`

		// LastName user last name
		LastName string `json:"last_name"`
	} `json:"sender"`
}

// GetListUserNotificationsParams defines parameters for GetListUserNotifications.
type GetListUserNotificationsParams struct {
	// Token the next_token is returned in previous query to get the next page
	Token string `form:"token" json:"token"`

	// Limit the limitation of records will be returned
	Limit int `form:"limit" json:"limit"`
}