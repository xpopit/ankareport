package models

import (
	"encoding/json"
	"time"
)

type Report struct {
	ID             string    `json:"id" db:"id"`
	TenantID       string    `json:"tenantId" db:"tenant_id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	Type           string    `json:"type" db:"type"` // DASHBOARD, PIXEL_REPORT, DETAIL_REPORT, HYBRID
	Status         string    `json:"status" db:"status"`
	DatasetID      string    `json:"datasetId" db:"dataset_id"`
	CurrentVersion int       `json:"currentVersion" db:"current_version"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`

	Definition *ReportDefinition `json:"definition,omitempty" db:"-"`
}

type ReportVersion struct {
	ID             string          `json:"id" db:"id"`
	ReportID       string          `json:"reportId" db:"report_id"`
	Version        int             `json:"version" db:"version"`
	DefinitionJSON json.RawMessage `json:"definitionJson" db:"definition_json"`
	CreatedBy      string          `json:"createdBy" db:"created_by"`
	CreatedAt      time.Time       `json:"createdAt" db:"created_at"`
}

type ReportDefinition struct {
	SchemaVersion int          `json:"schemaVersion"`
	Type          string       `json:"type"`
	Pages         []ReportPage `json:"pages"`
	Filters       []Filter     `json:"filters"`
	Settings      interface{}  `json:"settings"`
}

type ReportPage struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Mode       string                 `json:"mode"` // dashboard | pixel
	Widgets    []Widget               `json:"widgets,omitempty"`
	Ankareport map[string]interface{} `json:"ankareport,omitempty"`
}

type Widget struct {
	ID            string                 `json:"id"`
	Type          string                 `json:"type"`
	X             int                    `json:"x"`
	Y             int                    `json:"y"`
	Width         int                    `json:"width"`
	Height        int                    `json:"height"`
	DatasetID     string                 `json:"dataset_id,omitempty"`
	Configuration map[string]interface{} `json:"configuration"`
}

type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
	Scope    string      `json:"scope"`
}
