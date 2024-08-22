package db

import (
	"database/sql"
)

var searchEmojisQuery *sql.Stmt

type Emojis struct {
	Emoji        string
	Group        string
	Subgroups    string
	Annotation   string
	Tags         string
	OpenmojiTags string
}

func SearchEmojis() *[]Emojis {
	rows, err := GetDb().Query(`
    SELECT 
        emoji,
        [group],
        subgroups,
        annotation,
        tags,
        openmoji_tags
    FROM emojis 
    ORDER BY annotation ASC
    `)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var row Emojis
	var result []Emojis
	for rows.Next() {
		rows.Scan(&row.Emoji, &row.Group, &row.Subgroups, &row.Annotation, &row.Tags, &row.OpenmojiTags)
		result = append(result, row)
	}
	return &result
}
