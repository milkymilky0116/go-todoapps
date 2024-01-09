run:
	@bun run build
	@bun run tailwindcss -i ./views/static/css/input.css -o ./views/static/css/output.css 
	@templ generate
	@go build -o ./tmp/main ./cmd/web/main.go
