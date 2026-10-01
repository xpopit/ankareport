<template>
  <div class="wizard">
    <h2>Create New Report - Step {{ step }}</h2>

    <div v-if="step === 1">
      <h3>Select Report Type</h3>
      <select v-model="report.type">
        <option value="DASHBOARD">Dashboard</option>
        <option value="PIXEL_REPORT">Pixel Report</option>
        <option value="DETAIL_REPORT">Detail Report</option>
        <option value="HYBRID">Hybrid Report</option>
      </select>
      <button @click="step++">Next</button>
    </div>

    <div v-if="step === 2">
      <h3>Select Dataset</h3>
      <input
        v-model="report.datasetId"
        placeholder="Enter Dataset ID (e.g. ds-sales)"
      />
      <button @click="step--">Back</button>
      <button @click="step++">Next</button>
    </div>

    <div v-if="step === 3">
      <h3>Visual Layout / Template</h3>
      <p>Using Blank Template.</p>
      <button @click="step--">Back</button>
      <button @click="step++">Next</button>
    </div>

    <div v-if="step === 4">
      <h3>Pages</h3>
      <div v-for="(page, idx) in definition.pages" :key="idx">
        <input v-model="page.name" placeholder="Page Name" />
        <select v-model="page.mode">
          <option value="dashboard">Dashboard</option>
          <option value="pixel">Pixel Perfect (AnkaReport)</option>
        </select>
        <button @click="definition.pages.splice(idx, 1)">Remove</button>
      </div>
      <button @click="addPage">Add Page</button>
      <br /><br />
      <button @click="step--">Back</button>
      <button @click="step++">Next</button>
    </div>

    <div v-if="step === 5">
      <h3>Filters</h3>
      <p>Configure Global Filters</p>
      <!-- filters stub -->
      <button @click="step--">Back</button>
      <button @click="step++">Next</button>
    </div>

    <div v-if="step === 6">
      <h3>Review & Create</h3>
      <input v-model="report.name" placeholder="Report Name" />
      <textarea
        v-model="report.description"
        placeholder="Description"
      ></textarea>

      <p>Type: {{ report.type }}</p>
      <p>Dataset: {{ report.datasetId }}</p>
      <p>Pages: {{ definition.pages.length }}</p>

      <button @click="step--">Back</button>
      <button @click="submit" :disabled="!report.name">Create Report</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { useRouter } from "vue-router";
import { useReportStore } from "../stores/report";

const step = ref(1);
const router = useRouter();
const reportStore = useReportStore();

const report = reactive({
  name: "",
  description: "",
  type: "HYBRID",
  datasetId: "ds-demo",
});

const definition = reactive({
  schemaVersion: 1,
  type: "HYBRID",
  pages: [
    { id: "page-1", name: "Dashboard", mode: "dashboard", widgets: [] },
    { id: "page-2", name: "Print", mode: "pixel", ankareport: {} },
  ],
  filters: [],
  settings: {},
});

const addPage = () => {
  definition.pages.push({
    id: `page-${Date.now()}`,
    name: "New Page",
    mode: "dashboard",
    widgets: [],
  });
};

const submit = async () => {
  const newReport = await reportStore.createReport(report, definition);
  router.push(`/reports/${newReport.id}`);
};
</script>

<style scoped>
.wizard {
  max-width: 600px;
  margin: 0 auto;
}
input,
select,
textarea {
  display: block;
  margin-bottom: 10px;
  width: 100%;
  padding: 8px;
}
</style>
