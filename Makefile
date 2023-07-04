SERVER_PORT=8182
ADDRESS="localhost:$(SERVER_PORT)"
TEMP_FILE=tempfile.json

BUILD_VERSION = $$(git describe --tags)
BUILD_DATE = $$(date)
BUILD_COMMIT = $$(git rev-parse HEAD)

trace_fn = go tool pprof -http=\":9090\" -seconds=120 http://localhost:$(1)/debug/pprof/$(2)
trace_and_save_fn = curl -sK -v http://localhost:$(1)/debug/pprof/$(2)?seconds=300 > profiles/$(3) && go tool pprof -http=":9090" profiles/$(3)

coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

build-server:
	cd cmd/server && go build -buildvcs=false -v -ldflags="-X 'main.buildVersion=$(BUILD_VERSION)' -X 'main.buildDate=$(BUILD_DATE)' -X main.buildCommit=$(BUILD_COMMIT)" -o server

build-agent:
	cd cmd/agent && go build -buildvcs=false -v -ldflags="-X 'main.buildVersion=$(BUILD_VERSION)' -X 'main.buildDate=$(BUILD_DATE)' -X main.buildCommit=$(BUILD_COMMIT)" -o agent
clean:
	rm -f cmd/server/server && rm -f cmd/agent/agent && rm -f $(TEMP_FILE)

build: clean build-server build-agent
iter1: build
	devopstest -test.v -test.run=^TestIteration1$$ \
            -agent-binary-path=cmd/agent/agent
iter2: build
	devopstest -test.v -test.run=^TestIteration2[b]*$ \
            -source-path=. \
            -binary-path=cmd/server/server
iter3: build
	devopstest -test.v -test.run=^TestIteration3[b]*$ \
            -source-path=. \
            -binary-path=cmd/server/server
iter4: build
	devopstest -test.v -test.run=^TestIteration4$ \
            -source-path=. \
            -binary-path=cmd/server/server \
            -agent-binary-path=cmd/agent/agent
iter5: build
	devopstest -test.v -test.run=^TestIteration5$ \
            -source-path=. \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -server-port=$(SERVER_PORT)
iter6: build
	devopstest -test.v -test.run=^TestIteration6$ \
            -source-path=. \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -server-port=$(SERVER_PORT) \
            -database-dsn='postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable' \
            -file-storage-path=$(TEMP_FILE)
iter7: build
	devopstest -test.v -test.run=^TestIteration7$ \
            -source-path=. \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -server-port=$(SERVER_PORT) \
            -database-dsn='postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable' \
            -file-storage-path=$(TEMP_FILE)
iter8: build
	devopstest -test.v -test.run=^TestIteration8 \
            -source-path=. \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -server-port=$(SERVER_PORT) \
            -database-dsn='postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable' \
            -file-storage-path=$(TEMP_FILE)
iter9: build
	devopstest -test.v -test.run=^TestIteration9$ \
            -source-path=. \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -server-port=$(SERVER_PORT) \
            -file-storage-path=$(TEMP_FILE) \
            -database-dsn='postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable' \
            -key="$(TEMP_FILE)"
iter10: build
	devopstest -test.v -test.run=^TestIteration10[b]*$ \
            -source-path=. \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -server-port=$(SERVER_PORT) \
            -database-dsn='postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable' \
            -key="$(TEMP_FILE)"
iter11: build
	devopstest -test.v -test.run=^TestIteration11$ \
            -source-path=. \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -server-port=$(SERVER_PORT) \
            -database-dsn='postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable' \
            -key="$(TEMP_FILE)"
iter12: build
	devopstest -test.v -test.run=^TestIteration12$ \
            -source-path=. \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -server-port=$(SERVER_PORT) \
            -database-dsn='postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable' \
            -key="$(TEMP_FILE)"
iter13: build
	devopstest -test.v -test.run=^TestIteration13$ \
            -source-path=.
iter14: build
	devopstest -test.v -test.run=^TestIteration14$ \
            -source-path=. \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -server-port=$(SERVER_PORT) \
            -file-storage-path=$(TEMP_FILE) \
            -database-dsn='postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable' \
            -key="$(TEMP_FILE)"
iter14race: build
	go test -v -race ./...