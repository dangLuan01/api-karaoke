package models

type Song struct {
	Uuid	  	string `db:"uuid"`
	Id 			string `db:"_id"`
	Name	  	string `db:"name"`
	Namesort 	string `db:"name_sort"`
	Singer   	*string `db:"singer" goqu:"omitnil"`
	Author   	*string `db:"author" goqu:"omitnil"`
	Tone	 	*string `db:"tone" goqu:"omitnil"`
	Tune	 	*string `db:"tune" goqu:"omitnil"`
	Thumbail 	*string `db:"thumbnail" goqu:"omitnil"`
	Lyrics   	*string `db:"lyrics" goqu:"omitnil"`
	Type	  	*string `db:"type" goqu:"omitnil"`
}

type Image struct {
	Songuuid 	string `db:"song_uuid"`
	Id 			string `db:"_id"`
	Image 		string `db:"image"`
}