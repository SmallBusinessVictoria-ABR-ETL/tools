6931


-- SELECT count(*),pid FROM "sbv_abr"."location" where bus_locn_indy_clsn = '6931' group by pid having count(*) > 1 limit 10;

-- SELECT count(*) from (SELECT count(*),pid FROM "sbv_abr"."location" where bus_locn_indy_clsn = '6931' group by pid having count(*) > 1) as multi_location


SELECT * from agency_combined ac JOIN "sbv_abr"."location" bl ON ac.pid=bl.pid where bus_locn_indy_clsn = '6931' and ac.date = date('2018-09-03');
SELECT * from agency_combined ac JOIN "sbv_abr"."location" bl ON ac.pid=bl.pid where bus_locn_indy_clsn = '6932' and ac.date = date('2018-09-03');
SELECT * from agency_combined ac JOIN "sbv_abr"."location" bl ON ac.pid=bl.pid where bus_locn_indy_clsn = '7291' and ac.date = date('2018-09-03');
