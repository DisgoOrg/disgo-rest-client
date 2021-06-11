package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
)

type MultipartBuffer struct {
	*bytes.Buffer
}

func PayloadWithFiles(v interface{}, files ...File) (buffer *MultipartBuffer, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer func() {
		_ = writer.Close()
	}()

	var payload []byte
	payload, err = json.Marshal(v)
	if err != nil {
		return
	}

	part, err := writer.CreatePart(partHeader(`form-data; name="payload_json"`, "application/json"))
	if err != nil {
		return
	}

	if _, err = part.Write(payload); err != nil {
		return
	}

	for i, file := range files {
		var name string
		if file.Flags.Has(FileFlagSpoiler) {
			name = "SPOILER_" + file.Name
		} else {
			name = file.Name
		}
		var contentType string
		if file.ContentType == "" {
			contentType = "application/octet-stream"
		} else {
			contentType = file.ContentType
		}
		part, err = writer.CreatePart(partHeader(fmt.Sprintf(`form-data; name="file%d"; filename="%s"`, i, name), contentType))
		if err != nil {
			return
		}

		if _, err = io.Copy(part, file.Reader); err != nil {
			return
		}
	}

	buffer = &MultipartBuffer{Buffer: body}
	return
}

func partHeader(contentDisposition string, contentType string) textproto.MIMEHeader {
	return textproto.MIMEHeader{
		"Content-Disposition": []string{contentDisposition},
		"Content-Type":        []string{contentType},
	}
}

func NewFile(name string, reader io.Reader, contentType string, flags ...FileFlags) File {
	return File{
		Name:        name,
		Reader:      reader,
		ContentType: contentType,
		Flags:       FileFlagNone.Add(flags...),
	}
}

type File struct {
	Name        string
	Reader      io.Reader
	ContentType string
	Flags       FileFlags
}

type FileFlags int

const (
	FileFlagSpoiler FileFlags = 1 << iota
	FileFlagNone    FileFlags = 0
)

func (f FileFlags) Add(flags ...FileFlags) FileFlags {
	total := FileFlags(0)
	for _, flag := range flags {
		total |= flag
	}
	f |= total
	return f
}

func (f FileFlags) Remove(flags ...FileFlags) FileFlags {
	total := FileFlags(0)
	for _, flag := range flags {
		total |= flag
	}
	f &^= total
	return f
}

func (f FileFlags) HasAll(flags ...FileFlags) bool {
	for _, flag := range flags {
		if !f.Has(flag) {
			return false
		}
	}
	return true
}

func (f FileFlags) Has(flag FileFlags) bool {
	return (f & flag) == flag
}

func (f FileFlags) MissingAny(flags ...FileFlags) bool {
	for _, flag := range flags {
		if !f.Has(flag) {
			return true
		}
	}
	return false
}

func (f FileFlags) Missing(flag FileFlags) bool {
	return !f.Has(flag)
}
