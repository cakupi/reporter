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

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"coverage-worker/orchestrator"
	"coverage-worker/provider"
	"coverage-worker/provider/clover"
	"coverage-worker/provider/golang"
	"coverage-worker/provider/lcov"
	"coverage-worker/server"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func main() {
	configuration, err := ParseConfig()
	if err != nil {
		log.Fatalf("Parsing configuration failed! %s", err.Error())
		return
	}

	clickhouseOptions, err := clickhouse.ParseDSN(configuration.ClickhouseDsn)
	if err != nil {
		log.Fatalf("Failed to parse clickhouse dsn: %v", err)
		return
	}

	clickhouseConnection, err := clickhouse.Open(clickhouseOptions)
	if err != nil {
		log.Fatalf("Connecting to clickhouse: %v", err)
	}
	defer func() {
		log.Printf("Closing clickhouse connection")

		err := clickhouseConnection.Close()
		if err != nil {
			log.Printf("Closing clickhouse connection: %v", err)
		}
	}()

	var reporterProviders = []provider.Parser{
		clover.Parser{},
		golang.Parser{},
		lcov.Parser{},
	}

	orchestratorService, err := orchestrator.NewOrchestrator(orchestrator.Config{
		ClickhouseClient:  clickhouseConnection,
		ClickhouseCluster: configuration.ClickhouseCluster,
		ReporterProviders: reporterProviders,
	})
	if err != nil {
		log.Fatalln(err)
	}

	httpServer, err := server.NewHttpServer(server.Config{
		Hostname:            configuration.HttpHostname,
		Port:                configuration.HttpPort,
		TLSConfig:           nil,
		OrchestratorService: orchestratorService,
	})
	if err != nil {
		log.Fatalln(err)
	}

	exitSignal := make(chan os.Signal, 1)

	signal.Notify(exitSignal, os.Interrupt)

	go func() {
		log.Printf("HTTP server is listening on %s", httpServer.Addr)

		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed serving HTTP server: %s", err.Error())
		}
	}()

	<-exitSignal

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err = httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("Closing http server: %s", err.Error())
	}
}
