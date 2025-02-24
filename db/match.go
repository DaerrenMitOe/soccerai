package db

import "database/sql"

type Match []struct {
	MatchID     int    `json:"match_id"`
	MatchDate   string `json:"match_date"`
	KickOff     string `json:"kick_off"`
	Competition struct {
		CompetitionID   int    `json:"competition_id"`
		CountryName     string `json:"country_name"`
		CompetitionName string `json:"competition_name"`
	} `json:"competition"`
	Season struct {
		SeasonID   int    `json:"season_id"`
		SeasonName string `json:"season_name"`
	} `json:"season"`
	HomeTeam struct {
		HomeTeamID     int    `json:"home_team_id"`
		HomeTeamName   string `json:"home_team_name"`
		HomeTeamGender string `json:"home_team_gender"`
		HomeTeamGroup  string `json:"home_team_group"`
		Country        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		Managers []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Nickname string `json:"nickname"`
			Dob      string `json:"dob"`
			Country  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"country"`
		} `json:"managers"`
	} `json:"home_team"`
	AwayTeam struct {
		AwayTeamID     int    `json:"away_team_id"`
		AwayTeamName   string `json:"away_team_name"`
		AwayTeamGender string `json:"away_team_gender"`
		AwayTeamGroup  string `json:"away_team_group"`
		Country        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		Managers []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Nickname string `json:"nickname"`
			Dob      string `json:"dob"`
			Country  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"country"`
		} `json:"managers"`
	} `json:"away_team"`
	HomeScore      int    `json:"home_score"`
	AwayScore      int    `json:"away_score"`
	MatchStatus    string `json:"match_status"`
	MatchStatus360 string `json:"match_status_360"`
	LastUpdated    string `json:"last_updated"`
	LastUpdated360 string `json:"last_updated_360"`
	Metadata       struct {
		DataVersion         string `json:"data_version"`
		ShotFidelityVersion string `json:"shot_fidelity_version"`
		XyFidelityVersion   string `json:"xy_fidelity_version"`
	} `json:"metadata"`
	MatchWeek        int `json:"match_week"`
	CompetitionStage struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"competition_stage"`
	Stadium struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"stadium"`
	Referee struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"referee"`
}

func createMatchSQLTable(db *sql.DB) {
	// Create the table
	matchTableSQL := `CREATE TABLE IF NOT EXISTS match (
		"id" INTEGER PRIMARY KEY,
		"date" TEXT,
		"kick_off" TEXT,
		"competition_id" INTEGER,
		"season_id" INTEGER,
		"home_team_id" INTEGER,
		"away_team_id" INTEGER,
		"home_score" INTEGER,
		"away_score" INTEGER,
		"status" TEXT,
		"status_360" TEXT,
		"last_updated" TEXT,
		"last_updated_360" TEXT,
		"metadata_id" INTEGER,
		"week" INTEGER,
		"competition_stage_id" INTEGER,
		"stadium_id" INTEGER,
		"referee_id" INTEGER
	);`

	// Create the table
	competitionTableSQL := `CREATE TABLE IF NOT EXISTS competition (
		"id" INTEGER PRIMARY KEY,
		"name" TEXT,
		"country_name" TEXT
	);`

	// Create the table
	seasonTableSQL := `CREATE TABLE IF NOT EXISTS season (
		"id" INTEGER PRIMARY KEY,
		"name" TEXT
	);`

	teamTableSQL := `CREATE TABLE IF NOT EXISTS team (
		"id" INTEGER PRIMARY KEY,
		"name" TEXT,
		"gender" TEXT,
		"group" TEXT,
		"country_id" INTEGER,
		"managers_id" INTEGER
	);`

	countryTableSQL := `CREATE TABLE IF NOT EXISTS country (
		"id" INTEGER PRIMARY KEY,
		"name" TEXT
	);`

	managerTableSQL := `CREATE TABLE IF NOT EXISTS manager (
		"id" INTEGER PRIMARY KEY,
		"name" TEXT,
		"nickname" TEXT,
		"dob" TEXT,
		"country_id" INTEGER
	);`

	metadataTableSQL := `CREATE TABLE IF NOT EXISTS metadata (
		"id" INTEGER PRIMARY KEY,
		"data_version" TEXT,
		"shot_fidelity_version" INTEGER,
		"xy_fidelity_version" INTEGER
	);`

	competitionStageTableSQL := `CREATE TABLE IF NOT EXISTS competition_stage (
		"id" INTEGER PRIMARY KEY,
		"name" TEXT
	);`

	stadiumTableSQL := `CREATE TABLE IF NOT EXISTS stadium (
		"id" INTEGER PRIMARY KEY,
		"name" TEXT,
		"country_id" INTEGER
	);`

	refereeTableSQL := `CREATE TABLE IF NOT EXISTS referee (
		"id" INTEGER PRIMARY KEY,
		"name" TEXT,
		"country_id" INTEGER
	);`

	createSQLTable(db, matchTableSQL)
	createSQLTable(db, competitionTableSQL)
	createSQLTable(db, seasonTableSQL)
	createSQLTable(db, teamTableSQL)
	createSQLTable(db, countryTableSQL)
	createSQLTable(db, managerTableSQL)
	createSQLTable(db, metadataTableSQL)
	createSQLTable(db, competitionStageTableSQL)
	createSQLTable(db, stadiumTableSQL)
	createSQLTable(db, refereeTableSQL)
}

func insertMatchSQLData(db *sql.DB, matches *Match) {
	// Insert data into competition table if competition ID does not exist
	insertCompetitionSQL := `INSERT INTO competition (id, name, country_name) SELECT ?, ?, ? WHERE NOT EXISTS (
	SELECT 1 FROM competition WHERE id = ?
);`

	// Insert data into season table if season ID does not exist
	insertSeasonSQL := `INSERT INTO season (id, name) SELECT ?, ? WHERE NOT EXISTS (
	SELECT 1 FROM season WHERE id = ?
);`

	// Insert data into team table if team ID does not exist
	insertTeamSQL := `INSERT INTO team (id, name, gender, group, country_id) SELECT ?, ?, ?, ?, ? WHERE NOT EXISTS (
	SELECT 1 FROM team WHERE id = ?
);`

	// Insert data into country table if country ID does not exist
	insertCountrySQL := `INSERT INTO country (id, name) SELECT ?, ? WHERE NOT EXISTS (
	SELECT 1 FROM country WHERE id = ?
);`

	// Insert data into managers table if manager ID does not exist
	insertManagerSQL := `INSERT INTO manager (id, name, nickname, dob, country_id) SELECT ?, ?, ?, ?, ? WHERE NOT EXISTS (
	SELECT 1 FROM manager WHERE id = ?
);`

	// Insert data into metadata table if metadata ID does not exist
	insertMetadataSQL := `INSERT INTO metadata (id, data_version, shot_fidelity_version, xy_fidelity_version) SELECT ?, ?, ?, ? WHERE NOT EXISTS (
	SELECT 1 FROM metadata WHERE id = ?
);`

	// Insert data into competition stages table if competition stage ID does not exist
	insertCompetitionStageSQL := `INSERT INTO competition_stage (id, name) SELECT ?, ? WHERE NOT EXISTS (
	SELECT 1 FROM competition_stage WHERE id = ?
);`

	// Insert data into stadiums table if stadium ID does not exist
	insertStadiumSQL := `INSERT INTO stadium (id, name, country_id) SELECT ?, ?, ? WHERE NOT EXISTS (
	SELECT 1 FROM stadium WHERE id = ?
);`

	// Insert data into referees table if referee ID does not exist
	insertRefereeSQL := `INSERT INTO referee (id, name, country_id) SELECT ?, ?, ? WHERE NOT EXISTS (
	SELECT 1 FROM referee WHERE id = ?
);`

	// Insert the data into match table if match ID does not exist
	insertMatchSQL := `INSERT INTO match (
		id, date, kick_off, competition_id, season_id, home_team_id, away_team_id, home_score, away_score, status, status_360, last_updated, last_updated_360, week, competition_stage_id, stadium_id, referee_id
	) SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? WHERE NOT EXISTS (
	SELECT 1 FROM match WHERE match_id = ?
);`
	for _, match := range matches {
		insertSQLData(db, insertCompetitionSQL, match.Competition.CompetitionID, match.Competition.CompetitionName, match.Competition.CountryName, match.Competition.CompetitionID)
		insertSQLData(db, insertSeasonSQL, match.Season.SeasonID, match.Season.SeasonName, match.Season.SeasonID)
		insertSQLData(db, insertTeamSQL, match.HomeTeam.HomeTeamID, match.HomeTeam.HomeTeamName, match.HomeTeam.HomeTeamGender, match.HomeTeam.HomeTeamGroup, match.HomeTeam.Country.ID, match.HomeTeam.HomeTeamID)
		insertSQLData(db, insertTeamSQL, match.AwayTeam.AwayTeamID, match.AwayTeam.AwayTeamName, match.AwayTeam.AwayTeamGender, match.AwayTeam.AwayTeamGroup, match.AwayTeam.Country.ID, match.AwayTeam.AwayTeamID)
		insertSQLData(db, insertCountrySQL, match.HomeTeam.Country.ID, match.HomeTeam.Country.Name, match.HomeTeam.Country.ID)
		insertSQLData(db, insertCountrySQL, match.AwayTeam.Country.ID, match.AwayTeam.Country.Name, match.AwayTeam.Country.ID)
		for _, manager := range match.HomeTeam.Managers {
			insertSQLData(db, insertManagerSQL, manager.ID, manager.Name, manager.Nickname, manager.Dob, manager.Country.ID, manager.ID)
			insertSQLData(db, insertCountrySQL, manager.Country.ID, manager.Country.Name, manager.Country.ID)
		}
		for _, manager := range match.AwayTeam.Managers {
			insertSQLData(db, insertManagerSQL, manager.ID, manager.Name, manager.Nickname, manager.Dob, manager.Country.ID, manager.ID)
			insertSQLData(db, insertCountrySQL, manager.Country.ID, manager.Country.Name, manager.Country.ID)
		}
		insertSQLData(db, insertMetadataSQL, match.Metadata.DataVersion, match.Metadata.ShotFidelityVersion, match.Metadata.XyFidelityVersion, match.Metadata.DataVersion)
		insertSQLData(db, insertCompetitionStageSQL, match.CompetitionStage.ID, match.CompetitionStage.Name, match.CompetitionStage.ID)
		insertSQLData(db, insertStadiumSQL, match.Stadium.ID, match.Stadium.Name, match.Stadium.Country.ID, match.Stadium.ID)
		insertSQLData(db, insertCountrySQL, match.Stadium.Country.ID, match.Stadium.Country.Name, match.Stadium.Country.ID)
		insertSQLData(db, insertRefereeSQL, match.Referee.ID, match.Referee.Name, match.Referee.Country.ID, match.Referee.ID)
		insertSQLData(db, insertCountrySQL, match.Referee.Country.ID, match.Referee.Country.Name, match.Referee.Country.ID)
		insertSQLData(db, insertMatchSQL, match.MatchID, match.MatchDate, match.KickOff, match.Competition.CompetitionID, match.Season.SeasonID, match.HomeTeam.HomeTeamID, match.AwayTeam.AwayTeamID, match.HomeScore, match.AwayScore, match.MatchStatus, match.MatchStatus360, match.LastUpdated, match.LastUpdated360, match.MatchWeek, match.CompetitionStage.ID, match.Stadium.ID, match.Referee.ID, match.MatchID)
	}
}
