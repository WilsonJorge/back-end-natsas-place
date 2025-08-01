#!/bin/bash

echo "🌱 Running database seeds..."

# Compilar o seed
go build -o tmp/seed cmd/seed/main.go

# Executar o seed
./tmp/seed

echo "✅ Seed completed!" 