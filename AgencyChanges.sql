-- Load agency partitions
MSCK REPAIR TABLE agency_v2;

-- Get Samples
SELECT * from agency_v2 where "date"=DATE('2018-08-30') and state='previous' limit 10;
SELECT * from agency_v2 where "date"=DATE('2018-08-30') and state='update' limit 10;
SELECT * from agency_v2 where "date"=DATE('2018-08-30') and state='new' limit 10;


SELECT count(*) from agency_v2 where "date"=DATE('2018-08-30') and state='previous';
SELECT count(*) from agency_v2 where "date"=DATE('2018-08-30') and state='update';
SELECT count(*) from agency_v2 where "date"=DATE('2018-08-30') and state='new';

SELECT a.pid, a.org_nm as prev_org_nm, b.org_nm FROM agency_v2 a JOIN agency_v2 b ON a.pid = b.pid WHERE a."date"=DATE('2018-08-30') and a.state='previous' and b."date"=DATE('2018-08-30') and b.state='update' and a.org_nm != b.org_nm