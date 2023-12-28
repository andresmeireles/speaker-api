package entity

import "time"

type Invite struct {
	Id         int `db:"id"`
	Person     Person
	Theme      string    `db:"theme"`
	Time       int       `db:"time"`
	Date       time.Time `db:"date"`
	Accepted   bool      `db:"accepted"`
	Remembered bool      `db:"remembered"`
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
		"date":       i.Date.Format("2006-01-02 15:04:05"),
		"accepted":   i.Accepted,
		"remembered": i.Remembered,
	}
}
