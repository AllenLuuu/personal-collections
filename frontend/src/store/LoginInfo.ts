import { defineStore } from "pinia";

interface LoginInfo {
  userId: string;
  username: string;
  hasLoggedIn: boolean;
}

export const useLoginInfoStore = defineStore("login-info", {
  state: () => ({
    loginInfo: {
      userId: "",
      username: "",
      hasLoggedIn: false,
    },
  }),
  getters: {
    getLoginInfo(): LoginInfo {
      return this.loginInfo;
    },
    isLoggedIn(): boolean {
      return this.loginInfo.hasLoggedIn;
    }
  },
  actions: {
    setLoginInfo(loginInfo: LoginInfo) {
      this.loginInfo = loginInfo;
    },
  },
  persist: {
    key: "login-info",
    storage: sessionStorage,
  },
});
