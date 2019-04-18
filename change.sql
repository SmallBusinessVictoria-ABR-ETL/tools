SELECT n.pid, n.abn
n.ent_typ_cd
n.org_nm
n.nm_titl_cd
n.prsn_gvn_nm
n.prsn_othr_gvn_nm
n.prsn_fmly_nm
n.nm_sufx_cd
n.abn_regn_dt
n.abn_cancn_dt
n.mn_trdg_nm
n.son_addr_ln_1
n.son_addr_ln_2
n.son_sbrb
n.son_stt
n.son_pc
n.son_cntry_cd
n.son_dpid
n.mn_bus_addr_ln_1
n.mn_bus_addr_ln_2
n.mn_bus_sbrb
n.mn_bus_stt
n.mn_bus_pc
n.mn_bus_cntry_cd
n.mn_bus_dpid
n.ent_eml
n.prty_id_blnk
n.gst_regn_dt
n.gst_cancn_dt
n.mn_indy_clsn
n.mn_indy_clsn_descn
n.acn
n.sprsn_ind string
FROM abr_weekly_agency_data n
JOIN abr_weekly_agency_data n2 on n.pid=n2.pid

WHERE n.importdate='{{.Arg1}}'
AND n2.importdate='{{.Arg2}}'

and (n.abn != n2.abn OR
n.ent_typ_cd != n2.ent_typ_cd OR
n.org_nm != n2.org_nm OR
n.nm_titl_cd != n2.nm_titl_cd OR
n.prsn_gvn_nm != n2.prsn_gvn_nm OR
n.prsn_othr_gvn_nm != n2.prsn_othr_gvn_nm OR
n.prsn_fmly_nm != n2.prsn_fmly_nm OR
n.nm_sufx_cd != n2.nm_sufx_cd OR
n.abn_regn_dt != n2.abn_regn_dt OR
n.abn_cancn_dt != n2.abn_cancn_dt OR
n.mn_trdg_nm != n2.mn_trdg_nm OR
n.son_addr_ln_1 != n2.son_addr_ln_1 OR
n.son_addr_ln_2 != n2.son_addr_ln_2 OR
n.son_sbrb != n2.son_sbrb OR
n.son_stt != n2.son_stt OR
n.son_pc != n2.son_pc OR
n.son_cntry_cd != n2.son_cntry_cd OR
n.son_dpid != n2.son_dpid OR
n.mn_bus_addr_ln_1 != n2.mn_bus_addr_ln_1 OR
n.mn_bus_addr_ln_2 != n2.mn_bus_addr_ln_2 OR
n.mn_bus_sbrb != n2.mn_bus_sbrb OR
n.mn_bus_stt != n2.mn_bus_stt OR
n.mn_bus_pc != n2.mn_bus_pc OR
n.mn_bus_cntry_cd != n2.mn_bus_cntry_cd OR
n.mn_bus_dpid != n2.mn_bus_dpid OR
n.ent_eml != n2.ent_eml OR
n.prty_id_blnk != n2.prty_id_blnk OR
n.gst_regn_dt != n2.gst_regn_dt OR
n.gst_cancn_dt != n2.gst_cancn_dt OR
n.mn_indy_clsn != n2.mn_indy_clsn OR
n.mn_indy_clsn_descn != n2.mn_indy_clsn_descn OR
n.acn != n2.acn OR
n.sprsn_ind != n2.sprsn_ind)