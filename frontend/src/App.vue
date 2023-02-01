<template>
  <n-config-provider
    :theme="colorMode.isDarkMode ? darkTheme : null"
    :theme-overrides="
      colorMode.isDarkMode ? darkThemeOverrides : lightThemeOverrides
    "
  >
    <n-dialog-provider>
      <n-message-provider>
        <RouterView />
      </n-message-provider>
    </n-dialog-provider>
  </n-config-provider>
</template>

<script setup lang="ts">
import { useColorModeStore } from "./store/ColorMode";
import { darkTheme, GlobalThemeOverrides } from "naive-ui";

const colorMode = useColorModeStore();

const lightThemeOverrides: GlobalThemeOverrides = {
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
  },
  Tag: {
    colorBordered: "#FFFFF7FF",
  },
};

const darkThemeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: "#E6AD4FFF",
    primaryColorHover: "#EABD73FF",
    primaryColorPressed: "#B48535FF",
    primaryColorSuppl: "#EABD73FF",
  },
};
</script>
