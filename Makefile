
GOLIBS:=bluetoo.co/otflib
GOAPPS:=bluetoo.co/otf-export bluetoo.co/otf-import
GORUN:=src/bluetoo.co/otf-export/ort-export.go
GOCLEAN:=otf-export otf-import


# -------------------------------------------------
all:test build

# `make ci` to continuously run tests
# requires entr(1) http://entrproject.org/
#   `brew install entr` on OS X if you have homebrew
ci:
	@echo Monitoring for source changes...
	@find . -name '*.go' | entr -c make test

GO:=go
GO_ALL_TARGETS:=$(GOLIBS) $(GOAPPS)
GO_TESTS = $(foreach itr, $(GO_ALL_TARGETS), $(itr)_test)
GO_APPS = $(foreach itr, $(GOAPPS), $(itr)_apps)
GO_CLEAN = $(foreach itr, $(GO_ALL_TARGETS), $(itr)_clean)
GO_CLEAN_EXTRAS = $(foreach itr, $(GOCLEAN), $(itr)_extraclean)

.PHONY: $(GO_TESTS) $(GO_APPS) $(GO_RUN) $(GO_CLEAN) $(GO_CLEAN_EXTRAS)

test: export GOPATH = $(shell pwd)
test: $(GO_TESTS)

build: export GOPATH = $(shell pwd)
build: $(GO_APPS)

clean: export GOPATH = $(shell pwd)
clean: $(GO_CLEAN) $(GO_CLEAN_EXTRAS)
	
run: export GOPATH = $(shell pwd)
run:
	$(GO) run $(GORUN)
	
$(GO_TESTS): %_test:
	@echo [TEST] $*
	$(GO) test $* -test.v

$(GO_APPS): %_apps:
	$(GO) build $*

$(GO_CLEAN): %_clean:
	$(GO) clean $*

$(GO_CLEAN_EXTRAS): %_extraclean:
	@echo [CLEAN] $*
	@if [ -f $* ]; then rm $*; fi;