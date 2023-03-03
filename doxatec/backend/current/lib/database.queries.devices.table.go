package lib

func (s *Database) Query_tableDevices() error {
	query := `
		create table if not exists devices (
			id varchar(250) primary key,
			owner varchar(250),
			name varchar(250),
			tempsup decimal,
			tempmid decimal,
			tempsub decimal,
			created timestamp,
			modified timestamp
		)
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *Database) Query_allDevices() ([]*Device, error) {
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
