package sesagent

import (
	"bytes"
	"encoding/base64"
	"mime/multipart"
	"net/textproto"
	"strings"
)

type RAWEMAIL struct {
	Subject string
	Message string

	Attachments []ATTACHMENT
}

type ATTACHMENT struct {
	FileName    string
	FileContent []byte // base64 format
	ContentType string // ex : image/jpeg, text/csv, application/pdf
}

func (r *RAWEMAIL) BuildEmail() []byte {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	r.SetMainHeader(writer)
	r.SetBody(writer)
	r.SetAttachment(writer)
	err := writer.Close()
	if err != nil {
		panic(err)
	}

	// Strip boundary line before header (doesn't work with it present)
	s := buf.String()
	if strings.Count(s, "\n") < 2 {
		panic("invalid e-mail content")
	}
	s = strings.SplitN(s, "\n", 2)[1]
	return []byte(s)
}

func (r *RAWEMAIL) SetMainHeader(writer *multipart.Writer) {
	// email main header:
	h := make(textproto.MIMEHeader)
	/* // don't need
	h.Set("From", source)
	h.Set("To", destination)
	h.Set("Return-Path", source)
	*/
	h.Set("Subject", r.Subject)
	h.Set("Content-Language", "en-US")
	h.Set("Content-Type", "multipart/mixed; boundary=\""+writer.Boundary()+"\"")
	h.Set("MIME-Version", "1.0")
	_, err := writer.CreatePart(h)
	if err != nil {
		panic(err)
	}
}

func (r *RAWEMAIL) SetBody(writer *multipart.Writer) {

	h := make(textproto.MIMEHeader)
	h.Set("Content-Transfer-Encoding", "7bit")
	h.Set("Content-Type", "text/plain; charset=us-ascii")
	part, err := writer.CreatePart(h)
	if err != nil {
		panic(err)
	}
	_, err = part.Write([]byte(r.Message))
	if err != nil {
		panic(err)
	}
}

func (r *RAWEMAIL) SetAttachment(writer *multipart.Writer) {

	if len(r.Attachments) == 0 {
		return
	}
	for _, attachment := range r.Attachments {
		func(obj []byte) {
			_, err := base64.StdEncoding.DecodeString(string(obj))
			if err != nil {
				panic("ATTACHMENT.FileContent not base64 format")
			}
		}(attachment.FileContent)

		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", "attachment")
		h.Set("Content-Type", attachment.ContentType+"; name=\""+attachment.FileName+"\"")
		h.Set("Content-Transfer-Encoding", "base64")
		part, err := writer.CreatePart(h)
		if err != nil {
			panic(err)
		}
		_, err = part.Write(attachment.FileContent)
		if err != nil {
			panic(err)
		}
	}
}
