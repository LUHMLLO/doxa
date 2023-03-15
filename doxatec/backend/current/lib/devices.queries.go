package lib

import (
	"fmt"

	"github.com/google/uuid"
)

func (s *Database) devices_init() error {
	schema := []string{
		"id varchar(250) primary key",
		"owner varchar(250)",
		"name varchar(250)",
		"pin varchar(250)",
		"tempsup decimal",
		"tempmid decimal",
		"tempsub decimal",
		"created timestamp",
		"modified timestamp",
	}

	query := fmt.Sprintf(`create table if not exists devices (%s)`, StringToQuery(schema))

	_, err := s.db.Exec(query)
	return err
}

func (s *Database) devices_beforeInsert(d *Device) (*Device, error) {
	query := `select * from devices where (owner= $1 OR name=$2)`

	rows, err := s.db.Query(query, d.Owner, d.Name)
	if err != nil {
		return nil, err
	}

	device := &Device{}

	for rows.Next() {
		if err := rows.Scan(
			&device.ID,
			&device.Owner,
			&device.Name,
			&device.PIN,
			&device.TempSup,
			&device.TempMid,
			&device.TempSub,
			&device.Created,
			&device.Modified,
		); err != nil {
			return nil, err
		}

		if device.Owner == d.Owner && device.Name == d.Name {
			return nil, fmt.Errorf("device name already in use")
		}
	}

	return device, nil
}

func (s *Database) devices_insert(d *Device) error {
	schema := []string{
		"id",
		"owner",
		"name",
		"pin",
		"tempsup",
		"tempmid",
		"tempsub",
		"created",
		"modified",
	}

	cols := []string{}
	for i := 0; i < len(schema); i++ {
		cols = append(cols, fmt.Sprintf("$%d", i+1))
	}

	query := fmt.Sprintf(`insert into devices (%s) values (%s)`, StringToQuery(schema), StringToQuery(cols))

	if _, err := s.db.Query(
		query,
		&d.ID,
		&d.Owner,
		&d.Name,
		&d.PIN,
		&d.TempSup,
		&d.TempMid,
		&d.TempSub,
		&d.Created,
		&d.Modified,
	); err != nil {
		return err
	}

	return nil
}

func (s *Database) devices_readTable() ([]*Device, error) {
	query := `select * from devices`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	devices := []*Device{}

	for rows.Next() {
		device := &Device{}

		if err := rows.Scan(
			&device.ID,
			&device.Owner,
			&device.Name,
			&device.PIN,
			&device.TempSup,
			&device.TempMid,
			&device.TempSub,
			&device.Created,
			&device.Modified,
		); err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	return devices, err
}
func (s *Database) devices_readTableWhereOwner(owner string) ([]*Device, error) {
	query := `select * from devices where owner=$1`

	rows, err := s.db.Query(query, owner)
	if err != nil {
		return nil, err
	}

	devices := []*Device{}

	for rows.Next() {
		device := &Device{}

		if err := rows.Scan(
			&device.ID,
			&device.Owner,
			&device.Name,
			&device.PIN,
			&device.TempSup,
			&device.TempMid,
			&device.TempSub,
			&device.Created,
			&device.Modified,
		); err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	return devices, err
}

func (s *Database) devices_read(id uuid.UUID) (*Device, error) {
	query := `select * from devices where id=$1`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	device := &Device{}

	for rows.Next() {
		if err := rows.Scan(
			&device.ID,
			&device.Owner,
			&device.Name,
			&device.PIN,
			&device.TempSup,
			&device.TempMid,
			&device.TempSub,
			&device.Created,
			&device.Modified,
		); err != nil {
			return nil, err
		}
	}

	return device, nil
}

func (s *Database) devices_readCol(column string, value any) error {
	query := fmt.Sprintf(`select * from devices where %s=$1`, column)

	_, err := s.db.Exec(query, value)
	if err != nil {
		return err
	}

	return nil
}

func (s *Database) devices_update(id uuid.UUID, column string, value any) error {
	query := fmt.Sprintf(`update devices set %s=$2 where id=$1`, column)

	_, err := s.db.Exec(query, id, value)
	if err != nil {
		return err
	}

	return nil
}

func (s *Database) devices_delete(id uuid.UUID) error {
	query := `delete from devices where id=$1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
