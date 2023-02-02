<template>
  <div class="topic-bar">
    <div class="top-line">
      <NH1>{{ title }}</NH1>
      <NSpace v-if="showAdminButtons">
        <NButton
          text
          :focusable="false"
          @click="emit('edit')"
        >
          <template #icon>
            <NIcon size="20" class="small-button">
              <EditNoteOutlined />
            </NIcon>
          </template>
        </NButton>
        <NButton
          text
          :focusable="false"
          @click="emit('delete')"
        >
          <template #icon>
            <NIcon size="20" class="small-button">
              <DeleteOutlined />
            </NIcon>
          </template>
        </NButton>
      </NSpace>
    </div>
    <!-- <NH1>{{ title }}</NH1> -->
    <NBlockquote>
      {{ detail }}
    </NBlockquote>
  </div>
</template>

<script setup lang="ts">
import {
  EditNoteOutlined,
  DeleteOutlined,
} from "@vicons/material";

defineProps<{
  title: string;
  detail: string;
  showAdminButtons: boolean;
}>();

const emit = defineEmits<{
  (e: "edit"): void;
  (e: "delete"): void;
}>();
</script>

<style scoped>
.topic-bar {
  margin-top: 20px;
  margin-bottom: 30px;
}
.top-line {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}
</style>
