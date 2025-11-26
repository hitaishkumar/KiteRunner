#!/bin/bash

echo "📁 Creating KiteRunner TUI folder structure..."

# Base internal structure
mkdir -p internal/app
mkdir -p internal/ui/login
mkdir -p internal/ui/dashboard
mkdir -p internal/ui/watchlist
mkdir -p internal/ui/trades
mkdir -p internal/ui/components

mkdir -p internal/config
mkdir -p internal/kite
mkdir -p internal/util
mkdir -p internal/version

# CLI entrypoint
mkdir -p cmd/kiterunner

# Assets (ASCII banners, templates, etc.)
mkdir -p assets

echo "✅ Folder structure created successfully!"
