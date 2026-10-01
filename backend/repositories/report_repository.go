package repositories

import (
	"backend/db"
	"backend/models"
	"database/sql"
	"encoding/json"
	"errors"
)

func CreateReport(report *models.Report, definition *models.ReportDefinition) error {
	if db.DB == nil {
		return errors.New("database not connected")
	}

	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO reports (tenant_id, name, description, type, status, dataset_id, current_version)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`
	err = tx.QueryRowx(query, report.TenantID, report.Name, report.Description, report.Type, report.Status, report.DatasetID, 1).StructScan(report)
	if err != nil {
		return err
	}

	defJSON, _ := json.Marshal(definition)
	verQuery := `
		INSERT INTO report_versions (report_id, version, definition_json, created_by)
		VALUES ($1, 1, $2, 'system')
	`
	_, err = tx.Exec(verQuery, report.ID, defJSON)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func GetReports() ([]models.Report, error) {
	if db.DB == nil {
		return []models.Report{}, nil
	}
	var reports []models.Report
	err := db.DB.Select(&reports, "SELECT * FROM reports ORDER BY updated_at DESC")
	return reports, err
}

func GetReport(id string) (*models.Report, error) {
	if db.DB == nil {
		return nil, errors.New("database not connected")
	}

	var report models.Report
	err := db.DB.Get(&report, "SELECT * FROM reports WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	var version models.ReportVersion
	err = db.DB.Get(&version, "SELECT * FROM report_versions WHERE report_id = $1 AND version = $2", id, report.CurrentVersion)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == nil {
		var def models.ReportDefinition
		json.Unmarshal(version.DefinitionJSON, &def)
		report.Definition = &def
	}

	return &report, nil
}

func UpdateReport(id string, report *models.Report, definition *models.ReportDefinition) error {
	if db.DB == nil {
		return errors.New("database not connected")
	}

	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Increment version
	var currentVersion int
	err = tx.Get(&currentVersion, "SELECT current_version FROM reports WHERE id = $1", id)
	if err != nil {
		return err
	}
	newVersion := currentVersion + 1

	query := `
		UPDATE reports
		SET name = $1, description = $2, type = $3, status = $4, dataset_id = $5, current_version = $6, updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
	`
	_, err = tx.Exec(query, report.Name, report.Description, report.Type, report.Status, report.DatasetID, newVersion, id)
	if err != nil {
		return err
	}

	defJSON, _ := json.Marshal(definition)
	verQuery := `
		INSERT INTO report_versions (report_id, version, definition_json, created_by)
		VALUES ($1, $2, $3, 'system')
	`
	_, err = tx.Exec(verQuery, id, newVersion, defJSON)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func DeleteReport(id string) error {
	if db.DB == nil {
		return errors.New("database not connected")
	}
	_, err := db.DB.Exec("DELETE FROM reports WHERE id = $1", id)
	return err
}

func CloneReport(id string) (*models.Report, error) {
	if db.DB == nil {
		return nil, errors.New("database not connected")
	}
	original, err := GetReport(id)
	if err != nil {
		return nil, err
	}

	newReport := &models.Report{
		TenantID:    original.TenantID,
		Name:        original.Name + " (Clone)",
		Description: original.Description,
		Type:        original.Type,
		Status:      "DRAFT",
		DatasetID:   original.DatasetID,
	}

	err = CreateReport(newReport, original.Definition)
	if err != nil {
		return nil, err
	}

	return GetReport(newReport.ID)
}
