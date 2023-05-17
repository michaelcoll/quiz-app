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

package presentation

import "fmt"

type HttpStatusError struct {
	status  int
	message string
}

func (e *HttpStatusError) HTTPStatus() int {
	return e.status
}

func (e *HttpStatusError) Error() string {
	return e.message
}

func Errorf(status int, format string, a ...interface{}) error {
	return &HttpStatusError{status, fmt.Sprintf(format, a...)}
}

func fromError(err error) (int, bool) {
	if se, ok := err.(interface {
		HTTPStatus() int
	}); ok {
		return se.HTTPStatus(), true
	}

	return 0, false
}
