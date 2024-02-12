
.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh openapi ports ports

.PHONY: start
start: deps-start
	docker-compose --env-file=.env -f config/docker-compose.yml up -dV note-taker

deps-start:
	docker-compose --env-file=.env -f config/docker-compose.yml up -dV db && sleep 2

swagger-ui:
	docker-compose --env-file=.env -f config/docker-compose.yml up -dV swagger-ui

clean:
	docker-compose --env-file=.env -f config/docker-compose.yml down -v

