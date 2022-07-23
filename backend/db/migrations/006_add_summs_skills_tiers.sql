-- Modify the table
ALTER TABLE builds
ADD summoners varchar(64) [2],
    ADD skill_order INT [18],
    ADD tier INT;
-- Default Summoners:
-- Snowball and Flash
UPDATE builds
SET summoners = '{ "SummonerSnowball", "SummonerFlash" }';
-- Set summoners required
ALTER TABLE builds
ALTER COLUMN summoners
SET NOT NULL;