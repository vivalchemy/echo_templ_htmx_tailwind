PROJECT_NAME=github.com/vivalchemy/echo_templ_htmx_tailwind
PORT=3000

dev: ./main.go
	@air

build: ./main.go
	@go build -o ./build/${PROJECT_NAME}

tailwind: ./static/input.css
	@tailwindcss -i ./static/input.css -o ./build/output.css --watch

templ:
	@templ generate --watch --proxy http://localhost:${PORT}

clean: ./build/
	rm -rf tmp/*
