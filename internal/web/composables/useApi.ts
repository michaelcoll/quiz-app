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
  const token = await useAuthStore().getToken;

  return useFetch<T, R>(request, {
    ...options,

    headers: {
      Authorization: `Bearer ${token}`,
      ...options?.headers,
    },

    onResponseError({ response }) {
      console.log(response);
      useToast().danger(`Error : ${response._data.message}`);
    },
  });
}
