<template>
  <NGrid :cols="3">
    <NGi class="left" :span="2">
      <div class="inline">
        <NIcon v-if="!media.isMobile" :size="40">
          <BookOutlined />
        </NIcon>
        <NButton v-else text :focusable="false" @click="toggleDrawer">
          <template #icon>
            <NIcon size="25">
              <MenuRound />
            </NIcon>
          </template>
        </NButton>
        <h2 v-if="media.isMobile" style="overflow: hidden; white-space: nowrap">
          游逛者·书摘 {{ suffix }}
        </h2>
        <h1 v-else style="overflow: hidden; white-space: nowrap">
          游逛者·书摘 {{ suffix }}
        </h1>
      </div>
    </NGi>
    <NGi class="right">
      <NSpace size="large">
        <NButton
          text
          :focusable="false"
          v-if="showHome"
          @click="router.push('/admin')"
        >
          <template #icon>
            <NIcon size="20">
              <HomeOutlined />
            </NIcon>
          </template>
        </NButton>
        <NDropdown
          trigger="hover"
          :options="options"
          @select="handleFontSelect"
        >
          <NButton text :focusable="false">
            <NIcon size="20">
              <FontDownloadOutlined />
            </NIcon>
          </NButton>
        </NDropdown>
        <NButton text :focusable="false" @click="mode.revertColorMode">
          <template #icon>
            <NIcon size="20">
              <LightModeOutlined v-if="mode.isDarkMode" />
              <DarkModeOutlined v-else />
            </NIcon>
          </template>
        </NButton>
      </NSpace>
    </NGi>
  </NGrid>

  <NDrawer
    v-model:show="showDrawer"
    :trap-focus="false"
    placement="left"
    width="250px"
    :block-scroll="false"
    to="#main-content"
  >
    <SideMenu />
  </NDrawer>
</template>

<script setup lang="ts">
import {
  BookOutlined,
  DarkModeOutlined,
  FontDownloadOutlined,
  HomeOutlined,
  LightModeOutlined,
  MenuRound,
} from "@vicons/material";
import {
  DropdownDividerOption,
  DropdownGroupOption,
  DropdownOption,
  DropdownRenderOption,
  NDropdown,
  NIcon,
} from "naive-ui";
import { computed, h, ref } from "vue";
import { useRouter } from "vue-router";
import { fonts } from "../const";
import { useColorMode } from "../store/ColorMode";
import { useFontStore } from "../store/Font";
import { useMedia } from "../store/Media";
import SideMenu from "./SideMenu.vue";

const media = useMedia();
const router = useRouter();
const mode = useColorMode();
const font = useFontStore();

defineProps<{
  suffix?: string;
  showHome?: boolean;
}>();

const options = computed<
  Array<
    | DropdownOption
    | DropdownGroupOption
    | DropdownDividerOption
    | DropdownRenderOption
  >
>(() =>
  fonts.map((f) => ({
    key: f.cssName,
    label: f.name,
    icon: () =>
      f.cssName === font.cssName
        ? h(
            "p",
            {
              style: {
                fontFamily: "serif",
              },
            },
            "✔"
          )
        : null,
    props: {
      style: {
        fontFamily: f.cssName,
        // border: f.cssName === font.cssName ? "2px solid #ccc" : "none",
        // borderRadius: "5px",
      },
    },
  }))
);

const handleFontSelect = (key: string) => {
  font.setFont(fonts.find((f) => f.cssName === key)!);
};

const showDrawer = ref(false);
const toggleDrawer = () => {
  showDrawer.value = !showDrawer.value;
};
</script>

<style scoped>
.left {
  padding-left: 50px;
}
.right {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  padding-right: 50px;
}
.inline {
  font-family: "KingHwa_OldSong", serif;
  display: inline-flex;
  align-items: center;
  gap: 15px;
}

@media screen and (max-width: 768px) {
  .left {
    padding-left: 20px;
  }
  .right {
    padding-right: 20px;
  }
}
</style>
