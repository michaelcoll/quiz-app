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
	userWhereClause := "user_id = ?"
	classWhereClause := "class_uuid = ?"

	if arg.ClassId != "" && arg.UserId != "" {
		whereClause = "WHERE " + userWhereClause + " AND " + classWhereClause
		args = append(args, arg.UserId, arg.ClassId)
	} else if arg.ClassId != "" {
		whereClause = "WHERE " + classWhereClause
		args = append(args, arg.ClassId)
	} else if arg.UserId != "" {
		whereClause = "WHERE " + userWhereClause
		args = append(args, arg.UserId)
	}

	args = append(args, arg.Limit, arg.Offset)

	return fmt.Sprintf(findAllQuizSessions, whereClause), args
}
