// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This is an example API server that just statically serves a kernel,
// initrd and commandline. This is effectively the same as Pixiecore
// in static mode, only it's talking to an API server instead.
//
// This is not production-quality code. The focus is on being short
// and sweet, not robust and correct.
package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	//	"os"
	"path/filepath"
	"strconv"
)

var (
	port = flag.Int("port", 4242, "Port to listen on")
)

func main() {
	flag.Parse()
	http.HandleFunc("/v1/boot/", api)
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

func api(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving boot config for %s", filepath.Base(r.URL.Path))
	ipxeURL := os.Getenv("IPXEURL")
	resp := struct {
		K string `json:"ipxe-script"`
	}{
		K: ipxeURL,
	}
	/* resp := struct {
		K string   `json:"kernel"`
		I []string `json:"initrd"`
	}{
		K: "http://tinycorelinux.net/11.x/x86/release/distribution_files/vmlinuz64",
		I: []string{
			"http://tinycorelinux.net/11.x/x86/release/distribution_files/rootfs.gz",
			"http://tinycorelinux.net/11.x/x86/release/distribution_files/modules64.gz",
		}, */
	/* K: "http://mirrors.edge.kernel.org/ubuntu/dists/bionic/main/installer-amd64/current/images/netboot/ubuntu-installer/amd64/linux",
		I: []string{
			"http://mirrors.edge.kernel.org/ubuntu/dists/bionic/main/installer-amd64/current/images/netboot/ubuntu-installer/amd64/initrd.gz",
		},
	} */

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		panic(err)
	}
}
