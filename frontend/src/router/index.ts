import Index from "../pages/index.vue";
import LoginVue from "../pages/admin/login.vue";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    component: Index,
  },
  {
    path: "/admin/login",
    component: LoginVue,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});