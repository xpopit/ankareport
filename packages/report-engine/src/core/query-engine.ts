import { Dataset } from '../schema/dataset';

export interface QueryOptions {
  filters?: { field: string; operator: string; value: any }[];
  sort?: { field: string; direction: 'asc' | 'desc' }[];
  groupBy?: string[];
  aggregates?: { field: string; func: 'sum' | 'count' | 'avg' | 'min' | 'max' }[];
}

export class QueryEngine {
  constructor(private dataset: Dataset) {}

  execute(options: QueryOptions): Record<string, any>[] {
    let result = [...this.dataset.data];

    // Basic filtering
    if (options.filters && options.filters.length > 0) {
      result = result.filter(row => {
        return options.filters!.every(filter => {
          const val = row[filter.field];
          switch (filter.operator) {
            case 'eq': return val === filter.value;
            case 'neq': return val !== filter.value;
            case 'gt': return val > filter.value;
            case 'lt': return val < filter.value;
            case 'contains': return String(val).includes(String(filter.value));
            default: return true;
          }
        });
      });
    }

    // Basic sorting
    if (options.sort && options.sort.length > 0) {
      result.sort((a, b) => {
        for (const s of options.sort!) {
          if (a[s.field] < b[s.field]) return s.direction === 'asc' ? -1 : 1;
          if (a[s.field] > b[s.field]) return s.direction === 'asc' ? 1 : -1;
        }
        return 0;
      });
    }

    // Note: Group By and Aggregation would go here. Omitted for brevity in initial mock.

    return result;
  }
}
