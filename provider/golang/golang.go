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

package golang

import (
	"bufio"
	"io"
	"path"
	"strconv"
	"strings"

	"coverage-worker/provider"
)

// Some reference: https://stackoverflow.com/questions/31413281/golang-coverprofile-output-format

// Parser implements provider.Parser
type Parser struct{}

func (Parser) Parse(input io.Reader) (provider.Result, error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var totalStatements uint64
	var totalCoveredStatements uint64

	var files []provider.File
	var filesMapper = make(map[string]int)

	var currentIndex int = -1
	for scanner.Scan() {
		currentIndex += 1
		currentText := scanner.Text()
		// First line is "mode: foo", where foo is "set", "count", or "atomic".
		// Rest of file is in the format
		//	encoding/base64/base64.go:34.44,37.40 3 1
		// where the fields are: name.go:line.column,line.column numberOfStatements count
		if currentIndex == 0 {
			// Check if the first index really contains `mode: foo`.
			if !strings.HasPrefix(currentText, "mode: ") {
				return provider.Result{}, provider.ErrInvalidFormat
			}
			continue
		}

		// Create a tokenizer for the current text
		var accumulatedToken string
		var err error = nil
		// Create variables
		var filePath string
		var firstLineNumber uint64
		var secondLineNumber uint64
		var secondColumnNumberFilled bool
		var numberOfStatements uint64
		var count uint64

		// Iterate through tokens
		for _, token := range currentText {
			currentToken := string(token)

			if currentToken == ":" && filePath == "" {
				filePath = accumulatedToken
				accumulatedToken = ""
				continue
			}

			if currentToken == "." {
				// Make sure it's not a file extension
				if filePath == "" {
					accumulatedToken += currentToken
					continue
				}

				if firstLineNumber == 0 {
					firstLineNumber, err = strconv.ParseUint(accumulatedToken, 10, 64)
					if err != nil {
						return provider.Result{}, provider.ErrInvalidFormat
					}
					accumulatedToken = ""
					continue
				}

				if secondLineNumber == 0 {
					secondLineNumber, err = strconv.ParseUint(accumulatedToken, 10, 64)
					if err != nil {
						return provider.Result{}, provider.ErrInvalidFormat
					}
					accumulatedToken = ""
					continue
				}
			}

			if currentToken == "," {
				accumulatedToken = ""
				continue
			}

			if currentToken == " " {
				if !secondColumnNumberFilled {
					secondColumnNumberFilled = true
					accumulatedToken = ""
					continue
				}

				if numberOfStatements == 0 {
					numberOfStatements, err = strconv.ParseUint(accumulatedToken, 10, 64)
					if err != nil {
						return provider.Result{}, provider.ErrInvalidFormat
					}
					accumulatedToken = ""
					continue
				}
			}

			accumulatedToken += currentToken
		}

		count, err = strconv.ParseUint(accumulatedToken, 10, 64)
		if err != nil {
			return provider.Result{}, provider.ErrInvalidFormat
		}

		var linesReport []provider.Line
		for i := firstLineNumber; i <= secondLineNumber; i++ {
			// Global metrics
			totalStatements += 1
			if count > 0 {
				totalCoveredStatements += 1
			}

			linesReport = append(linesReport, provider.Line{
				Number:       i,
				Count:        count,
				Type:         provider.LineTypeStatement,
				FullyCovered: count > 0,
			})
		}

		// Check if filePath exists on the filesMapper
		if fileIndex, ok := filesMapper[filePath]; ok {
			// We'll just append
			var lines []provider.Line
			// Check if the lines on files[fileIndex] already has the same line number
			for lineIndex, line := range files[fileIndex].Lines {
				// The duplicate will only be on the first and second occurrence.
				// There will not be a duplicate in between.
				if line.Number == firstLineNumber && count == 0 {
					totalCoveredStatements -= 1
					files[fileIndex].Lines[lineIndex].FullyCovered = false
					lines = linesReport[1:]
					break
				}

				if line.Number == secondLineNumber && count == 0 {
					totalCoveredStatements -= 1
					files[fileIndex].Lines[lineIndex].FullyCovered = false
					lines = linesReport[:len(linesReport)-1]
					break
				}
			}

			files[fileIndex].Lines = append(files[fileIndex].Lines, lines...)
			continue
		}

		filesMapper[filePath] = len(files)

		// We'll create a new entry on files
		files = append(files, provider.File{
			Name:  path.Base(filePath),
			Path:  filePath,
			Lines: linesReport,
		})
	}

	return provider.Result{
		Files:                  files,
		TotalStatements:        totalStatements,
		TotalCoveredStatements: totalCoveredStatements,
		NumberOfFiles:          uint64(len(files)),
	}, nil
}
