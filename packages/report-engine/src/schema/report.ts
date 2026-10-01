export interface ReportDefinition {
  schemaVersion: number;
  type: "DASHBOARD" | "PIXEL_REPORT" | "DETAIL_REPORT" | "HYBRID";
  pages: ReportPage[];
  filters: Filter[];
  settings: Record<string, any>;
}

export interface ReportPage {
  id: string;
  name: string;
  mode: "dashboard" | "pixel";
  widgets?: Widget[];
  ankareport?: Record<string, any>;
}

export interface Widget {
  id: string;
  type: string;
  x: number;
  y: number;
  width: number;
  height: number;
  dataset_id?: string;
  configuration: Record<string, any>;
}

export interface Filter {
  field: string;
  operator: string;
  value: any;
  scope: string;
}
