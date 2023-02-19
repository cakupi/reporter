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

package clover_test

import (
	"strings"
	"testing"

	"coverage-worker/provider/clover"
)

func TestParser_Parse(t *testing.T) {
	var content = strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?>
<coverage generated="1676719029875" clover="3.2.0">
  <project timestamp="1676719029875" name="All files">
    <metrics statements="1678" coveredstatements="1553" conditionals="262" coveredconditionals="218" methods="86" coveredmethods="75" elements="2026" coveredelements="1846" complexity="0" loc="1678" ncloc="1678" packages="10" files="23" classes="23"/>
    <package name="Business">
      <metrics statements="17" coveredstatements="17" conditionals="7" coveredconditionals="7" methods="3" coveredmethods="3"/>
      <file name="WeatherReportService.ts" path="/home/user/repo/example-project/src/Business/WeatherReportService.ts">
        <metrics statements="17" coveredstatements="17" conditionals="7" coveredconditionals="7" methods="3" coveredmethods="3"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="9" count="5" type="cond" truecount="1" falsecount="0"/>
        <line num="10" count="5" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="13" count="3" type="cond" truecount="2" falsecount="0"/>
        <line num="14" count="2" type="stmt"/>
        <line num="15" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="16" count="3" type="stmt"/>
        <line num="17" count="5" type="cond" truecount="1" falsecount="0"/>
      </file>
    </package>
    <package name="Models">
      <metrics statements="91" coveredstatements="51" conditionals="4" coveredconditionals="4" methods="6" coveredmethods="2"/>
      <file name="SunModel.ts" path="/home/user/repo/example-project/src/Models/SunModel.ts">
        <metrics statements="32" coveredstatements="13" conditionals="2" coveredconditionals="2" methods="3" coveredmethods="1"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="cond" truecount="2" falsecount="0"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="stmt"/>
        <line num="13" count="1" type="stmt"/>
        <line num="14" count="0" type="stmt"/>
        <line num="15" count="0" type="stmt"/>
        <line num="16" count="0" type="stmt"/>
        <line num="17" count="0" type="stmt"/>
        <line num="18" count="0" type="stmt"/>
        <line num="19" count="0" type="stmt"/>
        <line num="20" count="0" type="stmt"/>
        <line num="21" count="0" type="stmt"/>
        <line num="22" count="0" type="stmt"/>
        <line num="23" count="0" type="stmt"/>
        <line num="24" count="0" type="stmt"/>
        <line num="25" count="0" type="stmt"/>
        <line num="26" count="0" type="stmt"/>
        <line num="27" count="0" type="stmt"/>
        <line num="28" count="0" type="stmt"/>
        <line num="29" count="0" type="stmt"/>
        <line num="30" count="0" type="stmt"/>
        <line num="31" count="0" type="stmt"/>
        <line num="32" count="0" type="stmt"/>
      </file>
      <file name="StatusCode.ts" path="/home/user/repo/example-project/src/Models/StatusCode.ts">
        <metrics statements="59" coveredstatements="38" conditionals="2" coveredconditionals="2" methods="3" coveredmethods="1"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="cond" truecount="2" falsecount="0"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="stmt"/>
        <line num="13" count="1" type="stmt"/>
        <line num="14" count="1" type="stmt"/>
        <line num="15" count="1" type="stmt"/>
        <line num="16" count="1" type="stmt"/>
        <line num="17" count="1" type="stmt"/>
        <line num="18" count="1" type="stmt"/>
        <line num="19" count="1" type="stmt"/>
        <line num="20" count="1" type="stmt"/>
        <line num="21" count="1" type="stmt"/>
        <line num="22" count="1" type="stmt"/>
        <line num="23" count="1" type="stmt"/>
        <line num="24" count="1" type="stmt"/>
        <line num="25" count="1" type="stmt"/>
        <line num="26" count="1" type="stmt"/>
        <line num="27" count="1" type="stmt"/>
        <line num="28" count="1" type="stmt"/>
        <line num="29" count="1" type="stmt"/>
        <line num="30" count="1" type="stmt"/>
        <line num="31" count="1" type="stmt"/>
        <line num="32" count="1" type="stmt"/>
        <line num="33" count="1" type="stmt"/>
        <line num="34" count="1" type="stmt"/>
        <line num="35" count="1" type="stmt"/>
        <line num="36" count="1" type="stmt"/>
        <line num="37" count="1" type="stmt"/>
        <line num="38" count="1" type="stmt"/>
        <line num="39" count="0" type="stmt"/>
        <line num="40" count="0" type="stmt"/>
        <line num="41" count="0" type="stmt"/>
        <line num="42" count="0" type="stmt"/>
        <line num="43" count="0" type="stmt"/>
        <line num="44" count="0" type="stmt"/>
        <line num="45" count="0" type="stmt"/>
        <line num="46" count="0" type="stmt"/>
        <line num="47" count="0" type="stmt"/>
        <line num="48" count="0" type="stmt"/>
        <line num="49" count="0" type="stmt"/>
        <line num="50" count="0" type="stmt"/>
        <line num="51" count="0" type="stmt"/>
        <line num="52" count="0" type="stmt"/>
        <line num="53" count="0" type="stmt"/>
        <line num="54" count="0" type="stmt"/>
        <line num="55" count="0" type="stmt"/>
        <line num="56" count="0" type="stmt"/>
        <line num="57" count="0" type="stmt"/>
        <line num="58" count="0" type="stmt"/>
        <line num="59" count="0" type="stmt"/>
      </file>
    </package>
    <package name="Models.Errors">
      <metrics statements="72" coveredstatements="72" conditionals="16" coveredconditionals="16" methods="15" coveredmethods="15"/>
      <file name="ArgumentError.ts" path="/home/user/repo/example-project/src/Models/Errors/ArgumentError.ts">
        <metrics statements="12" coveredstatements="12" conditionals="2" coveredconditionals="2" methods="2" coveredmethods="2"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="7" count="54" type="stmt"/>
        <line num="8" count="54" type="stmt"/>
        <line num="9" count="54" type="stmt"/>
        <line num="10" count="54" type="stmt"/>
        <line num="11" count="54" type="stmt"/>
        <line num="12" count="54" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="ClientError.ts" path="/home/user/repo/example-project/src/Models/Errors/ClientError.ts">
        <metrics statements="19" coveredstatements="19" conditionals="4" coveredconditionals="4" methods="3" coveredmethods="3"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="cond" truecount="2" falsecount="0"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="13" count="55" type="stmt"/>
        <line num="14" count="55" type="stmt"/>
        <line num="15" count="55" type="stmt"/>
        <line num="16" count="55" type="stmt"/>
        <line num="17" count="55" type="stmt"/>
        <line num="18" count="55" type="stmt"/>
        <line num="19" count="55" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="InvalidDatabaseRequestError.ts" path="/home/user/repo/example-project/src/Models/Errors/InvalidDatabaseRequestError.ts">
        <metrics statements="5" coveredstatements="5" conditionals="2" coveredconditionals="2" methods="2" coveredmethods="2"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="3" count="9" type="stmt"/>
        <line num="4" count="9" type="stmt"/>
        <line num="5" count="9" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="InvalidEnumIndexError.ts" path="/home/user/repo/example-project/src/Models/Errors/InvalidEnumIndexError.ts">
        <metrics statements="5" coveredstatements="5" conditionals="2" coveredconditionals="2" methods="2" coveredmethods="2"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="3" count="6" type="stmt"/>
        <line num="4" count="6" type="stmt"/>
        <line num="5" count="6" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="NotApplicableError.ts" path="/home/user/repo/example-project/src/Models/Errors/NotApplicableError.ts">
        <metrics statements="12" coveredstatements="12" conditionals="2" coveredconditionals="2" methods="2" coveredmethods="2"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="7" count="5" type="stmt"/>
        <line num="8" count="5" type="stmt"/>
        <line num="9" count="5" type="stmt"/>
        <line num="10" count="5" type="stmt"/>
        <line num="11" count="5" type="stmt"/>
        <line num="12" count="5" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="NotFoundError.ts" path="/home/user/repo/example-project/src/Models/Errors/NotFoundError.ts">
        <metrics statements="7" coveredstatements="7" conditionals="2" coveredconditionals="2" methods="2" coveredmethods="2"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="3" count="8" type="stmt"/>
        <line num="4" count="8" type="stmt"/>
        <line num="5" count="8" type="stmt"/>
        <line num="6" count="8" type="stmt"/>
        <line num="7" count="8" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="UnimplementedError.ts" path="/home/user/repo/example-project/src/Models/Errors/UnimplementedError.ts">
        <metrics statements="12" coveredstatements="12" conditionals="2" coveredconditionals="2" methods="2" coveredmethods="2"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="cond" truecount="1" falsecount="0"/>
      </file>
    </package>
    <package name="Repository.PostgresAdapter">
      <metrics statements="285" coveredstatements="244" conditionals="47" coveredconditionals="35" methods="10" coveredmethods="10"/>
      <file name="PostgresAdapter.ts" path="/home/user/repo/example-project/src/Repository/PostgresAdapter/PostgresAdapter.ts">
        <metrics statements="285" coveredstatements="244" conditionals="47" coveredconditionals="35" methods="10" coveredmethods="10"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="13" count="2" type="stmt"/>
        <line num="14" count="2" type="stmt"/>
        <line num="15" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="16" count="2" type="stmt"/>
        <line num="17" count="2" type="stmt"/>
        <line num="18" count="2" type="stmt"/>
        <line num="19" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="20" count="5" type="stmt"/>
        <line num="21" count="5" type="stmt"/>
        <line num="22" count="5" type="stmt"/>
        <line num="23" count="5" type="cond" truecount="2" falsecount="0"/>
        <line num="24" count="1" type="stmt"/>
        <line num="25" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="26" count="4" type="stmt"/>
        <line num="27" count="4" type="stmt"/>
        <line num="28" count="4" type="stmt"/>
        <line num="29" count="4" type="stmt"/>
        <line num="30" count="4" type="stmt"/>
        <line num="31" count="4" type="stmt"/>
        <line num="32" count="4" type="stmt"/>
        <line num="33" count="4" type="stmt"/>
        <line num="34" count="4" type="stmt"/>
        <line num="35" count="4" type="stmt"/>
        <line num="36" count="4" type="stmt"/>
        <line num="37" count="4" type="cond" truecount="1" falsecount="0"/>
        <line num="38" count="3" type="stmt"/>
        <line num="39" count="3" type="stmt"/>
        <line num="40" count="3" type="stmt"/>
        <line num="41" count="3" type="stmt"/>
        <line num="42" count="3" type="stmt"/>
        <line num="43" count="3" type="stmt"/>
        <line num="44" count="3" type="stmt"/>
        <line num="45" count="3" type="stmt"/>
        <line num="46" count="3" type="stmt"/>
        <line num="47" count="3" type="stmt"/>
        <line num="48" count="3" type="stmt"/>
        <line num="49" count="3" type="stmt"/>
        <line num="50" count="3" type="stmt"/>
        <line num="51" count="3" type="stmt"/>
        <line num="52" count="3" type="stmt"/>
        <line num="53" count="3" type="stmt"/>
        <line num="54" count="3" type="stmt"/>
        <line num="55" count="3" type="stmt"/>
        <line num="56" count="3" type="stmt"/>
        <line num="57" count="3" type="stmt"/>
        <line num="58" count="3" type="stmt"/>
        <line num="59" count="3" type="stmt"/>
        <line num="60" count="3" type="stmt"/>
        <line num="61" count="3" type="stmt"/>
        <line num="62" count="3" type="cond" truecount="1" falsecount="0"/>
        <line num="63" count="2" type="stmt"/>
        <line num="64" count="2" type="stmt"/>
        <line num="65" count="2" type="cond" truecount="0" falsecount="1"/>
        <line num="66" count="0" type="stmt"/>
        <line num="67" count="0" type="stmt"/>
        <line num="68" count="0" type="stmt"/>
        <line num="69" count="0" type="stmt"/>
        <line num="70" count="0" type="stmt"/>
        <line num="71" count="0" type="stmt"/>
        <line num="72" count="0" type="stmt"/>
        <line num="73" count="0" type="stmt"/>
        <line num="74" count="0" type="stmt"/>
        <line num="75" count="0" type="stmt"/>
        <line num="76" count="0" type="stmt"/>
        <line num="77" count="0" type="stmt"/>
        <line num="78" count="0" type="stmt"/>
        <line num="79" count="0" type="stmt"/>
        <line num="80" count="0" type="stmt"/>
        <line num="81" count="0" type="stmt"/>
        <line num="82" count="0" type="stmt"/>
        <line num="83" count="0" type="stmt"/>
        <line num="84" count="0" type="stmt"/>
        <line num="85" count="0" type="stmt"/>
        <line num="86" count="0" type="stmt"/>
        <line num="87" count="0" type="stmt"/>
        <line num="88" count="0" type="stmt"/>
        <line num="89" count="0" type="stmt"/>
        <line num="90" count="0" type="stmt"/>
        <line num="91" count="0" type="stmt"/>
        <line num="92" count="5" type="cond" truecount="1" falsecount="0"/>
        <line num="93" count="1" type="stmt"/>
        <line num="94" count="1" type="stmt"/>
        <line num="95" count="1" type="stmt"/>
        <line num="96" count="1" type="stmt"/>
        <line num="97" count="1" type="cond" truecount="0" falsecount="1"/>
        <line num="98" count="0" type="stmt"/>
        <line num="99" count="0" type="stmt"/>
        <line num="100" count="5" type="stmt"/>
        <line num="101" count="2" type="stmt"/>
        <line num="102" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="103" count="2" type="stmt"/>
        <line num="104" count="2" type="stmt"/>
        <line num="105" count="2" type="stmt"/>
        <line num="106" count="2" type="stmt"/>
        <line num="107" count="2" type="stmt"/>
        <line num="108" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="109" count="1" type="stmt"/>
        <line num="110" count="1" type="stmt"/>
        <line num="111" count="1" type="stmt"/>
        <line num="112" count="1" type="stmt"/>
        <line num="113" count="1" type="stmt"/>
        <line num="114" count="1" type="stmt"/>
        <line num="115" count="1" type="stmt"/>
        <line num="116" count="1" type="stmt"/>
        <line num="117" count="1" type="stmt"/>
        <line num="118" count="1" type="stmt"/>
        <line num="119" count="1" type="stmt"/>
        <line num="120" count="1" type="stmt"/>
        <line num="121" count="1" type="stmt"/>
        <line num="122" count="1" type="stmt"/>
        <line num="123" count="1" type="cond" truecount="0" falsecount="1"/>
        <line num="124" count="0" type="stmt"/>
        <line num="125" count="0" type="stmt"/>
        <line num="126" count="2" type="stmt"/>
        <line num="127" count="2" type="stmt"/>
        <line num="128" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="129" count="5" type="stmt"/>
        <line num="130" count="5" type="stmt"/>
        <line num="131" count="5" type="stmt"/>
        <line num="132" count="5" type="stmt"/>
        <line num="133" count="5" type="stmt"/>
        <line num="134" count="5" type="stmt"/>
        <line num="135" count="5" type="stmt"/>
        <line num="136" count="5" type="stmt"/>
        <line num="137" count="5" type="stmt"/>
        <line num="138" count="5" type="stmt"/>
        <line num="139" count="5" type="cond" truecount="2" falsecount="0"/>
        <line num="140" count="2" type="stmt"/>
        <line num="141" count="2" type="stmt"/>
        <line num="142" count="2" type="stmt"/>
        <line num="143" count="2" type="stmt"/>
        <line num="144" count="2" type="stmt"/>
        <line num="145" count="2" type="stmt"/>
        <line num="146" count="2" type="stmt"/>
        <line num="147" count="2" type="stmt"/>
        <line num="148" count="2" type="stmt"/>
        <line num="149" count="4" type="cond" truecount="1" falsecount="0"/>
        <line num="150" count="3" type="cond" truecount="2" falsecount="0"/>
        <line num="151" count="1" type="stmt"/>
        <line num="152" count="1" type="stmt"/>
        <line num="153" count="1" type="stmt"/>
        <line num="154" count="1" type="stmt"/>
        <line num="155" count="2" type="stmt"/>
        <line num="156" count="2" type="stmt"/>
        <line num="157" count="5" type="stmt"/>
        <line num="158" count="2" type="stmt"/>
        <line num="159" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="160" count="4" type="stmt"/>
        <line num="161" count="4" type="stmt"/>
        <line num="162" count="4" type="stmt"/>
        <line num="163" count="4" type="stmt"/>
        <line num="164" count="4" type="stmt"/>
        <line num="165" count="4" type="stmt"/>
        <line num="166" count="4" type="stmt"/>
        <line num="167" count="4" type="stmt"/>
        <line num="168" count="4" type="stmt"/>
        <line num="169" count="4" type="stmt"/>
        <line num="170" count="4" type="stmt"/>
        <line num="171" count="4" type="cond" truecount="1" falsecount="0"/>
        <line num="172" count="3" type="stmt"/>
        <line num="173" count="3" type="cond" truecount="1" falsecount="0"/>
        <line num="174" count="3" type="stmt"/>
        <line num="175" count="3" type="stmt"/>
        <line num="176" count="3" type="stmt"/>
        <line num="177" count="3" type="stmt"/>
        <line num="178" count="3" type="stmt"/>
        <line num="179" count="3" type="stmt"/>
        <line num="180" count="3" type="stmt"/>
        <line num="181" count="3" type="stmt"/>
        <line num="182" count="3" type="stmt"/>
        <line num="183" count="3" type="stmt"/>
        <line num="184" count="3" type="stmt"/>
        <line num="185" count="3" type="stmt"/>
        <line num="186" count="3" type="stmt"/>
        <line num="187" count="3" type="stmt"/>
        <line num="188" count="3" type="stmt"/>
        <line num="189" count="3" type="stmt"/>
        <line num="190" count="3" type="stmt"/>
        <line num="191" count="3" type="stmt"/>
        <line num="192" count="4" type="cond" truecount="1" falsecount="0"/>
        <line num="193" count="1" type="stmt"/>
        <line num="194" count="1" type="stmt"/>
        <line num="195" count="1" type="stmt"/>
        <line num="196" count="1" type="stmt"/>
        <line num="197" count="1" type="cond" truecount="0" falsecount="1"/>
        <line num="198" count="0" type="stmt"/>
        <line num="199" count="0" type="stmt"/>
        <line num="200" count="4" type="stmt"/>
        <line num="201" count="2" type="stmt"/>
        <line num="202" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="203" count="2" type="stmt"/>
        <line num="204" count="2" type="stmt"/>
        <line num="205" count="2" type="stmt"/>
        <line num="206" count="2" type="stmt"/>
        <line num="207" count="2" type="stmt"/>
        <line num="208" count="2" type="stmt"/>
        <line num="209" count="2" type="stmt"/>
        <line num="210" count="2" type="stmt"/>
        <line num="211" count="2" type="stmt"/>
        <line num="212" count="2" type="stmt"/>
        <line num="213" count="2" type="stmt"/>
        <line num="214" count="2" type="stmt"/>
        <line num="215" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="216" count="1" type="stmt"/>
        <line num="217" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="218" count="1" type="stmt"/>
        <line num="219" count="1" type="stmt"/>
        <line num="220" count="1" type="stmt"/>
        <line num="221" count="1" type="stmt"/>
        <line num="222" count="1" type="stmt"/>
        <line num="223" count="1" type="stmt"/>
        <line num="224" count="1" type="stmt"/>
        <line num="225" count="1" type="stmt"/>
        <line num="226" count="1" type="stmt"/>
        <line num="227" count="1" type="stmt"/>
        <line num="228" count="1" type="stmt"/>
        <line num="229" count="1" type="stmt"/>
        <line num="230" count="1" type="stmt"/>
        <line num="231" count="1" type="stmt"/>
        <line num="232" count="1" type="stmt"/>
        <line num="233" count="1" type="stmt"/>
        <line num="234" count="1" type="stmt"/>
        <line num="235" count="1" type="stmt"/>
        <line num="236" count="1" type="stmt"/>
        <line num="237" count="1" type="stmt"/>
        <line num="238" count="1" type="stmt"/>
        <line num="239" count="1" type="stmt"/>
        <line num="240" count="1" type="stmt"/>
        <line num="241" count="1" type="cond" truecount="0" falsecount="1"/>
        <line num="242" count="0" type="stmt"/>
        <line num="243" count="0" type="stmt"/>
        <line num="244" count="2" type="stmt"/>
        <line num="245" count="2" type="stmt"/>
        <line num="246" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="247" count="8" type="stmt"/>
        <line num="248" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="249" count="4" type="stmt"/>
        <line num="250" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="251" count="2" type="stmt"/>
        <line num="252" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="253" count="1" type="stmt"/>
        <line num="254" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="255" count="0" type="stmt"/>
        <line num="256" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="257" count="0" type="stmt"/>
        <line num="258" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="259" count="1" type="stmt"/>
        <line num="260" count="8" type="stmt"/>
        <line num="261" count="8" type="stmt"/>
        <line num="262" count="2" type="stmt"/>
        <line num="263" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="264" count="9" type="stmt"/>
        <line num="265" count="9" type="cond" truecount="1" falsecount="0"/>
        <line num="266" count="2" type="stmt"/>
        <line num="267" count="9" type="cond" truecount="1" falsecount="0"/>
        <line num="268" count="4" type="stmt"/>
        <line num="269" count="9" type="cond" truecount="1" falsecount="0"/>
        <line num="270" count="1" type="stmt"/>
        <line num="271" count="9" type="cond" truecount="0" falsecount="1"/>
        <line num="272" count="0" type="stmt"/>
        <line num="273" count="9" type="cond" truecount="0" falsecount="1"/>
        <line num="274" count="0" type="stmt"/>
        <line num="275" count="9" type="cond" truecount="0" falsecount="1"/>
        <line num="276" count="0" type="stmt"/>
        <line num="277" count="9" type="cond" truecount="0" falsecount="1"/>
        <line num="278" count="0" type="stmt"/>
        <line num="279" count="9" type="cond" truecount="0" falsecount="1"/>
        <line num="280" count="0" type="stmt"/>
        <line num="281" count="9" type="cond" truecount="1" falsecount="0"/>
        <line num="282" count="2" type="stmt"/>
        <line num="283" count="9" type="stmt"/>
        <line num="284" count="9" type="stmt"/>
        <line num="285" count="2" type="cond" truecount="1" falsecount="0"/>
      </file>
    </package>
    <package name="Repository.Logger">
      <metrics statements="44" coveredstatements="42" conditionals="5" coveredconditionals="4" methods="3" coveredmethods="3"/>
      <file name="Logger.ts" path="/home/user/repo/example-project/src/Repository/Logger/Logger.ts">
        <metrics statements="44" coveredstatements="42" conditionals="5" coveredconditionals="4" methods="3" coveredmethods="3"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="8" count="3" type="stmt"/>
        <line num="9" count="3" type="stmt"/>
        <line num="10" count="3" type="stmt"/>
        <line num="11" count="3" type="cond" truecount="1" falsecount="0"/>
        <line num="12" count="3" type="stmt"/>
        <line num="13" count="1" type="stmt"/>
        <line num="14" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="15" count="1001" type="stmt"/>
        <line num="16" count="1001" type="stmt"/>
        <line num="17" count="1001" type="cond" truecount="0" falsecount="1"/>
        <line num="18" count="0" type="stmt"/>
        <line num="19" count="0" type="stmt"/>
        <line num="20" count="1001" type="stmt"/>
        <line num="21" count="1001" type="stmt"/>
        <line num="22" count="1001" type="stmt"/>
        <line num="23" count="1001" type="stmt"/>
        <line num="24" count="1001" type="stmt"/>
        <line num="25" count="1001" type="stmt"/>
        <line num="26" count="1001" type="stmt"/>
        <line num="27" count="1001" type="stmt"/>
        <line num="28" count="1001" type="stmt"/>
        <line num="29" count="1001" type="stmt"/>
        <line num="30" count="1001" type="stmt"/>
        <line num="31" count="1001" type="stmt"/>
        <line num="32" count="1001" type="stmt"/>
        <line num="33" count="1001" type="stmt"/>
        <line num="34" count="1001" type="stmt"/>
        <line num="35" count="1001" type="stmt"/>
        <line num="36" count="1001" type="stmt"/>
        <line num="37" count="1001" type="stmt"/>
        <line num="38" count="1001" type="stmt"/>
        <line num="39" count="1001" type="stmt"/>
        <line num="40" count="1001" type="stmt"/>
        <line num="41" count="1001" type="stmt"/>
        <line num="42" count="1001" type="stmt"/>
        <line num="43" count="1001" type="stmt"/>
        <line num="44" count="3" type="cond" truecount="1" falsecount="0"/>
      </file>
    </package>
    <package name="Repository.ImageGenerator">
      <metrics statements="57" coveredstatements="57" conditionals="11" coveredconditionals="11" methods="2" coveredmethods="2"/>
      <file name="ImageGenerator.ts" path="/home/user/repo/example-project/src/Repository/ImageGenerator/ImageGenerator.ts">
        <metrics statements="57" coveredstatements="57" conditionals="11" coveredconditionals="11" methods="2" coveredmethods="2"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="stmt"/>
        <line num="13" count="1" type="stmt"/>
        <line num="14" count="1" type="stmt"/>
        <line num="15" count="1" type="stmt"/>
        <line num="16" count="1" type="stmt"/>
        <line num="17" count="1" type="stmt"/>
        <line num="18" count="1" type="stmt"/>
        <line num="19" count="1" type="stmt"/>
        <line num="20" count="1" type="stmt"/>
        <line num="21" count="1" type="stmt"/>
        <line num="22" count="1" type="stmt"/>
        <line num="23" count="1" type="stmt"/>
        <line num="24" count="1" type="stmt"/>
        <line num="25" count="1" type="stmt"/>
        <line num="26" count="1" type="stmt"/>
        <line num="27" count="1" type="stmt"/>
        <line num="28" count="1" type="stmt"/>
        <line num="29" count="1" type="stmt"/>
        <line num="30" count="1" type="stmt"/>
        <line num="31" count="1" type="stmt"/>
        <line num="32" count="1" type="stmt"/>
        <line num="33" count="1" type="stmt"/>
        <line num="34" count="1" type="stmt"/>
        <line num="35" count="1" type="stmt"/>
        <line num="36" count="1" type="stmt"/>
        <line num="37" count="1" type="stmt"/>
        <line num="38" count="1" type="stmt"/>
        <line num="39" count="1" type="stmt"/>
        <line num="40" count="1" type="stmt"/>
        <line num="41" count="1" type="stmt"/>
        <line num="42" count="1" type="stmt"/>
        <line num="43" count="1" type="stmt"/>
        <line num="44" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="45" count="6" type="cond" truecount="2" falsecount="0"/>
        <line num="46" count="6" type="cond" truecount="2" falsecount="0"/>
        <line num="47" count="6" type="cond" truecount="2" falsecount="0"/>
        <line num="48" count="6" type="cond" truecount="2" falsecount="0"/>
        <line num="49" count="6" type="stmt"/>
        <line num="50" count="6" type="stmt"/>
        <line num="51" count="6" type="cond" truecount="1" falsecount="0"/>
        <line num="52" count="40" type="stmt"/>
        <line num="53" count="40" type="stmt"/>
        <line num="54" count="40" type="stmt"/>
        <line num="55" count="6" type="stmt"/>
        <line num="56" count="6" type="stmt"/>
        <line num="57" count="1" type="cond" truecount="1" falsecount="0"/>
      </file>
    </package>
    <package name="Repository.Encryption">
      <metrics statements="12" coveredstatements="12" conditionals="3" coveredconditionals="3" methods="3" coveredmethods="3"/>
      <file name="Ed25519.ts" path="/home/user/repo/example-project/src/Repository/Encryption/Ed25519.ts">
        <metrics statements="12" coveredstatements="12" conditionals="3" coveredconditionals="3" methods="3" coveredmethods="3"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="6" count="2" type="stmt"/>
        <line num="7" count="2" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="10" count="2" type="stmt"/>
        <line num="11" count="2" type="stmt"/>
        <line num="12" count="2" type="cond" truecount="1" falsecount="0"/>
      </file>
    </package>
    <package name="Repository.WeatherProvider">
      <metrics statements="811" coveredstatements="793" conditionals="138" coveredconditionals="109" methods="20" coveredmethods="18"/>
      <file name="AsianProvider.ts" path="/home/user/repo/example-project/src/Repository/WeatherProvider/AsianProvider.ts">
        <metrics statements="255" coveredstatements="248" conditionals="43" coveredconditionals="34" methods="5" coveredmethods="4"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="stmt"/>
        <line num="13" count="1" type="stmt"/>
        <line num="14" count="1" type="stmt"/>
        <line num="15" count="1" type="stmt"/>
        <line num="16" count="1" type="stmt"/>
        <line num="17" count="1" type="stmt"/>
        <line num="18" count="1" type="stmt"/>
        <line num="19" count="1" type="stmt"/>
        <line num="20" count="1" type="stmt"/>
        <line num="21" count="1" type="stmt"/>
        <line num="22" count="1" type="stmt"/>
        <line num="23" count="1" type="stmt"/>
        <line num="24" count="1" type="stmt"/>
        <line num="25" count="1" type="stmt"/>
        <line num="26" count="1" type="stmt"/>
        <line num="27" count="1" type="stmt"/>
        <line num="28" count="1" type="stmt"/>
        <line num="29" count="1" type="stmt"/>
        <line num="30" count="1" type="stmt"/>
        <line num="31" count="1" type="stmt"/>
        <line num="32" count="1" type="stmt"/>
        <line num="33" count="1" type="stmt"/>
        <line num="34" count="1" type="stmt"/>
        <line num="35" count="1" type="stmt"/>
        <line num="36" count="1" type="stmt"/>
        <line num="37" count="1" type="stmt"/>
        <line num="38" count="1" type="stmt"/>
        <line num="39" count="1" type="stmt"/>
        <line num="40" count="1" type="stmt"/>
        <line num="41" count="1" type="stmt"/>
        <line num="42" count="1" type="stmt"/>
        <line num="43" count="1" type="stmt"/>
        <line num="44" count="1" type="stmt"/>
        <line num="45" count="1" type="stmt"/>
        <line num="46" count="1" type="stmt"/>
        <line num="47" count="1" type="stmt"/>
        <line num="48" count="1" type="stmt"/>
        <line num="49" count="1" type="stmt"/>
        <line num="50" count="1" type="stmt"/>
        <line num="51" count="1" type="stmt"/>
        <line num="52" count="1" type="stmt"/>
        <line num="53" count="1" type="stmt"/>
        <line num="54" count="1" type="stmt"/>
        <line num="55" count="1" type="stmt"/>
        <line num="56" count="1" type="stmt"/>
        <line num="57" count="1" type="stmt"/>
        <line num="58" count="1" type="stmt"/>
        <line num="59" count="1" type="stmt"/>
        <line num="60" count="1" type="stmt"/>
        <line num="61" count="1" type="stmt"/>
        <line num="62" count="1" type="stmt"/>
        <line num="63" count="1" type="stmt"/>
        <line num="64" count="1" type="stmt"/>
        <line num="65" count="1" type="stmt"/>
        <line num="66" count="1" type="stmt"/>
        <line num="67" count="1" type="stmt"/>
        <line num="68" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="69" count="22" type="stmt"/>
        <line num="70" count="22" type="stmt"/>
        <line num="71" count="22" type="stmt"/>
        <line num="72" count="22" type="stmt"/>
        <line num="73" count="22" type="stmt"/>
        <line num="74" count="22" type="stmt"/>
        <line num="75" count="22" type="cond" truecount="2" falsecount="0"/>
        <line num="76" count="22" type="cond" truecount="1" falsecount="0"/>
        <line num="77" count="22" type="stmt"/>
        <line num="78" count="1" type="stmt"/>
        <line num="79" count="1" type="stmt"/>
        <line num="80" count="0" type="stmt"/>
        <line num="81" count="0" type="stmt"/>
        <line num="82" count="1" type="stmt"/>
        <line num="83" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="84" count="8" type="cond" truecount="2" falsecount="0"/>
        <line num="85" count="8" type="cond" truecount="2" falsecount="0"/>
        <line num="86" count="6" type="stmt"/>
        <line num="87" count="6" type="stmt"/>
        <line num="88" count="6" type="stmt"/>
        <line num="89" count="6" type="stmt"/>
        <line num="90" count="6" type="stmt"/>
        <line num="91" count="6" type="stmt"/>
        <line num="92" count="6" type="stmt"/>
        <line num="93" count="6" type="stmt"/>
        <line num="94" count="6" type="stmt"/>
        <line num="95" count="6" type="stmt"/>
        <line num="96" count="6" type="stmt"/>
        <line num="97" count="6" type="stmt"/>
        <line num="98" count="6" type="stmt"/>
        <line num="99" count="6" type="stmt"/>
        <line num="100" count="6" type="stmt"/>
        <line num="101" count="6" type="stmt"/>
        <line num="102" count="6" type="stmt"/>
        <line num="103" count="6" type="stmt"/>
        <line num="104" count="6" type="stmt"/>
        <line num="105" count="6" type="stmt"/>
        <line num="106" count="6" type="stmt"/>
        <line num="107" count="6" type="stmt"/>
        <line num="108" count="6" type="stmt"/>
        <line num="109" count="6" type="stmt"/>
        <line num="110" count="6" type="stmt"/>
        <line num="111" count="6" type="stmt"/>
        <line num="112" count="6" type="stmt"/>
        <line num="113" count="6" type="stmt"/>
        <line num="114" count="6" type="stmt"/>
        <line num="115" count="6" type="stmt"/>
        <line num="116" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="117" count="8" type="stmt"/>
        <line num="118" count="8" type="cond" truecount="1" falsecount="1"/>
        <line num="119" count="8" type="stmt"/>
        <line num="120" count="8" type="stmt"/>
        <line num="121" count="8" type="cond" truecount="1" falsecount="1"/>
        <line num="122" count="8" type="stmt"/>
        <line num="123" count="8" type="stmt"/>
        <line num="124" count="8" type="stmt"/>
        <line num="125" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="126" count="1" type="stmt"/>
        <line num="127" count="1" type="stmt"/>
        <line num="128" count="1" type="stmt"/>
        <line num="129" count="1" type="stmt"/>
        <line num="130" count="1" type="stmt"/>
        <line num="131" count="1" type="stmt"/>
        <line num="132" count="1" type="stmt"/>
        <line num="133" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="134" count="1" type="stmt"/>
        <line num="135" count="1" type="stmt"/>
        <line num="136" count="1" type="stmt"/>
        <line num="137" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="138" count="1" type="stmt"/>
        <line num="139" count="1" type="stmt"/>
        <line num="140" count="1" type="stmt"/>
        <line num="141" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="142" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="143" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="144" count="1" type="stmt"/>
        <line num="145" count="1" type="stmt"/>
        <line num="146" count="1" type="stmt"/>
        <line num="147" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="148" count="1" type="stmt"/>
        <line num="149" count="1" type="stmt"/>
        <line num="150" count="1" type="stmt"/>
        <line num="151" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="152" count="1" type="stmt"/>
        <line num="153" count="1" type="stmt"/>
        <line num="154" count="8" type="stmt"/>
        <line num="155" count="1" type="stmt"/>
        <line num="156" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="157" count="12" type="cond" truecount="2" falsecount="0"/>
        <line num="158" count="11" type="stmt"/>
        <line num="159" count="11" type="stmt"/>
        <line num="160" count="11" type="stmt"/>
        <line num="161" count="11" type="stmt"/>
        <line num="162" count="11" type="stmt"/>
        <line num="163" count="11" type="stmt"/>
        <line num="164" count="11" type="stmt"/>
        <line num="165" count="11" type="stmt"/>
        <line num="166" count="11" type="stmt"/>
        <line num="167" count="11" type="stmt"/>
        <line num="168" count="11" type="stmt"/>
        <line num="169" count="11" type="stmt"/>
        <line num="170" count="11" type="stmt"/>
        <line num="171" count="11" type="stmt"/>
        <line num="172" count="11" type="stmt"/>
        <line num="173" count="11" type="stmt"/>
        <line num="174" count="11" type="stmt"/>
        <line num="175" count="11" type="stmt"/>
        <line num="176" count="11" type="stmt"/>
        <line num="177" count="11" type="stmt"/>
        <line num="178" count="11" type="stmt"/>
        <line num="179" count="11" type="stmt"/>
        <line num="180" count="11" type="cond" truecount="1" falsecount="0"/>
        <line num="181" count="10" type="stmt"/>
        <line num="182" count="10" type="stmt"/>
        <line num="183" count="10" type="stmt"/>
        <line num="184" count="10" type="stmt"/>
        <line num="185" count="12" type="cond" truecount="0" falsecount="1"/>
        <line num="186" count="12" type="stmt"/>
        <line num="187" count="12" type="stmt"/>
        <line num="188" count="12" type="stmt"/>
        <line num="189" count="12" type="stmt"/>
        <line num="190" count="12" type="cond" truecount="1" falsecount="1"/>
        <line num="191" count="12" type="stmt"/>
        <line num="192" count="12" type="stmt"/>
        <line num="193" count="12" type="stmt"/>
        <line num="194" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="195" count="8" type="stmt"/>
        <line num="196" count="8" type="stmt"/>
        <line num="197" count="8" type="stmt"/>
        <line num="198" count="8" type="stmt"/>
        <line num="199" count="8" type="stmt"/>
        <line num="200" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="201" count="1" type="stmt"/>
        <line num="202" count="1" type="stmt"/>
        <line num="203" count="1" type="stmt"/>
        <line num="204" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="205" count="1" type="stmt"/>
        <line num="206" count="1" type="stmt"/>
        <line num="207" count="1" type="stmt"/>
        <line num="208" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="209" count="1" type="stmt"/>
        <line num="210" count="1" type="stmt"/>
        <line num="211" count="1" type="stmt"/>
        <line num="212" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="213" count="1" type="stmt"/>
        <line num="214" count="1" type="stmt"/>
        <line num="215" count="1" type="stmt"/>
        <line num="216" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="217" count="1" type="stmt"/>
        <line num="218" count="1" type="stmt"/>
        <line num="219" count="1" type="stmt"/>
        <line num="220" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="221" count="1" type="stmt"/>
        <line num="222" count="1" type="stmt"/>
        <line num="223" count="1" type="stmt"/>
        <line num="224" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="225" count="1" type="stmt"/>
        <line num="226" count="1" type="stmt"/>
        <line num="227" count="1" type="stmt"/>
        <line num="228" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="229" count="1" type="stmt"/>
        <line num="230" count="1" type="stmt"/>
        <line num="231" count="8" type="stmt"/>
        <line num="232" count="8" type="stmt"/>
        <line num="233" count="8" type="stmt"/>
        <line num="234" count="8" type="stmt"/>
        <line num="235" count="8" type="stmt"/>
        <line num="236" count="8" type="stmt"/>
        <line num="237" count="8" type="stmt"/>
        <line num="238" count="8" type="stmt"/>
        <line num="239" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="240" count="1" type="stmt"/>
        <line num="241" count="1" type="stmt"/>
        <line num="242" count="1" type="stmt"/>
        <line num="243" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="244" count="1" type="stmt"/>
        <line num="245" count="1" type="stmt"/>
        <line num="246" count="1" type="stmt"/>
        <line num="247" count="12" type="cond" truecount="0" falsecount="1"/>
        <line num="248" count="0" type="stmt"/>
        <line num="249" count="0" type="stmt"/>
        <line num="250" count="0" type="stmt"/>
        <line num="251" count="12" type="cond" truecount="0" falsecount="1"/>
        <line num="252" count="0" type="stmt"/>
        <line num="253" count="0" type="stmt"/>
        <line num="254" count="12" type="stmt"/>
        <line num="255" count="22" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="EuropeanProvider.ts" path="/home/user/repo/example-project/src/Repository/WeatherProvider/EuropeanProvider.ts">
        <metrics statements="266" coveredstatements="257" conditionals="47" coveredconditionals="37" methods="5" coveredmethods="5"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="stmt"/>
        <line num="13" count="1" type="stmt"/>
        <line num="14" count="1" type="stmt"/>
        <line num="15" count="1" type="stmt"/>
        <line num="16" count="1" type="stmt"/>
        <line num="17" count="1" type="stmt"/>
        <line num="18" count="1" type="stmt"/>
        <line num="19" count="1" type="stmt"/>
        <line num="20" count="1" type="stmt"/>
        <line num="21" count="1" type="stmt"/>
        <line num="22" count="1" type="stmt"/>
        <line num="23" count="1" type="stmt"/>
        <line num="24" count="1" type="stmt"/>
        <line num="25" count="1" type="stmt"/>
        <line num="26" count="1" type="stmt"/>
        <line num="27" count="1" type="stmt"/>
        <line num="28" count="1" type="stmt"/>
        <line num="29" count="1" type="stmt"/>
        <line num="30" count="1" type="stmt"/>
        <line num="31" count="1" type="stmt"/>
        <line num="32" count="1" type="stmt"/>
        <line num="33" count="1" type="stmt"/>
        <line num="34" count="1" type="stmt"/>
        <line num="35" count="1" type="stmt"/>
        <line num="36" count="1" type="stmt"/>
        <line num="37" count="1" type="stmt"/>
        <line num="38" count="1" type="stmt"/>
        <line num="39" count="1" type="stmt"/>
        <line num="40" count="1" type="stmt"/>
        <line num="41" count="1" type="stmt"/>
        <line num="42" count="1" type="stmt"/>
        <line num="43" count="1" type="stmt"/>
        <line num="44" count="1" type="stmt"/>
        <line num="45" count="1" type="stmt"/>
        <line num="46" count="1" type="stmt"/>
        <line num="47" count="1" type="stmt"/>
        <line num="48" count="1" type="stmt"/>
        <line num="49" count="1" type="stmt"/>
        <line num="50" count="1" type="stmt"/>
        <line num="51" count="1" type="stmt"/>
        <line num="52" count="1" type="stmt"/>
        <line num="53" count="1" type="stmt"/>
        <line num="54" count="1" type="stmt"/>
        <line num="55" count="1" type="stmt"/>
        <line num="56" count="1" type="stmt"/>
        <line num="57" count="1" type="stmt"/>
        <line num="58" count="1" type="stmt"/>
        <line num="59" count="1" type="stmt"/>
        <line num="60" count="1" type="stmt"/>
        <line num="61" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="62" count="26" type="stmt"/>
        <line num="63" count="26" type="stmt"/>
        <line num="64" count="26" type="stmt"/>
        <line num="65" count="26" type="stmt"/>
        <line num="66" count="26" type="stmt"/>
        <line num="67" count="26" type="stmt"/>
        <line num="68" count="26" type="cond" truecount="2" falsecount="0"/>
        <line num="69" count="26" type="cond" truecount="1" falsecount="0"/>
        <line num="70" count="26" type="stmt"/>
        <line num="71" count="1" type="stmt"/>
        <line num="72" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="73" count="1" type="stmt"/>
        <line num="74" count="1" type="stmt"/>
        <line num="75" count="1" type="stmt"/>
        <line num="76" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="77" count="8" type="cond" truecount="2" falsecount="0"/>
        <line num="78" count="8" type="cond" truecount="2" falsecount="0"/>
        <line num="79" count="6" type="stmt"/>
        <line num="80" count="6" type="stmt"/>
        <line num="81" count="6" type="stmt"/>
        <line num="82" count="6" type="stmt"/>
        <line num="83" count="6" type="stmt"/>
        <line num="84" count="6" type="stmt"/>
        <line num="85" count="6" type="stmt"/>
        <line num="86" count="6" type="stmt"/>
        <line num="87" count="6" type="stmt"/>
        <line num="88" count="6" type="stmt"/>
        <line num="89" count="6" type="stmt"/>
        <line num="90" count="6" type="stmt"/>
        <line num="91" count="6" type="stmt"/>
        <line num="92" count="6" type="stmt"/>
        <line num="93" count="6" type="stmt"/>
        <line num="94" count="6" type="stmt"/>
        <line num="95" count="6" type="stmt"/>
        <line num="96" count="6" type="stmt"/>
        <line num="97" count="6" type="stmt"/>
        <line num="98" count="6" type="stmt"/>
        <line num="99" count="6" type="stmt"/>
        <line num="100" count="6" type="stmt"/>
        <line num="101" count="6" type="stmt"/>
        <line num="102" count="6" type="stmt"/>
        <line num="103" count="6" type="stmt"/>
        <line num="104" count="6" type="stmt"/>
        <line num="105" count="6" type="stmt"/>
        <line num="106" count="6" type="stmt"/>
        <line num="107" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="108" count="8" type="stmt"/>
        <line num="109" count="8" type="cond" truecount="1" falsecount="1"/>
        <line num="110" count="8" type="stmt"/>
        <line num="111" count="8" type="stmt"/>
        <line num="112" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="113" count="8" type="stmt"/>
        <line num="114" count="8" type="stmt"/>
        <line num="115" count="8" type="stmt"/>
        <line num="116" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="117" count="1" type="stmt"/>
        <line num="118" count="1" type="stmt"/>
        <line num="119" count="1" type="stmt"/>
        <line num="120" count="1" type="stmt"/>
        <line num="121" count="1" type="stmt"/>
        <line num="122" count="1" type="stmt"/>
        <line num="123" count="1" type="stmt"/>
        <line num="124" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="125" count="1" type="stmt"/>
        <line num="126" count="1" type="stmt"/>
        <line num="127" count="1" type="stmt"/>
        <line num="128" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="129" count="1" type="stmt"/>
        <line num="130" count="1" type="stmt"/>
        <line num="131" count="1" type="stmt"/>
        <line num="132" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="133" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="134" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="135" count="1" type="stmt"/>
        <line num="136" count="1" type="stmt"/>
        <line num="137" count="1" type="stmt"/>
        <line num="138" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="139" count="1" type="stmt"/>
        <line num="140" count="1" type="stmt"/>
        <line num="141" count="1" type="stmt"/>
        <line num="142" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="143" count="1" type="stmt"/>
        <line num="144" count="1" type="stmt"/>
        <line num="145" count="1" type="stmt"/>
        <line num="146" count="1" type="stmt"/>
        <line num="147" count="1" type="stmt"/>
        <line num="148" count="1" type="stmt"/>
        <line num="149" count="8" type="stmt"/>
        <line num="150" count="1" type="stmt"/>
        <line num="151" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="152" count="16" type="cond" truecount="2" falsecount="0"/>
        <line num="153" count="15" type="stmt"/>
        <line num="154" count="15" type="stmt"/>
        <line num="155" count="15" type="stmt"/>
        <line num="156" count="15" type="stmt"/>
        <line num="157" count="15" type="stmt"/>
        <line num="158" count="15" type="stmt"/>
        <line num="159" count="15" type="stmt"/>
        <line num="160" count="15" type="stmt"/>
        <line num="161" count="15" type="stmt"/>
        <line num="162" count="15" type="stmt"/>
        <line num="163" count="15" type="stmt"/>
        <line num="164" count="15" type="stmt"/>
        <line num="165" count="15" type="stmt"/>
        <line num="166" count="15" type="stmt"/>
        <line num="167" count="15" type="stmt"/>
        <line num="168" count="15" type="stmt"/>
        <line num="169" count="15" type="stmt"/>
        <line num="170" count="15" type="stmt"/>
        <line num="171" count="15" type="cond" truecount="1" falsecount="0"/>
        <line num="172" count="14" type="stmt"/>
        <line num="173" count="14" type="stmt"/>
        <line num="174" count="14" type="stmt"/>
        <line num="175" count="14" type="stmt"/>
        <line num="176" count="16" type="cond" truecount="0" falsecount="1"/>
        <line num="177" count="16" type="stmt"/>
        <line num="178" count="16" type="cond" truecount="1" falsecount="1"/>
        <line num="179" count="16" type="stmt"/>
        <line num="180" count="16" type="stmt"/>
        <line num="181" count="16" type="cond" truecount="0" falsecount="1"/>
        <line num="182" count="16" type="stmt"/>
        <line num="183" count="16" type="stmt"/>
        <line num="184" count="16" type="stmt"/>
        <line num="185" count="16" type="cond" truecount="1" falsecount="0"/>
        <line num="186" count="12" type="stmt"/>
        <line num="187" count="12" type="stmt"/>
        <line num="188" count="12" type="stmt"/>
        <line num="189" count="12" type="stmt"/>
        <line num="190" count="12" type="stmt"/>
        <line num="191" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="192" count="1" type="stmt"/>
        <line num="193" count="1" type="stmt"/>
        <line num="194" count="1" type="stmt"/>
        <line num="195" count="12" type="stmt"/>
        <line num="196" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="197" count="1" type="stmt"/>
        <line num="198" count="1" type="stmt"/>
        <line num="199" count="1" type="stmt"/>
        <line num="200" count="12" type="stmt"/>
        <line num="201" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="202" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="203" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="204" count="3" type="stmt"/>
        <line num="205" count="3" type="stmt"/>
        <line num="206" count="3" type="stmt"/>
        <line num="207" count="12" type="stmt"/>
        <line num="208" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="209" count="1" type="stmt"/>
        <line num="210" count="1" type="stmt"/>
        <line num="211" count="1" type="stmt"/>
        <line num="212" count="12" type="stmt"/>
        <line num="213" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="214" count="1" type="stmt"/>
        <line num="215" count="1" type="stmt"/>
        <line num="216" count="1" type="stmt"/>
        <line num="217" count="12" type="stmt"/>
        <line num="218" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="219" count="1" type="stmt"/>
        <line num="220" count="1" type="stmt"/>
        <line num="221" count="1" type="stmt"/>
        <line num="222" count="12" type="stmt"/>
        <line num="223" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="224" count="1" type="stmt"/>
        <line num="225" count="1" type="stmt"/>
        <line num="226" count="1" type="stmt"/>
        <line num="227" count="12" type="stmt"/>
        <line num="228" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="229" count="2" type="stmt"/>
        <line num="230" count="2" type="stmt"/>
        <line num="231" count="2" type="stmt"/>
        <line num="232" count="12" type="stmt"/>
        <line num="233" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="234" count="1" type="stmt"/>
        <line num="235" count="1" type="stmt"/>
        <line num="236" count="1" type="stmt"/>
        <line num="237" count="12" type="stmt"/>
        <line num="238" count="12" type="stmt"/>
        <line num="239" count="12" type="stmt"/>
        <line num="240" count="12" type="stmt"/>
        <line num="241" count="12" type="stmt"/>
        <line num="242" count="12" type="stmt"/>
        <line num="243" count="12" type="stmt"/>
        <line num="244" count="12" type="stmt"/>
        <line num="245" count="12" type="stmt"/>
        <line num="246" count="16" type="cond" truecount="1" falsecount="0"/>
        <line num="247" count="1" type="stmt"/>
        <line num="248" count="1" type="stmt"/>
        <line num="249" count="1" type="stmt"/>
        <line num="250" count="16" type="cond" truecount="1" falsecount="0"/>
        <line num="251" count="1" type="stmt"/>
        <line num="252" count="1" type="stmt"/>
        <line num="253" count="1" type="stmt"/>
        <line num="254" count="16" type="cond" truecount="0" falsecount="1"/>
        <line num="255" count="0" type="stmt"/>
        <line num="256" count="0" type="stmt"/>
        <line num="257" count="0" type="stmt"/>
        <line num="258" count="16" type="cond" truecount="0" falsecount="1"/>
        <line num="259" count="0" type="stmt"/>
        <line num="260" count="0" type="stmt"/>
        <line num="261" count="0" type="stmt"/>
        <line num="262" count="0" type="stmt"/>
        <line num="263" count="0" type="stmt"/>
        <line num="264" count="0" type="stmt"/>
        <line num="265" count="16" type="stmt"/>
        <line num="266" count="26" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="AmericanProvider.ts" path="/home/user/repo/example-project/src/Repository/WeatherProvider/AmericanProvider.ts">
        <metrics statements="141" coveredstatements="141" conditionals="24" coveredconditionals="19" methods="5" coveredmethods="5"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="stmt"/>
        <line num="13" count="1" type="stmt"/>
        <line num="14" count="1" type="stmt"/>
        <line num="15" count="1" type="stmt"/>
        <line num="16" count="1" type="stmt"/>
        <line num="17" count="1" type="stmt"/>
        <line num="18" count="1" type="stmt"/>
        <line num="19" count="1" type="stmt"/>
        <line num="20" count="1" type="stmt"/>
        <line num="21" count="1" type="stmt"/>
        <line num="22" count="1" type="stmt"/>
        <line num="23" count="1" type="stmt"/>
        <line num="24" count="1" type="stmt"/>
        <line num="25" count="1" type="stmt"/>
        <line num="26" count="1" type="stmt"/>
        <line num="27" count="1" type="stmt"/>
        <line num="28" count="1" type="stmt"/>
        <line num="29" count="1" type="stmt"/>
        <line num="30" count="1" type="stmt"/>
        <line num="31" count="1" type="stmt"/>
        <line num="32" count="1" type="stmt"/>
        <line num="33" count="1" type="stmt"/>
        <line num="34" count="1" type="stmt"/>
        <line num="35" count="1" type="stmt"/>
        <line num="36" count="1" type="stmt"/>
        <line num="37" count="1" type="stmt"/>
        <line num="38" count="1" type="stmt"/>
        <line num="39" count="1" type="stmt"/>
        <line num="40" count="1" type="stmt"/>
        <line num="41" count="1" type="stmt"/>
        <line num="42" count="1" type="stmt"/>
        <line num="43" count="1" type="stmt"/>
        <line num="44" count="1" type="stmt"/>
        <line num="45" count="1" type="stmt"/>
        <line num="46" count="1" type="stmt"/>
        <line num="47" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="48" count="12" type="stmt"/>
        <line num="49" count="12" type="stmt"/>
        <line num="50" count="12" type="stmt"/>
        <line num="51" count="12" type="stmt"/>
        <line num="52" count="12" type="stmt"/>
        <line num="53" count="12" type="stmt"/>
        <line num="54" count="12" type="cond" truecount="2" falsecount="0"/>
        <line num="55" count="12" type="cond" truecount="1" falsecount="0"/>
        <line num="56" count="12" type="stmt"/>
        <line num="57" count="1" type="stmt"/>
        <line num="58" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="59" count="1" type="stmt"/>
        <line num="60" count="1" type="stmt"/>
        <line num="61" count="1" type="stmt"/>
        <line num="62" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="63" count="8" type="cond" truecount="2" falsecount="0"/>
        <line num="64" count="8" type="cond" truecount="2" falsecount="0"/>
        <line num="65" count="6" type="stmt"/>
        <line num="66" count="6" type="stmt"/>
        <line num="67" count="6" type="stmt"/>
        <line num="68" count="6" type="stmt"/>
        <line num="69" count="6" type="stmt"/>
        <line num="70" count="6" type="stmt"/>
        <line num="71" count="6" type="stmt"/>
        <line num="72" count="6" type="stmt"/>
        <line num="73" count="6" type="stmt"/>
        <line num="74" count="6" type="stmt"/>
        <line num="75" count="6" type="stmt"/>
        <line num="76" count="6" type="stmt"/>
        <line num="77" count="6" type="stmt"/>
        <line num="78" count="6" type="stmt"/>
        <line num="79" count="6" type="stmt"/>
        <line num="80" count="6" type="stmt"/>
        <line num="81" count="6" type="stmt"/>
        <line num="82" count="6" type="stmt"/>
        <line num="83" count="6" type="stmt"/>
        <line num="84" count="6" type="stmt"/>
        <line num="85" count="6" type="stmt"/>
        <line num="86" count="6" type="stmt"/>
        <line num="87" count="6" type="stmt"/>
        <line num="88" count="6" type="stmt"/>
        <line num="89" count="6" type="stmt"/>
        <line num="90" count="6" type="stmt"/>
        <line num="91" count="6" type="stmt"/>
        <line num="92" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="93" count="8" type="stmt"/>
        <line num="94" count="8" type="cond" truecount="1" falsecount="1"/>
        <line num="95" count="8" type="stmt"/>
        <line num="96" count="8" type="stmt"/>
        <line num="97" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="98" count="8" type="stmt"/>
        <line num="99" count="8" type="stmt"/>
        <line num="100" count="8" type="stmt"/>
        <line num="101" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="102" count="1" type="stmt"/>
        <line num="103" count="1" type="stmt"/>
        <line num="104" count="1" type="stmt"/>
        <line num="105" count="1" type="stmt"/>
        <line num="106" count="1" type="stmt"/>
        <line num="107" count="1" type="stmt"/>
        <line num="108" count="1" type="stmt"/>
        <line num="109" count="1" type="stmt"/>
        <line num="110" count="1" type="stmt"/>
        <line num="111" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="112" count="1" type="stmt"/>
        <line num="113" count="1" type="stmt"/>
        <line num="114" count="1" type="stmt"/>
        <line num="115" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="116" count="1" type="stmt"/>
        <line num="117" count="1" type="stmt"/>
        <line num="118" count="1" type="stmt"/>
        <line num="119" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="120" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="121" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="122" count="1" type="stmt"/>
        <line num="123" count="1" type="stmt"/>
        <line num="124" count="1" type="stmt"/>
        <line num="125" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="126" count="1" type="stmt"/>
        <line num="127" count="1" type="stmt"/>
        <line num="128" count="1" type="stmt"/>
        <line num="129" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="130" count="1" type="stmt"/>
        <line num="131" count="1" type="stmt"/>
        <line num="132" count="1" type="stmt"/>
        <line num="133" count="1" type="stmt"/>
        <line num="134" count="1" type="stmt"/>
        <line num="135" count="1" type="stmt"/>
        <line num="136" count="8" type="stmt"/>
        <line num="137" count="1" type="stmt"/>
        <line num="138" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="139" count="1" type="stmt"/>
        <line num="140" count="1" type="stmt"/>
        <line num="141" count="12" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="AustralianProvider.ts" path="/home/user/repo/example-project/src/Repository/WeatherProvider/AustralianProvider.ts">
        <metrics statements="149" coveredstatements="147" conditionals="24" coveredconditionals="19" methods="5" coveredmethods="4"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="stmt"/>
        <line num="13" count="1" type="stmt"/>
        <line num="14" count="1" type="stmt"/>
        <line num="15" count="1" type="stmt"/>
        <line num="16" count="1" type="stmt"/>
        <line num="17" count="1" type="stmt"/>
        <line num="18" count="1" type="stmt"/>
        <line num="19" count="1" type="stmt"/>
        <line num="20" count="1" type="stmt"/>
        <line num="21" count="1" type="stmt"/>
        <line num="22" count="1" type="stmt"/>
        <line num="23" count="1" type="stmt"/>
        <line num="24" count="1" type="stmt"/>
        <line num="25" count="1" type="stmt"/>
        <line num="26" count="1" type="stmt"/>
        <line num="27" count="1" type="stmt"/>
        <line num="28" count="1" type="stmt"/>
        <line num="29" count="1" type="stmt"/>
        <line num="30" count="1" type="stmt"/>
        <line num="31" count="1" type="stmt"/>
        <line num="32" count="1" type="stmt"/>
        <line num="33" count="1" type="stmt"/>
        <line num="34" count="1" type="stmt"/>
        <line num="35" count="1" type="stmt"/>
        <line num="36" count="1" type="stmt"/>
        <line num="37" count="1" type="stmt"/>
        <line num="38" count="1" type="stmt"/>
        <line num="39" count="1" type="stmt"/>
        <line num="40" count="1" type="stmt"/>
        <line num="41" count="1" type="stmt"/>
        <line num="42" count="1" type="stmt"/>
        <line num="43" count="1" type="stmt"/>
        <line num="44" count="1" type="stmt"/>
        <line num="45" count="1" type="stmt"/>
        <line num="46" count="1" type="stmt"/>
        <line num="47" count="1" type="stmt"/>
        <line num="48" count="1" type="stmt"/>
        <line num="49" count="1" type="stmt"/>
        <line num="50" count="1" type="stmt"/>
        <line num="51" count="1" type="stmt"/>
        <line num="52" count="1" type="stmt"/>
        <line num="53" count="1" type="stmt"/>
        <line num="54" count="1" type="stmt"/>
        <line num="55" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="56" count="11" type="stmt"/>
        <line num="57" count="11" type="stmt"/>
        <line num="58" count="11" type="stmt"/>
        <line num="59" count="11" type="stmt"/>
        <line num="60" count="11" type="stmt"/>
        <line num="61" count="11" type="stmt"/>
        <line num="62" count="11" type="cond" truecount="2" falsecount="0"/>
        <line num="63" count="11" type="cond" truecount="1" falsecount="0"/>
        <line num="64" count="11" type="stmt"/>
        <line num="65" count="1" type="stmt"/>
        <line num="66" count="1" type="stmt"/>
        <line num="67" count="0" type="stmt"/>
        <line num="68" count="0" type="stmt"/>
        <line num="69" count="1" type="stmt"/>
        <line num="70" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="71" count="8" type="cond" truecount="2" falsecount="0"/>
        <line num="72" count="8" type="cond" truecount="2" falsecount="0"/>
        <line num="73" count="6" type="stmt"/>
        <line num="74" count="6" type="stmt"/>
        <line num="75" count="6" type="stmt"/>
        <line num="76" count="6" type="stmt"/>
        <line num="77" count="6" type="stmt"/>
        <line num="78" count="6" type="stmt"/>
        <line num="79" count="6" type="stmt"/>
        <line num="80" count="6" type="stmt"/>
        <line num="81" count="6" type="stmt"/>
        <line num="82" count="6" type="stmt"/>
        <line num="83" count="6" type="stmt"/>
        <line num="84" count="6" type="stmt"/>
        <line num="85" count="6" type="stmt"/>
        <line num="86" count="6" type="stmt"/>
        <line num="87" count="6" type="stmt"/>
        <line num="88" count="6" type="stmt"/>
        <line num="89" count="6" type="stmt"/>
        <line num="90" count="6" type="stmt"/>
        <line num="91" count="6" type="stmt"/>
        <line num="92" count="6" type="stmt"/>
        <line num="93" count="6" type="stmt"/>
        <line num="94" count="6" type="stmt"/>
        <line num="95" count="6" type="stmt"/>
        <line num="96" count="6" type="stmt"/>
        <line num="97" count="6" type="stmt"/>
        <line num="98" count="6" type="stmt"/>
        <line num="99" count="6" type="stmt"/>
        <line num="100" count="6" type="stmt"/>
        <line num="101" count="6" type="stmt"/>
        <line num="102" count="6" type="stmt"/>
        <line num="103" count="6" type="stmt"/>
        <line num="104" count="6" type="stmt"/>
        <line num="105" count="6" type="stmt"/>
        <line num="106" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="107" count="8" type="stmt"/>
        <line num="108" count="8" type="cond" truecount="1" falsecount="1"/>
        <line num="109" count="8" type="stmt"/>
        <line num="110" count="8" type="stmt"/>
        <line num="111" count="8" type="cond" truecount="1" falsecount="1"/>
        <line num="112" count="8" type="stmt"/>
        <line num="113" count="8" type="stmt"/>
        <line num="114" count="8" type="stmt"/>
        <line num="115" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="116" count="1" type="stmt"/>
        <line num="117" count="1" type="stmt"/>
        <line num="118" count="1" type="stmt"/>
        <line num="119" count="1" type="stmt"/>
        <line num="120" count="1" type="stmt"/>
        <line num="121" count="1" type="stmt"/>
        <line num="122" count="1" type="stmt"/>
        <line num="123" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="124" count="1" type="stmt"/>
        <line num="125" count="1" type="stmt"/>
        <line num="126" count="1" type="stmt"/>
        <line num="127" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="128" count="1" type="stmt"/>
        <line num="129" count="1" type="stmt"/>
        <line num="130" count="1" type="stmt"/>
        <line num="131" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="132" count="8" type="cond" truecount="0" falsecount="1"/>
        <line num="133" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="134" count="1" type="stmt"/>
        <line num="135" count="1" type="stmt"/>
        <line num="136" count="1" type="stmt"/>
        <line num="137" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="138" count="1" type="stmt"/>
        <line num="139" count="1" type="stmt"/>
        <line num="140" count="1" type="stmt"/>
        <line num="141" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="142" count="1" type="stmt"/>
        <line num="143" count="1" type="stmt"/>
        <line num="144" count="8" type="stmt"/>
        <line num="145" count="1" type="stmt"/>
        <line num="146" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="147" count="1" type="stmt"/>
        <line num="148" count="1" type="stmt"/>
        <line num="149" count="11" type="cond" truecount="1" falsecount="0"/>
      </file>
    </package>
    <package name="Repository.Security">
      <metrics statements="197" coveredstatements="190" conditionals="19" coveredconditionals="17" methods="15" coveredmethods="13"/>
      <file name="KVv2.ts" path="/home/user/repo/example-project/src/Repository/Security/KVv2.ts">
        <metrics statements="128" coveredstatements="124" conditionals="11" coveredconditionals="10" methods="6" coveredmethods="6"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="stmt"/>
        <line num="13" count="1" type="stmt"/>
        <line num="14" count="1" type="stmt"/>
        <line num="15" count="1" type="stmt"/>
        <line num="16" count="1" type="stmt"/>
        <line num="17" count="1" type="stmt"/>
        <line num="18" count="1" type="stmt"/>
        <line num="19" count="1" type="stmt"/>
        <line num="20" count="1" type="stmt"/>
        <line num="21" count="1" type="stmt"/>
        <line num="22" count="1" type="stmt"/>
        <line num="23" count="1" type="stmt"/>
        <line num="24" count="1" type="stmt"/>
        <line num="25" count="1" type="stmt"/>
        <line num="26" count="1" type="stmt"/>
        <line num="27" count="1" type="stmt"/>
        <line num="28" count="1" type="stmt"/>
        <line num="29" count="1" type="stmt"/>
        <line num="30" count="1" type="stmt"/>
        <line num="31" count="1" type="stmt"/>
        <line num="32" count="1" type="stmt"/>
        <line num="33" count="1" type="stmt"/>
        <line num="34" count="1" type="stmt"/>
        <line num="35" count="1" type="stmt"/>
        <line num="36" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="37" count="5" type="stmt"/>
        <line num="38" count="5" type="stmt"/>
        <line num="39" count="5" type="stmt"/>
        <line num="40" count="5" type="cond" truecount="1" falsecount="0"/>
        <line num="41" count="5" type="stmt"/>
        <line num="42" count="1" type="stmt"/>
        <line num="43" count="1" type="stmt"/>
        <line num="44" count="1" type="stmt"/>
        <line num="45" count="1" type="stmt"/>
        <line num="46" count="1" type="stmt"/>
        <line num="47" count="1" type="stmt"/>
        <line num="48" count="1" type="stmt"/>
        <line num="49" count="1" type="stmt"/>
        <line num="50" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="51" count="10" type="stmt"/>
        <line num="52" count="10" type="stmt"/>
        <line num="53" count="10" type="stmt"/>
        <line num="54" count="10" type="stmt"/>
        <line num="55" count="10" type="stmt"/>
        <line num="56" count="10" type="stmt"/>
        <line num="57" count="10" type="stmt"/>
        <line num="58" count="10" type="stmt"/>
        <line num="59" count="10" type="cond" truecount="2" falsecount="0"/>
        <line num="60" count="6" type="stmt"/>
        <line num="61" count="6" type="stmt"/>
        <line num="62" count="8" type="cond" truecount="1" falsecount="0"/>
        <line num="63" count="4" type="stmt"/>
        <line num="64" count="4" type="stmt"/>
        <line num="65" count="4" type="stmt"/>
        <line num="66" count="4" type="stmt"/>
        <line num="67" count="4" type="cond" truecount="0" falsecount="1"/>
        <line num="68" count="0" type="stmt"/>
        <line num="69" count="0" type="stmt"/>
        <line num="70" count="0" type="stmt"/>
        <line num="71" count="0" type="stmt"/>
        <line num="72" count="10" type="stmt"/>
        <line num="73" count="1" type="stmt"/>
        <line num="74" count="1" type="stmt"/>
        <line num="75" count="1" type="stmt"/>
        <line num="76" count="1" type="stmt"/>
        <line num="77" count="1" type="stmt"/>
        <line num="78" count="1" type="stmt"/>
        <line num="79" count="1" type="stmt"/>
        <line num="80" count="1" type="stmt"/>
        <line num="81" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="82" count="3" type="stmt"/>
        <line num="83" count="3" type="stmt"/>
        <line num="84" count="3" type="stmt"/>
        <line num="85" count="3" type="stmt"/>
        <line num="86" count="3" type="stmt"/>
        <line num="87" count="3" type="stmt"/>
        <line num="88" count="3" type="stmt"/>
        <line num="89" count="3" type="stmt"/>
        <line num="90" count="3" type="stmt"/>
        <line num="91" count="3" type="stmt"/>
        <line num="92" count="3" type="stmt"/>
        <line num="93" count="1" type="stmt"/>
        <line num="94" count="1" type="stmt"/>
        <line num="95" count="1" type="stmt"/>
        <line num="96" count="1" type="stmt"/>
        <line num="97" count="1" type="stmt"/>
        <line num="98" count="1" type="stmt"/>
        <line num="99" count="1" type="stmt"/>
        <line num="100" count="1" type="stmt"/>
        <line num="101" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="102" count="3" type="stmt"/>
        <line num="103" count="3" type="stmt"/>
        <line num="104" count="3" type="stmt"/>
        <line num="105" count="3" type="stmt"/>
        <line num="106" count="3" type="stmt"/>
        <line num="107" count="3" type="stmt"/>
        <line num="108" count="3" type="stmt"/>
        <line num="109" count="3" type="stmt"/>
        <line num="110" count="3" type="stmt"/>
        <line num="111" count="3" type="stmt"/>
        <line num="112" count="3" type="stmt"/>
        <line num="113" count="3" type="stmt"/>
        <line num="114" count="3" type="stmt"/>
        <line num="115" count="3" type="stmt"/>
        <line num="116" count="1" type="stmt"/>
        <line num="117" count="1" type="stmt"/>
        <line num="118" count="1" type="stmt"/>
        <line num="119" count="1" type="stmt"/>
        <line num="120" count="1" type="stmt"/>
        <line num="121" count="1" type="stmt"/>
        <line num="122" count="1" type="stmt"/>
        <line num="123" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="124" count="3" type="stmt"/>
        <line num="125" count="3" type="stmt"/>
        <line num="126" count="3" type="stmt"/>
        <line num="127" count="3" type="stmt"/>
        <line num="128" count="5" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="SecurityClient.ts" path="/home/user/repo/example-project/src/Repository/Security/SecurityClient.ts">
        <metrics statements="69" coveredstatements="66" conditionals="8" coveredconditionals="7" methods="9" coveredmethods="7"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="stmt"/>
        <line num="13" count="1" type="stmt"/>
        <line num="14" count="1" type="stmt"/>
        <line num="15" count="1" type="stmt"/>
        <line num="16" count="1" type="stmt"/>
        <line num="17" count="0" type="stmt"/>
        <line num="18" count="1" type="stmt"/>
        <line num="19" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="20" count="2" type="stmt"/>
        <line num="21" count="2" type="stmt"/>
        <line num="22" count="2" type="stmt"/>
        <line num="23" count="2" type="stmt"/>
        <line num="24" count="2" type="stmt"/>
        <line num="25" count="2" type="stmt"/>
        <line num="26" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="27" count="2" type="cond" truecount="0" falsecount="1"/>
        <line num="28" count="2" type="stmt"/>
        <line num="29" count="2" type="stmt"/>
        <line num="30" count="2" type="stmt"/>
        <line num="31" count="2" type="stmt"/>
        <line num="32" count="2" type="stmt"/>
        <line num="33" count="2" type="stmt"/>
        <line num="34" count="2" type="stmt"/>
        <line num="35" count="2" type="stmt"/>
        <line num="36" count="2" type="stmt"/>
        <line num="37" count="2" type="stmt"/>
        <line num="38" count="2" type="stmt"/>
        <line num="39" count="2" type="stmt"/>
        <line num="40" count="2" type="stmt"/>
        <line num="41" count="2" type="stmt"/>
        <line num="42" count="2" type="stmt"/>
        <line num="43" count="2" type="stmt"/>
        <line num="44" count="2" type="stmt"/>
        <line num="45" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="46" count="2" type="stmt"/>
        <line num="47" count="2" type="stmt"/>
        <line num="48" count="2" type="stmt"/>
        <line num="49" count="2" type="stmt"/>
        <line num="50" count="2" type="stmt"/>
        <line num="51" count="2" type="stmt"/>
        <line num="52" count="2" type="stmt"/>
        <line num="53" count="2" type="stmt"/>
        <line num="54" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="55" count="2" type="stmt"/>
        <line num="56" count="2" type="stmt"/>
        <line num="57" count="2" type="stmt"/>
        <line num="58" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="59" count="1" type="stmt"/>
        <line num="60" count="1" type="stmt"/>
        <line num="61" count="2" type="stmt"/>
        <line num="62" count="2" type="stmt"/>
        <line num="63" count="0" type="stmt"/>
        <line num="64" count="0" type="stmt"/>
        <line num="65" count="2" type="stmt"/>
        <line num="66" count="2" type="cond" truecount="1" falsecount="0"/>
        <line num="67" count="1" type="stmt"/>
        <line num="68" count="1" type="stmt"/>
        <line num="69" count="2" type="cond" truecount="1" falsecount="0"/>
      </file>
    </package>
    <package name="Repository.Security.Auth">
      <metrics statements="92" coveredstatements="75" conditionals="12" coveredconditionals="12" methods="9" coveredmethods="6"/>
      <file name="Token.ts" path="/home/user/repo/example-project/src/Repository/Security/Auth/Token.ts">
        <metrics statements="32" coveredstatements="32" conditionals="6" coveredconditionals="6" methods="3" coveredmethods="3"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="13" count="4" type="stmt"/>
        <line num="14" count="4" type="stmt"/>
        <line num="15" count="4" type="stmt"/>
        <line num="16" count="4" type="stmt"/>
        <line num="17" count="4" type="cond" truecount="2" falsecount="0"/>
        <line num="18" count="4" type="cond" truecount="1" falsecount="0"/>
        <line num="19" count="4" type="stmt"/>
        <line num="20" count="1" type="stmt"/>
        <line num="21" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="22" count="2" type="stmt"/>
        <line num="23" count="2" type="stmt"/>
        <line num="24" count="2" type="stmt"/>
        <line num="25" count="2" type="stmt"/>
        <line num="26" count="2" type="stmt"/>
        <line num="27" count="2" type="stmt"/>
        <line num="28" count="2" type="stmt"/>
        <line num="29" count="2" type="stmt"/>
        <line num="30" count="2" type="stmt"/>
        <line num="31" count="2" type="stmt"/>
        <line num="32" count="4" type="cond" truecount="1" falsecount="0"/>
      </file>
      <file name="LDAP.ts" path="/home/user/repo/example-project/src/Repository/Security/Auth/LDAP.ts">
        <metrics statements="30" coveredstatements="13" conditionals="0" coveredconditionals="0" methods="3" coveredmethods="0"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="stmt"/>
        <line num="12" count="0" type="stmt"/>
        <line num="13" count="0" type="stmt"/>
        <line num="14" count="0" type="stmt"/>
        <line num="15" count="0" type="stmt"/>
        <line num="16" count="0" type="stmt"/>
        <line num="17" count="0" type="stmt"/>
        <line num="18" count="0" type="stmt"/>
        <line num="19" count="1" type="stmt"/>
        <line num="20" count="1" type="stmt"/>
        <line num="21" count="0" type="stmt"/>
        <line num="22" count="0" type="stmt"/>
        <line num="23" count="0" type="stmt"/>
        <line num="24" count="0" type="stmt"/>
        <line num="25" count="0" type="stmt"/>
        <line num="26" count="0" type="stmt"/>
        <line num="27" count="0" type="stmt"/>
        <line num="28" count="0" type="stmt"/>
        <line num="29" count="0" type="stmt"/>
        <line num="30" count="0" type="stmt"/>
      </file>
      <file name="UserPass.ts" path="/home/user/repo/example-project/src/Repository/Security/Auth/UserPass.ts">
        <metrics statements="30" coveredstatements="30" conditionals="6" coveredconditionals="6" methods="3" coveredmethods="3"/>
        <line num="1" count="1" type="stmt"/>
        <line num="2" count="1" type="stmt"/>
        <line num="3" count="1" type="stmt"/>
        <line num="4" count="1" type="stmt"/>
        <line num="5" count="1" type="stmt"/>
        <line num="6" count="1" type="stmt"/>
        <line num="7" count="1" type="stmt"/>
        <line num="8" count="1" type="stmt"/>
        <line num="9" count="1" type="stmt"/>
        <line num="10" count="1" type="stmt"/>
        <line num="11" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="12" count="4" type="stmt"/>
        <line num="13" count="4" type="stmt"/>
        <line num="14" count="4" type="stmt"/>
        <line num="15" count="4" type="stmt"/>
        <line num="16" count="4" type="cond" truecount="2" falsecount="0"/>
        <line num="17" count="4" type="cond" truecount="1" falsecount="0"/>
        <line num="18" count="4" type="stmt"/>
        <line num="19" count="1" type="stmt"/>
        <line num="20" count="1" type="cond" truecount="1" falsecount="0"/>
        <line num="21" count="2" type="stmt"/>
        <line num="22" count="2" type="stmt"/>
        <line num="23" count="2" type="stmt"/>
        <line num="24" count="2" type="stmt"/>
        <line num="25" count="2" type="stmt"/>
        <line num="26" count="2" type="stmt"/>
        <line num="27" count="2" type="stmt"/>
        <line num="28" count="2" type="stmt"/>
        <line num="29" count="2" type="stmt"/>
        <line num="30" count="4" type="cond" truecount="1" falsecount="0"/>
      </file>
    </package>
  </project>
</coverage>`)

	parsedCoverage, err := (clover.Parser{}).Parse(content)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if parsedCoverage.TotalCoveredStatements != 1553 {
		t.Errorf("expecting TotalCoveredStatements to be 1553, instead got %d", parsedCoverage.TotalCoveredStatements)
	}

	if parsedCoverage.NumberOfFiles != 23 {
		t.Errorf("expecting NumberOfFiles to be 23, instead got %d", parsedCoverage.NumberOfFiles)
	}

	if parsedCoverage.TotalStatements != 1678 {
		t.Errorf("expecting TotalStatements to be 1678, instead got %d", parsedCoverage.TotalStatements)
	}
}

func TestParser_Parse2(t *testing.T) {
	_, err := (clover.Parser{}).Parse(strings.NewReader("{\\INVALID!!!}"))
	if err == nil {
		t.Error("expecting an error, got nil instead")
	}
}
