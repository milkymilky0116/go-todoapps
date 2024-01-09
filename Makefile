run:
	@bun run tailwindcss -i ./views/static/input.css -o ./views/static/output.css 
	@templ generate
	@go run cmd/web/main.go
