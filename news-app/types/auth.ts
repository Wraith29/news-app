export type AuthRequest = {
  username: string;
  password: string;
};

export const authStore = reactive({
  loggedIn: false,
  authToken: "",
});
