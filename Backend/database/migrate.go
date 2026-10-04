package database

import (
	"fmt"
)

func CreateTable() {
	fmt.Println("Database table created")

	table := `CREATE TABLE IF NOT EXISTS users (
    email TEXT PRIMARY KEY,
    password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    refresh_token TEXT DEFAULT '',
	roles TEXT[] DEFAULT '{user}'
);`

	userInfoTable := `CREATE TABLE IF NOT EXISTS user_info (
	email TEXT PRIMARY KEY,
	user_name TEXT UNIQUE NOT NULL,
	age INT,
	gender TEXT,
	level INT,
	profile_picture_url TEXT,
	profile_picture_id TEXT,
	playing_style TEXT[] DEFAULT '{}',                                                               
	games TEXT[] DEFAULT '{}',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

	gameTable := `CREATE TABLE IF NOT EXISTS games (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    cover_id INT,
    cover_image_id TEXT,
    cover_url TEXT,
    categories TEXT[] DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

	seedGames := `INSERT INTO games (id, name, cover_id, cover_image_id, cover_url, categories) VALUES
(5589, 'Clash of Clans', 75572, 'co1mb8', '//images.igdb.com/igdb/image/upload/t_thumb/co1mb8.jpg', ARRAY['Real Time Strategy (RTS)', 'Strategy']),
(13166, 'The Witcher 3: Wild Hunt - Blood and Wine', 121128, 'co2lgo', '//images.igdb.com/igdb/image/upload/t_thumb/co2lgo.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(119171, 'Baldur''s Gate III', 289025, 'co670h', '//images.igdb.com/igdb/image/upload/t_thumb/co670h.jpg', ARRAY['Role-playing (RPG)', 'Strategy', 'Turn-based strategy (TBS)']),
(2985, 'Persona 4 Golden', 76533, 'co1n1x', '//images.igdb.com/igdb/image/upload/t_thumb/co1n1x.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(141540, 'Disco Elysium: The Final Cut', 133993, 'co2ve1', '//images.igdb.com/igdb/image/upload/t_thumb/co2ve1.jpg', ARRAY['Role-playing (RPG)', 'Adventure', 'Indie']),
(114283, 'Persona 5 Royal', 587629, 'coclf1', '//images.igdb.com/igdb/image/upload/t_thumb/coclf1.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(14647, 'Bloodborne: The Old Hunters', 525224, 'cob99k', '//images.igdb.com/igdb/image/upload/t_thumb/cob99k.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(1942, 'The Witcher 3: Wild Hunt', 480513, 'coaarl', '//images.igdb.com/igdb/image/upload/t_thumb/coaarl.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(298526, 'Kingdom Come: Deliverance II', 407767, 'co8qmv', '//images.igdb.com/igdb/image/upload/t_thumb/co8qmv.jpg', ARRAY['Role-playing (RPG)', 'Simulator', 'Adventure']),
(119133, 'Elden Ring', 212094, 'co4jni', '//images.igdb.com/igdb/image/upload/t_thumb/co4jni.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(6036, 'The Last of Us Remastered', 279388, 'co5zks', '//images.igdb.com/igdb/image/upload/t_thumb/co5zks.jpg', ARRAY['Shooter', 'Adventure']),
(25076, 'Red Dead Redemption 2', 80403, 'co1q1f', '//images.igdb.com/igdb/image/upload/t_thumb/co1q1f.jpg', ARRAY['Shooter', 'Role-playing (RPG)', 'Adventure']),
(1802, 'Chrono Trigger', 173156, 'co3plw', '//images.igdb.com/igdb/image/upload/t_thumb/co3plw.jpg', ARRAY['Role-playing (RPG)']),
(204350, 'The Last of Us Part I', 468458, 'coa1gq', '//images.igdb.com/igdb/image/upload/t_thumb/coa1gq.jpg', ARRAY['Shooter', 'Adventure']),
(1009, 'The Last of Us', 81915, 'co1r7f', '//images.igdb.com/igdb/image/upload/t_thumb/co1r7f.jpg', ARRAY['Shooter', 'Adventure']),
(126098, 'Half-Life: Alyx', 383452, 'co87vg', '//images.igdb.com/igdb/image/upload/t_thumb/co87vg.jpg', ARRAY['Shooter', 'Puzzle', 'Adventure']),
(119388, 'The Legend of Zelda: Tears of the Kingdom', 274264, 'co5vmg', '//images.igdb.com/igdb/image/upload/t_thumb/co5vmg.jpg', ARRAY['Puzzle', 'Adventure']),
(379, 'Metal Gear Solid 3: Snake Eater', 252077, 'co5ei5', '//images.igdb.com/igdb/image/upload/t_thumb/co5ei5.jpg', ARRAY['Shooter', 'Tactical', 'Adventure']),
(26472, 'Disco Elysium', 83503, 'co1sfj', '//images.igdb.com/igdb/image/upload/t_thumb/co1sfj.jpg', ARRAY['Role-playing (RPG)', 'Adventure', 'Indie']),
(115289, 'Hollow Knight: Silksong', 531786, 'cobebu', '//images.igdb.com/igdb/image/upload/t_thumb/cobebu.jpg', ARRAY['Platform', 'Adventure', 'Indie']),
(7346, 'The Legend of Zelda: Breath of the Wild', 172453, 'co3p2d', '//images.igdb.com/igdb/image/upload/t_thumb/co3p2d.jpg', ARRAY['Puzzle', 'Adventure']),
(9927, 'Persona 5', 81906, 'co1r76', '//images.igdb.com/igdb/image/upload/t_thumb/co1r76.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(45181, 'Mass Effect Trilogy', 199846, 'co4a7a', '//images.igdb.com/igdb/image/upload/t_thumb/co4a7a.jpg', ARRAY['Shooter', 'Role-playing (RPG)', 'Adventure']),
(1103, 'Super Metroid', 265426, 'co5osy', '//images.igdb.com/igdb/image/upload/t_thumb/co5osy.jpg', ARRAY['Shooter', 'Platform', 'Adventure']),
(3349, 'Paper Mario: The Thousand-Year Door', 100900, 'co25us', '//images.igdb.com/igdb/image/upload/t_thumb/co25us.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(479, 'Metal Gear Solid 3: Subsistence', 87306, 'co1vd6', '//images.igdb.com/igdb/image/upload/t_thumb/co1vd6.jpg', ARRAY['Shooter', 'Tactical', 'Adventure']),
(1128, 'Castlevania: Symphony of the Night', 237968, 'co53m8', '//images.igdb.com/igdb/image/upload/t_thumb/co53m8.jpg', ARRAY['Platform', 'Role-playing (RPG)', 'Adventure']),
(277143, 'The Last of Us Part II Remastered', 468459, 'coa1gr', '//images.igdb.com/igdb/image/upload/t_thumb/coa1gr.jpg', ARRAY['Shooter', 'Adventure']),
(19560, 'God of War', 540186, 'cobkt6', '//images.igdb.com/igdb/image/upload/t_thumb/cobkt6.jpg', ARRAY['Role-playing (RPG)', 'Hack and slash/Beat ''em up', 'Adventure']),
(1026, 'The Legend of Zelda: A Link to the Past', 181427, 'co3vzn', '//images.igdb.com/igdb/image/upload/t_thumb/co3vzn.jpg', ARRAY['Puzzle', 'Adventure']),
(28337, '13 Sentinels: Aegis Rim', 132976, 'co2uls', '//images.igdb.com/igdb/image/upload/t_thumb/co2uls.jpg', ARRAY['Simulator', 'Strategy', 'Tactical', 'Adventure']),
(12503, 'The Witcher 3: Wild Hunt - Hearts of Stone', 84518, 'co1t7q', '//images.igdb.com/igdb/image/upload/t_thumb/co1t7q.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(14593, 'Hollow Knight', 533941, 'cobfzp', '//images.igdb.com/igdb/image/upload/t_thumb/cobfzp.jpg', ARRAY['Platform', 'Adventure', 'Indie']),
(305152, 'Clair Obscur: Expedition 33', 441022, 'co9gam', '//images.igdb.com/igdb/image/upload/t_thumb/co9gam.jpg', ARRAY['Role-playing (RPG)', 'Adventure', 'Indie']),
(1070, 'Super Mario World', 401336, 'co8lo8', '//images.igdb.com/igdb/image/upload/t_thumb/co8lo8.jpg', ARRAY['Platform', 'Adventure']),
(1454, 'Suikoden II', 90946, 'co1y6a', '//images.igdb.com/igdb/image/upload/t_thumb/co1y6a.jpg', ARRAY['Role-playing (RPG)', 'Strategy', 'Turn-based strategy (TBS)']),
(1105, 'Metroid Prime', 181616, 'co3w4w', '//images.igdb.com/igdb/image/upload/t_thumb/co3w4w.jpg', ARRAY['Shooter', 'Platform', 'Adventure']),
(26916, 'Dark Souls III: The Ringed City', 104246, 'co28fq', '//images.igdb.com/igdb/image/upload/t_thumb/co28fq.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(1029, 'The Legend of Zelda: Ocarina of Time', 170637, 'co3nnx', '//images.igdb.com/igdb/image/upload/t_thumb/co3nnx.jpg', ARRAY['Puzzle', 'Adventure']),
(7334, 'Bloodborne', 525225, 'cob99l', '//images.igdb.com/igdb/image/upload/t_thumb/cob99l.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(119295, 'Microsoft Flight Simulator', 111116, 'co2dqk', '//images.igdb.com/igdb/image/upload/t_thumb/co2dqk.jpg', ARRAY['Simulator']),
(9577, 'Persona 3', 262662, 'co5mo6', '//images.igdb.com/igdb/image/upload/t_thumb/co5mo6.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(227211, 'S.T.A.L.K.E.R. G.A.M.M.A.', 271147, 'co5t7v', '//images.igdb.com/igdb/image/upload/t_thumb/co5t7v.jpg', ARRAY['Shooter', 'Tactical', 'Adventure', 'Indie']),
(7194, 'Steins;Gate', 253267, 'co5ff7', '//images.igdb.com/igdb/image/upload/t_thumb/co5ff7.jpg', ARRAY['Adventure', 'Visual Novel']),
(74, 'Mass Effect 2', 93684, 'co20ac', '//images.igdb.com/igdb/image/upload/t_thumb/co20ac.jpg', ARRAY['Shooter', 'Role-playing (RPG)', 'Adventure']),
(72, 'Portal 2', 82660, 'co1rs4', '//images.igdb.com/igdb/image/upload/t_thumb/co1rs4.jpg', ARRAY['Platform', 'Puzzle', 'Adventure']),
(481, 'Silent Hill 2', 134728, 'co2vyg', '//images.igdb.com/igdb/image/upload/t_thumb/co2vyg.jpg', ARRAY['Puzzle', 'Adventure']),
(215769, 'Cyberpunk 2077: Phantom Liberty', 333294, 'co7566', '//images.igdb.com/igdb/image/upload/t_thumb/co7566.jpg', ARRAY['Shooter', 'Role-playing (RPG)', 'Adventure']),
(112875, 'God of War Ragnarök', 526297, 'coba3d', '//images.igdb.com/igdb/image/upload/t_thumb/coba3d.jpg', ARRAY['Role-playing (RPG)', 'Hack and slash/Beat ''em up', 'Adventure']),
(227, 'World of Warcraft: The Burning Crusade', 90582, 'co1xw6', '//images.igdb.com/igdb/image/upload/t_thumb/co1xw6.jpg', ARRAY['Role-playing (RPG)', 'Adventure']),
(228530, 'Death Stranding 2: On the Beach', 444165, 'co9ipx', '//images.igdb.com/igdb/image/upload/t_thumb/co9ipx.jpg', ARRAY['Shooter', 'Adventure'])
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    cover_id = EXCLUDED.cover_id,
    cover_image_id = EXCLUDED.cover_image_id,
    cover_url = EXCLUDED.cover_url,
    categories = EXCLUDED.categories,
    updated_at = CURRENT_TIMESTAMP;`

	if _, err := Conn.Exec(Ctx, table); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := Conn.Exec(Ctx, userInfoTable); err != nil {
		fmt.Println(err)
	}

	if _, err := Conn.Exec(Ctx, gameTable); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := Conn.Exec(Ctx, seedGames); err != nil {
		fmt.Println("Error seeding games:", err)
	} else {
		fmt.Println("Games table created and seeded successfully!")
	}

	fmt.Println("Database setup complete")

}
