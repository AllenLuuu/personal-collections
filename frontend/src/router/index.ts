import Index from "../pages/index.vue";
import Login from "../pages/admin/login.vue";
import Home from "../pages/admin/home.vue";
import Editor from "../pages/admin/editor.vue";
import TopicEditor from "../pages/admin/topicEditor.vue";
import { useLoginInfoStore } from "../store/LoginInfo";
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
    name: "login",
    component: Login,
  },
  {
    path: "/admin",
    component: Home,
  },
  {
    path: "/admin/:tid",
    component: Home,
    props: true,
  },
  {
    path: "/admin/add",
    component: Editor,
    props: (route) => ({
      author: route.query.author,
      book: route.query.book,
      tid: route.query.tid,
    }),
  },
  {
    path: "/admin/edit/:cid",
    component: Editor,
    props: true,
  },
  {
    path: "/admin/topic/add",
    component: TopicEditor,
  },
  {
    path: "/admin/topic/edit/:tid",
    component: TopicEditor,
    props: true,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const loginInfoStore = useLoginInfoStore();
  if (to.path.startsWith("/admin") && to.name !== "login" && !loginInfoStore.isLoggedIn) {
    next("/admin/login");
  } else {
    next();
  }
});
