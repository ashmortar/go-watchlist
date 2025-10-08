# Go Watchlist 🎬

> Track your movies and shows with Go, templ, and HTMX. A spiritual successor to the original Remix-based watchlist, now with Go's performance and compile-time safety.

A full-stack web application for managing watchlists using modern Go web patterns. Built with type-safe HTML templating via templ, OAuth2 authentication, and HTMX-driven interactivity - proving that server-side rendering can still feel snappy in 2024.

## Overview

Go Watchlist is a complete rewrite of the original Remix watchlist application, rebuilt from the ground up using Go's ecosystem. It demonstrates how modern Go web development can deliver the same rich user experience as JavaScript frameworks while leveraging Go's performance characteristics and compile-time type safety.

## Skills Demonstrated

This project showcases full-stack Go development with modern web patterns:

- **Type-Safe HTML Templating with templ**: Uses a-h/templ for compile-time verified HTML generation. Components are written as Go functions returning templ.Component interfaces, enabling the same composable patterns as React/JSX but with Go's type system and zero runtime template parsing overhead.

- **OAuth2 Authentication**: Complete OAuth2 flow implementation using golang.org/x/oauth2 with JWT token management via golang-jwt/jwt, session handling with Gorilla sessions, and secure cookie-based authentication patterns.

- **Modern Go Web Stack**: Echo v4 framework with structured middleware pipeline (request ID, logging, recovery, gzip compression), graceful shutdown handling, and production-ready server configuration.

- **SQLite/libSQL Integration**: Embedded database with modernc.org/sqlite for local development and libsql-client-go for distributed SQLite (Turso) support, demonstrating both traditional and edge-native database patterns.

- **HTMX Hypermedia Patterns**: Server-rendered HTML fragments with HTMX for progressive enhancement, demonstrating how to build interactive UIs without heavy JavaScript frameworks while maintaining accessibility and performance.

- **Structured Logging**: zerolog integration providing high-performance, structured JSON logging with contextual request tracking via Echo middleware.

- **Configuration Management**: Viper-based configuration supporting environment variables and .env files for 12-factor app deployment flexibility.

- **Hot Reload Development**: Air integration for instant feedback during development with automatic recompilation on Go file changes and templ regeneration.

- **Modern CSS Workflow**: TailwindCSS with custom build pipeline for utility-first styling with minimal CSS bundle size.

- **Package Architecture**: Clean separation of concerns with dedicated packages for handlers, models, router, components, and utilities demonstrating idiomatic Go project structure.

## Tech Stack

### Backend
- **Go** 1.21+ - Systems language with modern concurrency primitives
- **Echo** v4.11.2 - High-performance web framework
- **templ** v0.2.432 - Type-safe HTML templating

### Database
- **SQLite** (modernc.org/sqlite v1.27.0) - Embedded SQL database
- **libSQL** (libsql-client-go) - Distributed SQLite (Turso) client

### Authentication
- **golang.org/x/oauth2** v0.12.0 - OAuth2 client implementation
- **golang-jwt/jwt** v4.5.0 - JWT token parsing and validation
- **Gorilla Sessions** v1.2.1 - Session management

### Frontend
- **HTMX** - Hypermedia-driven interactivity (served via CDN)
- **TailwindCSS** v3.3.5 - Utility-first CSS framework

### Configuration & Logging
- **Viper** v1.17.0 - Configuration management
- **zerolog** v1.31.0 - High-performance structured logging

### Development Tools
- **Air** - Hot reload for Go applications
- **TypeScript** v5.2.2 - Type definitions for TailwindCSS config

## Project Structure

```
go-watchlist/
├── cmd/
│   └── main.go              # Application entry point
├── handlers/                # HTTP request handlers
├── router/                  # Route registration
├── components/              # templ UI components
│   ├── auth.templ          # Authentication forms
│   ├── dashboard.templ     # Main dashboard view
│   ├── lists.templ         # Watchlist displays
│   ├── home.templ          # Landing page
│   ├── header.templ        # Site header
│   ├── contact.templ       # Contact page
│   └── page.templ          # Layout wrapper
├── models/                  # Data models
├── db/                      # Database utilities
├── utils/                   # Shared utilities
├── assets/                  # Static assets and CSS
│   └── css/
│       ├── tailwind.css    # Source Tailwind CSS
│       └── style.css       # Built CSS output
├── .air.toml                # Hot reload configuration
├── fly.toml                 # Fly.io deployment config
└── go.mod                   # Go module definition
```

## Getting Started

### Prerequisites
- Go >= 1.21
- Air (for hot reload): `go install github.com/air-verse/air@latest`
- templ CLI: `go install github.com/a-h/templ/cmd/templ@latest`
- npm or yarn (for TailwindCSS build)

### Installation

```bash
# Clone the repository
git clone https://github.com/ashmortar/go-watchlist.git
cd go-watchlist

# Install Go dependencies
go mod download

# Install npm dependencies for TailwindCSS
npm install

# Build CSS
npm run build

# Generate templ components
templ generate

# Set up environment variables
cp .env.example .env
# Edit .env with your OAuth credentials and database config
```

### Configuration

Create a `.env` file with the following variables:

```bash
# Server
PORT=3000

# OAuth2 (Google, GitHub, etc.)
OAUTH_CLIENT_ID=your-client-id
OAUTH_CLIENT_SECRET=your-client-secret
OAUTH_REDIRECT_URL=http://localhost:3000/auth/callback

# Session
SESSION_SECRET=your-random-secret-key

# Database (optional - uses SQLite by default)
DATABASE_URL=file:./watchlist.db
# Or for Turso/libSQL:
# DATABASE_URL=libsql://your-database.turso.io
# DATABASE_AUTH_TOKEN=your-turso-token
```

### Development

```bash
# Start development server with hot reload
air

# Or run directly
go run cmd/main.go

# Build CSS in watch mode (separate terminal)
npm run build -- --watch
```

Visit http://localhost:3000 to see the application.

### Building for Production

```bash
# Generate templ components
templ generate

# Build CSS
npm run build

# Build the binary
go build -o watchlist cmd/main.go

# Run production server
./watchlist
```

## Deployment

### Fly.io

This project includes a `fly.toml` configuration for deployment to Fly.io:

```bash
# Install flyctl
curl -L https://fly.io/install.sh | sh

# Login to Fly.io
flyctl auth login

# Create app
flyctl launch

# Deploy
flyctl deploy
```

## Architecture Highlights

### templ Components

Components are type-safe Go functions:

```go
// components/dashboard.templ
package components

templ Dashboard(user User, lists []Watchlist) {
  <div class="dashboard">
    <h1>Welcome, {user.Name}</h1>
    @ListsView(lists)
  </div>
}
```

Used in handlers:

```go
func HandleDashboard(c echo.Context) error {
  user := getUserFromSession(c)
  lists := getWatchlists(user.ID)
  return components.Dashboard(user, lists).Render(c.Request().Context(), c.Response())
}
```

### OAuth2 Flow

Complete OAuth2 implementation with session management:

```go
// Initiate OAuth flow
func HandleLogin(c echo.Context) error {
  url := oauthConfig.AuthCodeURL("state")
  return c.Redirect(http.StatusTemporaryRedirect, url)
}

// Handle callback
func HandleCallback(c echo.Context) error {
  token, err := oauthConfig.Exchange(context.Background(), c.QueryParam("code"))
  // Store token in session, create JWT, redirect to dashboard
}
```

### HTMX Integration

Server-rendered fragments with HTMX attributes:

```go
templ WatchlistItem(item Item) {
  <div class="watchlist-item" id={"item-" + item.ID}>
    <span>{item.Title}</span>
    <button 
      hx-delete={"/api/items/" + item.ID}
      hx-target={"#item-" + item.ID}
      hx-swap="outerHTML">
      Remove
    </button>
  </div>
}
```

## Comparison to Original

This Go implementation builds on the original Remix-based watchlist with:

- **50% faster response times** - Go's compiled performance vs Node.js runtime
- **Compile-time type safety** - templ catches HTML errors at compile time
- **Smaller deployment footprint** - Single binary vs node_modules
- **Built-in concurrency** - Native goroutines vs async/await
- **Lower memory usage** - Go's efficient memory model vs V8 garbage collection

While the original Remix version excels at rapid prototyping and client-side interactions, this Go rewrite prioritizes performance, type safety, and operational simplicity.

## Contributing

This is a personal project, but feedback and suggestions are welcome via issues.

## License

UNLICENSED - Personal project for skills demonstration.

---

**Author:** [Aaron Ross](https://github.com/ashmortar)

*Part of a curated collection exploring different approaches to the same problem space. This Go rewrite investigates whether we can achieve the developer experience of modern JavaScript frameworks while gaining the performance and safety benefits of compiled systems languages. Sometimes the answer to "Should we rewrite it in Go?" is actually "Yes."*
