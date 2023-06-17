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

import { getApi } from "@/api/common-api";
import { QuizSession } from "@/api/model";

const pageSize = 15;

export const useQuizSessionStore = defineStore("quiz-session", {
  state: () => {
    return {
      quizSessions: null,
      page: 0,
    };
  },
  getters: {
    getQuizSessions(): QuizSession[] {
      return this.quizSessions;
    },
    getCurrentPage(): number {
      return this.page;
    },
  },
  actions: {
    fetchQuizSessions() {
      const start = pageSize * this.page;
      const end = pageSize * (this.page + 1) - 1;

      getApi()
        .then((axiosInstance) =>
          axiosInstance.get<QuizSession[]>(`/api/v1/quiz-session`, {
            headers: {
              Range: `quiz-session=${start}-${end}`,
            },
          })
        )
        .then(({ data }) => {
          this.quizSessions = data;
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
