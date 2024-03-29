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

// Code generated by sqlc-addon. DO NOT EDIT.
// versions:
//   sqlc-addon v1.9.4

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(readOnly bool, databaseLocation string) *sql.DB {
	_, err := os.Stat(databaseLocation)
	if err != nil {
		err := os.MkdirAll(databaseLocation, 0700)
		if err != nil {
			log.Fatalf("Can't create database folder : %v (%v)\n", databaseLocation, err)
			os.Exit(-1)
		}
	}

	db, err := sql.Open("sqlite3", getDBUrl(readOnly, databaseLocation))
	if err != nil {
		log.Fatalf("Can't open database %v", err)
	}

	return db
}

func getDBUrl(readOnly bool, databaseLocation string) string {

	var options string
	if readOnly {
		options = "cache=shared&mode=ro"
	} else {
		options = "cache=shared&mode=rwc&_auto_vacuum=full&_journal_mode=WAL"
	}

	return fmt.Sprintf("file:%s/%s?%s", databaseLocation, "data.db", options)
}

func Init(databaseLocation string) *sql.DB {
	logLocation(databaseLocation)
	conn := Connect(false, databaseLocation)
	activateForeignKeys(conn)
	New(conn).Migrate()

	// Handles the CTRL+C properly
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		conn.Close()
		os.Exit(0)
	}()

	return conn
}

func GetDatabaseFullPath(databaseLocation string) string {
	fullPath := ""
	if strings.HasPrefix(databaseLocation, "/") {
		fullPath = fmt.Sprintf("%s/data.db", databaseLocation)
	} else {
		path, err := os.Getwd()
		if err != nil {
			log.Fatalf("Can't read current working directory (%v)\n", err)
			os.Exit(-1)
		}
		fullPath = fmt.Sprintf("%s/%s/data.db", path, databaseLocation)
	}

	return fullPath
}

func logLocation(databaseLocation string) {
	fmt.Printf("%s Storing database in %s\n",
		color.HiBlueString("i"),
		color.BlueString(GetDatabaseFullPath(databaseLocation)))
}

func activateForeignKeys(conn *sql.DB) {
	conn.Exec("PRAGMA foreign_keys = ON")
}
