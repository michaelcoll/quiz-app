import { NitroFetchRequest } from "nitropack";

export function usePostApi<T = unknown, R extends NitroFetchRequest = NitroFetchRequest>(
  request: Parameters<typeof useApi<T, R>>[0],
  options?: Partial<Parameters<typeof useApi<T, R>>[1]>,
) {
  return useApi<T, R>(request, {
    ...options,
    method: "POST",
  });
}
