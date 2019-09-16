CREATE DATABASE lair;

CREATE TABLE lair.users (
    -- ID
	id BIGINT NOT NULL AUTO_INCREMENT,

    -- FIELDS
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,

    -- TIME STAMPS
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,

    -- PRIMARY KEY
    PRIMARY KEY (id)
);

CREATE TABLE lair.campaigns (
    -- IDS
    id BIGINT NOT NULL AUTO_INCREMENT,
    dm_id BIGINT,

    -- FIELDS
    name VARCHAR(255) NOT NULL,

    -- TIME STAMPS
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,

    -- KEYS
    PRIMARY KEY (id),
    FOREIGN KEY (dm_id) REFERENCES lair.users(id)
);

CREATE TABLE lair.players (
    -- IDS
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT,
    campaign_id BIGINT,

    -- TIME STAMPS
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,

    -- KEYS
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES lair.users(id),
    FOREIGN KEY (campaign_id) REFERENCES lair.campaigns(id)
);

CREATE TABLE lair.characters (
    -- IDS
    id BIGINT NOT NULL AUTO_INCREMENT,
    player_id BIGINT,
    
    -- FIELDS
    name VARCHAR(255) NOT NULL,
    race VARCHAR(255) NOT NULL,
    level INT NOT NULL,
    size VARCHAR(255),
    alignment VARCHAR(255),
    ac INT NOT NULL,
    armor VARCHAR(255),
    hit_points INT NOT NULL,
    hit_die VARCHAR(255) NOT NULL,
    speed INT,
    initiative INT NOT NULL,
    str VARCHAR(255) NOT NULL,
    dex VARCHAR(255) NOT NULL,
    con VARCHAR(255) NOT NULL,
    itl VARCHAR(255) NOT NULL,
    wis VARCHAR(255) NOT NULL,
    cha VARCHAR(255) NOT NULL,

    -- TIME STAMPS
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,

    -- PRIMARY AND FOREIGN KEYS
    PRIMARY KEY (id),
    FOREIGN KEY (player_id) REFERENCES lair.players(id)
);