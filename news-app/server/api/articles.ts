import { Article } from "@/types/article";

export default defineEventHandler(async (event): Promise<Article[]> => {
  const config = useRuntimeConfig(event);
  const url = `${config.apiBaseUrl}/articles`;

  console.log(getQuery(event));

  const authToken = getRequestHeader(event, "Authorization");

  if (!authToken) {
    console.error("Missing Authorization Header");
    throw createError({
      statusCode: 401,
      message: 'Missing required header "Authorization"',
    });
  }

  return await $fetch<Article[]>(url, {
    method: "GET",
    responseType: "json",
    headers: {
      Authorization: authToken,
    },
  });
});
