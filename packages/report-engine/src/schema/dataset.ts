export interface DatasetColumn {
  name: string;
  type: "string" | "number" | "date" | "boolean";
  label?: string;
}

export interface Dataset {
  id: string;
  name: string;
  columns: DatasetColumn[];
  data: Record<string, any>[];
}
