<template>
  <n-config-provider :theme="colorMode.isDarkMode ? darkTheme : null">
    <RouterView />
  </n-config-provider>
</template>

<script setup lang="ts">
import { onMounted } from "vue";
import { useColorModeStore } from "./store/ColorMode";
import { darkTheme } from "naive-ui";

const colorMode = useColorModeStore();

onMounted(() => {
  let media = window.matchMedia("(prefers-color-scheme: dark)");
  if (media.matches) {
    colorMode.setColorMode("dark");
  } else {
    colorMode.setColorMode("light");
  }

  console.log(colorMode.isDarkMode);

  media.addEventListener("change", (e) => {
    if (e.matches) {
      colorMode.setColorMode("dark");
    } else {
      colorMode.setColorMode("light");
    }
  });
});
</script>
