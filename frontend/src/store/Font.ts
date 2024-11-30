import { defineStore } from "pinia";

export const useFontStore = defineStore<
  string,
  Font,
  {},
  {
    setFont(font: Font): void;
  }
>("font", {
  state: () => ({
    name: "默认",
    cssName: "Lato",
  }),
  actions: {
    setFont(font: Font) {
      this.name = font.name;
      this.cssName = font.cssName;
    },
  },
  persist: {
    key: "font",
    storage: localStorage,
  },
});
