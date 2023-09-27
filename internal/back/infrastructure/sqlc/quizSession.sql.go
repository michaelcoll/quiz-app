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

package sqlc

import (
	"context"
	"fmt"
)

const findAllQuizSessions = `
SELECT quiz_sha1, quiz_name, quiz_filename, quiz_version, quiz_duration, quiz_created_at, session_uuid, user_id, user_name, user_picture, class_uuid, class_name, remaining_sec, checked_answers, results
FROM quiz_session_view 
%s
LIMIT ? OFFSET ?
`

type FindAllQuizSessionsParams struct {
	Limit   int64  `db:"limit"`
	Offset  int64  `db:"offset"`
	ClassId string `db:"classId"`
	UserId  string `db:"userId"`
}

func (q *Queries) FindAllQuizSessions(ctx context.Context, arg FindAllQuizSessionsParams) ([]QuizSessionView, error) {
	req, args := findAllQuizSessionsQueryWithWhereClause(arg)
	rows, err := q.db.QueryContext(ctx, req, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []QuizSessionView
	for rows.Next() {
		var i QuizSessionView
		if err := rows.Scan(
			&i.QuizSha1,
			&i.QuizName,
			&i.QuizFilename,
			&i.QuizVersion,
			&i.QuizDuration,
			&i.QuizCreatedAt,
			&i.SessionUuid,
			&i.UserID,
			&i.UserName,
			&i.UserPicture,
			&i.ClassUuid,
			&i.ClassName,
			&i.RemainingSec,
			&i.CheckedAnswers,
			&i.Results,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func findAllQuizSessionsQueryWithWhereClause(arg FindAllQuizSessionsParams) (req string, args []interface{}) {
	whereClause := ""
	classWhereClause := "class_uuid = ?"

	if arg.ClassId != "" {
		whereClause = "WHERE " + classWhereClause
		args = append(args, arg.ClassId)
	}

	args = append(args, arg.Limit, arg.Offset)

	return fmt.Sprintf(findAllQuizSessions, whereClause), args
}

const findAllQuizSessionsForUser = `
SELECT q.sha1                                                                  AS quiz_sha1,
       q.name                                                                  AS quiz_name,
       q.filename                                                              AS quiz_filename,
       q.version                                                               AS quiz_version,
       q.duration                                                              AS quiz_duration,
       q.created_at                                                            AS quiz_created_at,
       CASE WHEN s.uuid IS NULL THEN '' ELSE s.uuid END                        AS session_uuid,
       CASE WHEN s.user_id IS NULL THEN '' ELSE s.user_id END                  AS user_id,
       CASE WHEN u.name IS NULL THEN '' ELSE u.name END                        AS user_name,
       CASE WHEN u.picture IS NULL THEN '' ELSE u.picture END                  AS user_picture,
       CASE WHEN sc.uuid IS NULL THEN '' ELSE sc.uuid END                      AS class_uuid,
       CASE WHEN sc.name IS NULL THEN '' ELSE sc.name END                      AS class_name,
       CASE WHEN sv.remaining_sec IS NULL THEN 0 ELSE sv.remaining_sec END     AS remaining_sec,
       CASE WHEN sv.checked_answers IS NULL THEN 0 ELSE sv.checked_answers END AS checked_answers,
       CASE WHEN sv.results IS NULL THEN 0 ELSE sv.results END                 AS results
FROM quiz q
         JOIN quiz_class_visibility qcv ON q.sha1 = qcv.quiz_sha1
         JOIN student_class sc ON qcv.class_uuid = sc.uuid
         JOIN user u ON qcv.class_uuid = u.class_uuid
         LEFT JOIN session s ON q.sha1 = s.quiz_sha1 AND s.user_id = u.id
         LEFT JOIN session_view sv ON sv.uuid = s.uuid
WHERE q.active = TRUE
AND u.id = ?
LIMIT ? OFFSET ?
`

func (q *Queries) FindAllQuizSessionsForUser(ctx context.Context, arg FindAllQuizSessionsParams) ([]QuizSessionView, error) {
	rows, err := q.db.QueryContext(ctx, findAllQuizSessionsForUser, arg.UserId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []QuizSessionView
	for rows.Next() {
		var i QuizSessionView
		if err := rows.Scan(
			&i.QuizSha1,
			&i.QuizName,
			&i.QuizFilename,
			&i.QuizVersion,
			&i.QuizDuration,
			&i.QuizCreatedAt,
			&i.SessionUuid,
			&i.UserID,
			&i.UserName,
			&i.UserPicture,
			&i.ClassUuid,
			&i.ClassName,
			&i.RemainingSec,
			&i.CheckedAnswers,
			&i.Results,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
