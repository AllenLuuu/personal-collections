<template>
  <NModal :show="show" :auto-focus="false">
    <n-card
      title="导出为图片"
      style="width: min-content"
      content-style="padding: 0"
      closable
      @close="emit('close')"
    >
      <div
        class="pic-wrapper"
        :style="{ 'font-size': media.isMobile ? '1rem' : '2rem' }"
      >
        <div ref="picture" class="pic-back">
          <div class="pic-border">
            <NIcon class="pic-left-quote" size="1em" color="#ccc">
              <FormatQuoteRound />
            </NIcon>
            <NIcon class="pic-right-quote" size="1em" color="#ccc">
              <FormatQuoteRound />
            </NIcon>
            <div class="text">{{ content }}</div>
            <div class="right text" style="margin-top: 1.5em">
              {{ `${author + (book ? `《${book}》` : "")}` }}
            </div>
          </div>
          <div v-if="showPicFooter" class="pic-footer">
            <div class="footer-logo">
              <NIcon size="1em" color="#A07C18FF"><BookOutlined /></NIcon>
              <span class="text">游逛者·书摘</span>
            </div>
            <img style="width: 2em; height: 2em" :src="qrcode" />
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
import { BookOutlined, FormatQuoteRound } from "@vicons/material";
import html2canvas from "html2canvas";
import { NButton, NIcon, NModal, NSwitch } from "naive-ui";
import { ref } from "vue";
import qrcode from "../assets/qrcode.png";
import { useMedia } from "../store/Media";

const media = useMedia();

defineProps<{
  show: boolean;
  content: string;
  author?: string;
  book?: string;
}>();

const emit = defineEmits(["close"]);

const showPicFooter = ref(false);

const picture = ref<HTMLDivElement | null>(null);

const exportImage = () => {
  if (picture.value === null) return;
  html2canvas(picture.value, {
    scale: media.isMobile ? 4 : 2,
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
.right {
  text-align: right;
}
.text {
  font-size: 0.8em;
}
.pic-wrapper {
  border: 1px solid #ccc;
}
.pic-back {
  width: 16em;
  white-space: pre-line;
  color: black;
  background-color: white;
  padding: 0.8em;
}
.pic-border {
  padding: 1em;
  min-height: 6em;
  border: 0.1em solid #ccc;
  position: relative;

  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.pic-left-quote {
  transform: rotate(180deg);
  position: absolute;
  top: 0.5em;
  left: -0.5em;
  background-color: white;
}
.pic-right-quote {
  position: absolute;
  bottom: 0.5em;
  right: -0.5em;
  background-color: white;
}
.pic-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;

  margin-top: 1em;

  font-family: "KingHwa_OldSong", serif;
}
.footer-logo {
  display: flex;
  align-items: center;
  height: 1em;
  gap: 0.2em;
}
.export-footer {
  margin-top: 1em;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
