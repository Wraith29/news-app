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

export function logIn(authToken: string): void {
  authStore.loggedIn = true;
  authStore.authToken = authToken;

  if (localStorage)
    localStorage.setItem("authToken", authToken);
}

export function logOut(): void {
  authStore.loggedIn = false;
  authStore.authToken = "";

  if (localStorage)
    localStorage.removeItem("authToken");
}
