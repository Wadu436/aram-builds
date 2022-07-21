-- Modify the table
ALTER TABLE builds
ADD game_version varchar(8);
-- Fill the new column
UPDATE builds
SET game_version = game_version_major::varchar || '.' || game_version_minor::varchar;
-- Drop old primary key
ALTER TABLE builds DROP CONSTRAINT builds_pkey;
-- Change primary key
ALTER TABLE builds
ADD PRIMARY KEY (champion, game_version);
-- Drop previous game version field
ALTER TABLE builds DROP game_version_major,
    DROP game_version_minor;