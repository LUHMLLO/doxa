package storages

import (
	"doxatec/types"
	"doxatec/utils"
	"fmt"

	"github.com/google/uuid"
)

func (s *PostgresStore) Query_CreateDeviceTable() error {
	query := `
		create table if not exists devices (
			ID varchar(250) primary key,
			Name varchar(250),
			TempSup varchar(250),
			TempMid varchar(250),
			TempSub varchar(250),

			Owner varchar(250),

			created timestamp,
			modified timestamp
		)
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) Query_CreateDevice(u *types.Device) error {
	query := (`
		insert into devices (
			ID, 
			Name,
			TempSup,
			TempMid,
			TempSub,

			Owner,

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
			$7,
			$8
		)
	`)

	_, err := s.db.Query(
		query,
		u.ID,
		u.Name,
		u.TempSup,
		u.TempMid,
		u.TempSub,

		u.Owner,

		u.Created,
		u.Modified,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) Query_ReadDevices() ([]*types.Device, error) {
	rows, err := s.db.Query(`select * from devices`)
	if err != nil {
		return nil, err
	}

	devices := []*types.Device{}
	for rows.Next() {
		device, err := utils.ScanIntoDevices(rows)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}

func (s *PostgresStore) Query_ReadDeviceByID(id uuid.UUID) (*types.Device, error) {
	rows, err := s.db.Query("select * from devices where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return utils.ScanIntoDevices(rows)
	}
	return nil, fmt.Errorf("device %s not found", id)
}

func (s *PostgresStore) Query_UpdateDevice(*types.Device) error {
	return nil
}

func (s *PostgresStore) Query_DeleteDevice(id uuid.UUID) error {
	_, err := s.db.Query("delete from devices where id = $1", id)
	return err
}
