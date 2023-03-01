package lib

import (
	"github.com/google/uuid"
)

func (s *Database) Devices_CreateTable() error {
	query := `
		create table if not exists devices (
			id varchar(250) primary key,
			name varchar(250),
			owner varchar(250),
			tempsup varchar(250),
			tempmid varchar(250),
			tempsub varchar(250),
			created timestamp,
			modified timestamp
		)
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *Database) Devices_ReadFromTable() ([]*Device, error) {
	query := `select * from devices`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	devices := []*Device{}
	for rows.Next() {
		device := &Device{}

		err := rows.Scan(&device.ID, &device.Owner, &device.Name, &device.TempSup, &device.TempMid, &device.TempSub, &device.Created, &device.Modified)
		if err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	return devices, err
}

func (s *Database) Devices_InsertToTable(d *Device) error {
	query := `
		insert into devices (
			id,
			owner,
			name,
			tempsup,
			tempmid,
			tempsub,
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
	`

	_, err := s.db.Query(query, &d.ID, &d.Owner, &d.Name, &d.TempSup, &d.TempMid, &d.TempSub, &d.Created, &d.Modified)
	if err != nil {
		return err
	}

	return nil
}

func (s *Database) Devices_ReadFromTableByID(id uuid.UUID) (*Device, error) {
	query := `select * from devices where id = $1`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	device := &Device{}
	for rows.Next() {
		err := rows.Scan(&device.ID, &device.Owner, &device.Name, &device.TempSup, &device.TempMid, &device.TempSub, &device.Created, &device.Modified)
		if err != nil {
			return nil, err
		}
	}

	return device, err
}

func (s *Database) Devices_UpdateFromTableByID(id uuid.UUID, d *Device) error {
	query := `update devices set owner=$2, name=$3, tempsup=$4, tempmid=$5, tempsub=$6, modified=$7 where id = $1`

	_, err := s.db.Exec(query, id, &d.Owner, &d.Name, &d.TempSup, &d.TempMid, &d.TempSub, &d.Modified)
	if err != nil {
		return err
	}

	return nil
}

func (s *Database) Devices_DeleteFromTableByID(id uuid.UUID) (uuid.UUID, error) {
	query := `delete from devices where id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return id, err
	}

	return id, err
}
