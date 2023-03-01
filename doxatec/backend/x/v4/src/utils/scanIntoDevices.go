package utils

import (
	"database/sql"
	"doxatec/types"
)

func ScanIntoDevices(rows *sql.Rows) (*types.Device, error) {
	device := new(types.Device)
	err := rows.Scan(
		&device.ID,
		&device.Name,
		&device.TempSup,
		&device.TempMid,
		&device.TempSub,

		&device.Created,
		&device.Modified,
	)

	return device, err
}
