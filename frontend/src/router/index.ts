import Index from "../pages/index.vue";
import Login from "../pages/admin/login.vue";
import Home from "../pages/admin/home.vue";
import Editor from "../pages/admin/editor.vue";
import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: Index,
  },
  {
    path: "/:tid",
    component: Index,
    props: true,
  },
  {
    path: "/admin/login",
    component: Login,
  },
  {
    path: "/admin",
    component: Home,
  },
  {
    path: "/admin/add",
    component: Editor,
    props: (route) => ({
      author: route.query.author,
      book: route.query.book,
    }),
  },
  {
    path: "/admin/edit/:cid",
    component: Editor,
    props: true,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
