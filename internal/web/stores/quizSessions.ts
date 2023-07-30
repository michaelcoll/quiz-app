/*
 * Copyright (c) 2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { defineStore } from "pinia";

import { getApi } from "~/api/common-api";
import { QuizSession, Session } from "~/api/model";

const pageSize = 8;

export type QuizSessionState = {
  quizSessions: QuizSession[] | null;
  page: number;
  lastPage: number;
  total: number;
};
export const useQuizSessionsStore = defineStore("quiz-sessions", {
  state: () =>
    ({
      quizSessions: null,
      page: 0,
      total: 0,
    }) as QuizSessionState,

  getters: {
    getQuizSessions(): QuizSession[] | null {
      return this.quizSessions;
    },
    getCurrentPage(): number {
      return this.page;
    },
    getLastPage(): number {
      return this.total / pageSize;
    },
    getTotal(): number {
      return this.total;
    },
    hasMoreThanOnePage(): boolean {
      return this.total > pageSize;
    },
  },
  actions: {
    fetchQuizSessions() {
      const start = pageSize * this.page;
      const end = pageSize * (this.page + 1) - 1;

      getApi()
        .get<QuizSession[]>(`/api/v1/quiz-session`, {
          headers: {
            Range: `quiz-session=${start}-${end}`,
          },
        })
        .then((res) => {
          const contentRangeHeader = res.headers["content-range"];
          const split = contentRangeHeader.split("/");

          this.quizSessions = res.data;
          this.total = parseInt(split[1]);
        });
    },
    startQuiz(sha1: string): Promise<Session> {
      return getApi()
        .post<Session>(
          `/api/v1/session`,
          {},
          {
            params: {
              quizSha1: sha1,
            },
          },
        )
        .then((res) => {
          return res.data;
        });
    },
    setPage(page: number) {
      this.page = page;
      this.fetchQuizSessions();
    },
    incrementPage() {
      this.page++;
      this.fetchQuizSessions();
    },
    decrementPage() {
      this.page--;
      if (this.page < 0) {
        this.page = 0;
      }
      this.fetchQuizSessions();
    },
  },
});
