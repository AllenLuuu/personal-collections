<template>
  <NLayout class="layout">
    <NLayoutHeader
      bordered
      :style="{ height: media.isMobile ? '68.5px' : '80px' }"
    >
      <Header></Header>
    </NLayoutHeader>
    <NLayout
      id="main-content"
      has-sider
      position="absolute"
      :style="{ top: media.isMobile ? '68.5px' : '80px' }"
    >
      <NLayoutSider
        v-if="!media.isMobile"
        collapse-mode="width"
        :show-collapsed-content="false"
        :collapsed-width="20"
        show-trigger="arrow-circle"
        @update:collapsed="hideSider = $event"
        :bordered="!hideSider"
        :class="hideSider ? 'transparent' : null"
      >
        <SideMenu />
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
import { ref } from "vue";
import { LayoutInst } from "naive-ui";
import CollectionPage from "../components/CollectionPage.vue";
import SideMenu from "../components/SideMenu.vue";
import { useMedia } from "../store/Media";

const media = useMedia();

defineProps<{
  tid?: string;
}>();

const hideSider = ref(false);

const contentRef = ref<LayoutInst | null>(null);
</script>

<style scoped>
.transparent {
  background-color: transparent;
}
.content {
  padding: 10px 20px;
}
.layout {
  height: 100vh;
}

@media screen and (min-width: 768px) {
  .content {
    padding-right: 80px;
  }
}
</style>
