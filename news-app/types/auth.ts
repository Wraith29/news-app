import { reactive } from "vue";

export type AuthRequest = {
  username: string;
  password: string;
};

export type AuthResponse = {
  authToken: string;
};

export const authStore = reactive({
  loggedIn: false,
  authToken: "",
});
