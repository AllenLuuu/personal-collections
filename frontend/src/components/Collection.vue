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
  <NModal v-model:show="showExportModal" :auto-focus="false">
    <n-card
      title="导出为图片"
      style="width: min-content"
      content-style="padding: 0"
      closable
      @close="showExportModal = false"
    >
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
              {{ `${author + (book ? `《${book}》` : "")}` }}
            </div>
          </div>
          <div v-if="showPicFooter" class="pic-footer">
            <div class="footer-logo">
              <NIcon size="1rem" color="#A07C18FF"><BookOutlined /></NIcon>
              <span>游逛者·书摘</span>
            </div>
            <img style="width: 2rem; height: 2rem" :src="qrcode" />
          </div>
        </div>
      </div>
      <template #footer>
        <div class="export-footer">
          <div>二维码：<NSwitch v-model:value="showPicFooter" /></div>
          <NButton ghost round type="primary" @click="exportImage">
            保存
          </NButton>
        </div>
      </template>
    </n-card>
  </NModal>
</template>

<script setup lang="ts">
import {
  BookOutlined,
  ContentCopyRound,
  DeleteOutlined,
  EditNoteOutlined,
  FormatQuoteRound,
  IosShareOutlined,
} from "@vicons/material";
import html2canvas from "html2canvas";
import { NModal, useMessage } from "naive-ui";
import { computed, ref } from "vue";
import qrcode from "../assets/qrcode.png";
import { useMedia } from "../store/Media";
import StarButton from "./StarButton.vue";

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

const showPicFooter = ref(false);

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
}
.content {
  font-size: 1.3rem;
  padding: 0 50px;
  white-space: pre-line;
}
.pic-wrapper {
  border: 1px solid #ccc;
}
.pic-back {
  font-size: 0.8rem;
  width: 16rem;
  white-space: pre-line;
  color: black;
  background-color: white;
  padding: 0.8rem;
}
.pic-border {
  padding: 1rem;
  min-height: 6rem;
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
.pic-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;

  margin-top: 0.5rem;

  font-family: "KingHwa_OldSong", serif;
}
.footer-logo {
  display: flex;
  align-items: center;
  height: 1rem;
  gap: 0.2rem;
}
.export-footer {
  margin-top: 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
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
