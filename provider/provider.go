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

package provider

import (
	"errors"
	"io"
)

type LineType uint8

const (
	LineTypeUnspecified LineType = iota
	LineTypeStatement
	LineTypeConditional
)

type Line struct {
	Number       uint64
	Count        uint64
	Type         LineType
	FullyCovered bool
}

type File struct {
	Name  string
	Path  string
	Lines []Line
}

type Result struct {
	Files                  []File
	TotalStatements        uint64
	TotalCoveredStatements uint64
	NumberOfFiles          uint64
}

type Parser interface {
	Parse(input io.Reader) (Result, error)
}

var ErrInvalidFormat = errors.New("invalid format")
