-- SELECT type,id,start_date,start_date_unix,name,distance,total_elevation_gain,client_id,start_date_unix FROM db_activities LIMIT 4;
-- SELECT id,start_date,name,distance,total_elevation_gain,client_id,start_date_unix FROM db_activities ORDER BY start_date_unix DESC LIMIT 4;
-- SELECT id,start_date,name,distance,total_elevation_gain,client_id,start_date_unix FROM db_activities WHERE start_date_unix>1442326646 LIMIT 4;
-- SELECT client_id, token FROM db_clients;
ALTER TABLE "db_activities" DROP "athlet_id";
ALTER TABLE "db_activities" DROP "athlete_id";
ALTER TABLE "db_activities" ADD "athlete_id" bigint;
UPDATE db_activities SET athlete_id = 7845894;
ALTER TABLE "db_activities" ALTER COLUMN "athlete_id" bigint NOT NULL;
SELECT id,athlete_id FROM db_activities ORDER BY start_date_unix DESC LIMIT 4;


DROP TABLE "db_athletes";
DROP TABLE "db_clients";