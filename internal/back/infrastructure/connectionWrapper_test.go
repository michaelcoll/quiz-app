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

package infrastructure

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewConnectionWrapperForTest(dbLocation string, c *sql.DB) *ConnectionWrapper {
	return &ConnectionWrapper{dbLocation: dbLocation, c: c}
}

func TestQueries_WhenConnectionIsClosed_ShouldReinitializeConnection(t *testing.T) {
	// Given
	wrapper := NewConnectionWrapper("test_db_location")
	err := wrapper.Close()
	if err != nil {
		assert.Fail(t, "Error while closing connection", err)
	}

	assert.True(t, wrapper.isClosed)

	// When
	result := wrapper.queries()

	// Then
	assert.NotNil(t, result)
	assert.False(t, wrapper.isClosed)
}

func TestQueries_WhenConnectionIsOpen_ShouldNotReinitializeConnection(t *testing.T) {
	// Given
	wrapper := NewConnectionWrapper("test_db_location")

	// When
	result := wrapper.queries()

	// Then
	assert.NotNil(t, result)
	assert.False(t, wrapper.isClosed)
}
