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

package domain

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var quizNameRegexp = regexp.MustCompile(`^# (?P<quizName>.*) \(duration: (?P<duration>[0-9]+)min\)`)
var quizQuestionRegexp = regexp.MustCompile(`^# .*\n`)
var quizquestionCodeRegexp = regexp.MustCompile("```(?P<language>.*)\\n(?s)(?P<code>.*?)\\n```")
var quizAnswersRegexp = regexp.MustCompile(`(- \[[ xX]] .*\n)+`)
var quizAnswerRegexp = regexp.MustCompile(`- \[[ xX]] .*`)
var quizValidAnswerRegexp = regexp.MustCompile(`- \[[xX]] .*`)

// Parse parse the content of a quiz file
func (s *QuizService) Parse(filename string, content string) (*Quiz, error) {

	name, duration, err := extractQuizNameAndDuration(content)
	if err != nil {
		return nil, err
	}

	questions, err := extractQuestions(content)
	if err != nil {
		return nil, err
	}

	return &Quiz{
		Sha1:      getSha1(content),
		Name:      name,
		Filename:  filename,
		CreatedAt: time.Now().Format(time.RFC3339),
		Version:   1,
		Duration:  duration,
		Questions: questions,
	}, nil
}

func getSha1(content string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(content))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func extractQuizNameAndDuration(content string) (string, int, error) {
	subMatch := quizNameRegexp.FindStringSubmatch(content)

	if len(subMatch) < 3 {
		return "", 0, fmt.Errorf("quiz name or quiz duration not found. The first line must be '# <Name> (duration: <duration>min)")
	}

	name := subMatch[1]
	durationMin, err := strconv.ParseInt(subMatch[2], 10, 32)
	if err != nil {
		return "", 0, err
	}

	return name, int(durationMin) * 60, nil
}

func extractQuestions(content string) (map[string]QuizQuestion, error) {
	quizName := quizQuestionRegexp.FindString(content)
	questionsStr := strings.ReplaceAll(content, quizName, "")
	questionsUnParsed := strings.Split(questionsStr, "---\n")

	questions := map[string]QuizQuestion{}

	for i, s := range questionsUnParsed {
		question, err := extractQuestion(s)
		if err != nil {
			return nil, err
		}

		question.Position = i + 1
		questions[question.Sha1] = question
	}

	return questions, nil
}

func extractQuestion(content string) (QuizQuestion, error) {

	answersStr := quizAnswersRegexp.FindString(content)

	questionContent, code, language := extractQuestionCode(strings.ReplaceAll(content, answersStr, ""))

	answers, err := extractAnswers(answersStr)
	if err != nil {
		return QuizQuestion{}, err
	}

	return QuizQuestion{
		Sha1:         getSha1(content),
		Content:      strings.Trim(questionContent, " \n"),
		Code:         code,
		CodeLanguage: language,
		Answers:      answers,
	}, nil
}

func extractQuestionCode(questionContent string) (content string, code string, language string) {
	subMatch := quizquestionCodeRegexp.FindStringSubmatch(questionContent)

	if len(subMatch) < 3 {
		return questionContent, "", ""
	}

	language = subMatch[1]
	code = subMatch[2]
	content = strings.ReplaceAll(questionContent, subMatch[0], "")

	return content, code, language
}

func extractAnswers(answersStr string) (map[string]QuizQuestionAnswer, error) {

	answersStrSplit := quizAnswerRegexp.FindAllString(answersStr, 10)

	answers := map[string]QuizQuestionAnswer{}

	for _, s := range answersStrSplit {
		valid := quizValidAnswerRegexp.MatchString(s)
		sha1Str := getSha1(s)

		answers[sha1Str] = QuizQuestionAnswer{
			Sha1:    sha1Str,
			Content: string([]rune(s)[6:]),
			Valid:   valid,
		}
	}

	return answers, nil
}
