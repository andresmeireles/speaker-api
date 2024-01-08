package invite

type InvitePost struct {
	PersonId   int    `json:"person_id"`
	Theme      string `json:"theme"`
	Time       int    `json:"time"`
	Date       string `json:"date"`
	References string `json:"references"`
	Accepted   bool   `json:"accepted"`
	Remembered bool   `json:"remembered"`
}

type UpdateInviteData struct {
	Theme      string `json:"theme"`
	Time       int    `json:"time"`
	Date       string `json:"date"`
	References string `json:"references"`
	Accepted   bool   `json:"accepted"`
	Remembered bool   `json:"remembered"`
}
