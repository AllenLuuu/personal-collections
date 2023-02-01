<template>
  <n-button
    type="primary"
    dashed
    class="addButton"
    @click="emitClick"
    :style="{ height: height ? height + 'px' : null }"
  >
    <template #icon>
      <n-icon>
        <PlusOutlined />
      </n-icon>
    </template>
  </n-button>
</template>

<script setup lang="ts">
import { PlusOutlined } from "@vicons/material";

defineProps<{
  height?: number;
}>();

const emit = defineEmits(["click"]);

const emitClick = () => {
  emit("click");
};
</script>

<style scoped>
.addButton {
  width: calc(100% - 20px);
  margin: 10px;
}
</style>
