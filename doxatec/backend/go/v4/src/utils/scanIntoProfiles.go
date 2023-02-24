package utils

import (
	"database/sql"
	"doxatec/types"
)

func ScanIntoProfiles(rows *sql.Rows) (*types.Profile, error) {
	profile := new(types.Profile)
	err := rows.Scan(
		&profile.ID,
		&profile.Avatar,
		&profile.Name,
		&profile.Email,
		&profile.Phone,

		&profile.Created,
		&profile.Modified,
	)

	return profile, err
}
