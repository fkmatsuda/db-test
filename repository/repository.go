package repository

import "time"

// TimeRecord records the time of operations.
type TimeRecord struct {
	// StartTime marks the start of operations.
	StartTime time.Time

	// RowsTime marks operations throughout the process.
	RowsTime map[int64]time.Duration

	// EndTime marks the end of operations.
	EndTime time.Time
}

// Repository is the interface that wraps the database operations
type Repository interface {
	// Close closes the database connection.
	Close() error

	// Migrate performs the database migrations.
	Migrate() (TimeRecord, error)

	// Insert performs data insertion operations.
	Insert(checkpoint int64) (TimeRecord, error)

	// InsertUnique performs non-repeating data insertion operations.
	InsertUnique(checkpoint int64) (TimeRecord, error)

	// Query executes queries on the entered data.
	Query(checkpoint int64) (TimeRecord, error)
}
