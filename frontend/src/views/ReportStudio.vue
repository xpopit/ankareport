<template>
  <div class="studio-layout" v-if="report">
    <div class="studio-sidebar">
      <h3>Pages</h3>
      <ul>
        <li
          v-for="page in definition.pages"
          :key="page.id"
          :class="{ active: selectedPageId === page.id }"
          @click="selectedPageId = page.id"
        >
          {{ page.name }} ({{ page.mode }})
        </li>
      </ul>
    </div>

    <div class="studio-main">
      <div class="studio-header">
        <h2>{{ report.name }}</h2>
        <div class="actions">
          <button @click="saveReport">Save</button>
          <button>Preview</button>
          <button>Execute</button>
        </div>
      </div>

      <div class="studio-canvas">
        <template v-if="currentPage?.mode === 'pixel'">
          <AnkaReportAdapter
            :initial-layout="currentPage.ankareport"
            @change="onAnkaReportChange"
          />
        </template>
        <template v-else>
          <div class="dashboard-canvas">
            <h4>Dashboard Canvas (Grid system)</h4>
            <!-- Phase 2 dashboard implementation here -->
          </div>
        </template>
      </div>
    </div>
  </div>
  <div v-else>Loading...</div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { useRoute } from "vue-router";
import { useReportStore } from "../stores/report";
import AnkaReportAdapter from "../components/AnkaReportAdapter.vue";

const route = useRoute();
const reportStore = useReportStore();

const report = ref<any>(null);
const definition = ref<any>(null);
const selectedPageId = ref<string>("");

onMounted(async () => {
  const data = await reportStore.fetchReport(route.params.id as string);
  report.value = data;
  definition.value = data.definition;

  if (definition.value.pages && definition.value.pages.length > 0) {
    selectedPageId.value = definition.value.pages[0].id;
  }
});

const currentPage = computed(() => {
  if (!definition.value) return null;
  return definition.value.pages.find((p: any) => p.id === selectedPageId.value);
});

const onAnkaReportChange = (layout: any) => {
  if (currentPage.value) {
    currentPage.value.ankareport = layout;
  }
};

const saveReport = async () => {
  await reportStore.updateReport(
    report.value.id,
    report.value,
    definition.value,
  );
  alert("Report saved successfully!");
};
</script>

<style scoped>
.studio-layout {
  display: flex;
  height: calc(100vh - 40px);
}
.studio-sidebar {
  width: 250px;
  background: #f8f9fa;
  border-right: 1px solid #ddd;
  padding: 10px;
}
.studio-sidebar ul {
  list-style: none;
  padding: 0;
}
.studio-sidebar li {
  padding: 10px;
  cursor: pointer;
  border-bottom: 1px solid #eee;
}
.studio-sidebar li.active {
  background: #007bff;
  color: white;
}
.studio-main {
  flex: 1;
  display: flex;
  flex-direction: column;
}
.studio-header {
  padding: 10px 20px;
  border-bottom: 1px solid #ddd;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.studio-canvas {
  flex: 1;
  padding: 20px;
  overflow: auto;
  position: relative;
}
.dashboard-canvas {
  border: 2px dashed #ccc;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
