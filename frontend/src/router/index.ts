import { createRouter, createWebHistory } from "vue-router";
import Dashboard from "../views/Dashboard.vue";
import ReportList from "../views/ReportList.vue";
import ReportWizard from "../views/ReportWizard.vue";
import ReportStudio from "../views/ReportStudio.vue";

const routes = [
  { path: "/", component: Dashboard },
  { path: "/reports", component: ReportList },
  { path: "/reports/new", component: ReportWizard },
  { path: "/reports/:id", component: ReportStudio },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
