<!-- psql "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
goose -dir sql/schema postgres "postgres://ravindra_choudhary:@localhost:5432/gator?sslmode=disable" down // to delete or reset database manually // apply the same url above here .

go run . reset // -- direct common for complete reset


rc5091119-pixel/Blog_Aggregator

goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up

NOTE --> Middleware is a way to wrap a function with additional functionality. It is a common pattern that allows us to write DRY code. -->

# Gator – RSS Feed Aggregator CLI

Gator is a command-line RSS feed aggregator written in Go.
It allows users to follow RSS feeds, periodically scrape posts from them, store those posts in a PostgreSQL database, and browse the latest posts directly from the terminal.

This project demonstrates how to build a backend CLI tool using Go, PostgreSQL, and SQL queries generated with sqlc.

---

## Features

* Register and manage users
* Add and follow RSS feeds
* Periodically scrape feeds for new posts
* Store posts in a PostgreSQL database
* Browse posts from followed feeds directly from the terminal

---

## Requirements

Before running the program, make sure you have the following installed:

* **Go (1.21 or later)**
* **PostgreSQL**

You can verify installations with:

```bash
go version
psql --version
```

---

## Installing the CLI

To install the Gator CLI, run:

```bash
go install github.com/your-github-username/gator@latest
```

This will compile the program and place the `gator` binary in your `$GOPATH/bin` directory.

Make sure that directory is in your system `PATH`.

After installation you should be able to run:

```bash
gator
```

---

## Database Setup

Create a PostgreSQL database:

```bash
createdb gator
```

Set your database connection URL:

```bash
export DB_URL="postgres://username:password@localhost:5432/gator?sslmode=disable"
```

Run database migrations (using goose):

```bash
goose postgres $DB_URL up
```

This will create the required tables:

* users
* feeds
* feed_follows
* posts

---

## Configuration File

Create a configuration file at:

```
~/.gatorconfig.json
```

Example configuration:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

This file stores the database connection and the currently active user.

---

## Running the Program

During development you can run:

```bash
go run .
```

However, once installed you should use:

```bash
gator
```

---

## Example Commands

### Register a User

```bash
gator register ravindra
```

Creates a new user and sets them as the current user.

---

### Add a Feed

```bash
gator addfeed "TechCrunch" https://techcrunch.com/feed/
```

Adds a new RSS feed to the database and follows it.

---

### View Available Feeds

```bash
gator feeds
```

Lists all feeds in the database.

---

### Follow a Feed

```bash
gator follow https://techcrunch.com/feed/
```

Follow an existing feed.

---

### Start the Feed Scraper

```bash
gator agg 10s
```

This command continuously scrapes feeds every 10 seconds and saves posts to the database.

---

### Browse Posts

```bash
gator browse
```

Shows the latest 2 posts from feeds you follow.

You can specify a limit:

```bash
gator browse 5
```

This will display the 5 most recent posts.

---

## Development Notes

Go programs are statically compiled binaries.

After building the project with:

```bash
go build
```

or installing it with:

```bash
go install
```

you can run the program directly without needing the Go toolchain.

`go run .` should only be used during development.

---

## Example RSS Feeds for Testing

You can test the program using the following feeds:

* https://techcrunch.com/feed/
* https://news.ycombinator.com/rss
* https://www.boot.dev/blog/index.xml

---

## Repository

GitHub repository:

```
https://github.com/your-github-username/gator
```

---

## Author

Ravindra Choudhary
