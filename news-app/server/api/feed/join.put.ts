type Request = { feedId: number };

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig(event);
  const request = await readBody<Request>(event);
  const headers = event.node.req.headers;

  const url = `${config.apiBaseUrl}/feed/join`;

  return await $fetch(url, {
    method: "PUT",
    credentials: "include",
    body: JSON.stringify(request),
    headers: {
      Cookie: headers.cookie || "",
    },
  });
});
