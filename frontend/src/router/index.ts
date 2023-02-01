import Index from "../pages/index.vue";
import Login from "../pages/admin/login.vue";
import Home from "../pages/admin/home.vue";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    component: Index,
  },
  {
    path: "/admin/login",
    component: Login,
  },
  {
    path: "/admin",
    component: Home,
  }
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});