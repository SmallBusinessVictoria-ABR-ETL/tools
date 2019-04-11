-- noinspection SqlNoDataSourceInspectionForFile

SELECT
ass.*

FROM agency a
  JOIN associates ass ON a.pid=ass.pid
  JOIN (select distinct(pid)
            FROM location
            WHERE locn_typ_cd = '010'
                AND bus_locn_indy_clsn = '{{.Arg1}}'
    ) bl ON a.pid = bl.pid
 WHERE
       a.ent_typ_cd in ('IND', 'LPT', 'OIE', 'PRV', 'PTR', 'PTT', 'PUB', 'STR', 'UIE')

