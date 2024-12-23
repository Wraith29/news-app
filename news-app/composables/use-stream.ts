export async function useStream<T>(body: ReadableStream): Promise<T> {
  const reader = body.pipeThrough(new TextDecoderStream()).getReader();

  let messageBody = "";
  while (true) {
    const { done, value } = await reader.read();
    if (done) break;

    messageBody += value;
  }

  const result = JSON.parse(messageBody);

  return result as T;
}
