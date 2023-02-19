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

package clover

import (
	"encoding/xml"
	"fmt"
	"io"

	"coverage-worker/provider"
)

type Metrics struct {
	Statements          int64 `xml:"statements,attr"`
	CoveredStatements   int64 `xml:"coveredstatements,attr"`
	Conditionals        int64 `xml:"conditionals,attr"`
	CoveredConditionals int64 `xml:"coveredconditionals,attr"`
	Methods             int64 `xml:"methods,attr"`
	CoveredMethods      int64 `xml:"coveredmethods,attr"`
	Files               int64 `xml:"files,attr"`
	Packages            int64 `xml:"packages,attr"`
}

type Line struct {
	Number     int64  `xml:"number,attr"`
	Count      int64  `xml:"count,attr"`
	Type       string `xml:"type,attr"`
	TrueCount  int64  `xml:"truecount,attr"`
	FalseCount int64  `xml:"falsecount,attr"`
}

type Coverage struct {
	Generated int64  `xml:"generated,attr"`
	Clover    string `xml:"clover,attr"`
	Project   []struct {
		Timestamp int64   `xml:"timestamp,attr"`
		Name      string  `xml:"name,attr"`
		Metrics   Metrics `xml:"metrics"`
		Package   []struct {
			Name    string  `xml:"name,attr"`
			Metrics Metrics `xml:"metrics"`
			File    []struct {
				Name    string  `xml:"name,attr"`
				Path    string  `xml:"path,attr"`
				Metrics Metrics `xml:"metrics"`
				Line    []Line  `xml:"line"`
			} `xml:"file"`
		} `xml:"package"`
	} `xml:"project"`
}

// Parser implements provider.Parser
type Parser struct{}

func (Parser) Parse(input io.Reader) (provider.Result, error) {
	var coverage Coverage
	err := xml.NewDecoder(input).Decode(&coverage)
	if err != nil {
		return provider.Result{}, fmt.Errorf("%w: parsing xml file: %s", provider.ErrInvalidFormat, err.Error())
	}

	var totalStatements uint64
	var totalCoveredStatements uint64
	var numberOfFiles uint64
	var files []provider.File
	for _, project := range coverage.Project {
		for _, pkg := range project.Package {
			for _, file := range pkg.File {
				totalStatements += uint64(file.Metrics.Statements)
				totalCoveredStatements += uint64(file.Metrics.CoveredStatements)
				numberOfFiles += 1

				var lines []provider.Line
				for _, line := range file.Line {
					var fullyCovered bool = true
					var lineType = provider.LineTypeUnspecified
					if line.Type == "cond" {
						lineType = provider.LineTypeConditional
						if line.TrueCount == 0 || line.FalseCount == 0 {
							fullyCovered = false
						}
					} else if line.Type == "stmt" {
						lineType = provider.LineTypeStatement
					}

					lines = append(lines, provider.Line{
						Number:       uint64(line.Number),
						Count:        uint64(line.Count),
						Type:         lineType,
						FullyCovered: fullyCovered,
					})
				}

				files = append(files, provider.File{
					Name:  file.Name,
					Path:  file.Path,
					Lines: lines,
				})
			}
		}
	}

	return provider.Result{
		Files:                  files,
		TotalStatements:        totalStatements,
		TotalCoveredStatements: totalCoveredStatements,
		NumberOfFiles:          numberOfFiles,
	}, nil
}
