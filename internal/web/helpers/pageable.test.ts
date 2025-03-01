import { describe, expect, it } from "vitest";

import {
  extractTotalFromCententRangeHeaderValue,
  toRangeHeader,
} from "~/helpers/pageable";

describe("toRangeHeader", () => {
  it("returns correct range header for given item name, page, and page size", () => {
    expect(toRangeHeader("items", 1, 10)).toBe("items=0-9");
    expect(toRangeHeader("items", 2, 10)).toBe("items=10-19");
  });

  it("handles edge case where page is 0", () => {
    expect(toRangeHeader("items", 0, 10)).toBe("items=0-9");
  });

  it("handles edge case where page size is 0", () => {
    expect(toRangeHeader("items", 1, 0)).toBe("items=0-4");
  });
});

describe("extractTotalFromCententRangeHeaderValue", () => {
  it("extracts total from content-range header", () => {
    expect(extractTotalFromCententRangeHeaderValue("items 0-9/100")).toBe(100);
  });

  it("returns 0 if content-range header is missing", () => {
    expect(extractTotalFromCententRangeHeaderValue(null)).toBe(0);
  });

  it("returns 0 if content-range header is malformed", () => {
    expect(extractTotalFromCententRangeHeaderValue("invalid-header")).toBe(0);
  });
});
