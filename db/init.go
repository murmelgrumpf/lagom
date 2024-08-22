package db

import "os/exec"

func InitDB() {

	emojiTableName := QueryPanic("SELECT name FROM sqlite_master WHERE type='table' AND name='emojis'")
	if !emojiTableName.Next() {
		PreparePanic(`CREATE TABLE emojis (
            emoji VARCHAR(40) ,
            hexcode VARCHAR(40) PRIMARY KEY,
            "group" VARCHAR(100),
            subgroups VARCHAR(100),
            annotation VARCHAR(40),
            tags VARCHAR(200),
            openmoji_tags VARCHAR(200)
        )`).Exec()
		cmd := exec.Command("./import_emojis")
		_, err := cmd.Output()
		if err != nil {
			panic(err)
		}
	}

	// Games
	PreparePanic(`CREATE TABLE IF NOT EXISTS games (
        id INTEGER PRIMARY KEY, 
        name VARCHAR(200) NOT NULL
    )`).Exec()

	// Places
	PreparePanic(`CREATE TABLE IF NOT EXISTS places (
        id INTEGER PRIMARY KEY,
        name VARCHAR(200) NOT NULL
    )`).Exec()

	PreparePanic(`CREATE TABLE IF NOT EXISTS game_places (
        game_id INTEGER NOT NULL,
        place_id INTEGER NOT NULL,
        PRIMARY KEY(game_id, place_id),
        FOREIGN KEY(game_id) REFERENCES games(id),
        FOREIGN KEY(place_id) REFERENCES places(id)
    )`).Exec()

	// Resources
	PreparePanic(`CREATE TABLE IF NOT EXISTS resources (
        id INTEGER PRIMARY KEY,
        name VARCHAR(200) NOT NULL,
        expiration INT,
        emoji VARCHAR(20)
    )`).Exec()

	PreparePanic(`CREATE TABLE IF NOT EXISTS game_resources (
        game_id INTEGER NOT NULL,
        resource_id INTEGER NOT NULL,
        PRIMARY KEY(game_id, resource_id),
        FOREIGN KEY(game_id) REFERENCES games(id),
        FOREIGN KEY(resource_id) REFERENCES resources(id)
    )`).Exec()

	// Feature
	//PreparePanic(`CREATE TABLE IF NOT EXISTS features (
	//    id INTEGER PRIMARY KEY,
	//    name VARCHAR(200) NOT NULL,
	//)`).Exec()

	//PreparePanic(`CREATE TABLE IF NOT EXISTS game_resources (
	//    game_id INTEGER NOT NULL,
	//    resource_id INTEGER NOT NULL,
	//    PRIMARY KEY(game_id, resource_id),
	//    FOREIGN KEY(game_id) REFERENCES games(id),
	//    FOREIGN KEY(resource_id) REFERENCES resources(id)
	//)`).Exec()
	// rows, _ := Query("SELECT id, firstname, lastname, socialsecurity FROM people")
	// var id int
	// var firstname string
	// var lastname string
	// var socialsecurity string
	//
	//	for rows.Next() {
	//		rows.Scan(&id, &firstname, &lastname, &socialsecurity)
	//		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname + " " + socialsecurity)
	//	}
}
