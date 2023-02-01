<template>
  <div class="container">
    <NCard hoverable>
      <div class="top-line">
        <NIcon size="40" class="left-quote" color="rgba(150, 150, 150, 0.5)">
          <FormatQuoteRound />
        </NIcon>
        <NButton text :focusable="false" @click="copy">
          <template #icon>
            <NIcon size="20" class="copy-button">
              <ContentCopyRound />
            </NIcon>
          </template>
        </NButton>
      </div>
      <div class="content">{{ content }}</div>
      <div class="right">
        <NIcon size="40" color="rgba(150, 150, 150, 0.5)">
          <FormatQuoteRound />
        </NIcon>
      </div>
      <div class="source" style="margin-top: 1.5rem">——{{ source }}</div>
      <template #footer v-if="tags">
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
import { FormatQuoteRound, ContentCopyRound } from "@vicons/material";
import { useMessage } from "naive-ui";

const message = useMessage();

const props = defineProps<{
  content: string;
  author: string;
  book: string;
  tags: string[];
}>();

const source = computed(() => {
  if (props.author && props.book) {
    return `${props.author}《${props.book}》`;
  } else if (props.author) {
    return props.author;
  } else if (props.book) {
    return `《${props.book}》`;
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
  padding: 10px;
}
.n-card {
  border-radius: 10px;
}
.top-line {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}
.copy-button {
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
  padding-inline-start: 50px;
  padding-inline-end: 50px;
}
</style>
