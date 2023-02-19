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

package golang_test

import (
	"errors"
	"strings"
	"testing"

	"coverage-worker/provider"
	"coverage-worker/provider/golang"
)

func TestParse_Count(t *testing.T) {
	var content = strings.NewReader(`mode: count
github.com/aldy505/cheapcash/append.go:18.56,20.16 2 2
github.com/aldy505/cheapcash/append.go:20.16,22.3 1 0
github.com/aldy505/cheapcash/append.go:24.2,24.12 1 2
github.com/aldy505/cheapcash/append.go:24.12,26.3 1 1
github.com/aldy505/cheapcash/append.go:28.2,29.16 2 1
github.com/aldy505/cheapcash/append.go:29.16,31.3 1 0
github.com/aldy505/cheapcash/append.go:32.2,35.16 3 1
github.com/aldy505/cheapcash/append.go:35.16,37.3 1 0
github.com/aldy505/cheapcash/append.go:39.2,39.12 1 1
github.com/aldy505/cheapcash/cheapcash.go:32.23,36.2 1 17
github.com/aldy505/cheapcash/cheapcash.go:49.30,50.16 1 2
github.com/aldy505/cheapcash/cheapcash.go:50.16,51.24 1 1
github.com/aldy505/cheapcash/cheapcash.go:54.2,54.35 1 1
github.com/aldy505/cheapcash/cheapcash.go:54.35,56.3 1 1
github.com/aldy505/cheapcash/cheapcash.go:58.2,60.3 1 1
github.com/aldy505/cheapcash/check.go:20.50,22.16 2 63
github.com/aldy505/cheapcash/check.go:22.16,23.37 1 35
github.com/aldy505/cheapcash/check.go:23.37,25.4 1 35
github.com/aldy505/cheapcash/check.go:27.3,27.20 1 0
github.com/aldy505/cheapcash/check.go:29.2,31.18 2 28
github.com/aldy505/cheapcash/check.go:41.34,46.16 3 26
github.com/aldy505/cheapcash/check.go:46.16,47.18 1 25
github.com/aldy505/cheapcash/check.go:47.18,49.4 1 25
github.com/aldy505/cheapcash/check.go:53.2,55.38 2 1
github.com/aldy505/cheapcash/check.go:55.38,56.25 1 3
github.com/aldy505/cheapcash/check.go:56.25,58.12 2 1
github.com/aldy505/cheapcash/check.go:59.9,61.4 1 2
github.com/aldy505/cheapcash/check.go:63.3,64.17 2 2
github.com/aldy505/cheapcash/check.go:64.17,65.35 1 1
github.com/aldy505/cheapcash/check.go:65.35,66.13 1 1
github.com/aldy505/cheapcash/check.go:68.4,68.14 1 0
github.com/aldy505/cheapcash/check.go:72.2,72.12 1 1
github.com/aldy505/cheapcash/delete.go:15.42,17.16 2 10
github.com/aldy505/cheapcash/delete.go:17.16,19.3 1 0
github.com/aldy505/cheapcash/delete.go:21.2,21.12 1 10
github.com/aldy505/cheapcash/delete.go:21.12,23.3 1 4
github.com/aldy505/cheapcash/delete.go:25.2,26.16 2 6
github.com/aldy505/cheapcash/delete.go:26.16,28.3 1 0
github.com/aldy505/cheapcash/delete.go:30.2,30.12 1 6
github.com/aldy505/cheapcash/reader.go:19.50,21.16 2 10
github.com/aldy505/cheapcash/reader.go:21.16,23.3 1 0
github.com/aldy505/cheapcash/reader.go:25.2,25.12 1 10
github.com/aldy505/cheapcash/reader.go:25.12,27.3 1 2
github.com/aldy505/cheapcash/reader.go:29.2,30.16 2 8
github.com/aldy505/cheapcash/reader.go:30.16,32.3 1 0
github.com/aldy505/cheapcash/reader.go:33.2,36.16 3 8
github.com/aldy505/cheapcash/reader.go:36.16,38.3 1 0
github.com/aldy505/cheapcash/reader.go:40.2,40.18 1 8
github.com/aldy505/cheapcash/rename.go:16.47,18.16 2 8
github.com/aldy505/cheapcash/rename.go:18.16,20.3 1 0
github.com/aldy505/cheapcash/rename.go:22.2,23.16 2 8
github.com/aldy505/cheapcash/rename.go:23.16,25.3 1 0
github.com/aldy505/cheapcash/rename.go:27.2,28.16 2 8
github.com/aldy505/cheapcash/rename.go:28.16,30.3 1 0
github.com/aldy505/cheapcash/rename.go:32.2,32.15 1 8
github.com/aldy505/cheapcash/rename.go:32.15,34.3 1 5
github.com/aldy505/cheapcash/rename.go:36.2,36.14 1 3
github.com/aldy505/cheapcash/rename.go:36.14,38.3 1 1
github.com/aldy505/cheapcash/rename.go:40.2,41.16 2 2
github.com/aldy505/cheapcash/rename.go:41.16,43.3 1 0
github.com/aldy505/cheapcash/rename.go:45.2,45.12 1 2
github.com/aldy505/cheapcash/sanitize.go:13.39,25.31 2 122
github.com/aldy505/cheapcash/sanitize.go:25.31,27.3 1 976
github.com/aldy505/cheapcash/sanitize.go:28.2,28.13 1 122
github.com/aldy505/cheapcash/writer.go:16.55,18.16 2 18
github.com/aldy505/cheapcash/writer.go:18.16,20.3 1 0
github.com/aldy505/cheapcash/writer.go:22.2,23.16 2 18
github.com/aldy505/cheapcash/writer.go:23.16,25.3 1 0
github.com/aldy505/cheapcash/writer.go:27.2,27.11 1 18
github.com/aldy505/cheapcash/writer.go:27.11,29.17 2 4
github.com/aldy505/cheapcash/writer.go:29.17,31.4 1 0
github.com/aldy505/cheapcash/writer.go:34.2,35.16 2 18
github.com/aldy505/cheapcash/writer.go:35.16,37.3 1 0
github.com/aldy505/cheapcash/writer.go:38.2,41.16 3 18
github.com/aldy505/cheapcash/writer.go:41.16,43.3 1 0
github.com/aldy505/cheapcash/writer.go:45.2,45.12 1 18`)

	parsedCoverage, err := (golang.Parser{}).Parse(content)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if parsedCoverage.TotalStatements != 198 {
		t.Errorf("expecting TotalStatements to be 198, instead got %d", parsedCoverage.TotalStatements)
	}

	if parsedCoverage.NumberOfFiles != 8 {
		t.Errorf("expecting NumberOfFiles to be 8, instead got %d", parsedCoverage.NumberOfFiles)
	}

	if parsedCoverage.TotalCoveredStatements != 140 {
		t.Errorf("expecting TotalCoveredStatements to be 140, instead got %d", parsedCoverage.TotalCoveredStatements)
	}
}

func TestParse_Atomic(t *testing.T) {
	var content = strings.NewReader(`mode: atomic
github.com/aldy505/cheapcash/append.go:18.56,20.16 2 2
github.com/aldy505/cheapcash/append.go:20.16,22.3 1 0
github.com/aldy505/cheapcash/append.go:24.2,24.12 1 2
github.com/aldy505/cheapcash/append.go:24.12,26.3 1 1
github.com/aldy505/cheapcash/append.go:28.2,29.16 2 1
github.com/aldy505/cheapcash/append.go:29.16,31.3 1 0
github.com/aldy505/cheapcash/append.go:32.2,35.16 3 1
github.com/aldy505/cheapcash/append.go:35.16,37.3 1 0
github.com/aldy505/cheapcash/append.go:39.2,39.12 1 1
github.com/aldy505/cheapcash/cheapcash.go:32.23,36.2 1 17
github.com/aldy505/cheapcash/cheapcash.go:49.30,50.16 1 2
github.com/aldy505/cheapcash/cheapcash.go:50.16,51.24 1 1
github.com/aldy505/cheapcash/cheapcash.go:54.2,54.35 1 1
github.com/aldy505/cheapcash/cheapcash.go:54.35,56.3 1 1
github.com/aldy505/cheapcash/cheapcash.go:58.2,60.3 1 1
github.com/aldy505/cheapcash/check.go:20.50,22.16 2 63
github.com/aldy505/cheapcash/check.go:22.16,23.37 1 35
github.com/aldy505/cheapcash/check.go:23.37,25.4 1 35
github.com/aldy505/cheapcash/check.go:27.3,27.20 1 0
github.com/aldy505/cheapcash/check.go:29.2,31.18 2 28
github.com/aldy505/cheapcash/check.go:41.34,46.16 3 26
github.com/aldy505/cheapcash/check.go:46.16,47.18 1 25
github.com/aldy505/cheapcash/check.go:47.18,49.4 1 25
github.com/aldy505/cheapcash/check.go:53.2,55.38 2 1
github.com/aldy505/cheapcash/check.go:55.38,56.25 1 3
github.com/aldy505/cheapcash/check.go:56.25,58.12 2 1
github.com/aldy505/cheapcash/check.go:59.9,61.4 1 2
github.com/aldy505/cheapcash/check.go:63.3,64.17 2 2
github.com/aldy505/cheapcash/check.go:64.17,65.35 1 1
github.com/aldy505/cheapcash/check.go:65.35,66.13 1 1
github.com/aldy505/cheapcash/check.go:68.4,68.14 1 0
github.com/aldy505/cheapcash/check.go:72.2,72.12 1 1
github.com/aldy505/cheapcash/delete.go:15.42,17.16 2 10
github.com/aldy505/cheapcash/delete.go:17.16,19.3 1 0
github.com/aldy505/cheapcash/delete.go:21.2,21.12 1 10
github.com/aldy505/cheapcash/delete.go:21.12,23.3 1 4
github.com/aldy505/cheapcash/delete.go:25.2,26.16 2 6
github.com/aldy505/cheapcash/delete.go:26.16,28.3 1 0
github.com/aldy505/cheapcash/delete.go:30.2,30.12 1 6
github.com/aldy505/cheapcash/reader.go:19.50,21.16 2 10
github.com/aldy505/cheapcash/reader.go:21.16,23.3 1 0
github.com/aldy505/cheapcash/reader.go:25.2,25.12 1 10
github.com/aldy505/cheapcash/reader.go:25.12,27.3 1 2
github.com/aldy505/cheapcash/reader.go:29.2,30.16 2 8
github.com/aldy505/cheapcash/reader.go:30.16,32.3 1 0
github.com/aldy505/cheapcash/reader.go:33.2,36.16 3 8
github.com/aldy505/cheapcash/reader.go:36.16,38.3 1 0
github.com/aldy505/cheapcash/reader.go:40.2,40.18 1 8
github.com/aldy505/cheapcash/rename.go:16.47,18.16 2 8
github.com/aldy505/cheapcash/rename.go:18.16,20.3 1 0
github.com/aldy505/cheapcash/rename.go:22.2,23.16 2 8
github.com/aldy505/cheapcash/rename.go:23.16,25.3 1 0
github.com/aldy505/cheapcash/rename.go:27.2,28.16 2 8
github.com/aldy505/cheapcash/rename.go:28.16,30.3 1 0
github.com/aldy505/cheapcash/rename.go:32.2,32.15 1 8
github.com/aldy505/cheapcash/rename.go:32.15,34.3 1 5
github.com/aldy505/cheapcash/rename.go:36.2,36.14 1 3
github.com/aldy505/cheapcash/rename.go:36.14,38.3 1 1
github.com/aldy505/cheapcash/rename.go:40.2,41.16 2 2
github.com/aldy505/cheapcash/rename.go:41.16,43.3 1 0
github.com/aldy505/cheapcash/rename.go:45.2,45.12 1 2
github.com/aldy505/cheapcash/sanitize.go:13.39,25.31 2 122
github.com/aldy505/cheapcash/sanitize.go:25.31,27.3 1 976
github.com/aldy505/cheapcash/sanitize.go:28.2,28.13 1 122
github.com/aldy505/cheapcash/writer.go:16.55,18.16 2 18
github.com/aldy505/cheapcash/writer.go:18.16,20.3 1 0
github.com/aldy505/cheapcash/writer.go:22.2,23.16 2 18
github.com/aldy505/cheapcash/writer.go:23.16,25.3 1 0
github.com/aldy505/cheapcash/writer.go:27.2,27.11 1 18
github.com/aldy505/cheapcash/writer.go:27.11,29.17 2 4
github.com/aldy505/cheapcash/writer.go:29.17,31.4 1 0
github.com/aldy505/cheapcash/writer.go:34.2,35.16 2 18
github.com/aldy505/cheapcash/writer.go:35.16,37.3 1 0
github.com/aldy505/cheapcash/writer.go:38.2,41.16 3 18
github.com/aldy505/cheapcash/writer.go:41.16,43.3 1 0
github.com/aldy505/cheapcash/writer.go:45.2,45.12 1 18`)

	parsedCoverage, err := (golang.Parser{}).Parse(content)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if parsedCoverage.TotalStatements != 198 {
		t.Errorf("expecting TotalStatements to be 198, instead got %d", parsedCoverage.TotalStatements)
	}

	if parsedCoverage.NumberOfFiles != 8 {
		t.Errorf("expecting NumberOfFiles to be 8, instead got %d", parsedCoverage.NumberOfFiles)
	}

	if parsedCoverage.TotalCoveredStatements != 140 {
		t.Errorf("expecting TotalCoveredStatements to be 140, instead got %d", parsedCoverage.TotalCoveredStatements)
	}
}

func TestParse_Set(t *testing.T) {
	var content = strings.NewReader(`mode: set
github.com/aldy505/cheapcash/append.go:18.56,20.16 2 1
github.com/aldy505/cheapcash/append.go:20.16,22.3 1 0
github.com/aldy505/cheapcash/append.go:24.2,24.12 1 1
github.com/aldy505/cheapcash/append.go:24.12,26.3 1 1
github.com/aldy505/cheapcash/append.go:28.2,29.16 2 1
github.com/aldy505/cheapcash/append.go:29.16,31.3 1 0
github.com/aldy505/cheapcash/append.go:32.2,35.16 3 1
github.com/aldy505/cheapcash/append.go:35.16,37.3 1 0
github.com/aldy505/cheapcash/append.go:39.2,39.12 1 1
github.com/aldy505/cheapcash/cheapcash.go:32.23,36.2 1 1
github.com/aldy505/cheapcash/cheapcash.go:49.30,50.16 1 1
github.com/aldy505/cheapcash/cheapcash.go:50.16,51.24 1 1
github.com/aldy505/cheapcash/cheapcash.go:54.2,54.35 1 1
github.com/aldy505/cheapcash/cheapcash.go:54.35,56.3 1 1
github.com/aldy505/cheapcash/cheapcash.go:58.2,60.3 1 1
github.com/aldy505/cheapcash/check.go:20.50,22.16 2 1
github.com/aldy505/cheapcash/check.go:22.16,23.37 1 1
github.com/aldy505/cheapcash/check.go:23.37,25.4 1 1
github.com/aldy505/cheapcash/check.go:27.3,27.20 1 0
github.com/aldy505/cheapcash/check.go:29.2,31.18 2 1
github.com/aldy505/cheapcash/check.go:41.34,46.16 3 1
github.com/aldy505/cheapcash/check.go:46.16,47.18 1 1
github.com/aldy505/cheapcash/check.go:47.18,49.4 1 1
github.com/aldy505/cheapcash/check.go:53.2,55.38 2 1
github.com/aldy505/cheapcash/check.go:55.38,56.25 1 1
github.com/aldy505/cheapcash/check.go:56.25,58.12 2 1
github.com/aldy505/cheapcash/check.go:59.9,61.4 1 1
github.com/aldy505/cheapcash/check.go:63.3,64.17 2 1
github.com/aldy505/cheapcash/check.go:64.17,65.35 1 1
github.com/aldy505/cheapcash/check.go:65.35,66.13 1 1
github.com/aldy505/cheapcash/check.go:68.4,68.14 1 0
github.com/aldy505/cheapcash/check.go:72.2,72.12 1 1
github.com/aldy505/cheapcash/delete.go:15.42,17.16 2 1
github.com/aldy505/cheapcash/delete.go:17.16,19.3 1 0
github.com/aldy505/cheapcash/delete.go:21.2,21.12 1 1
github.com/aldy505/cheapcash/delete.go:21.12,23.3 1 1
github.com/aldy505/cheapcash/delete.go:25.2,26.16 2 1
github.com/aldy505/cheapcash/delete.go:26.16,28.3 1 0
github.com/aldy505/cheapcash/delete.go:30.2,30.12 1 1
github.com/aldy505/cheapcash/reader.go:19.50,21.16 2 1
github.com/aldy505/cheapcash/reader.go:21.16,23.3 1 0
github.com/aldy505/cheapcash/reader.go:25.2,25.12 1 1
github.com/aldy505/cheapcash/reader.go:25.12,27.3 1 1
github.com/aldy505/cheapcash/reader.go:29.2,30.16 2 1
github.com/aldy505/cheapcash/reader.go:30.16,32.3 1 0
github.com/aldy505/cheapcash/reader.go:33.2,36.16 3 1
github.com/aldy505/cheapcash/reader.go:36.16,38.3 1 0
github.com/aldy505/cheapcash/reader.go:40.2,40.18 1 1
github.com/aldy505/cheapcash/rename.go:16.47,18.16 2 1
github.com/aldy505/cheapcash/rename.go:18.16,20.3 1 0
github.com/aldy505/cheapcash/rename.go:22.2,23.16 2 1
github.com/aldy505/cheapcash/rename.go:23.16,25.3 1 0
github.com/aldy505/cheapcash/rename.go:27.2,28.16 2 1
github.com/aldy505/cheapcash/rename.go:28.16,30.3 1 0
github.com/aldy505/cheapcash/rename.go:32.2,32.15 1 1
github.com/aldy505/cheapcash/rename.go:32.15,34.3 1 1
github.com/aldy505/cheapcash/rename.go:36.2,36.14 1 1
github.com/aldy505/cheapcash/rename.go:36.14,38.3 1 1
github.com/aldy505/cheapcash/rename.go:40.2,41.16 2 1
github.com/aldy505/cheapcash/rename.go:41.16,43.3 1 0
github.com/aldy505/cheapcash/rename.go:45.2,45.12 1 1
github.com/aldy505/cheapcash/sanitize.go:13.39,25.31 2 1
github.com/aldy505/cheapcash/sanitize.go:25.31,27.3 1 1
github.com/aldy505/cheapcash/sanitize.go:28.2,28.13 1 1
github.com/aldy505/cheapcash/writer.go:16.55,18.16 2 1
github.com/aldy505/cheapcash/writer.go:18.16,20.3 1 0
github.com/aldy505/cheapcash/writer.go:22.2,23.16 2 1
github.com/aldy505/cheapcash/writer.go:23.16,25.3 1 0
github.com/aldy505/cheapcash/writer.go:27.2,27.11 1 1
github.com/aldy505/cheapcash/writer.go:27.11,29.17 2 1
github.com/aldy505/cheapcash/writer.go:29.17,31.4 1 0
github.com/aldy505/cheapcash/writer.go:34.2,35.16 2 1
github.com/aldy505/cheapcash/writer.go:35.16,37.3 1 0
github.com/aldy505/cheapcash/writer.go:38.2,41.16 3 1
github.com/aldy505/cheapcash/writer.go:41.16,43.3 1 0
github.com/aldy505/cheapcash/writer.go:45.2,45.12 1 1`)

	parsedCoverage, err := (golang.Parser{}).Parse(content)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if parsedCoverage.TotalStatements != 198 {
		t.Errorf("expecting TotalStatements to be 198, instead got %d", parsedCoverage.TotalStatements)
	}

	if parsedCoverage.NumberOfFiles != 8 {
		t.Errorf("expecting NumberOfFiles to be 8, instead got %d", parsedCoverage.NumberOfFiles)
	}

	if parsedCoverage.TotalCoveredStatements != 140 {
		t.Errorf("expecting TotalCoveredStatements to be 140, instead got %d", parsedCoverage.TotalCoveredStatements)
	}
}

func TestParse_InvalidFormat(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		_, err := (golang.Parser{}).Parse(strings.NewReader("\n\n"))
		if !errors.Is(err, provider.ErrInvalidFormat) {
			t.Errorf("exepcting an error of reporter.ErrInvalidFormat, instead got %s", err.Error())
		}
	})

	t.Run("Invalid firstLineNumber", func(t *testing.T) {
		var content = strings.NewReader(`mode: atomic
foobar:lorem.ipsum,lorem.ipsum lorem ipsum`)

		_, err := (golang.Parser{}).Parse(content)
		if !errors.Is(err, provider.ErrInvalidFormat) {
			t.Errorf("exepcting an error of reporter.ErrInvalidFormat, instead got %s", err.Error())
		}
	})

	t.Run("Invalid secondLineNumber", func(t *testing.T) {
		var content = strings.NewReader(`mode: atomic
foobar:10.10,lorem.ipsum lorem ipsum`)

		_, err := (golang.Parser{}).Parse(content)
		if !errors.Is(err, provider.ErrInvalidFormat) {
			t.Errorf("exepcting an error of reporter.ErrInvalidFormat, instead got %s", err.Error())
		}
	})

	t.Run("Invalid numberOfStatements", func(t *testing.T) {
		var content = strings.NewReader(`mode: atomic
foobar:10.10,12.15 lorem ipsum`)

		_, err := (golang.Parser{}).Parse(content)
		if !errors.Is(err, provider.ErrInvalidFormat) {
			t.Errorf("exepcting an error of reporter.ErrInvalidFormat, instead got %s", err.Error())
		}
	})

	t.Run("Invalid count", func(t *testing.T) {
		var content = strings.NewReader(`mode: atomic
foobar:10.10,12.15 10 ipsum`)

		_, err := (golang.Parser{}).Parse(content)
		if !errors.Is(err, provider.ErrInvalidFormat) {
			t.Errorf("exepcting an error of reporter.ErrInvalidFormat, instead got %s", err.Error())
		}
	})
}
