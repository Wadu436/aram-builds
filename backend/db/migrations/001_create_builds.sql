CREATE TABLE IF NOT EXISTS builds (
    -- Metadata
    champion VARCHAR(32) NOT NULL,
    game_version VARCHAR(8) NOT NULL,
    mtime TIMESTAMP NOT NULL,
    -- Runes
    runesPrimaryKey VARCHAR(16) NOT NULL,
    runesPrimarySelections int [4] NOT NULL,
    runesSecondaryKey VARCHAR(16) NOT NULL,
    runesSecondarySelections int [3] NOT NULL,
    runesStats int [3] NOT NULL,
    -- Items
    itemsStart VARCHAR(8) [] NOT NULL,
    itemsFullBuild VARCHAR(8) [] NOT NULL,
    PRIMARY KEY (champion, game_version)
);