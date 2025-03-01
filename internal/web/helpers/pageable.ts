import type { FetchResponse } from "ofetch";

export function toRangeHeader(itemName: string, page: number, pageSize: number): string {
  if (page === 0) {
    return `${itemName}=0-${pageSize - 1}`;
  }
  if (pageSize === 0) {
    return `${itemName}=${5 * (page - 1)}-${5 * page - 1}`;
  }
  return `${itemName}=${pageSize * (page - 1)}-${pageSize * page - 1}`;
}

export function extractTotalFromHeader(response: FetchResponse<any>): number {
  const contentRangeHeader = response.headers.get("content-range");
  return extractTotalFromCententRangeHeaderValue(contentRangeHeader);
}

export function extractTotalFromCententRangeHeaderValue(
  contentRangeHeader: string | null,
): number {
  if (contentRangeHeader != null) {
    const split = contentRangeHeader.split("/");
    if (split.length === 2) {
      return parseInt(split[1]);
    }
  }

  return 0;
}
