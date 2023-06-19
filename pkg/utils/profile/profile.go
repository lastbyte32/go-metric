/*
Copyright 2021 The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// To use profile in a go program over http
// import this package and call ProfileIfEnabled()
// in your main function.
// Please set the environment variable PPROF_ENABLE=true to enable/disable it runtime.
// To customize host and port you can set PPROF_HOST and PPROF_PORT environment variables.
//	$ PPROF_ENABLE=true PPROF_HOST=localhost PPROF_PORT=6060 go run myprogram.go

package profile

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func ProfileIfEnabled() {
	enablePprof := os.Getenv("PPROF_ENABLED")
	if enablePprof != "true" {
		return
	}
	pprofPort := os.Getenv("PPROF_PORT")
	if pprofPort == "" {
		pprofPort = "6060"
	}

	go func() {
		fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", pprofPort), nil))
	}()

}
