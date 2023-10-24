<template>
  <div class="topic-bar">
    <div class="top-line">
      <NH2 v-if="media.isMobile" style="margin-bottom: 10px">{{ title }}</NH2>
      <NH1 v-else>{{ title }}</NH1>
      <NSpace v-if="showAdminButtons">
        <NButton text :focusable="false" @click="emit('edit')">
          <template #icon>
            <NIcon size="20" class="small-button">
              <EditNoteOutlined />
            </NIcon>
          </template>
        </NButton>
        <NButton text :focusable="false" @click="emit('delete')">
          <template #icon>
            <NIcon size="20" class="small-button">
              <DeleteOutlined />
            </NIcon>
          </template>
        </NButton>
      </NSpace>
    </div>
    <NBlockquote class="topic-detail">
      {{ detail }}
    </NBlockquote>
  </div>
</template>

<script setup lang="ts">
import { EditNoteOutlined, DeleteOutlined } from "@vicons/material";
import { useMedia } from "../store/Media";

const media = useMedia();

defineProps<{
  title: string;
  detail: string;
  showAdminButtons?: boolean;
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
  font-family: "楷体", "楷体_GB2312", serif;
}
.top-line {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}
.topic-detail {
  white-space: pre-line;
  font-size: 1.2rem;
}

@media screen and (max-width: 768px) {
  .topic-bar {
    margin-top: 15px;
    margin-bottom: 20px;
  }
  .topic-detail {
    white-space: pre-line;
    font-size: 14px;
  }
}
</style>
