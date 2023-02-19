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

package lcov

import (
	"bufio"
	"io"
	"path"
	"strconv"
	"strings"

	"coverage-worker/provider"
)

// Some reference: https://github.com/linux-test-project/lcov/issues/113#issuecomment-762335134
// TN: test name
// SF: source file path
// FN: line number,function name
// FNF:  number functions found
// FNH: number hit
// BRDA: branch data: line, block, (expressions,count)+
// BRF: branches found
// BRH: branches hit
// DA: line number, hit count
// LF: lines found
// LH:  lines hit.

// Another reference: https://ltp.sourceforge.net/coverage/lcov/geninfo.1.php
// Following is a quick description of the tracefile  format  as  used  by
// genhtml, geninfo and lcov.
//
// A tracefile is made up of several human-readable lines of text, divided
// into sections. If available, a tracefile begins with the testname which
// is stored in the following format:
//
//   TN:<test name>
//
// For  each  source  file  referenced in the .da file, there is a section
// containing filename and coverage data:
//
//   SF:<absolute path to the source file>
//
// Following is a list of line numbers for each function name found in the
// source file:
//
//   FN:<line number of function start>,<function name>
//
// Next,  there  is a list of execution counts for each instrumented func‐
// tion:
//
//   FNDA:<execution count>,<function name>
//
// This list is followed by two lines containing the number  of  functions
// found and hit:
//
//   FNF:<number of functions found>
//   FNH:<number of function hit>
//
// Branch coverage information is stored which one line per branch:
//
//   BRDA:<line number>,<block number>,<branch number>,<taken>
//
// Block  number  and  branch  number are gcc internal IDs for the branch.
// Taken is either '-' if the basic block containing the branch was  never
// executed or a number indicating how often that branch was taken.
//
// Branch coverage summaries are stored in two lines:
//
//   BRF:<number of branches found>
//   BRH:<number of branches hit>
//
// Then  there  is  a  list of execution counts for each instrumented line
// (i.e. a line which resulted in executable code):
//
//   DA:<line number>,<execution count>[,<checksum>]
//
// Note that there may be an optional checksum present  for  each  instru‐
// mented  line.  The  current  geninfo implementation uses an MD5 hash as
// checksumming algorithm.
//
// At the end of a section, there is a summary about how many  lines  were
// found and how many were actually instrumented:
//
//   LH:<number of lines with a non-zero execution count>
//   LF:<number of instrumented lines>
//
// Each sections ends with:
//
//   end_of_record

// Parser implements provider.Parser
type Parser struct{}

func (Parser) Parse(input io.Reader) (provider.Result, error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var totalStatements uint64
	var totalCoveredStatements uint64
	var numberOfFiles uint64

	var files []provider.File
	var currentFile provider.File
	for scanner.Scan() {
		currentText := scanner.Text()

		if currentText == "end_of_record" {
			// Process current file
			files = append(files, currentFile)
			currentFile = provider.File{
				Name:  "",
				Path:  "",
				Lines: nil,
			}
			continue
		}

		parts := strings.SplitN(currentText, ":", 2)
		switch parts[0] {
		case "SF":
			numberOfFiles += 1
			currentFile.Path = parts[1]
			currentFile.Name = path.Base(parts[1])
		case "DA":
			linePart := strings.SplitN(parts[1], ",", 2)
			lineNumber, err := strconv.ParseUint(linePart[0], 10, 64)
			if err != nil {
				return provider.Result{}, provider.ErrInvalidFormat
			}

			hitsCount, err := strconv.ParseUint(linePart[1], 10, 64)
			if err != nil {
				return provider.Result{}, provider.ErrInvalidFormat
			}

			currentFile.Lines = append(currentFile.Lines, provider.Line{
				Number:       lineNumber,
				Count:        hitsCount,
				Type:         provider.LineTypeStatement,
				FullyCovered: hitsCount > 0,
			})
		case "LF":
			linesFound, err := strconv.ParseUint(parts[1], 10, 64)
			if err != nil {
				return provider.Result{}, provider.ErrInvalidFormat
			}

			totalStatements += linesFound
		case "LH":
			linesHit, err := strconv.ParseUint(parts[1], 10, 64)
			if err != nil {
				return provider.Result{}, provider.ErrInvalidFormat
			}

			totalCoveredStatements += linesHit
		}

		continue
	}

	return provider.Result{
		Files:                  files,
		TotalStatements:        totalStatements,
		TotalCoveredStatements: totalCoveredStatements,
		NumberOfFiles:          numberOfFiles,
	}, nil
}
