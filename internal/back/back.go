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
	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/db"
	"github.com/michaelcoll/quiz-app/internal/back/presentation"
)

type Module struct {
	quizServ domain.QuizService
	authServ domain.AuthService
	quizCtrl presentation.ApiController
}

func (m *Module) GetPhotoController() *presentation.ApiController {
	return &m.quizCtrl
}

func (m *Module) GetService() *domain.QuizService {
	return &m.quizServ
}

func New() Module {
	connection := db.Init("data")

	authRepository := infrastructure.NewAuthRepository(connection)
	classRepository := infrastructure.NewClassRepository(connection)
	quizRepository := infrastructure.NewQuizRepository(connection)
	userRepository := infrastructure.NewUserRepository(connection)

	authService := domain.NewAuthService(authRepository, userRepository)
	classService := domain.NewClassService(classRepository)
	quizService := domain.NewQuizService(quizRepository)
	userService := domain.NewUserService(userRepository)

	return Module{
		quizServ: quizService,
		authServ: authService,
		quizCtrl: presentation.NewApiController(&authService, &classService, &quizService, &userService),
	}
}
