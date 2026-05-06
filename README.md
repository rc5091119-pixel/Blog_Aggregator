# Gator

A fast, terminal-native RSS feed aggregator written in Go — follow feeds, scrape posts automatically, and browse content without leaving your terminal.

---

## What is Gator?

Gator is a CLI tool that lets you manage RSS feeds from the command line. It runs a background scraper on a configurable interval, stores all posts in PostgreSQL, and lets you browse the latest content across all feeds you follow — no browser required.

---

## Features

- **User management** — Register and switch between multiple users from the CLI
- **Feed management** — Add, list, follow, and unfollow any RSS feed by URL
- **Automated scraping** — Background aggregator polls feeds at a custom interval and stores new posts
- **ETL pipeline** — Full fetch → parse → transform → store pipeline with duplicate prevention
- **Normalized schema** — PostgreSQL with migrations, foreign keys, and no duplicate records
- **Browse from terminal** — View recent posts with configurable limit directly in your shell

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go (Golang) |
| Database | PostgreSQL |
| Migrations | Goose |
| Config | JSON (`~/.gatorconfig.json`) |
| Data pipeline | Custom ETL (fetch → parse → store) |

---

## Requirements

- **Go** 1.20 or newer
- **PostgreSQL** 14 or newer
- **Goose** (for running migrations)

Verify your setup:

```bash
go version
psql --version
```

---

## Installation

Install the CLI directly with `go install`:

```bash
go install github.com/rc5091119-pixel/gator@latest
```

This compiles and places the `gator` binary in your `$GOPATH/bin`. Once installed, run it from anywhere:

```bash
gator
```

> For development, use `go run .` inside the project directory instead.

---

## Setup

### 1. Create the database

```bash
createdb gator
```

### 2. Run migrations

```bash
goose postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" up
```

### 3. Create the config file

Create `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

This file stores your database connection string and tracks the currently active user across sessions.

---

## Usage

### User Commands

```bash
# Register a new user (sets them as current user)
gator register ravindra

# Login as an existing user
gator login ravindra

# List all registered users
gator users
```

### Feed Commands

```bash
# Add a new RSS feed
gator addfeed "TechCrunch" https://techcrunch.com/feed/

# List all feeds in the database
gator feeds

# Follow a feed (by URL)
gator follow https://techcrunch.com/feed/

# Unfollow a feed
gator unfollow https://techcrunch.com/feed/

# List feeds you currently follow
gator following
```

### Aggregator

```bash
# Start scraping feeds every 10 seconds
gator agg 10s

# Scrape every 1 minute
gator agg 1m
```

The aggregator runs continuously in the foreground, fetching new posts and storing them. Use `Ctrl+C` to stop.

### Browse Posts

```bash
# Show recent posts from followed feeds (default limit)
gator browse

# Show a specific number of posts
gator browse 10
```

---

## Example Feeds to Get Started

```
https://techcrunch.com/feed/
https://news.ycombinator.com/rss
https://www.boot.dev/blog/index.xml
```

---

## How the Aggregator Works

```
gator agg 10s
      │
      ▼
Fetch next feed from DB (least recently fetched)
      │
      ▼
HTTP GET → parse XML/RSS
      │
      ▼
Transform items into post structs
      │
      ▼
INSERT into posts (ON CONFLICT DO NOTHING — no duplicates)
      │
      ▼
Wait interval → repeat
```

Each cycle picks the feed that was scraped longest ago, ensuring fair distribution across all followed feeds.

---

## Project Structure

```
gator/
├── main.go                  # Entry point + command router
├── config.go                # Config file read/write (~/.gatorconfig.json)
├── commands.go              # CLI command definitions
├── handler_user.go          # User register/login/list handlers
├── handler_feed.go          # Add/list/follow/unfollow feed handlers
├── handler_agg.go           # Aggregator loop + scraping logic
├── handler_posts.go         # Browse posts handler
├── internal/
│   └── database/            # SQLC-generated DB layer
├── sql/
│   ├── queries/             # Raw SQL queries
│   └── schema/              # Goose migration files
├── go.mod
├── go.sum
└── README.md
```

---

## Author

**Ravindra Choudhary**
B.Tech — Electronics and Communication Engineering, NIT Agartala | GPA: 8.73

- 📧 rc5091119@gmail.com
- 🐙 [github.com/rc5091119-pixel](https://github.com/rc5091119-pixel)

---

> Gator — because your RSS feeds deserve a proper home.