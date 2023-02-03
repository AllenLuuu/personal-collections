<template>
  <NLayout class="layout">
    <NLayoutHeader bordered style="height: 80px">
      <Header suffix="管理端"></Header>
    </NLayoutHeader>
    <NLayoutContent
      position="absolute"
      style="top: 80px; height: calc(100vh - 80px)"
    >
      <div class="content">
        <NCard title="登录" style="max-width: 400px">
          <NForm ref="formRef" show-require-mark :model="model" :rules="rules">
            <NFormItem label="用户名" path="username">
              <NInput v-model:value="model.username" />
            </NFormItem>
            <NFormItem label="密码" path="password">
              <NInput
                type="password"
                show-password-on="click"
                v-model:value="model.password"
              />
            </NFormItem>
            <NFormItem class="center">
              <NButton type="primary" @click="onSubmit"> 登录 </NButton>
            </NFormItem>
          </NForm>
        </NCard>
      </div>
    </NLayoutContent>
  </NLayout>
</template>

<script setup lang="ts">
import Header from "../../components/Header.vue";
import { ref, reactive } from "vue";
import type { FormInst, FormRules } from "naive-ui";
import { login } from "../../utils/login";
import { useRouter } from "vue-router";

const router = useRouter();

const formRef = ref<FormInst | null>(null);

const model = reactive({
  username: "",
  password: "",
});

const rules: FormRules = {
  username: {
    required: true,
    message: "请输入用户名",
    trigger: "blur",
  },
  password: {
    required: true,
    message: "请输入密码",
    trigger: "blur",
  },
};

const onSubmit = () => {
  formRef.value?.validate(async (errors) => {
    if (!errors) {
      const loggedIn = await login(model.username, model.password);
      if (loggedIn) {
        router.push("/admin");
      }
    }
  });
};
</script>

<style scoped>
.layout {
  height: 100vh;
}
.content {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}
.center {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
