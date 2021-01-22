.DEFAULT_GOAL := run

PROJ := goldberg
SRCS := $(shell find $(CURDIR) -name '*.go')

$(PROJ): $(SRCS)
	go build -o $(PROJ) ./cmd/...

.PHONY: run
run: $(PROJ)
	$(CURDIR)/$(PROJ) --server

.PHONY: test
test:
	go test -v -race ./...


.PHONY: clean
clean:
	go clean
	rm -f $(CURDIR)/$(PROJ)
