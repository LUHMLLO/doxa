package storages

import (
	"doxatec/types"
	"doxatec/utils"
	"fmt"

	"github.com/google/uuid"
)

func (s *PostgresStore) Query_CreateProfileTable() error {
	query := `
		create table if not exists profiles (
			ID varchar(250) primary key,
			Avatar varchar(250),
			Name varchar(250),
			Email varchar(250),
			Phone varchar(250),

			created timestamp,
			modified timestamp
		)
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) Query_CreateProfile(u *types.Profile) error {
	query := (`
		insert into profiles (
			ID, 
			Avatar,
			Name,
			Email,
			Phone,

			created,
			modified 
		)
		values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
	`)

	_, err := s.db.Query(
		query,
		u.ID,
		u.Avatar,
		u.Name,
		u.Email,
		u.Phone,

		u.Created,
		u.Modified,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) Query_ReadProfiles() ([]*types.Profile, error) {
	rows, err := s.db.Query(`select * from profiles`)
	if err != nil {
		return nil, err
	}

	profiles := []*types.Profile{}
	for rows.Next() {
		profile, err := utils.ScanIntoProfiles(rows)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (s *PostgresStore) Query_ReadProfileByID(id uuid.UUID) (*types.Profile, error) {
	rows, err := s.db.Query("select * from profiles where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return utils.ScanIntoProfiles(rows)
	}
	return nil, fmt.Errorf("profile %s not found", id)
}

func (s *PostgresStore) Query_UpdateProfile(*types.Profile) error {
	return nil
}

func (s *PostgresStore) Query_DeleteProfile(id uuid.UUID) error {
	_, err := s.db.Query("delete from Profiles where id = $1", id)
	return err
}
