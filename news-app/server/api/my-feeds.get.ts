import type { Feed } from "@/types/feed";

export default defineEventHandler(async (event): Promise<Feed[]> => {
  const config = useRuntimeConfig(event);
  const url = `${config.apiBaseUrl}/feed`;

  const headers = event.node.req.headers;

  return await $fetch<Feed[]>(url, {
    method: "GET",
    responseType: "json",
    credentials: "include",
    headers: {
      Cookie: headers.cookie || "",
    },
  });
});
