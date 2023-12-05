package entity

type Invite struct {
	id         int
	person     Person
	theme      string
	time       int
	date       int
	accepted   bool
	remembered bool
}

func (i Invite) GetId() int {
	return i.id
}

func (i Invite) Table() string {
	return "invites"
}
