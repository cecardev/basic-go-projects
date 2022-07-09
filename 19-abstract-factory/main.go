package main

import "fmt"

//Notifications -> SMS, Push notificaction, Email

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChanell() string
}

//SMS Notification
type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sendint via SMS")
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct{}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChanell() string {
	return "Twilio"
}

//Email notification
type EmailNotification struct{}

func (EmailNotification) SendNotification() {
	fmt.Println("Sendint via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct{}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChanell() string {
	return "Gmail"
}

//Notificaction factory
func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case ("SMS"):
		return SMSNotification{}, nil
	case ("EMAIL"):
		return EmailNotification{}, nil
	default:
		return nil, fmt.Errorf("Notification type not supported")
	}
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("EMAIL")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
