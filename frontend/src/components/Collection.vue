<template>
  <div class="container">
    <NCard
      :segmented="{
        footer: 'soft',
      }"
      hoverable
      size="small"
    >
      <div class="top-line">
        <NIcon
          :size="media.isMobile ? 25 : 40"
          class="left-quote"
          color="rgba(150, 150, 150, 0.5)"
        >
          <FormatQuoteRound />
        </NIcon>
        <NSpace>
          <StarButton
            v-if="showAdminButtons"
            :starred="starred"
            @click="emit('star')"
            :size="media.isMobile ? 15 : 20"
          />
          <NButton
            v-if="showAdminButtons"
            text
            :focusable="false"
            @click="emit('edit')"
          >
            <template #icon>
              <NIcon :size="media.isMobile ? 15 : 20" class="small-button">
                <EditNoteOutlined />
              </NIcon>
            </template>
          </NButton>
          <NButton
            v-if="showAdminButtons"
            text
            :focusable="false"
            @click="emit('delete')"
          >
            <template #icon>
              <NIcon :size="media.isMobile ? 15 : 20" class="small-button">
                <DeleteOutlined />
              </NIcon>
            </template>
          </NButton>
          <NButton text :focusable="false" @click="copy">
            <template #icon>
              <NIcon :size="media.isMobile ? 15 : 20" class="small-button">
                <ContentCopyRound />
              </NIcon>
            </template>
          </NButton>
          <NButton text :focusable="false" @click="showExportModal = true">
            <template #icon>
              <NIcon :size="media.isMobile ? 15 : 20" class="small-button">
                <IosShareOutlined />
              </NIcon>
            </template>
          </NButton>
        </NSpace>
      </div>
      <div class="content">{{ content }}</div>
      <div class="right">
        <NIcon
          :size="media.isMobile ? 25 : 40"
          color="rgba(150, 150, 150, 0.5)"
        >
          <FormatQuoteRound />
        </NIcon>
      </div>
      <div class="source" style="margin-top: 1rem">{{ source }}</div>
      <template #footer v-if="tags && tags.length">
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
  <!-- 导出图片弹窗 -->
  <NModal v-model:show="showExportModal">
    <n-card title="导出为图片" style="width: min-content">
      <div class="pic-wrapper">
        <div ref="picture" class="pic-back">
          <div class="pic-border">
            <NIcon class="pic-left-quote" size="1rem" color="#ccc">
              <FormatQuoteRound />
            </NIcon>
            <NIcon class="pic-right-quote" size="1rem" color="#ccc">
              <FormatQuoteRound />
            </NIcon>
            {{ content }}
            <div class="right" style="margin-top: 1.5em">
              {{ `${author} · ${book}` }}
            </div>
          </div>
        </div>
      </div>
      <div class="export-footer">
        <NButton @click="exportImage">保存</NButton>
      </div>
    </n-card>
  </NModal>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import {
  FormatQuoteRound,
  ContentCopyRound,
  EditNoteOutlined,
  DeleteOutlined,
  IosShareOutlined,
} from "@vicons/material";
import { NModal, useMessage } from "naive-ui";
import StarButton from "./StarButton.vue";
import { useMedia } from "../store/Media";
import html2canvas from "html2canvas";

const media = useMedia();
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
  navigator.clipboard.writeText(props.content + "\n" + source.value);
  message.success("已复制到剪贴板");
}

const showExportModal = ref(false);

const picture = ref<HTMLDivElement | null>(null);
const exportImage = () => {
  if (picture.value === null) return;
  html2canvas(picture.value, {
    scale: 3,
  }).then((canvas) => {
    const data = canvas.toDataURL("image/png");
    const link = document.createElement("a");
    link.download = "quote.png";
    link.href = data;
    link.click();
  });
};
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
  font-family: "楷体", "楷体_GB2312", serif;
}
.content {
  font-size: 1.3rem;
  font-family: "楷体", "楷体_GB2312", serif;
  padding: 0 50px;
  white-space: pre-line;
}
.pic-wrapper {
  border: 1px solid #ccc;
}
.pic-back {
  font-size: 0.8rem;
  font-family: "LXGW WenKai", serif;
  width: 16rem;
  white-space: pre-line;
  color: black;
  background-color: white;
  padding: 0.8rem;
}
.pic-border {
  padding: 1rem;
  min-height: 8rem;
  border: 0.1rem solid #ccc;
  position: relative;

  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.pic-left-quote {
  transform: rotate(180deg);
  position: absolute;
  top: 0.5rem;
  left: -0.5rem;
  background-color: white;
}
.pic-right-quote {
  position: absolute;
  bottom: 0.5rem;
  right: -0.5rem;
  background-color: white;
}
.export-footer {
  display: flex;
  justify-content: flex-end;
  margin-top: 1rem;
  padding: 0.5rem 1rem;
}

@media screen and (max-width: 768px) {
  .content {
    padding: 0 25px;
    font-size: medium;
  }
  .source {
    font-size: medium;
  }
}
</style>
