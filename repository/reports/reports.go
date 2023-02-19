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

package reports

import (
	"context"
	"fmt"
	"strings"
	"time"

	"coverage-worker/provider"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Repository struct {
	clickhouseClient  driver.Conn
	clickhouseCluster string
}

func (r Repository) Insert(ctx context.Context, repository string, commitSha string, commitBranch string, results provider.Result) error {
	err := r.clickhouseClient.AsyncInsert(
		context.Background(),
		fmt.Sprintf(`INSERT INTO 
			cakupi.coverage_reports_summary
				(
					repository_name,
					commit_hash,
					commit_branch,
					total_statements,
					total_covered_statements,
					number_of_files,
					timestamp
				)
			VALUES 
			    (
			     	'%s',
			     	'%s',
			     	'%s',
			     	%d,
			     	%d,
			     	%d,
			     	'%s'
			    )`,
			strings.ReplaceAll(repository, "'", "\\'"),
			strings.ReplaceAll(commitSha, "'", "\\'"),
			strings.ReplaceAll(commitBranch, "'", "\\'"),
			results.TotalStatements,
			results.TotalCoveredStatements,
			results.NumberOfFiles,
			time.Now().Format("2006-01-02 15:04:05.99999999"),
		),
		false,
	)
	if err != nil {
		return fmt.Errorf("executing async insert for coverage report summary: %w", err)
	}

	for _, file := range results.Files {
		for _, line := range file.Lines {
			err := r.clickhouseClient.AsyncInsert(
				context.Background(),
				fmt.Sprintf(`INSERT INTO 
			cakupi.coverage_reports_detail
				(
					repository_name,
					commit_hash,
					commit_branch,
				 	file_name,
				 	file_path,
				 	line_number,
				 	hits_count,
				 	fully_covered,
					timestamp
				)
			VALUES 
			    (
			     	'%s',
			     	'%s',
			     	'%s',
			     	'%s',
			     	'%s',
			     	%d,
			     	%d,
			     	%t,
			     	'%s'
			    )`,
					strings.ReplaceAll(repository, "'", "\\'"),
					strings.ReplaceAll(commitSha, "'", "\\'"),
					strings.ReplaceAll(commitBranch, "'", "\\'"),
					strings.ReplaceAll(file.Name, "'", "\\'"),
					strings.ReplaceAll(file.Path, "'", "\\'"),
					line.Number,
					line.Count,
					line.FullyCovered,
					time.Now().Format("2006-01-02 15:04:05.99999999"),
				),
				false,
			)
			if err != nil {
				return fmt.Errorf("executing async insert for coverage report summary: %w", err)
			}
		}
	}

	return nil
}

func NewReportsRepository(clickhouseClient driver.Conn, clickhouseCluster string) (*Repository, error) {
	if clickhouseClient == nil {
		return &Repository{}, fmt.Errorf("clickhouse client is nil")
	}

	return &Repository{
		clickhouseClient:  clickhouseClient,
		clickhouseCluster: clickhouseCluster,
	}, nil
}
