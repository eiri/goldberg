.DEFAULT_GOAL := run

PROJ := goldberg
SRCS := $(shell find $(CURDIR) -name '*.go')

$(PROJ): $(SRCS)
	go build -o $(PROJ) ./...

.PHONY: run
run: $(PROJ)
	$(CURDIR)/$(PROJ) --server

.PHONY: clean
clean:
	go clean
	rm -f $(CURDIR)/$(PROJ)
