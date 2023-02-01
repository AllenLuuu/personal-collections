import { createApp } from "vue";
import App from "./App.vue";
import { router } from "./router";
import { createPinia } from "pinia";
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import "vfonts/FiraSans.css";
import "vfonts/FiraCode.css";

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

createApp(App).use(pinia).use(router).mount("#app");
