<template>
  <div class="anka-container" ref="container"></div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from "vue";

const props = defineProps<{
  initialLayout?: any;
}>();

const emit = defineEmits<{
  (e: "change", layout: any): void;
}>();

const container = ref<HTMLElement | null>(null);

onMounted(async () => {
  if (container.value) {
    const AnkaReport = await import("ankareport").catch(
      () => (window as any).AnkaReport,
    );

    if (AnkaReport && AnkaReport.designer) {
      AnkaReport.designer({
        element: container.value,
        dataSource: [
          { label: "Example Header", field: "header" },
          { label: "Example Content", field: "content" },
        ],
        layout: props.initialLayout || {},
        onSaveButtonClick: (layout: any) => {
          emit("change", layout);
        },
      });
    }
  }
});

onBeforeUnmount(() => {
  if (container.value) {
    container.value.innerHTML = "";
  }
});
</script>

<style scoped>
.anka-container {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}
</style>
