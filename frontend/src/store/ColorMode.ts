import { defineStore } from "pinia";

export const useColorMode = defineStore("color-mode", {
  state: () => ({
    colorMode: "light",
  }),
  getters: {
    isDarkMode(): boolean {
      return this.colorMode === "dark";
    },
  },
  actions: {
    setColorMode(colorMode: string) {
      this.colorMode = colorMode;
    },
    revertColorMode() {
      this.colorMode = this.colorMode === "dark" ? "light" : "dark";
    },
  },
  persist: true,
});
