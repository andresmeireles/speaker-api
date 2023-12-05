package repository

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/database/entity"
)

// Splits the entity into keys, interrogations and values.
//
// Keys are comma separated values with entity table column name.
//
// Interrogations are comma separated values with "?" equivalent to keys.
//
// Values are the values of the entity.
func Split[T entity.Entity](en T) (string, string, []interface{}) {
	keys := ""
	interrogations := ""
	values := []interface{}{}

	for key, value := range en.ToJson() {
		keys = keys + key + ","
		interrogations = interrogations + "?" + ","
		values = append(values, fmt.Sprintf("%v", value))
	}

	keys = keys[:len(keys)-1]
	interrogations = interrogations[:len(interrogations)-1]

	return keys, interrogations, values
}
