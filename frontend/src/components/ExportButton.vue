<template>
  <NPopselect :delay="500" v-model:value="exportType" :options="exportOptions" trigger="hover">
    <NButton
      text
      class="button"
      :style="{
        backgroundColor: colorMode.isDarkMode ? '#48484e' : '#FFFFF7',
      }"
      :focusable="false"
      :bordered="false"
      @click="handleExport"
      :loading="loading"
    >
      <template #icon>
        <NIcon size="30">
          <FileDownloadRound />
        </NIcon>
      </template>
    </NButton>
  </NPopselect>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { FileDownloadRound } from "@vicons/material";
import { useColorMode } from "../store/ColorMode";
import { exportCollections } from "../utils/collection";

const colorMode = useColorMode();

const props = defineProps<{
  collections: CollectionType[];
}>();

const loading = ref(false);
const exportType = ref<"docx" | "md">("docx");
const exportOptions = [
  { label: "导出为 Word", value: "docx" },
  { label: "导出为 Markdown", value: "md" },
];
const handleExport = async () => {
  loading.value = true;
  await exportCollections(props.collections, exportType.value);
  loading.value = false;
};
</script>

<style scoped>
.button {
  position: absolute;
  bottom: 105px;
  right: 20px;
  width: 44px;
  height: 44px;
  z-index: 2;
  border-radius: 50%;
  box-shadow: rgba(0, 0, 0, 0.12) 0 2px 8px 0;
}
</style>
