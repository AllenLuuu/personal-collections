import { post } from "./request";
import { useLoginInfoStore } from "../store/LoginInfo";

export async function login(username: string, password: string) {
  const loginInfoStore = useLoginInfoStore();
  try {
    const res = await post<{
      id: string;
      username: string;
      password: string;
    }>(
      "/login",
      {
        username,
        password,
      },
      "登录"
    );
    
    loginInfoStore.setLoginInfo({
      userId: res.id,
      username: res.username,
      hasLoggedIn: true,
    });

    return true;
  } catch (error: any) {
    return false;
  }
}
