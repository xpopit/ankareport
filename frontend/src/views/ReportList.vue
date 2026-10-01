<template>
  <div>
    <div class="header">
      <h2>Reports</h2>
      <button @click="$router.push('/reports/new')">New Report</button>
    </div>

    <table class="report-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Type</th>
          <th>Status</th>
          <th>Version</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="report in reportStore.reports" :key="(report as any).id">
          <td>{{ (report as any).name }}</td>
          <td>{{ (report as any).type }}</td>
          <td>{{ (report as any).status }}</td>
          <td>{{ (report as any).currentVersion }}</td>
          <td class="actions">
            <button @click="$router.push(`/reports/${(report as any).id}`)">
              Design
            </button>
            <button @click="cloneReport((report as any).id)">Clone</button>
            <button @click="deleteReport((report as any).id)" class="danger">
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from "vue";
import { useReportStore } from "../stores/report";

const reportStore = useReportStore();

onMounted(async () => {
  await reportStore.fetchReports();
});

const cloneReport = async (id: string) => {
  await reportStore.cloneReport(id);
  await reportStore.fetchReports();
};

const deleteReport = async (id: string) => {
  if (confirm("Are you sure you want to delete this report?")) {
    await reportStore.deleteReport(id);
    await reportStore.fetchReports();
  }
};
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}
.report-table {
  width: 100%;
  border-collapse: collapse;
}
.report-table th,
.report-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}
.actions button {
  margin-right: 5px;
}
.danger {
  color: white;
  background-color: #dc3545;
}
</style>
