/*
 * g2z - Zabbix module adapter for Go
 * Copyright (C) 2015 - Ryan Armstrong <ryan@cavaliercoder.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 */

/*
Package g2z provides Go bindings for creating a native Zabbix module.

An example module is provided with the g2z sources (in dummy/dummy.go) which
is a Go implementation of the dummy C module published by Zabbix.

To build a shared library compatible with Zabbix v2.2.0+, this project takes
advantage of the c-shared build mode introduced in Go v1.5.0.

	package main

	import (
		"github.com/cavaliercoder/g2z"
		"strings"
	)

	func main() {
	    panic("THIS_SHOULD_NEVER_HAPPEN")
	}

	func init() {
	    g2z.RegisterStringItem("go.echo", "Hello world!", Echo)
	}

	func Echo(request *g2z.AgentRequest) (string, error) {
	    return strings.Join(request.Params, " "), nil
	}

Compile your shared library with:

	$ go build -buildmode=c-shared

Load your `.so` module file into Zabbix as per the Zabbix manual

	https://www.zabbix.com/documentation/2.2/manual/config/items/loadablemodules#configuration_parameters

Check that your module items are loaded with:

	$ zabbix_agentd -t go.echo[hello,world]

For more thorough performance testing, use zabbix_agent_bench

	https://github.com/cavaliercoder/zabbix_agent_bench

*/
package g2z
