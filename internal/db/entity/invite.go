package entity

type Invite struct {
	Id         int
	Person     Person
	Theme      string
	Time       int
	Date       int
	Accepted   bool
	Remembered bool
}

func (i Invite) GetId() int {
	return i.Id
}

func (i Invite) Table() string {
	return "invites"
}

func (i Invite) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"person":     i.Person.GetId(),
		"theme":      i.Theme,
		"time":       i.Time,
		"date":       i.Date,
		"accepted":   i.Accepted,
		"remembered": i.Remembered,
	}
}
