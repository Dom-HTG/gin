package utils

//initialize and return database connection.
func initDatabase() (gorm.DB, error) {
	if conn, err := gorm.open(postgres.open(), gorm.options{}); err != nil {
		return nil, err
	}
	return conn, nil
}
