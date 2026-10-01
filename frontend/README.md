# @xerp/report-studio-vue

`@xerp/report-studio-vue` is the Vue 3 component library and application layer for the XERP Enterprise Reporting Platform. It wraps `@xerp/report-engine` and the `ankareport` designer to provide a complete Report Studio interface.

## Installation

```bash
npm install @xerp/report-studio-vue
```

## Features

- **Report Studio UI**: Drag-and-drop designer powered by AnkaReport.
- **Multi-page Layouts**: Support for Hybrid dashboards combining interactive metrics and pixel-perfect print designs.
- **Vue Components**: `ReportStudio`, `ReportWizard`, `ReportList` are fully exportable and usable in any Vue 3 Vite application.
- **Pinia State**: Built-in state management using `useReportStore`.

## Usage

### Integrating into a Vue 3 Application

In your `main.ts`:

```typescript
import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import { createRouter, createWebHistory } from "vue-router";
import {
  ReportStudio,
  ReportWizard,
  ReportList,
} from "@xerp/report-studio-vue";

// Optional: you can route directly to the provided components
const routes = [
  { path: "/reports", component: ReportList },
  { path: "/reports/new", component: ReportWizard },
  { path: "/reports/:id", component: ReportStudio },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

const app = createApp(App);
app.use(createPinia());
app.use(router);
app.mount("#app");
```

### Local Storage Backend

Currently, the `@xerp/report-studio-vue` library uses the browser's `localStorage` (via the `useReportStore` Pinia store) to simulate a complete backend flow. Reports are stored, cloned, and executed completely in the browser for demonstration and offline capabilities.
