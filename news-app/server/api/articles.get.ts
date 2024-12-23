import { Article } from "@/types/article";

export default defineEventHandler(async (event): Promise<Article[]> => {
  const config = useRuntimeConfig(event);
  const url = `${config.apiBaseUrl}/articles`;

  const headers = event.node.req.headers;

  return await $fetch<Article[]>(url, {
    method: "GET",
    responseType: "json",
    credentials: "include",
    headers: {
      Cookie: headers.cookie || "",
    },
  });
});
