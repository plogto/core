package db

// Open opens a database specified by the data source name.
// Format: host=foo port=5432 user=bar password=baz dbname=qux sslmode=disable"
// func Open(dataSourceName string) (*sql.DB, error) {
// 	conn, err := pgx.Connect(context.Background(), dataSourceName)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(context.Background())

// 	q := New(conn)

// 	author, err := q.GetAuthor(context.Background(), 1)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(author.Name)
// }
