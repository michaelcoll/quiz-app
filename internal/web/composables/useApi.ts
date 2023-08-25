import { NitroFetchRequest } from "nitropack";
import { useToast } from "tailvue";

import { useAuthStore } from "~/stores/auth";

export async function useApi<
  T = unknown,
  R extends NitroFetchRequest = NitroFetchRequest,
>(
  request: Parameters<typeof useFetch<T, R>>[0],
  options?: Partial<Parameters<typeof useFetch<T, R>>[1]>,
) {
  const config = useRuntimeConfig();
  const token = await useAuthStore().getToken;

  return useFetch<T, R>(request, {
    ...options,
    baseURL: config.public.apiBase,

    headers: {
      Authorization: `Bearer ${token}`,
      ...options?.headers,
    },

    onResponseError({ response }) {
      useToast().danger(`Error : ${response._data.message}`);
    },
  });
}
