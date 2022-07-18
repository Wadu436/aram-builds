ALTER TABLE builds
ALTER COLUMN runesPrimarySelections DROP NOT NULL;
ALTER TABLE builds
ALTER COLUMN runesSecondarySelections DROP NOT NULL;
ALTER TABLE builds
ALTER COLUMN runesStats DROP NOT NULL;