<template>
  <RouterView />
</template>

<script setup lang="ts">
import { useMessage } from "naive-ui";
import { onMounted } from "vue";
import { useMedia } from "./store/Media";

const media = useMedia();

onMounted(() => {
  window.$message = useMessage();
  let windowWidth = document.body.clientWidth;
  if (windowWidth < 768) {
    media.setMedia("mobile");
  } else {
    media.setMedia("desktop");
  }
  window.addEventListener("resize", () => {
    let windowWidth = document.body.clientWidth;
    if (windowWidth < 768) {
      media.setMedia("mobile");
    } else {
      media.setMedia("desktop");
    }
  });
});
</script>
