import { defineStore } from "pinia";

export const useMedia = defineStore("media", {
    state: () => ({
      isMobile: false,
    }),
    actions: {
      setMedia(media: "mobile" | "desktop") {
        this.isMobile = media === "mobile";
      },
    },
  });