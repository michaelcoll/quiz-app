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

import { AxiosResponse } from "axios";
import { defineStore } from "pinia";

import { getApi } from "~/api/common-api";
import { Message, QuizSessionDetail } from "~/api/model";

export type QuizSessionState = {
  quizSession: QuizSessionDetail | null;
};
export const useQuizSessionStore = defineStore("quiz-session", {
  state: () =>
    ({
      quizSession: null,
    }) as QuizSessionState,

  getters: {
    getQuizSession(): QuizSessionDetail | null {
      return this.quizSession;
    },
  },
  actions: {
    fetchQuizSession(sessionUuid: string) {
      getApi()
        .get<QuizSessionDetail>(`/api/v1/quiz-session/${sessionUuid}`)
        .then(({ data }) => {
          this.quizSession = data;
        });
    },
    saveAnswer(
      sessionUuid: string,
      questionSha1: string,
      answerSha1: string,
      checked: boolean,
    ): Promise<AxiosResponse<Message>> {
      return getApi().post<Message>(`/api/v1/session/${sessionUuid}/answer`, {
        questionSha1,
        answerSha1,
        checked,
      });
    },
  },
});
