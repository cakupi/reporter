// Cakupi - An open-source, self-hostable code coverage platform
// Copyright (C) 2023  Reinaldy Rafli <aldy505@proton.me>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package reports_test

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"coverage-worker/provider"
	"coverage-worker/repository/reports"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var clickhouseClient driver.Conn
var clickhouseCluster string

func TestMain(m *testing.M) {
	clickhouseDsn, ok := os.LookupEnv("CLICKHOUSE_DSN")
	if !ok {
		clickhouseDsn = "clickhouse://default:@localhost:9000/default?dial_timeout=250ms&max_execution_time=60&debug=false"
	}

	clickhouseCluster, ok = os.LookupEnv("CLICKHOUSE_CLUSTER")
	if !ok {
		clickhouseCluster = ""
	}

	clickhouseOptions, err := clickhouse.ParseDSN(clickhouseDsn)
	if err != nil {
		log.Fatalf("Failed to parse clickhouse dsn: %s", err.Error())
		return
	}

	clickhouseConnection, err := clickhouse.Open(clickhouseOptions)
	if err != nil {
		log.Fatalf("Connecting to clickhouse: %s", err.Error())
	}

	setupCtx, setupCancel := context.WithTimeout(context.Background(), time.Minute)
	defer setupCancel()

	if clickhouseCluster != "" {
		err := clickhouseConnection.Exec(setupCtx, `CREATE DATABASE IF NOT EXISTS "cakupi" ON CLUSTER `+strconv.Quote(clickhouseCluster))
		if err != nil {
			setupCancel()
			if err := clickhouseConnection.Close(); err != nil {
				log.Printf("closing clickhouse connection: %s", err.Error())
			}

			log.Fatalf("executing migration: %s", err.Error())
			return
		}

		err = clickhouseConnection.Exec(
			setupCtx,
			`CREATE TABLE IF NOT EXISTS cakupi.coverage_reports_summary_cluster ON CLUSTER `+
				strconv.Quote(clickhouseCluster)+
				`(
					repository_name String NOT NULL,
					commit_hash String NOT NULL,
					commit_branch String NOT NULL,
					total_statements UInt64 NOT NULL,
					total_covered_statements UInt64 NOT NULL,
					number_of_files UInt64 NOT NULL,
					timestamp DateTime64(3, 'UTC')
				)
				ENGINE = ReplicatedMergeTree('/clickhouse/tables/01/cakupi/coverage_summary', '{replica}')
				PRIMARY KEY (repository_name, commit_hash)`,
		)
		if err != nil {
			setupCancel()
			if err := clickhouseConnection.Close(); err != nil {
				log.Printf("closing clickhouse connection: %s", err.Error())
			}

			log.Fatalf("executing migration: %s", err.Error())
			return
		}

		err = clickhouseConnection.Exec(
			setupCtx,
			`CREATE TABLE IF NOT EXISTS cakupi.coverage_reports_summary ON CLUSTER `+
				strconv.Quote(clickhouseCluster)+
				`(
					repository_name String NOT NULL,
					commit_hash String NOT NULL,
					commit_branch String NOT NULL,
					total_statements UInt64 NOT NULL,
					total_covered_statements UInt64 NOT NULL,
					number_of_files UInt64 NOT NULL,
					timestamp DateTime64(3, 'UTC')
				)
				ENGINE = Distributed(`+strconv.Quote(clickhouseCluster)+`, "cakupi", "coverage_summary")`,
		)
		if err != nil {
			setupCancel()
			if err := clickhouseConnection.Close(); err != nil {
				log.Printf("closing clickhouse connection: %s", err.Error())
			}

			log.Fatalf("executing migration: %s", err.Error())
			return
		}

		err = clickhouseConnection.Exec(
			setupCtx,
			`CREATE TABLE IF NOT EXISTS cakupi.coverage_reports_detail_cluster ON CLUSTER `+
				strconv.Quote(clickhouseCluster)+
				`(
					repository_name String NOT NULL,
					commit_hash String NOT NULL,
					commit_branch String NOT NULL,
					file_name String NOT NULL,
					file_path String NOT NULL,
					line_number UInt64 NOT NULL,
					hits_count UInt64 NOT NULL,
					fully_covered Bool DEFAULT TRUE,
					timestamp DateTime64(3, 'UTC')
				)
				ENGINE = ReplicatedMergeTree('/clickhouse/tables/01/cakupi/coverage_detail', '{replica}')
				PRIMARY KEY (repository_name, commit_hash)`,
		)
		if err != nil {
			setupCancel()
			if err := clickhouseConnection.Close(); err != nil {
				log.Printf("closing clickhouse connection: %s", err.Error())
			}

			log.Fatalf("executing migration: %s", err.Error())
			return
		}

		err = clickhouseConnection.Exec(
			setupCtx,
			`CREATE TABLE IF NOT EXISTS cakupi.coverage_reports_summary ON CLUSTER `+
				strconv.Quote(clickhouseCluster)+
				`(
							repository_name String NOT NULL,
							commit_hash String NOT NULL,
							commit_branch String NOT NULL,
							file_name String NOT NULL,
							file_path String NOT NULL,
							line_number UInt64 NOT NULL,
							hits_count UInt64 NOT NULL,
							fully_covered Bool DEFAULT TRUE,
							timestamp DateTime64(3, 'UTC')
						)
						ENGINE = Distributed(`+strconv.Quote(clickhouseCluster)+`, "cakupi", "coverage_detail")`,
		)
		if err != nil {
			setupCancel()
			if err := clickhouseConnection.Close(); err != nil {
				log.Printf("closing clickhouse connection: %s", err.Error())
			}

			log.Fatalf("executing migration: %s", err.Error())
			return
		}
	} else {
		err := clickhouseConnection.Exec(setupCtx, `CREATE DATABASE IF NOT EXISTS "cakupi"`)
		if err != nil {
			setupCancel()
			if err := clickhouseConnection.Close(); err != nil {
				log.Printf("closing clickhouse connection: %s", err.Error())
			}

			log.Fatalf("executing migration: %s", err.Error())
			return
		}

		err = clickhouseConnection.Exec(
			setupCtx,
			`CREATE TABLE IF NOT EXISTS cakupi.coverage_reports_summary
				(
					repository_name String NOT NULL,
					commit_hash String NOT NULL,
					commit_branch String NOT NULL,
					total_statements UInt64 NOT NULL,
					total_covered_statements UInt64 NOT NULL,
					number_of_files UInt64 NOT NULL,
					timestamp DateTime64(3, 'UTC')
				)
				ENGINE = MergeTree()
				PRIMARY KEY (repository_name, commit_hash)`,
		)
		if err != nil {
			setupCancel()
			if err := clickhouseConnection.Close(); err != nil {
				log.Printf("closing clickhouse connection: %s", err.Error())
			}

			log.Fatalf("executing migration: %s", err.Error())
			return
		}

		err = clickhouseConnection.Exec(
			setupCtx,
			`CREATE TABLE IF NOT EXISTS cakupi.coverage_reports_summary 
						(
							repository_name String NOT NULL,
							commit_hash String NOT NULL,
							commit_branch String NOT NULL,
							file_name String NOT NULL,
							file_path String NOT NULL,
							line_number UInt64 NOT NULL,
							hits_count UInt64 NOT NULL,
							fully_covered Bool DEFAULT TRUE,
							timestamp DateTime64(3, 'UTC')
						)
						ENGINE = MergeTree()
						PRIMARY KEY (repository_name, commit_hash)`,
		)
		if err != nil {
			setupCancel()
			if err := clickhouseConnection.Close(); err != nil {
				log.Printf("closing clickhouse connection: %s", err.Error())
			}

			log.Fatalf("executing migration: %s", err.Error())
			return
		}
	}

	clickhouseClient = clickhouseConnection

	exitCode := m.Run()

	err = clickhouseConnection.Close()
	if err != nil {
		log.Printf("Closing clickhouse connection: %v", err)
	}

	os.Exit(exitCode)
}

func TestNewReportsRepository(t *testing.T) {
	t.Run("Empty clickhouse client", func(t *testing.T) {
		_, err := reports.NewReportsRepository(nil, "")
		if err.Error() != "clickhouse client is nil" {
			t.Errorf("expecting an error of 'clickhouse client is nil', instead got %s", err.Error())
		}
	})
}

func TestRepository_Insert(t *testing.T) {
	reportsRepository, err := reports.NewReportsRepository(clickhouseClient, clickhouseCluster)
	if err != nil {
		t.Fatalf("creating reports repository: %s", err.Error())
	}

	err = reportsRepository.Insert(
		context.Background(),
		"cakupi/reporter",
		"b5f2ace8121ffd1482218e7ff72e9117b49cb829",
		"master",
		provider.Result{
			Files: []provider.File{
				{
					Name: "main.go",
					Path: "./main.go",
					Lines: []provider.Line{
						{
							Number:       1,
							Count:        1,
							Type:         provider.LineTypeStatement,
							FullyCovered: true,
						},
						{
							Number:       2,
							Count:        2,
							Type:         provider.LineTypeStatement,
							FullyCovered: true,
						},
						{
							Number:       3,
							Count:        0,
							Type:         provider.LineTypeStatement,
							FullyCovered: false,
						},
					},
				},
				{
					Name: "provider.go",
					Path: "provider/provider.go",
					Lines: []provider.Line{
						{
							Number:       1,
							Count:        1,
							Type:         provider.LineTypeStatement,
							FullyCovered: true,
						},
						{
							Number:       2,
							Count:        2,
							Type:         provider.LineTypeStatement,
							FullyCovered: true,
						},
						{
							Number:       3,
							Count:        0,
							Type:         provider.LineTypeStatement,
							FullyCovered: false,
						},
					},
				},
				{
					Name: "reports.go",
					Path: "reports/reports.go",
					Lines: []provider.Line{
						{
							Number:       1,
							Count:        1,
							Type:         provider.LineTypeStatement,
							FullyCovered: true,
						},
						{
							Number:       2,
							Count:        2,
							Type:         provider.LineTypeStatement,
							FullyCovered: true,
						},
						{
							Number:       3,
							Count:        0,
							Type:         provider.LineTypeStatement,
							FullyCovered: false,
						},
					},
				},
			},
			TotalStatements:        100,
			TotalCoveredStatements: 89,
			NumberOfFiles:          3,
		},
	)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
}
