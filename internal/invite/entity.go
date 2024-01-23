package invite

import (
	"time"

	"github.com/andresmeireles/speaker/internal/person"
)

const (
	STATUS_WAIT_CONFIRMATION = iota
	STATUS_CONFIRMED
	STATUS_REJECTED
	STATUS_WAIT_REMEMBER
	STATUS_REMEMBERED
	STATUS_DONE
	STATUS_NOT_DONE
)

// TODO: remove accepted and remembered
type Invite struct {
	Id         int           `db:"id" json:"id"`
	PersonId   int           `db:"person_id" json:"person_id"`
	Person     person.Person `db:"-" json:"person"`
	Theme      string        `db:"theme" json:"theme"`
	Time       int           `db:"time" json:"time"`
	Date       time.Time     `db:"date" json:"date"`
	References string        `db:"references" json:"references"`
	Status     int           `db:"status" json:"status"`
}

func (i Invite) StatusList() []int {
	return []int{
		STATUS_WAIT_CONFIRMATION,
		STATUS_CONFIRMED,
		STATUS_REJECTED,
		STATUS_WAIT_REMEMBER,
		STATUS_REMEMBERED,
		STATUS_DONE,
		STATUS_NOT_DONE,
	}
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
		"status":     i.Status,
	}
}
