BIN=binghuang
BIN_SRC=cmd/main.go
DIST_DIR=.dist
COVERAGE_FILE=.coverage
LOG_FILE=*.log*

##################################
#######       Setup       ########
##################################
.PHONY: ensure tidy


tidy:
	@go mod tidy

ensure: tidy
	@go mod download

##################################
#######        Tool       ########
##################################
.PHONY: fmt lint build clean api rpc

fmt:
	@golangci-lint run --fix

lint:
	@golangci-lint run ./...

build:
	@go build -o ${DIST_DIR}/${BIN} ${BIN_SRC}

clean:
	@git clean -fdx ${DIST_DIR} ${COVERAGE_FILE} ${LOG_FILE}


API_PATH = "./service/api/binghuang.api"
API_DIR = "./service/api"
API_TEMPLATE = "./service/api/.goctl"
SWAGGER_GEN_DIR = "./gen/swagger"
ifeq ($(shell uname),Darwin)
    CMD = gsed
else
    CMD = sed
endif

api:
	go install github.com/zeromicro/go-zero/tools/goctl@latest
	go install github.com/zeromicro/goctl-swagger@latest
	goctl api format --dir ${API_DIR} --declare
	goctl api go --api ${API_PATH} --dir ${API_DIR} --home ${API_TEMPLATE}
	goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api ${API_PATH} -dir ${SWAGGER_GEN_DIR}
	find * -type f -exec ${CMD} -i 's/^\s\+$$//g' {} +

##################################
#######      Server       ########
##################################
start_http:
	@go run ${BIN_SRC} start_http --env $(ENV)

start_worker:
	@go run ${BIN_SRC} start_worker --env $(ENV)

start_cron:
	@go run ${BIN_SRC} start_cron -- env $(ENV)

##################################
#######     Coverage      ########
##################################
.PHONY: test coverage coverage_html

TEST_FLAGS = -v -race -failfast -covermode=atomic
MINIMAL_COVERAGE = 46

test:
	@CGO_ENABLED=1 ROOT=${PWD} ENV=testing go test ${TEST_FLAGS} -coverprofile=${COVERAGE_FILE} -cover -timeout=40s `go list ./service/...`

coverage: test
	@go tool cover -func=${COVERAGE_FILE}
	@COVERAGE=$$(go tool cover -func=${COVERAGE_FILE} | grep total | awk '{print $$3}' | sed 's/\%//g'); \
	echo "Current coverage is $${COVERAGE}%, minimal is ${MINIMAL_COVERAGE}."; \
	awk "BEGIN {exit ($${COVERAGE} < ${MINIMAL_COVERAGE})}"

coverage_html: test
	@go tool cover -html ${COVERAGE_FILE}

##################################
#######     Migrate       ########
##################################
.PHONY: migrations migrate sqlmigrate

SCHEMA_DIR=./spec/schema

migrations:
	go install entgo.io/ent/cmd/ent@v0.10.0
	ent generate --feature sql/upsert --target gen/entschema ${SCHEMA_DIR}

migrate:
	@go run ${BIN_SRC} migrate

sqlmigrate:
	@go run ${BIN_SRC} migrate --fake

##################################
#######      CI/CD        ########
##################################

.PHONY: init ci_test compile

ci_test: lint coverage

init: ensure build migrations

compile: build

cyclic:
	go install github.com/elza2/go-cyclic@latest
	go-cyclic run --dir . filter *_test.go
