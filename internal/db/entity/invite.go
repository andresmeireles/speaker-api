package entity

import "time"

type Invite struct {
	Id         int       `db:"id" json:"id"`
	PersonId   int       `db:"person_id" json:"person_id"`
	Person     Person    `db:"-" json:"person"`
	Theme      string    `db:"theme" json:"theme"`
	Time       int       `db:"time" json:"time"`
	Date       time.Time `db:"date" json:"date"`
	References string    `db:"references" json:"references"`
	Accepted   bool      `db:"accepted" json:"accepted"`
	Remembered bool      `db:"remembered" json:"remembered"`
}

func (i Invite) GetId() int {
	return i.Id
}

func (i Invite) Table() string {
	return "invites"
}

func (i Invite) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"person_id":  i.PersonId,
		"theme":      i.Theme,
		"time":       i.Time,
		"date":       i.Date.Format("2006-01-02 15:04:05"),
		"references": i.References,
		"accepted":   i.Accepted,
		"remembered": i.Remembered,
	}
}
