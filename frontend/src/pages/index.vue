<template>
  <NLayout class="layout">
    <NLayoutHeader bordered style="height: 80px">
      <Header></Header>
    </NLayoutHeader>
    <NLayout has-sider position="absolute" style="top: 80px">
      <NLayoutSider
        collapse-mode="width"
        :show-collapsed-content="false"
        :collapsed-width="20"
        show-trigger="arrow-circle"
        @update:collapsed="hideSider = $event"
        :bordered="!hideSider"
        :class="hideSider ? 'transparent' : null"
      >
        <SideMenu v-model:menuValue="menuValue" />
      </NLayoutSider>
      <NLayoutContent ref="contentRef" :native-scrollbar="false">
        <div class="content">
          <CollectionPage :container-ref="contentRef!" :tid="tid" />
        </div>
      </NLayoutContent>
    </NLayout>
  </NLayout>
</template>

<script setup lang="ts">
import Header from "../components/Header.vue";
import { ref, onMounted } from "vue";
import { LayoutInst } from "naive-ui";
import CollectionPage from "../components/CollectionPage.vue";
import SideMenu from "../components/SideMenu.vue";

const props = defineProps<{
  tid?: string;
}>();

onMounted(async () => {
  menuValue.value = props.tid ?? "all";
});

const menuValue = ref<string>("all");

const hideSider = ref(false);

const contentRef = ref<LayoutInst | null>(null);
</script>

<style scoped>
.transparent {
  background-color: transparent;
}
.content {
  padding: 10px 80px 10px 20px;
}
.layout {
  height: 100vh;
}
</style>
