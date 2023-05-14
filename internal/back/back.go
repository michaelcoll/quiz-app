/*
 * Copyright (c) 2022-2023 Michaël COLL.
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

package back

import (
	"github.com/school-by-hiit/quiz-app/internal/back/domain/service"
	"github.com/school-by-hiit/quiz-app/internal/back/infrastructure/infra_repository"
	"github.com/school-by-hiit/quiz-app/internal/back/presentation"
)

type Module struct {
	s        service.QuizService
	quizCtrl presentation.ApiController
}

func (m *Module) GetPhotoController() *presentation.ApiController {
	return &m.quizCtrl
}

func (m *Module) GetService() *service.QuizService {
	return &m.s
}

func New() Module {
	repository := infra_repository.New()
	quizService := service.New(repository)

	return Module{
		s:        quizService,
		quizCtrl: presentation.NewApiController(&quizService),
	}
}
