import { AuthRequest } from "~/types/auth";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig(event);
  const request = await readBody<AuthRequest>(event);

  const url = `${config.apiBaseUrl}/auth/login`;

  return await $fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(request),
    responseType: "json",
    credentials: "include",
    async onResponse({ response }) {
      setResponseHeader(event, "Set-Cookie", response.headers.getSetCookie());
    },
  });
});
