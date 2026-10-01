package db

import (
	"log"
)

var schema = `
CREATE TABLE IF NOT EXISTS reports (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id VARCHAR(255) DEFAULT 'default',
    name VARCHAR(255) NOT NULL,
    description TEXT,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(50) DEFAULT 'DRAFT',
    dataset_id VARCHAR(255),
    current_version INTEGER DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS report_versions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    report_id UUID REFERENCES reports(id) ON DELETE CASCADE,
    version INTEGER NOT NULL,
    definition_json JSONB NOT NULL,
    created_by VARCHAR(255) DEFAULT 'system',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- We are keeping the actual page/widget structures in the definition_json,
-- but these tables could be used in a later normalization phase.
-- For now, the prompt allows us to use versioned JSON definitions.
CREATE TABLE IF NOT EXISTS report_pages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    report_version_id UUID REFERENCES report_versions(id) ON DELETE CASCADE,
    name VARCHAR(255),
    page_order INTEGER,
    page_type VARCHAR(50),
    definition_json JSONB
);

CREATE TABLE IF NOT EXISTS report_widgets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    report_page_id UUID REFERENCES report_pages(id) ON DELETE CASCADE,
    widget_type VARCHAR(50),
    position_json JSONB,
    configuration_json JSONB
);
`

func Migrate() error {
	if DB == nil {
		return nil
	}
	_, err := DB.Exec(schema)
	if err != nil {
		return err
	}
	log.Println("Database migration completed")
	return nil
}
