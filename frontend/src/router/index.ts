import Index from "../pages/index.vue";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    component: Index,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});