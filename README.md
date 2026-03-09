# Gator – RSS Feed Aggregator CLI

Gator is a command-line RSS feed aggregator written in Go.
It allows users to follow RSS feeds, scrape posts from them periodically, store them in a PostgreSQL database, and browse posts directly from the terminal.

---

# Requirements

To run this program you must have the following installed:

* Go (version 1.20 or newer)
* PostgreSQL

You can verify installation with:

```bash
go version
psql --version
```

---

# Installing the Gator CLI

You can install the CLI using `go install`:

```bash
go install github.com/YOUR_GITHUB_USERNAME/gator@latest
```

This will compile the program and install the `gator` binary on your system.

Once installed, you should be able to run:

```bash
gator
```

---

# Database Setup

Create a PostgreSQL database:

```bash
createdb gator
```

Then set the database URL environment variable:

```bash
export DB_URL="postgres://username:password@localhost:5432/gator?sslmode=disable"
```

Run the database migrations:

```bash
goose postgres $DB_URL up
```

This will create the required tables.

---

# Config File Setup

Create a config file at:

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

# Running the Program

During development you can run:

```bash
go run .
```

However, this is **only for development**.

For production usage you should run the compiled CLI:

```bash
gator
```

Go programs are statically compiled binaries. After running `go build` or `go install`, the program can run without the Go toolchain.

---

# Example Commands

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

Adds an RSS feed to the database.

---

### View Feeds

```bash
gator feeds
```

Lists all available feeds.

---

### Start the Feed Aggregator

```bash
gator agg 10s
```

This command scrapes feeds every 10 seconds and stores new posts in the database.

---

### Browse Posts

```bash
gator browse
```

Shows the most recent posts from feeds you follow.

You can specify how many posts to show:

```bash
gator browse 5
```

---

# Example Feeds for Testing

You can test the program with these RSS feeds:

* https://techcrunch.com/feed/
* https://news.ycombinator.com/rss
* https://www.boot.dev/blog/index.xml

---

# Repository

After pushing your project to GitHub, your repo link should look like:

```
https://github.com/YOUR_GITHUB_USERNAME/gator
```

Replace `YOUR_GITHUB_USERNAME` with your actual GitHub username.

---

# Author

Ravindra Choudhary
