<template>
  <div class="container">
    <NCard hoverable>
      <div class="top-line">
        <NIcon size="40" class="left-quote" color="rgba(150, 150, 150, 0.5)">
          <FormatQuoteRound />
        </NIcon>
        <NSpace>
          <StarButton v-if="showAdminButtons" :starred="starred" @click="emit('star')" />
          <NButton v-if="showAdminButtons" text :focusable="false" @click="emit('edit')">
            <template #icon>
              <NIcon size="20" class="small-button">
                <EditNoteOutlined />
              </NIcon>
            </template>
          </NButton>
          <NButton v-if="showAdminButtons" text :focusable="false" @click="emit('delete')">
            <template #icon>
              <NIcon size="20" class="small-button">
                <DeleteOutlined />
              </NIcon>
            </template>
          </NButton>
          <NButton text :focusable="false" @click="copy">
            <template #icon>
              <NIcon size="20" class="small-button">
                <ContentCopyRound />
              </NIcon>
            </template>
          </NButton>
        </NSpace>
      </div>
      <div class="content">{{ content }}</div>
      <div class="right">
        <NIcon size="40" color="rgba(150, 150, 150, 0.5)">
          <FormatQuoteRound />
        </NIcon>
      </div>
      <div class="source" style="margin-top: 1.5rem">{{ source }}</div>
      <template #footer v-if="tags && tags.length">
        <NDivider />
        <NSpace :size="[0, 5]">
          <NTag
            v-for="(tag, index) in tags"
            :key="tag"
            :type="computeType(index)"
            style="margin-right: 0.5rem"
          >
            {{ tag }}
          </NTag>
        </NSpace>
      </template>
    </NCard>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import {
  FormatQuoteRound,
  ContentCopyRound,
  EditNoteOutlined,
  DeleteOutlined,
} from "@vicons/material";
import { useMessage } from "naive-ui";
import StarButton from "./StarButton.vue";

const message = useMessage();

const props = defineProps<{
  showAdminButtons?: boolean;
  starred: boolean;
  content: string;
  author: string;
  book: string;
  tags: string[];
}>();

const emit = defineEmits(["edit", "delete", "star"]);

const source = computed(() => {
  if (props.author && props.book) {
    return `—— ${props.author}《${props.book}》`;
  } else if (props.author) {
    return "—— " + props.author;
  } else if (props.book) {
    return `——《${props.book}》`;
  } else {
    return "";
  }
});

function computeType(index: number) {
  switch (index % 4) {
    case 0:
      return "info";
    case 1:
      return "success";
    case 2:
      return "warning";
    case 3:
      return "error";
    default:
      return "info";
  }
}

function copy() {
  navigator.clipboard.writeText(props.content + "\n——" + source.value);
  message.success("已复制到剪贴板");
}
</script>

<style scoped>
.container {
  padding: 10px 0;
}
.n-card {
  border-radius: 10px;
}
.top-line {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}
.small-button {
  opacity: 0.7;
}
.left-quote {
  transform: rotate(180deg);
}
.right {
  text-align: right;
}
.source {
  text-align: right;
  font-size: 1.3rem;
  font-weight: bold;
  font-family: v-sans, v-mono, "Times New Roman", Times, serif;
}
.content {
  font-size: 1.3rem;
  font-weight: bold;
  font-family: v-sans, v-mono, "Times New Roman", Times, serif;
  padding: 0 50px;
  white-space: pre-line;
}
</style>
