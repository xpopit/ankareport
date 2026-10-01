# @xerp/report-engine

`@xerp/report-engine` is a pure TypeScript library designed to process, query, filter, and structure dataset logic for enterprise reporting platforms. It operates independently of any frontend framework (like Vue or React) and does not rely on a backend.

## Features

- **QueryEngine**: In-memory filtering and sorting of dataset records.
- **Data Schemas**: Strict TypeScript definitions for Datasets, Reports, Pages, and Widgets.
- **Zero Dependencies**: Pure TypeScript core.

## Installation

```bash
npm install @xerp/report-engine
```

## Basic Usage

### Defining a Dataset

```typescript
import { Dataset } from "@xerp/report-engine";

const salesDataset: Dataset = {
  id: "sales-2026",
  name: "Sales Data 2026",
  columns: [
    { name: "product", type: "string" },
    { name: "amount", type: "number" },
  ],
  data: [
    { product: "Laptop", amount: 1500 },
    { product: "Desk", amount: 300 },
  ],
};
```

### Using the Query Engine

```typescript
import { QueryEngine } from "@xerp/report-engine";

const engine = new QueryEngine(salesDataset);

const results = engine.execute({
  filters: [{ field: "amount", operator: "gt", value: 500 }],
  sort: [{ field: "amount", direction: "desc" }],
});

console.log(results);
// Output: [{ product: 'Laptop', amount: 1500 }]
```
