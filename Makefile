GO ?= GO111MODULE=on go
INFO_COLOR=\033[1;34m
RESET=\033[0m
BOLD=\033[1m

depsdev: ## Installing dependencies for development
	$(GO) get -u golang.org/x/lint/golint
	$(GO) get -u github.com/git-chglog/git-chglog/cmd/git-chglog

changelog:
	git-chglog -o CHANGELOG.md

test: ## Run test
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Testing$(RESET)"
	$(GO) test -v -timeout=30s -parallel=4
	$(GO) test -race

lint: ## Exec golint
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Linting$(RESET)"
	cd $(PACKAGE_DIR) && golint -min_confidence 1.1 -set_exit_status


ci: depsdev test lint

release_major: releasedeps
	git semv major --bump

release_minor: releasedeps
	git semv minor --bump

release_patch: releasedeps
	git semv patch --bump

releasedeps: git-semv

git-semv:
	which git-semv > /dev/null || brew tap linyows/git-semv
	which git-semv > /dev/null || brew install git-semv
