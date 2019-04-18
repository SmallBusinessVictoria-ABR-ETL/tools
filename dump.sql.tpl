-- noinspection SqlNoDataSourceInspectionForFile

SELECT

a.*,
bl.*

FROM agency_tmp a
  JOIN location bl ON a.pid = bl.pid

