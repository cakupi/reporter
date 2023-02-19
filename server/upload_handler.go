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

package server

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"

	"coverage-worker/orchestrator"
	"coverage-worker/primitive"
)

type UploadResponse struct {
	Message string `json:"message"`
}

func (d *Dependency) UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, fmt.Sprintf("Parsing multipart/form-data body: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	repositoryName := r.FormValue("repository_name")
	commitHash := r.FormValue("commit_hash")
	commitBranch := r.FormValue("commit_branch")
	baseGroup := r.FormValue("base_group")

	coverageFile, _, err := r.FormFile("coverage")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed parsing coverage file: %s", err.Error()), http.StatusBadRequest)
		return
	}

	coverageFileContents, err := io.ReadAll(coverageFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed reading coverage file: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	var coverageFiles []primitive.File

	uncompressedCoverageContents, err := gzip.NewReader(bytes.NewBuffer(coverageFileContents))
	if err != nil {
		http.Error(w, "Decompressing gzip file", http.StatusInternalServerError)
		return
	}

	coverageTarReader := tar.NewReader(uncompressedCoverageContents)
	for {
		fileHeader, err := coverageTarReader.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			http.Error(w, "Reading next entry of the tar archive", http.StatusInternalServerError)
			return
		}

		fileContent, err := io.ReadAll(coverageTarReader)
		if err != nil {
			http.Error(w, "Reading content of the tar archive", http.StatusInternalServerError)
			return
		}

		coverageFiles = append(coverageFiles, primitive.File{
			Path:        fileHeader.Name,
			Size:        fileHeader.Size,
			Extension:   path.Ext(fileHeader.Name),
			ContentType: "text/plain",
			Content:     fileContent,
		})
	}

	err = d.orchestratorService.ProcessReport(
		r.Context(),
		orchestrator.UploadReportRequest{
			RepositoryName: repositoryName,
			CommitHash:     commitHash,
			CommitBranch:   commitBranch,
			BaseGroup:      baseGroup,
			CoverageFiles:  coverageFiles,
		},
	)
	if err != nil {
		if errors.Is(err, orchestrator.ErrUnableToParse) {
			http.Error(w, "Our parsers cannot parse your coverage input", http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, "Uploading report to upstream", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
