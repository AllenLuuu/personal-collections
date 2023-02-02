<template>
  <NLayout class="layout">
    <NLayoutHeader bordered style="height: 80px">
      <Header suffix="管理端" show-home></Header>
    </NLayoutHeader>
    <NLayout
      :native-scrollbar="false"
      position="absolute"
      style="top: 80px; padding: 10px"
    >
      <NCard>
        <NForm ref="formRef" :model="target" :rules="rules">
          <NFormItem path="title" label="标题">
            <NInput v-model:value="target.title" placeholder="请输入专题标题" />
          </NFormItem>
          <NFormItem path="detail" label="介绍">
            <NInput
              type="textarea"
              v-model:value="target.detail"
              placeholder="请输入内容"
            />
          </NFormItem>
          <NSpace justify="end">
            <NButton @click="cancel"> 取消 </NButton>
            <NButton type="primary" @click="submit"> 发布 </NButton>
          </NSpace>
        </NForm>
      </NCard>
    </NLayout>
  </NLayout>
</template>

<script setup lang="ts">
import type { FormInst, FormRules } from "naive-ui";
import { useMessage } from "naive-ui";
import { onMounted, ref } from "vue";
import Header from "../../components/Header.vue";
import { useRouter } from "vue-router";

const router = useRouter();
const message = useMessage();

const props = defineProps<{
  tid?: string;
}>();

onMounted(() => {
  if (props.tid) {
    setTopic(props.tid);
  }
});

const formRef = ref<FormInst | null>(null);

const target = ref({
  id: "",
  title: "",
  detail: "",
  collections: [],
});

const rules: FormRules = {
  title: {
    required: true,
    message: "请输入专题标题",
    trigger: "blur",
  },
  detail: {
    required: true,
    message: "请输入专题简介",
    trigger: "blur",
  },
};

function setTopic(tid: string) {
  target.value = {
    id: tid,
    title: "test",
    detail: "test",
    collections: [],
  };
}

function cancel() {
  router.back();
}

function submit() {
  formRef.value?.validate((errors) => {
    if (!errors) {
      message.success("发布成功");
      router.back();
    }
  });
}
</script>

<style scoped>
.layout {
  height: 100vh;
}
</style>
