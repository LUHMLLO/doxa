package utils

import (
	"fmt"
	"strings"
)

func StringToQuery(slice []string) string {
	lastKey := slice[len(slice)-1]

	var keys strings.Builder

	for _, value := range slice {
		if value == lastKey {
			keys.WriteString(fmt.Sprintf("%s\n", value))
		} else {
			keys.WriteString(fmt.Sprintf("%s,\n", value))
		}
	}

	return keys.String()
}
