APP_DIR = example/app/

lint:
	make lint.ci-lint ; make lint.gofactory ; make lint.gopublicfield

lint.ci-lint:
	golangci-lint run $(APP_DIR)... -c example/app/.golangci.yml

# Встраиваю его в golangci-lint https://github.com/golangci/golangci-lint/pull/4196
lint.gofactory: lint.gofactory.install lint.gofactory.run

lint.gofactory.install:
	go install github.com/maranqz/go-factory-lint/v2/cmd/go-factory-lint@latest

lint.gofactory.run:
	go-factory-lint --packageGlobs="ddd/example/app/domain/**" --onlyPackageGlobs=true ./$(APP_DIR)...

lint.gopublicfield: lint.gopublicfield.install lint.gopublicfield.run

lint.gopublicfield.install:
	go install github.com/maranqz/gopublicfield/cmd/gopublicfield@latest

lint.gopublicfield.run:
	gopublicfield --packageGlobs="ddd/example/app/domain/**" --onlyPackageGlobs=true ./$(APP_DIR)...

lint.tags: lint.tags.install lint.tags.run

lint.tags.install:
	go install -v github.com/quasilyte/go-ruleguard/cmd/ruleguard@latest

lint.tags.run:
	ruleguard -rules tags/rules.go -fix ./$(APP_DIR)...