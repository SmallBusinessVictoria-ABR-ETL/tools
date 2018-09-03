
-- Main diffing query

SELECT
  a.pid,
  a.ABN,
  a.Ent_Typ_Cd,

-- Org Name
  a.org_nm as prev_org_nm,
  coalesce( b.org_nm, a.org_nm) as org_nm,
  if(b.org_nm is not null, b.date, null) as org_nm_change_dt,

-- Person Name
  coalesce (c.Nm_Titl_Cd, a.Nm_Titl_Cd) as Nm_Titl_Cd,
  coalesce (c.Prsn_Gvn_Nm, a.Prsn_Gvn_Nm) as Prsn_Gvn_Nm,
  coalesce (c.Prsn_Othr_Gvn_Nm, a.Prsn_Othr_Gvn_Nm) as Prsn_Othr_Gvn_Nm,
  coalesce (c.Prsn_Fmly_Nm, a.Prsn_Fmly_Nm) as Prsn_Fmly_Nm,
  coalesce (c.Nm_Sufx_Cd, a.Nm_Sufx_Cd) as Nm_Sufx_Cd,

  a.Nm_Titl_Cd as prev_Nm_Titl_Cd,
  a.Prsn_Gvn_Nm as prev_Prsn_Gvn_Nm,
  a.Prsn_Othr_Gvn_Nm as prev_Prsn_Othr_Gvn_Nm,
  a.Prsn_Fmly_Nm as prev_Prsn_Fmly_Nm,
  a.Nm_Sufx_Cd as prev_Nm_Sufx_Cd,
  if(c.Nm_Titl_Cd is not null, c.date, null) as Nm_change_dt,

  a.ABN_Regn_Dt,
  a.ABN_Cancn_Dt,

-- Trading Name
  coalesce (d.Mn_Trdg_Nm, a.Mn_Trdg_Nm) as Mn_Trdg_Nm,
  a.Mn_Trdg_Nm as prev_Mn_Trdg_Nm,
  if(d.Nm_Titl_Cd is not null, d.date, null) as Mn_Trdg_Nm_change_dt,

-- SON Address
  coalesce (e.SON_Addr_Ln_1, a.SON_Addr_Ln_1) as SON_Addr_Ln_1,
  coalesce (e.SON_Addr_Ln_2, a.SON_Addr_Ln_2) as SON_Addr_Ln_2,
  coalesce (e.SON_Sbrb, a.SON_Sbrb) as SON_Sbrb,
  coalesce (e.SON_Stt, a.SON_Stt) as SON_Stt,
  coalesce (e.SON_Pc, a.SON_Pc) as SON_Pc,
  coalesce (e.SON_Cntry_Cd, a.SON_Cntry_Cd) as SON_Cntry_Cd,
  coalesce (e.SON_DPID, a.SON_DPID) as SON_DPID,


  a.SON_Addr_Ln_1 as prev_SON_Addr_Ln_1,
  a.SON_Addr_Ln_2 as prev_SON_Addr_Ln_2,
  a.SON_Sbrb as prev_SON_Sbrb,
  a.SON_Stt as prev_SON_Stt,
  a.SON_Pc as prev_SON_Pc,
  a.SON_Cntry_Cd as prev_SON_Cntry_Cd,
  a.SON_DPID as prev_SON_DPID,

  if(e.SON_Addr_Ln_1 is not null, e.date, null) as SON_change_dt,

-- Mn_Bus Address
  coalesce (f.Mn_Bus_Addr_Ln_1, a.Mn_Bus_Addr_Ln_1) as Mn_Bus_Addr_Ln_1,
  coalesce (f.Mn_Bus_Addr_Ln_2, a.Mn_Bus_Addr_Ln_2) as Mn_Bus_Addr_Ln_2,
  coalesce (f.Mn_Bus_Sbrb, a.Mn_Bus_Sbrb) as Mn_Bus_Sbrb,
  coalesce (f.Mn_Bus_Stt, a.Mn_Bus_Stt) as Mn_Bus_Stt,
  coalesce (f.Mn_Bus_Pc, a.Mn_Bus_Pc) as Mn_Bus_Pc,
  coalesce (f.Mn_Bus_Cntry_Cd, a.Mn_Bus_Cntry_Cd) as Mn_Bus_Cntry_Cd,
  coalesce (f.Mn_Bus_DPID, a.Mn_Bus_DPID) as Mn_Bus_DPID,


  a.Mn_Bus_Addr_Ln_1 as prev_Mn_Bus_Addr_Ln_1,
  a.Mn_Bus_Addr_Ln_2 as prev_Mn_Bus_Addr_Ln_2,
  a.Mn_Bus_Sbrb as prev_Mn_Bus_Sbrb,
  a.Mn_Bus_Stt as prev_Mn_Bus_Stt,
  a.Mn_Bus_Pc as prev_Mn_Bus_Pc,
  a.Mn_Bus_Cntry_Cd as prev_Mn_Bus_Cntry_Cd,
  a.Mn_Bus_DPID as prev_Mn_Bus_DPID,

  if(f.Mn_Bus_Addr_Ln_1 is not null, f.date, null) as Mn_Bus_change_dt,

-- Email
  coalesce (g.Ent_Eml, a.Ent_Eml) as Ent_Eml,
  a.Ent_Eml as prev_Ent_Eml,
  if(g.Nm_Titl_Cd is not null, g.date, null) as Ent_Eml_change_dt,


  a.Prty_Id_Blnk,
  a.GST_Regn_Dt,
  a.GST_Cancn_Dt,

-- Industry
  coalesce (f.Mn_Indy_Clsn, a.Mn_Indy_Clsn) as Mn_Indy_Clsn,
  coalesce (f.Mn_Indy_Clsn_Descn, a.Mn_Indy_Clsn_Descn) as Mn_Indy_Clsn_Descn,

  a.Mn_Indy_Clsn as prev_Mn_Indy_Clsn,
  a.Mn_Indy_Clsn_Descn as Mn_Indy_Clsn_Descn,

  if(f.Mn_Indy_Clsn is not null, f.date, null) as Indy_Clsn_change_dt,

  a.ACN,
  a.Sprsn_Ind

FROM agency_v2 a -- Previous
LEFT JOIN agency_v2 b -- Updates (org_nm)
  ON a.pid = b.pid AND a."date"=DATE('2018-07-16') and a.state='full' and b."date"=DATE('2018-08-30') and b.state='update' and a.org_nm != b.org_nm
LEFT JOIN agency_v2 c -- Updates (name change)
  ON a.pid = c.pid AND a."date"=DATE('2018-07-16') and a.state='full' and c."date"=DATE('2018-08-30') and c.state='update' and (
    a.Nm_Titl_Cd != c.Nm_Titl_Cd
    OR a.Prsn_Gvn_Nm != c.Prsn_Gvn_Nm
    OR a.Prsn_Othr_Gvn_Nm != c.Prsn_Othr_Gvn_Nm
    OR a.Prsn_Fmly_Nm != c.Prsn_Fmly_Nm
    OR a.Nm_Sufx_Cd != c.Nm_Sufx_Cd
  )
LEFT JOIN agency_v2 d -- Updates (Mn_Trdg_Nm Change)
  ON a.pid = d.pid AND a."date"=DATE('2018-07-16') and a.state='full' and d."date"=DATE('2018-08-30') and d.state='update' and a.Mn_Trdg_Nm != d.Mn_Trdg_Nm
LEFT JOIN agency_v2 e -- Updates (SON change)
  ON a.pid = e.pid AND a."date"=DATE('2018-07-16') and a.state='full' and e."date"=DATE('2018-08-30') and e.state='update' and (
    a.SON_Addr_Ln_1 != e.SON_Addr_Ln_1
    OR a.SON_Addr_Ln_2 != e.SON_Addr_Ln_2
    OR a.SON_Sbrb != e.SON_Sbrb
    OR a.SON_Stt != e.SON_Stt
    OR a.SON_Pc != e.SON_Pc
    OR a.SON_Cntry_Cd != e.SON_Cntry_Cd
    OR a.SON_DPID != e.SON_DPID
  )
LEFT JOIN agency_v2 f -- Updates (Mn_Bus change)
  ON a.pid = f.pid AND a."date"=DATE('2018-07-16') and a.state='full' and f."date"=DATE('2018-08-30') and f.state='update' and (
    a.Mn_Bus_Addr_Ln_1 != f.Mn_Bus_Addr_Ln_1
    OR a.Mn_Bus_Addr_Ln_2 != f.Mn_Bus_Addr_Ln_2
    OR a.Mn_Bus_Sbrb != f.Mn_Bus_Sbrb
    OR a.Mn_Bus_Stt != f.Mn_Bus_Stt
    OR a.Mn_Bus_Pc != f.Mn_Bus_Pc
    OR a.Mn_Bus_Cntry_Cd != f.Mn_Bus_Cntry_Cd
    OR a.Mn_Bus_DPID != f.Mn_Bus_DPID
  )
LEFT JOIN agency_v2 g -- Updates (Email Change)
  ON a.pid = g.pid AND a."date"=DATE('2018-07-16') and a.state='full' and g."date"=DATE('2018-08-30') and g.state='update' and a.Ent_Eml != g.Ent_Eml
LEFT JOIN agency_v2 h -- Updates (Industry change)
  ON a.pid = f.pid AND a."date"=DATE('2018-07-16') and a.state='full' and f."date"=DATE('2018-08-30') and f.state='update' and (
    a.Mn_Indy_Clsn != h.Mn_Indy_Clsn
    OR a.Mn_Indy_Clsn_Descn != h.Mn_Indy_Clsn_Descn
  )
