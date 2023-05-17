package helpers

type Email struct {
	ID    int64  `bun:",pk,autoincrement"`
	Email string `json:"email" form:"email" validate:"required,email,max=254"`
}

type OutgoingEmail struct {
	ID        int64  `bun:",pk,autoincrement"`
	Sender    string `json:"sender" form:"sender" validate:"required,email,max=254"`
	Recipient string `json:"recipient" form:"recipient" validate:"required,email,max=254"`
	Subject   string `json:"subject" form:"subject" validate:"required, max=256"`
	HtmlBody  string `json:"htmlbody" form:"htmlbody" validate:"required"`
	TextBody  string `json:"textbody" form:"textbody" validate:"required"`
	CharSet   string `json:"charset" form:"charset" validate:"required"`
}

type Message struct {
	Message string
}
