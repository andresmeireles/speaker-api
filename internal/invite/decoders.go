package invite

type InvitePost struct {
	PersonId   int    `json:"person_id"`
	Theme      string `json:"theme"`
	Time       int    `json:"time"`
	Date       string `json:"date"`
	References string `json:"references"`
}

type UpdateInviteData struct {
	Theme      string `json:"theme"`
	Time       int    `json:"time"`
	Date       string `json:"date"`
	References string `json:"references"`
}

type WasDone struct {
	Done bool `json:"done"`
}
