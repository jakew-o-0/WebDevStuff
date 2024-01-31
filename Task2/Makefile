run: build
	go run ./cmd/.

build:
	npx tailwindcss -i ./internal/base.css -o ./static/style.css
	templ generate ./templates
