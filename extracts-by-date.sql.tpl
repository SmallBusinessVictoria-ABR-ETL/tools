-- noinspection SqlNoDataSourceInspectionForFile

SELECT

a.pid,
a.abn,
a.ent_typ_cd,
a.org_nm,
a.nm_titl_cd,
a.prsn_gvn_nm,
a.prsn_othr_gvn_nm,
a.prsn_fmly_nm,
a.nm_sufx_cd,
a.abn_regn_dt,
a.abn_cancn_dt,
a.mn_trdg_nm,
a.son_addr_ln_1,
a.son_addr_ln_2,
a.son_sbrb,
a.son_stt,
a.son_pc,
a.son_cntry_cd,
a.son_dpid,
a.mn_bus_addr_ln_1,
a.mn_bus_addr_ln_2,
a.mn_bus_sbrb,
a.mn_bus_stt,
a.mn_bus_pc,
a.mn_bus_cntry_cd,
a.mn_bus_dpid,
a.ent_eml,
a.prty_id_blnk,
a.gst_regn_dt,
a.gst_cancn_dt,
a.mn_indy_clsn,
a.mn_indy_clsn_descn,
a.acn,
a.sprsn_ind,

bl.pid,
bl.locn_typ_cd,
bl.locn_strt_dt,
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
bl.bus_locn_eml,
bl.bus_locn_indy_clsn,
bl.bus_locn_indy_clsn_descn,
bl.lga_code_2016,
bl.state_code_2016,
bl.state_name_2016,
bl.area_albers_sqkm,
bl.local_government_area,
bl.vgbo_region,
bl.like_councils,

c.rltnshp_cd,
c.assoc_org_nm,
c.assoc_titl_cd,
c.assoc_gvn_nm,
c.assoc_othr_gvn_nms,
c.assoc_fmly_nm,
c.assoc_nm_sufx_cd

FROM agency_tmp a
  JOIN location bl ON a.pid = bl.pid
  JOIN associates c ON a.pid = c.pid
WHERE
  a.mn_bus_stt = 'VIC'
  AND strpos(coalesce(a.ent_eml, bl.bus_locn_eml),'@') > 1
  AND length(a.abn_regn_dt) = 8
  AND date_parse(a.abn_regn_dt, '%Y%m%d') >= date '{{.Arg1}}' AND date_parse(a.abn_regn_dt, '%Y%m%d') <= date '{{.Arg2}}'
