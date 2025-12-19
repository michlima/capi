# CAPI - Command-Line Database Utility

**Project Status:** This is an incomplete project and is still ongoing.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)](https://golang.org)
[![SQLite](https://img.shields.io/badge/SQLite-3.x-lightgrey)](https://sqlite.org)

A lightweight command-line utility for managing SQLite databases with a simple key-value oriented interface, built with Go and Cobra CLI.

## Features

- **Create Tables**: Easily create tables with specified columns (first column is primary key)
- **Insert Data**: Add records to database tables
- **View Data**: Display table contents and list all tables in a database
- **Interactive Mode**: Use the `use` command for an interactive experience to manage your tables.
- **SQLite3 Support**: Uses SQLite3 with WAL (Write-Ahead Logging) mode enabled for better performance
- **Simple CLI**: Intuitive command-line interface with helpful commands

## Installation

```bash
git clone https://github.com/michlima/capi.git
cd capi
# Build the application
go build -o capi
# Or install globally
go install
```

## Usage

### Basic Command Structure

```bash
capi [command] [arguments] [flags]
```

### Available Commands

#### 1. Create Table (`ct`)

Create a new table in the database.

```bash
capi ct <databaseName> -t <tableName> -c <columns>

# Example:
capi ct myApp -t users -c id,name,email,phone
# Creates: users table with columns: id (PRIMARY KEY), name, email, phone
```

**Flags:**

- `-t, --table`: Specify table name (default: "defaultTable")
- `-c, --cols`: Comma-separated column names (first is primary key, default: "key,value")

#### 2. Insert Data (`insert`)

Insert data into an existing table.

```bash
capi insert <database> <table> <columns> <values>

# Example:
capi insert myApp.db users id,name,email,phone 1,John,john@example.com,123456
```

**Note:** Columns and values must be in the same order and comma-separated.

#### 3. View Data (`view`)

View tables or table contents.

```bash
# View all tables in a database
capi view <database>

# View specific table contents
capi view -t <tableName> <database>

# View filtered rows from a table
capi view -t <tableName> -f "column=value" <database>
```

**Flags:**

- `-t, --table`: Specify table to view (optional)
- `-f, --filter`: Filter table to see specific rows (e.g., "id=1,name=John")

#### 4. Delete Data (`delete`)

Delete tables or rows from a database.

```bash
# Delete an entire table
capi delete <database> -t <tableName>

# Delete specific rows from a table (using filter)
capi delete <database> -t <tableName> -f "column=value"

# Example:
capi delete myApp -t users
# Deletes the 'users' table from 'myApp.db'

capi delete myApp -t users -f "id=1"
# Deletes rows where id is '1' from the 'users' table in 'myApp.db'
```

**Flags:**

- `-t, --table`: Table to delete from (required)
- `-f, --filter`: Filter to select rows for deletion (e.g., "id=1,name=John")

#### 5. List Databases (`list-db`)

List all available SQLite databases (`.db` files) in the `storage` directory.

```bash
capi list-db
```

#### 6. Use Database (`use`)

Enter an interactive mode to manage a database.

```bash
capi use <database>
```

This command will present a list of tables in the database. After selecting a table, you'll be presented with an interactive menu with the following options:
- **View**: View the content of the selected table.
- **Edit**: (Not yet implemented)
- **Add**: Add a new row to the table.
- **Delete**: Delete rows from the table.
- **Back**: Go back to the table selection menu.

**_still incomplete_**

## Project Structure

```
capi/
├── cmd/
│   ├── root.go          # Root command
│   ├── ct.go            # Create table command
│   ├── insert.go        # Insert command
│   ├── view.go          # View command
│   ├── delete.go        # Delete command (tables or rows)
│   ├── list-db.go       # List databases command
│   └── use.go           # Interactive mode command
├── core/
│   ├── commands.go      # Database operations (CreateTable, Insert, 
│   ├── utils.go         # Utility functions (OpenDatabase, PrintHeaders, 
│   └── store/
│       └── sqlite.go    # Database connection and higher-level operations 
├── interact/
│   └── interactions.go  # Interactive mode logic
└── main.go              # Application entry point
```

## Database Schema

- All columns are created as `TEXT` type
- First specified column becomes the `PRIMARY KEY`
- Tables are created with `IF NOT EXISTS` clause
- SQLite WAL mode is automatically enabled for better concurrency

## Dependencies

- **Go**: >= 1.21
- **github.com/spf13/cobra**: CLI framework
- **github.com/AlecAivazis/survey/v2**: Interactive prompts
- **github.com/mattn/go-sqlite3**: SQLite3 driver

## Limitations

- Currently only supports SQLite databases
- All data is stored as TEXT type
- Simple comma-separated parsing (no handling of commas within values)
- First column is always the primary key
- The "Edit" functionality in interactive mode is not yet implemented.

### Prerequisites

- Go 1.21 or higher

## License

Copyright © 2025 Michael Lima

## Notes

- Database files automatically get `.db` extension appended
- The utility is designed for simplicity and quick database operations
- For production use, consider adding validation and error handling improvements