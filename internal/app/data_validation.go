package app

func (app App) CheckIfContains(table string, needed string, column string) (bool, error) {
	var data bool

	err := app.db.QueryRow(`SELECT exists (SELECT 1 FROM knowledge WHERE $1 = $2 LIMIT 1)`, column, needed).Scan(&data)

	return data, err
}
