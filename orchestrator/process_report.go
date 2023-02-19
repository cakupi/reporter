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
	"bytes"
	"context"
	"fmt"
	"log"
	"path"

	"coverage-worker/provider"
)

func (o *orchestrator) ProcessReport(ctx context.Context, request UploadReportRequest) error {
	var coverageResults []provider.Result

	for _, parser := range o.reporterProviders {
		for _, coverageFile := range request.CoverageFiles {
			coverageResult, err := parser.Parse(bytes.NewReader(coverageFile.Content))
			if err != nil {
				log.Printf("An error occured during parsing the coverage: %s", err.Error())
				continue
			}

			coverageResults = append(coverageResults, coverageResult)
		}
	}

	if len(coverageResults) == 0 {
		return ErrUnableToParse
	}

	var finalCoverageResult provider.Result
	for _, result := range coverageResults {
		finalCoverageResult.TotalCoveredStatements += result.TotalCoveredStatements
		finalCoverageResult.NumberOfFiles += result.NumberOfFiles
		finalCoverageResult.TotalStatements += result.TotalStatements

		for _, file := range result.Files {
			finalCoverageResult.Files = append(finalCoverageResult.Files, provider.File{
				Name:  file.Name,
				Path:  path.Join(request.BaseGroup, file.Path),
				Lines: file.Lines,
			})
		}
	}

	err := o.reportsRepository.Insert(
		ctx,
		request.RepositoryName,
		request.CommitHash,
		request.CommitBranch,
		finalCoverageResult,
	)
	if err != nil {
		return fmt.Errorf("inserting record into reports repository: %s", err.Error())
	}

	return nil
}
