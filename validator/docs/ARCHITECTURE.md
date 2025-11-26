# GTFS Validator Architecture

## Overview

The GTFS Validator is a Go application that validates GTFS (General Transit Feed Specification) data files. It uses SQLite for efficient, low-memory data storage and provides a modular validation system.

## Architecture Components

### 1. Data Storage Layer

#### SQLite-Based Storage
- **Location**: `services/gtfs_sqlite.go`, `types/gtfs-raw.go`
- **Purpose**: Stores GTFS data in SQLite instead of in-memory slices
- **Benefits**: 
  - Low memory footprint (can handle 100M+ rows)
  - Persistent storage during validation
  - Efficient querying with indexes

#### Database Operations
- **Location**: `services/database/operations.go`
- **Purpose**: Pure database operations (table creation, sanitization, queries)
- **Separation**: Keeps database logic separate from business logic

### 2. Validation System

#### Validation Registry
- **Location**: `validations/registry.go`
- **Purpose**: Auto-registration system for validation packages
- **How it works**: Each validation package registers itself via `init()` functions
- **Benefits**: No hardcoded table lists, easy to add new validations

#### Validation Packages
- **Structure**: `validations/{table_name}/main.go`
- **Pattern**: Each package exports `RunValidations(gtfs types.Gtfs, rules *types.GtfsRules)`
- **Registration**: Auto-registers via `init()` function

#### Validation Context
- **Location**: `lib/validation_context.go`
- **Purpose**: Encapsulates common validation patterns
- **Features**: Message creation, severity checking, rule application

### 3. Configuration

#### Constants
- **Location**: `config/config.go`
- **Purpose**: Centralized configuration constants
- **Includes**: Batch sizes, issue limits, progress thresholds

#### GTFS File Definitions
- **Location**: `config/gtfs_files.go`
- **Purpose**: Single source of truth for GTFS file/table names
- **Benefits**: Eliminates duplication across codebase

### 4. Utilities

#### Progress Tracking
- **Location**: `lib/progress.go`
- **Purpose**: Centralized progress logging during validations
- **Features**: Automatic percentage logging, row-based thresholds

#### Test Helpers
- **Location**: `lib/test_helpers.go`
- **Purpose**: Common test utilities to reduce duplication
- **Features**: Assertion helpers, test message service interface

## Data Flow

### 1. Import Phase
```
GTFS ZIP File
    ↓
Extract & Parse CSV Files
    ↓
Stream Rows to SQLite (Batch Inserts)
    ↓
Create ID Map Table (for primary key lookups)
    ↓
GTFS Instance with SQLite Connection
```

### 2. Validation Phase
```
For each GTFS table:
    ↓
Check if table exists in database
    ↓
Get registered validation function
    ↓
Iterate rows using iterator methods
    ↓
Run validation functions
    ↓
Collect messages via MessageService
```

### 3. Output Phase
```
MessageService Summary
    ↓
JSON Output (or file)
```

## Key Design Patterns

### Iterator Pattern
- **Purpose**: Process large datasets without loading into memory
- **Implementation**: `IterateStops()`, `IterateTrips()`, etc.
- **Benefits**: Constant memory usage regardless of dataset size

### Registry Pattern
- **Purpose**: Auto-discovery of validation functions
- **Implementation**: `validations.Register()` in each package's `init()`
- **Benefits**: No manual registration, easy to extend

### Dependency Injection (Partial)
- **Purpose**: Improve testability
- **Implementation**: `MessageServiceInterface` for message service
- **Future**: Full DI for all services

## Error Handling

### Validation Errors
- **Location**: `types/validation_error.go`
- **Purpose**: Structured error types with context
- **Features**: Field, file, row, validation ID context

### Database Errors
- **Location**: `types/validation_error.go`
- **Purpose**: Database operation errors with context
- **Features**: Operation, table name context

## Performance Optimizations

### Caching
- **Trip Stop Sequences**: Pre-computed min/max sequences per trip
- **Location**: `types/cache.go`
- **Purpose**: Avoid N+1 queries during validation

### Batch Operations
- **SQLite Inserts**: Batch size of 2000 rows per transaction
- **Purpose**: Reduce transaction overhead
- **Configuration**: `config.BatchSize`

### Database Settings
- **WAL Mode**: Write-Ahead Logging for better concurrency
- **Synchronous**: NORMAL mode for balance between durability and performance
- **Connection Pool**: Single connection to serialize writes

## Extension Points

### Adding a New Validation

1. Create package: `validations/{table_name}/main.go`
2. Implement `RunValidations()` function
3. Register in `init()`: `registry.Register("table_name", RunValidations)`
4. Add table to `config/gtfs_files.go` if new GTFS file

### Adding a New Validation Function

1. Create file: `validations/{table_name}/validations/{field}_validation.go`
2. Follow existing validation function pattern
3. Call from `main.go`'s `RunValidations()`

### Customizing Validation Rules

- Rules are defined in `types/gtfs-raw.go` (GtfsRules struct)
- Each validation checks rules and applies severity accordingly
- Rules can be loaded from JSON/YAML configuration files

## Testing

### Test Helpers
- **Location**: `lib/test_helpers.go`
- **Purpose**: Reduce test code duplication
- **Features**: 
  - Assertion helpers for errors/warnings
  - Test message service interface
  - Test GTFS rules creation

### Test Structure
- Each validation package has a `tests/` directory
- Tests use `services.AppMessageService` (can be replaced with test instance)
- Common pattern: Clear → Run → Assert

## Future Improvements

1. **Full Dependency Injection**: Make all services injectable
2. **Validation Error Returns**: Standardize error returns from validation functions
3. **Parallel Validation**: Run independent validations in parallel
4. **Incremental Validation**: Only validate changed rows
5. **Validation Plugins**: External validation packages

## Dependencies

- **SQLite**: `modernc.org/sqlite` (pure Go SQLite driver)
- **CSV Parsing**: Standard library `encoding/csv`
- **ZIP Handling**: Standard library `archive/zip`
- **Table Output**: `github.com/olekukonko/tablewriter`

## File Organization

```
validator/
├── config/              # Configuration constants
├── docs/                # Documentation
├── lib/                 # Shared utilities
├── services/            # Core services
│   └── database/        # Database operations
├── types/               # Type definitions
├── validations/         # Validation packages
│   ├── registry.go      # Auto-registration
│   └── {table}/         # Per-table validations
└── main.go             # Entry point
```

