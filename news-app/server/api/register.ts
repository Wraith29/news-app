import { AuthRequest, AuthResponse } from "~/types/auth";

export default defineEventHandler(async (event): Promise<AuthResponse> => {
  const config = useRuntimeConfig(event);
  const request = await readBody<AuthRequest>(event);

  const url = `${config.apiBaseUrl}/auth/register`;

  return await $fetch(url, {
    method: "POST",
    body: JSON.stringify(request),
  });
});
