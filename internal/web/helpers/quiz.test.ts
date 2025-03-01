import { describe, expect, it } from "vitest";

import type { SessionResult } from "~/api/model";
import { toDurationStr, toPercent } from "~/helpers/quiz";

describe("toDurationStr", () => {
  it("returns correct duration string for given seconds", () => {
    expect(toDurationStr(125)).toBe("2 min 5 sec");
    expect(toDurationStr(60)).toBe("1 min");
    expect(toDurationStr(0)).toBe("");
  });

  it("returns empty string for undefined input", () => {
    expect(toDurationStr(undefined)).toBe("");
  });
});

describe("toPercent", () => {
  it("returns correct percentage for given session result", () => {
    const result = { totalAnswer: 10, goodAnswer: 7 } as SessionResult;
    expect(toPercent(result)).toBe(70);
  });

  it("returns 0 if totalAnswer is 0", () => {
    const result = { totalAnswer: 0, goodAnswer: 0 } as SessionResult;
    expect(toPercent(result)).toBe(0);
  });

  it("returns 0 if result is null or undefined", () => {
    expect(toPercent(null)).toBe(0);
    expect(toPercent(undefined)).toBe(0);
  });

  it("handles case where goodAnswer is undefined", () => {
    const result = { totalAnswer: 10, goodAnswer: undefined } as SessionResult;
    expect(toPercent(result)).toBe(0);
  });
});
