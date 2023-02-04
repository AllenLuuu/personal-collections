import { post } from "./request";

export async function login(username: string, password: string) {
  try {
    await post<{
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
    return true;
  } catch (error: any) {
    return false;
  }
}
