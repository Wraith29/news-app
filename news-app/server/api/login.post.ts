import { AuthRequest } from "~/types/auth";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig(event);
  const request = await readBody<AuthRequest>(event);

  const url = `${config.apiBaseUrl}/auth/login`;

  const result = await $fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(request),
    responseType: "json",
    credentials: "include",
    async onResponse({ response }) {
      console.log("Login Response: ", response.headers);
      // event.node.res.setHeader("Set-Cookie", response.headers.getSetCookie());
      setResponseHeader(event, "Set-Cookie", response.headers.getSetCookie());
    },
  });

  return result;
});
