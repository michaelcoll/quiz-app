import { FetchResponse } from "ofetch";

export function toRangeHeader(itemName: string, page: number, pageSize: number): string {
  return `${itemName}=${pageSize * (page - 1)}-${pageSize * page - 1}`;
}

export function extractTotalFromHeader(response: FetchResponse<any>): number {
  const contentRangeHeader = response.headers.get("content-range");
  if (contentRangeHeader != null) {
    const split = contentRangeHeader.split("/");
    return parseInt(split[1]);
  }

  return 0;
}
