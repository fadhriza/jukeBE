package database

// examples.go
//
// This package handles the database connection logic.
// It wraps the sql.Open call and ensures the connection is valid with a Ping.
//
// Example Usage:
//
// func main() {
//     cfg := config.LoadConfig()
//     db, err := database.Connect(cfg)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer db.Close()
//
//     // Use db for queries...
// }
