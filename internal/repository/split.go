package repository

import (
	"fmt"
	"strconv"

	"github.com/andresmeireles/speaker/internal/db"
)

// Splits the entity into keys, interrogations and values.
//
// Keys are comma separated values with entity table column name.
//
// Interrogations are comma separated values with "?" equivalent to keys.
//
// Values are the values of the entity.
func (r Repository) split(en db.Entity) (string, string, []any) {
	keys := ""
	interrogations := ""
	values := make([]any, 0)
	index := 0

	for key, value := range en.ToJson() {
		index++

		keys = keys + "\"" + key + "\"" + ","
		interrogations = interrogations + "$" + strconv.Itoa(index) + ","

		values = append(values, fmt.Sprintf("%v", value))
	}

	keys = keys[:len(keys)-1]
	interrogations = interrogations[:len(interrogations)-1]

	return keys, interrogations, values
}
