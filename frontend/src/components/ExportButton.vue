<template>
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
</template>

<script setup lang="ts">
import { ref } from "vue";
import { FileDownloadRound } from "@vicons/material";
import { useColorModeStore } from "../store/ColorMode";
import { exportCollections } from "../utils/collection";

const colorMode = useColorModeStore();

const props = defineProps<{
  collections: CollectionType[];
}>();

const loading = ref(false);
const handleExport = async () => {
    loading.value = true;
    await exportCollections(props.collections);
    loading.value = false;
};
</script>

<style scoped>
.button {
  position: absolute;
  bottom: 80px;
  right: 20px;
  width: 44px;
  height: 44px;
  z-index: 2;
  border-radius: 50%;
  box-shadow: rgba(0, 0, 0, 0.12) 0 2px 8px 0;
}
</style>
