# CAPI - Command-Line Database Utility

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)](https://golang.org)
[![SQLite](https://img.shields.io/badge/SQLite-3.x-lightgrey)](https://sqlite.org)

A lightweight command-line utility for managing SQLite databases with a simple key-value oriented interface, built with Go and Cobra CLI.

## Features

- **Create Tables**: Easily create tables with specified columns (first column is primary key)
- **Insert Data**: Add records to database tables
- **View Data**: Display table contents and list all tables in a database
- **SQLite3 Support**: Uses SQLite3 with WAL (Write-Ahead Logging) mode enabled for better performance
- **Simple CLI**: Intuitive command-line interface with helpful commands

## Installation

```bash
# Clone the repository
git clone <repository-url>
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

- `-t, --table`: Table name (default: "defaultTable")
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
```

**Flags:**

- `-t, --table`: Specify table to view (optional)

### Examples

#### Complete Workflow Example

```bash
# 1. Create a database table
capi ct mydatabase -t products -c id,name,price,category

# 2. Insert data
capi insert mydatabase.db products id,name,price,category 101,Laptop,999.99,Electronics
capi insert mydatabase.db products id,name,price,category 102,Chair,149.99,Furniture

# 3. View all tables
capi view mydatabase.db

# 4. View specific table
capi view -t products mydatabase.db
```

## Project Structure

```
capi/
├── cmd/
│   ├── root.go          # Root command
│   ├── ct.go            # Create table command
│   ├── insert.go        # Insert command
│   └── view.go          # View command
├── data/
│   ├── data.go          # Database operations (CreateTable, Insert, ViewTable)
│   └── store/
│       └── store.go     # Database connection and higher-level operations
└── main.go              # Application entry point
```

## Database Schema

- All columns are created as `TEXT` type
- First specified column becomes the `PRIMARY KEY`
- Tables are created with `IF NOT EXISTS` clause
- SQLite WAL mode is automatically enabled for better concurrency

## Dependencies

- **Go**: >= 1.16
- **github.com/spf13/cobra**: CLI framework
- **github.com/mattn/go-sqlite3**: SQLite3 driver

## Development

```bash
# Run tests (if available)
go test ./...

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o capi-linux

# Clean build
go clean && go build
```

## Limitations

- Currently only supports SQLite databases
- All data is stored as TEXT type
- Simple comma-separated parsing (no handling of commas within values)
- First column is always the primary key

## 📦 Installation

### Prerequisites

- Go 1.16 or higher

### From Source

# Clone the repository

```bash
git clone https://github.com/michlima/capi.git
cd capi
```

# Build the application

```bash
go build -o capi
```

# Or install globally

```bash
go install
```

## License

Copyright © 2025 [Name Here] <[Email Address]>

## Notes

- Database files automatically get `.db` extension appended
- The utility is designed for simplicity and quick database operations
- For production use, consider adding validation and error handling improvements
