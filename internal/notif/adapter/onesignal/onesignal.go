package onesignal

import (
	"context"

	onesignal "github.com/OneSignal/onesignal-go-api"
	"github.com/sirupsen/logrus"
	"github.com/tribefintech/microservices/internal/notif/domain"
)

type OneSignal struct {
	c      *onesignal.APIClient
	logger *logrus.Entry
	apiKey string
	appId  string
}

var template = map[domain.EventCode]string{
	domain.SocialMessageCreatedCode:          "c5c9891f-ac84-405c-b49a-008509a10a1e",
	domain.EndorsementRequestCreatedCode:     "b18dfacc-3559-400e-8cd6-49cd5d87ab36",
	domain.EducationCourseProgressRemindCode: "292b2f3c-cd52-4526-81cd-6611a4bb1b4d",
}

func NewOneSignal(apiKey string, appId string, logger *logrus.Entry) *OneSignal {
	return &OneSignal{
		c: onesignal.NewAPIClient(onesignal.NewConfiguration()),
		logger: logger.WithFields(logrus.Fields{
			"adapter": "one-signal",
		}),
		apiKey: apiKey,
		appId:  appId,
	}
}

func (o *OneSignal) Push(ecode domain.EventCode, userId string) {
	notif := *onesignal.NewNotification(o.appId)
	notif.SetIncludeExternalUserIds([]string{userId})
	notif.SetIsIos(false)

	if _, ok := template[ecode]; !ok {
		o.logger.Errorf("template id not found")
		return
	}

	notif.SetTemplateId(template[ecode])

	appAuth := context.WithValue(context.Background(), onesignal.AppAuth, o.apiKey)
	resp, r, err := o.c.DefaultApi.CreateNotification(appAuth).Notification(notif).Execute()
	if err != nil {
		o.logger.Errorf("can not push: %+v: %+v", err, r)
		return
	}

	if resp.Errors != nil {
		o.logger.Warnf("one signal error: %+v", resp.Errors.ArrayOfString)
		o.logger.Warnf("invalid identity: %+v", resp.Errors.InvalidIdentifierError)
		return
	}

	o.logger.Infof("sent to one signal: %+v", resp)
}
