package modules

import "log"

func (m *Module) InsertMovieLog(reqString string) error {
	if m.Cfg.DB.IsMock {
		log.Println(reqString)
		return nil
	}

	_, err := m.Db.Exec("insert into movie_log(payload, created_at) values ($1, current_timestamp)", reqString)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
