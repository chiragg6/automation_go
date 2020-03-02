package storage

import (
	"database/sql"
	"log"
)

func GetScript(alertType string) (string, error) {
	var script string
	err := db.QueryRow(`SELECT script FROM mapping WHERE alert_type = $1`,
		alertType).Scan(&script)

	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return script, err
}

func InsertScript(alertType, script string) error {
	_, err := GetScript(alertType)
	if err != nil {
		sqlStatement := `INSERT INTO mapping (alert_type, script)
						VALUES ($1, $2)`
		_, err = db.Exec(sqlStatement, alertType, script)
		return err

	}

	return err
}

func GetMappings() ([]*Mapping, error) {
	maps := make([]*Mapping, 0)

	rows, err := db.Query(
		`SELECT alert_type, script FROM mapping`)
	if err != nil {
		return maps, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			alertType string
			script    string
		)
		if err := rows.Scan(&alertType, &script); err != nil {
			return maps, err
		}

		maps = append(maps, &Mapping{alertType, script})
	}

	return maps, nil
}

func DeleteMapping(alertType string) error {
	_, err := db.Exec("DELETE FROM mapping WHERE alert_type = S1", alertType)
	if err != nil {
		log.Println(err)
	}

	return err
}

type Mapping struct {
	AlertType string
	Script    string
}
