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

package orchestrator

import (
	"context"
	"errors"

	"coverage-worker/primitive"
	"coverage-worker/provider"
	"coverage-worker/repository/reports"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Config struct {
	ClickhouseClient  driver.Conn
	ClickhouseCluster string
	ReporterProviders []provider.Parser
}

type UploadReportRequest struct {
	RepositoryName string
	CommitHash     string
	CommitBranch   string
	BaseGroup      string
	CoverageFiles  []primitive.File
}

type Orchestrator interface {
	ProcessReport(ctx context.Context, request UploadReportRequest) error
}

type orchestrator struct {
	reportsRepository *reports.Repository
	reporterProviders []provider.Parser
}

func NewOrchestrator(config Config) (Orchestrator, error) {
	if config.ClickhouseClient == nil {
		return nil, errors.New("clickhouse client must not be nil")
	}

	reportsRepository, err := reports.NewReportsRepository(config.ClickhouseClient, config.ClickhouseCluster)
	if err != nil {
		return nil, err
	}

	return &orchestrator{
		reportsRepository: reportsRepository,
		reporterProviders: config.ReporterProviders,
	}, nil
}

var ErrUnableToParse = errors.New("unable to parse")
