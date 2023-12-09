package invite

type InvitePost struct {
	PersonId int    `json:"person_id"`
	Theme    string `json:"theme"`
	Time     int    `json:"time"`
	Date     string `json:"date"`
}

type InviteSender struct {
	InvoiceId  int
	TemplateId int
}
