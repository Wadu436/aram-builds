CREATE TABLE IF NOT EXISTS builds (
    -- Metadata
    champion VARCHAR(32) NOT NULL,
    game_version_major INT NOT NULL,
    game_version_minor INT NOT NULL,
    mtime TIMESTAMP NOT NULL,
    -- Runes
    runesPrimaryKey VARCHAR(16) NOT NULL,
    runesPrimarySelections INT [4] NOT NULL,
    runesSecondaryKey VARCHAR(16) NOT NULL,
    runesSecondarySelections INT [3] NOT NULL,
    runesStats INT [3] NOT NULL,
    -- Items
    itemsStart VARCHAR(8) [] NOT NULL,
    itemsFullBuild VARCHAR(8) [] NOT NULL,
    PRIMARY KEY (champion, game_version_major, game_version_minor)
);