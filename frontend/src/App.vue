<template>
  <n-config-provider
    :theme="colorMode.isDarkMode ? darkTheme : null"
    :theme-overrides="
      colorMode.isDarkMode ? darkThemeOverrides : lightThemeOverrides
    "
  >
    <n-global-style />
    <n-dialog-provider>
      <n-message-provider>
        <InterLayer />
      </n-message-provider>
    </n-dialog-provider>
  </n-config-provider>
</template>

<script setup lang="ts">
import { darkTheme, GlobalThemeOverrides } from "naive-ui";
import { computed } from "vue";
import InterLayer from "./InterLayer.vue";
import { useColorMode } from "./store/ColorMode";
import { useFontStore } from "./store/Font";

const colorMode = useColorMode();
const font = useFontStore();

const lightThemeOverrides = computed<GlobalThemeOverrides>(() => {
  return {
    common: {
      baseColor: "#FFFFF7FF",
      primaryColor: "#A07C18FF",
      primaryColorHover: "#AD9136FF",
      primaryColorPressed: "#7A600CFF",
      primaryColorSuppl: "#AD9136FF",
      bodyColor: "#FFFFF0FF",
      modalColor: "#FFFFF7FF",
      cardColor: "#FFFFF7FF",
      popoverColor: "#FFFFF7FF",
      tableColor: "#FFFFF7FF",
      inputColor: "#FFFFF7FF",
      fontFamily: font.cssName,
    },
    Tag: {
      colorBordered: "#FFFFF7FF",
    },
  };
});

const darkThemeOverrides = computed<GlobalThemeOverrides>(() => {
  return {
    common: {
      baseColor: "#1A1A1A",
      primaryColor: "#E6AD4FFF",
      primaryColorHover: "#EABD73FF",
      primaryColorPressed: "#B48535FF",
      primaryColorSuppl: "#EABD73FF",
      fontFamily: font.cssName,
    },
    Tag: {
      colorBordered: "#1A1A1A",
    },
  };
});
</script>
