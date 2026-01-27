package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Use the - directive to hide it from json output
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"` // Add the omitempty directive
	// Use the Runtime type instead of int 32. Note that the omitempty directive will
	// still work on this: if the Runtime field has the underlying value 0, then it will
	// be considered empty and omitted -- and the MarshalJSON() method we jus made
	// won't be called at all.
	Runtime Runtime  `json:"runtime,omitempty"` // Add the omitempty directive
	Geners  []string `json:"geners,omitempty"`  // Add the omitempty directive
	Version int32    `json:"version"`
	// The version number starts at 1 and will be incremented each
	// time the movie information is updated
}
