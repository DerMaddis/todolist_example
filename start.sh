#!/bin/bash
npx tailwindcss -i ./internal/templates/tailwind.css -o ./internal/templates/assets/styles.css && templ generate && go run ./cmd/todolist/main.go
