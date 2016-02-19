package commitmessage

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const (
	//WhatTheCommitSiteURL - the url to whatthecommit
	WhatTheCommitSiteURL = "http://whatthecommit.com"
)

//MessageBody - the message body structure
type MessageBody struct {
	Message string `json:"message"`
}

//fetchCommitMessage - fetches a commit message from one source
func fetchCommitMessage() (string, error) {
	var message string
	doc, err := goquery.NewDocument(WhatTheCommitSiteURL)
	if err != nil {
		return message, err
	}
	doc.Find("body #content p:nth-child(1)").Each(func(i int, s *goquery.Selection) {
		message = s.Text()
	})
	return message, nil
}

//NewMessage - creates a new message structure
func NewMessage() (MessageBody, error) {
	var messageBody MessageBody
	msg, err := fetchCommitMessage()
	if err != nil {
		return messageBody, fmt.Errorf("NewMessage error: %s", err)
	}
	return MessageBody{Message: msg}, nil
}

//ToString - message body just as a string
func (m *MessageBody) ToString() (string, error) {
	return m.Message, nil
}

//ToJSON - message body output in json
func (m *MessageBody) ToJSON() ([]byte, error) {
	js, err := json.Marshal(m)
	if err != nil {
		err = fmt.Errorf("ToJSON error: %s", err)
	}
	return js, err
}

//ToXML - message body output in xml
func (m *MessageBody) ToXML() ([]byte, error) {
	xm, err := xml.Marshal(m)
	if err != nil {
		err = fmt.Errorf("ToXML error: %s", err)
	}
	return xm, err
}
