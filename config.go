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

import "os"

type Config struct {
	Environment string

	ClickhouseDsn     string
	ClickhouseCluster string

	HttpHostname string
	HttpPort     string
}

func ParseConfig() (Config, error) {
	var config = Config{
		Environment:       "development",
		ClickhouseDsn:     "clickhouse://default:@clickhouse:9000/default?dial_timeout=200ms&max_execution_time=60&debug=false",
		ClickhouseCluster: "",
		HttpHostname:      "0.0.0.0",
		HttpPort:          "22967",
	}

	if environment, ok := os.LookupEnv("ENVIRONMENT"); ok {
		config.Environment = environment
	}

	if clickhouseDsn, ok := os.LookupEnv("CLICKHOUSE_DSN"); ok {
		config.ClickhouseDsn = clickhouseDsn
	}

	if clickhouseCluster, ok := os.LookupEnv("CLICKHOUSE_CLUSTER"); ok {
		config.ClickhouseCluster = clickhouseCluster
	}

	if httpHostname, ok := os.LookupEnv("HTTP_HOSTNAME"); ok {
		config.HttpHostname = httpHostname
	}

	if httpPort, ok := os.LookupEnv("HTTP_PORT"); ok {
		config.HttpPort = httpPort
	}

	return config, nil
}
