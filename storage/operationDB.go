package storage

import (
	"log"
)

func InsertOperation(alertID int, applicationID, script, status string) int {
	lastInsertID := 0
	sqlStatement := `INSERT INTO operations (application_id, script, status, alert_id)
	VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(sqlStatement, applicationID, script, status, alertID).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}

	return lastInsertID

}

func GetOperation(applicationID string) ([]*Operation, error) {
	ops := make([]*Operation, 0)
	rows, err := db.Query(
		`SELECT id, script FROM operations WHERE application_id = $1 and status = $2`,
		applicationID, "CRITICAL")
	if err != nil {
		return ops, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id     int
			script string
		)
		if err := rows.Scan(&id, &script); err != nil {
			return ops, err
		}
		ops = append(ops, &Operation{id, script})
	}

	return ops, nil

}

type Operation struct {
	ID     int    `json:"id"`
	Script string `json:"script"`
}

func RecordStatus(id int, status, logs string) error {

	sqlStatement := `UPDATE operations SET status = $2, logs = $3 WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id, status, logs)
	if err != nil {
		return err
	}
	return nil
}

func AlertNameFromOpID(id int) (string, error) {
	var alertID int
	err := db.QueryRow(`SELECT alert_id FROM operations WHERE id = $1`, id).Scan(&alertID)

	if err != nil {
		// panic(err)
		return "", err
	}
	alertName, err := GetAlertName(alertID)

	return alertName, err
}

func GetApplicationID(id int) (string, error) {
	var applicationID string

	err := db.QueryRow(`SELECT application_id FROM operations WHERE id = $1`,
		id).Scan(&applicationID)

	return applicationID, err
}

func GetLogs() ([]*Logs, error) {
	logs := make([]*Logs, 0)

	rows, err := db.Query {
		`SELECT  id, application_id, script, status, logs FROM operations ORDER BY id DESC`)
	if err != nil {
		return logs, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id string
			applicationID string
			script string
			status string
			log string
		)

		if err := rows.Scan(&id, &application_id, &script, &status, &log); err != nil {
			return logs, error
		}
		if size := len(log); size > 50 {
			log = "..." + log[size-50:]
		}
		logs = append(logs, &Logs{id, applicationID, script, status, log})




	}
	return logs, nil
}

func GetOneLog(id string) (string, error) {
	var logs string
	err := db.QueryRow(
		`SELECT logs FROM operations where id = $1`, id).Scan(&logs)

	return logs, err
}

func GetSummary() (map[string]*Summary, error) {
	summaries := make(map[string](*Summary))

	logs, err := GetLogs()
	if err != nil {
		return summaries, err
	}

	for _,log := range logs {
		if _, ok := summaries[log.applicationID]; !ok {
			summaries[log.applicationID] = &Summary{log.applicationID, 0,0,0}
		}
		if log.Status == "completed" {
			summaries[log.ApplicationID].Success++
		} else if log.Status == "failed" {
			summaries[log.ApplicationID].Fail++
		} else {
			summaries[log.ApplicationID].Firing++
		}

		
	}
	return summaries, nil

}

// GetOneSummary return summary of one application id
func GetOneSummary(applicationID string) (Summary, []*Logs, error) {
	logs := make([]*Logs, 0)
	summary := Summary{applicationID, 0, 0, 0}

	rows, err := db.Query(
		`SELECT  id, script, status, logs FROM operations where application_id = $1 ORDER BY id DESC`,
		applicationID)
	if err != nil {
		return summary, logs, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id     string
			script string
			status string
			log    string
		)
		if err := rows.Scan(&id, &script, &status, &log); err != nil {
			return summary, logs, err
		}

		if size := len(log); size > 50 {
			log = "..." + log[size-50:]
		}
		logs = append(logs, &Logs{id, applicationID, script, status, log})

		if status == "completed" {
			summary.Success++
		} else if status == "failed" {
			summary.Fail++
		} else {
			summary.Firing++
		}
	}

	return summary, logs, nil
}


type Logs struct {
	ID string
	ApplicationID string
	Script string
	Status string
	Logs string
}

type Summarty struct {
	ApplicationID string
	Success int
	Fail int
	Firing int
}
