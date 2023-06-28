-- SELECT type,id,start_date,start_date_unix,name,distance,total_elevation_gain,client_id,start_date_unix FROM db_activities LIMIT 4;
SELECT id,start_date,name,distance,total_elevation_gain,client_id,start_date_unix FROM db_activities ORDER BY start_date_unix DESC LIMIT 4;
-- SELECT id,start_date,name,distance,total_elevation_gain,client_id,start_date_unix FROM db_activities WHERE start_date_unix>1442326646 LIMIT 4;
-- SELECT client_id, token FROM db_clients;