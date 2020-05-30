package main

import (
	"app/sesagent"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type EMAIL struct {
	From string   `json:"from"`
	To   []string `json:"to"`
}

func main() {
	var email *EMAIL
	b, _ := ioutil.ReadFile("./email.json")
	json.Unmarshal(b, &email)

	if email == nil {
		panic("You must set email.json")
	}

	//
	fmt.Println(`
	Select 1 = Simple, 2 = Raw
	【Simple】: Just send a email
	【Raw】: Attached jpg to email
	`)
	fmt.Print("Your Input : ")
	var EMAIL_TYPE int
	fmt.Scanf("%v", &EMAIL_TYPE)

	if EMAIL_TYPE == 1 {
		Simple_Email(email.From, email.To)
	} else if EMAIL_TYPE == 2 {
		Raw_Email(email.From, email.To)
	} else {
		fmt.Println("Do nothing")
	}

	//
}

func Simple_Email(from string, to []string) {
	var (
		// Replace sender@example.com with your "From" address.
		// This address must be verified with Amazon SES.
		From = from

		// Replace recipient@example.com with a "To" address. If your account
		// is still in the sandbox, this address must be verified.
		To = to

		// Specify a configuration set. To use a configuration
		// set, comment the next line and line 92.
		//ConfigurationSet = "ConfigSet"

		// The subject line for the email.
		Subject = "Amazon SES Test (AWS SDK for Go)"

		// The HTML body for the email.
		HtmlBody = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
			"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
			"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

		//The email body for recipients with non-HTML email clients.
		TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."

		// The character encoding for the email.
		CharSet = "UTF-8"
	)

	agent := sesagent.New_SES_SIMPLE_AGENT(From, To)
	agent.SetEmail(Subject, HtmlBody, TextBody, CharSet)
	agent.SendSimpleEmail()
}

func Raw_Email(from string, to []string) {
	//
	email := func() (email_data []byte) {
		//
		attachment, err := ioutil.ReadFile("./aws-ses-logo.jpg")
		if err != nil {
			panic(err)
		}
		RawEmail := sesagent.RAWEMAIL{
			Subject: "Test Email Subject",
			Message: "Hello,\r\nPlease see the attached file for a list of customers to contact.",
		}
		RawEmail.Attachments = []sesagent.ATTACHMENT{
			0: {
				FileName:    "SES_1.jpg",
				FileContent: []byte(base64.StdEncoding.EncodeToString([]byte(attachment))),
				ContentType: "image/jpeg",
			},
			1: {
				FileName:    "SES_2.jpg",
				FileContent: []byte(base64.StdEncoding.EncodeToString([]byte(attachment))),
				ContentType: "image/jpeg",
			},
		}
		email_data = RawEmail.BuildEmail()
		return
	}()

	//
	ses_raw_agent := sesagent.New_SES_RAW_AGENT(from, to)
	SendEmailOutput, err := ses_raw_agent.SendRawEmial(email)
	if err != nil {
		panic(err)
	}
	fmt.Println(SendEmailOutput)
}
