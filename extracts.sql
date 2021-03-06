-- noinspection SqlNoDataSourceInspectionForFile

SELECT

a.pid,
a.abn,
a.ent_typ_cd,
a.org_nm,
a.prev_org_nm,
try(date_format(a.org_nm_change_date,'%d%m%Y')) as org_nm_change_date,
a.nm_titl_cd,
a.prsn_gvn_nm,
a.prsn_othr_gvn_nm,
a.prsn_fmly_nm,
a.nm_sufx_cd,
a.prev_nm_titl_cd,
a.prev_prsn_gvn_nm,
a.prev_prsn_othr_gvn_nm,
a.prev_prsn_fmly_nm,
a.prev_nm_sufx_cd,
try(date_format(a.nm_change_date,'%d%m%Y')) as nm_change_date,
try(date_format(date_parse(a.abn_regn_dt, '%Y%m%d'),'%d%m%Y')) as abn_regn_dt,
try(date_format(date_parse(a.abn_cancn_dt, '%Y%m%d'),'%d%m%Y')) as abn_cancn_dt,
if(a.abn_cancn_dt = '','Active','Cancelled') as ABN_Status__c,
a.mn_trdg_nm,
a.prev_mn_trdg_nm,
try(date_format(a.mn_trdg_change_date,'%d%m%Y')) as mn_trdg_change_date,
a.son_addr_ln_1,
a.son_addr_ln_2,
a.son_sbrb,
a.son_stt,
a.son_pc,
a.son_cntry_cd,
a.son_dpid,
a.prev_son_addr_ln_1,
a.prev_son_addr_ln_2,
a.prev_son_sbrb,
a.prev_son_stt,
a.prev_son_pc,
a.prev_son_cntry_cd,
a.prev_son_dpid,
try(date_format(a.son_change_date,'%d%m%Y')) as son_change_date,
a.mn_bus_addr_ln_1,
a.mn_bus_addr_ln_2,
a.mn_bus_sbrb,
a.mn_bus_stt,
a.mn_bus_pc,
a.mn_bus_cntry_cd,
a.mn_bus_dpid,
a.prev_mn_bus_addr_ln_1,
a.prev_mn_bus_addr_ln_2,
a.prev_mn_bus_sbrb,
a.prev_mn_bus_stt,
a.prev_mn_bus_pc,
a.prev_mn_bus_cntry_cd,
a.prev_mn_bus_dpid,
try(date_format(a.mn_bus_change_date,'%d%m%Y')) as mn_bus_change_date,
a.ent_eml,
a.prev_ent_eml,
try(date_format(a.ent_eml_change_date,'%d%m%Y')) as ent_eml_change_date,
a.prty_id_blnk,
try(date_format(date_parse(a.gst_regn_dt, '%Y%m%d'),'%d%m%Y')) as gst_regn_dt,
try(date_format(date_parse(a.gst_cancn_dt, '%Y%m%d'),'%d%m%Y')) as gst_cancn_dt,
a.mn_indy_clsn,
a.mn_indy_clsn_descn,
a.prev_mn_indy_clsn,
a.prev_mn_indy_clsn_descn,
try(date_format(a.mn_indy_change_date,'%d%m%Y')) as mn_indy_change_date,
try_cast(a.acn as integer) as acn,
a.sprsn_ind,


bl.locn_typ_cd,
try(date_format(date_parse(bl.locn_strt_dt, '%Y%m%d'),'%d%m%Y')) as locn_strt_dt,
bl.bus_locn_addr_ln_1,
bl.bus_locn_addr_ln_2,
bl.bus_locn_sbrb,
bl.bus_locn_stt,
bl.bus_locn_pc,
bl.bus_locn_cntry_cd,
bl.bus_locn_dpid,
bl.bus_locn_ltd,
bl.bus_locn_lngtd,
bl.bus_locn_msh_blk,
bl.bus_locn_gnaf_pid,
bl.bus_locn_posnl_rlblty,
bl.bus_locn_ph_area_cd,
bl.bus_locn_ph_num,
bl.bus_locn_ph_area_cd_mbl,
bl.bus_locn_ph_num_mbl,
concat(coalesce (bl.bus_locn_ph_area_cd,''),coalesce (bl.bus_locn_ph_num,'')) as Phone,
concat(coalesce (bl.bus_locn_ph_area_cd_mbl,''),coalesce (bl.bus_locn_ph_num_mbl,'')) as Emergency_mobile_phone,
bl.bus_locn_eml,
bl.bus_locn_indy_clsn,
bl.bus_locn_indy_clsn_descn,
-- bl.lga_code_2016,
-- bl.state_code_2016,
-- bl.state_name_2016,
-- bl.area_albers_sqkm,
-- bl.local_government_area,
-- bl.vgbo_region,
-- bl.like_councils,

bn.bus_nm


FROM agency_combined a
  JOIN location bl ON a.pid = bl.pid
  JOIN (SELECT array_join(array_agg(bus_nm), ', ') bus_nm,pid from businessname group by pid) bn ON a.pid = bn.pid
 WHERE
       a.ent_typ_cd in ('IND', 'LPT', 'OIE', 'PRV', 'PTR', 'PTT', 'PUB', 'STR', 'UIE')
   AND bl.locn_typ_cd = '010'
  AND bl.bus_locn_indy_clsn = '6931'

limit 1;