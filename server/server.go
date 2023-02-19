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
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"coverage-worker/orchestrator"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Dependency struct {
	orchestratorService orchestrator.Orchestrator
}

type Config struct {
	Hostname  string
	Port      string
	TLSConfig *tls.Config

	OrchestratorService orchestrator.Orchestrator
}

func NewHttpServer(config Config) (*http.Server, error) {
	dependency := &Dependency{
		orchestratorService: config.OrchestratorService,
	}
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Minute * 5))
	router.Post("/api/v1/report", dependency.UploadHandler)

	return &http.Server{
		Addr:                         net.JoinHostPort(config.Hostname, config.Port),
		Handler:                      router,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    config.TLSConfig,
		ReadTimeout:                  time.Minute,
		ReadHeaderTimeout:            time.Minute,
	}, nil
}
