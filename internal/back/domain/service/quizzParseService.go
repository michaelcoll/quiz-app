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

package service

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"github.com/school-by-hiit/quizz-app/internal/back/domain/model"
)

// Parse parse the content of a quizz file
func (s *QuizzService) Parse(filename string, content string) (model.Quizz, error) {

	name, err := extractQuizzName(content)
	if err != nil {
		return model.Quizz{}, err
	}

	questions, err := extractQuestions(content)
	if err != nil {
		return model.Quizz{}, err
	}

	return model.Quizz{
		Sha1:      getSha1(content),
		Name:      name,
		Filename:  filename,
		Questions: questions,
	}, nil
}

func getSha1(content string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(content))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func extractQuizzName(content string) (string, error) {
	r := regexp.MustCompile(`^# .*`)

	if r.MatchString(content) {
		return string([]rune(r.FindString(content))[2:]), nil
	} else {
		return "", fmt.Errorf("no quizz name found")
	}
}

func extractQuestions(content string) ([]model.QuizzQuestion, error) {
	r := regexp.MustCompile(`^# .*\n`)

	quizzName := r.FindString(content)
	questionsStr := strings.ReplaceAll(content, quizzName, "")
	questionsUnParsed := strings.Split(questionsStr, "---\n")

	questions := make([]model.QuizzQuestion, len(questionsUnParsed))

	for i, s := range questionsUnParsed {
		question, err := extractQuestion(s)
		if err != nil {
			return nil, err
		}

		questions[i] = question
	}

	return questions, nil
}

func extractQuestion(content string) (model.QuizzQuestion, error) {

	r := regexp.MustCompile(`(- \[[ xX]] .*\n)+`)
	answersStr := r.FindString(content)

	questionContent := strings.ReplaceAll(content, answersStr, "")

	answers, err := extractAnswers(answersStr)
	if err != nil {
		return model.QuizzQuestion{}, err
	}

	return model.QuizzQuestion{
		Sha1:    getSha1(content),
		Content: strings.Trim(questionContent, " \n"),
		Answers: answers,
	}, nil
}

func extractAnswers(answersStr string) ([]model.QuizzQuestionAnswer, error) {

	r := regexp.MustCompile(`- \[[ xX]] .*`)
	validTestRegex := regexp.MustCompile(`- \[[xX]] .*`)
	answersStrSplit := r.FindAllString(answersStr, 10)

	answers := make([]model.QuizzQuestionAnswer, len(answersStrSplit))

	for i, s := range answersStrSplit {
		valid := validTestRegex.MatchString(s)

		answers[i] = model.QuizzQuestionAnswer{
			Sha1:    getSha1(s),
			Content: string([]rune(s)[6:]),
			Valid:   valid,
		}
	}

	return answers, nil
}
