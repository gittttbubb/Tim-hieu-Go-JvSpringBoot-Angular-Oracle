-- Start of DDL Script for Table MYHEALTH.DM_BSC_DICHVU
-- Generated 02-Jul-2026 15:49:54 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_bsc_dichvu
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_448949".nextval NOT NULL,
    ten_dv                         VARCHAR2(255 BYTE),
    dvt                            VARCHAR2(255 BYTE),
    cong_thuc                      VARCHAR2(255 BYTE),
    tom_tat                        VARCHAR2(255 BYTE),
    trang_thai                     NUMBER(1,0) DEFAULT 1)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for DM_BSC_DICHVU

ALTER TABLE dm_bsc_dichvu
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/


-- End of DDL Script for Table MYHEALTH.DM_BSC_DICHVU

-- Start of DDL Script for Table MYHEALTH.DM_CHUNG
-- Generated 02-Jul-2026 15:50:10 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_chung
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_449297".nextval NOT NULL,
    ma                             VARCHAR2(50 BYTE) NOT NULL,
    ten                            VARCHAR2(255 BYTE) NOT NULL,
    mota                           VARCHAR2(4000 BYTE),
    trang_thai                     NUMBER(1,0) DEFAULT 1,
    macha                          VARCHAR2(50 BYTE),
    muc_duoi                       NUMBER(10,0),
    muc_tren                       NUMBER(10,0),
    dongia                         NUMBER(20,2),
    activated                      NUMBER(1,0),
    tungay                         DATE,
    denngay                        DATE)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Indexes for DM_CHUNG

CREATE INDEX idx_ma_dm_chung ON dm_chung
  (
    ma                              ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/



-- Constraints for DM_CHUNG

ALTER TABLE dm_chung
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/


-- End of DDL Script for Table MYHEALTH.DM_CHUNG

-- Start of DDL Script for Table MYHEALTH.DM_CSYT
-- Generated 02-Jul-2026 15:50:15 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_csyt
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_448906".nextval NOT NULL,
    ma_bv                          VARCHAR2(500 BYTE) NOT NULL,
    ten_bv                         VARCHAR2(500 BYTE) NOT NULL,
    tuyen_bv                       VARCHAR2(500 BYTE),
    hang_bv                        VARCHAR2(500 BYTE),
    dia_chi                        VARCHAR2(500 BYTE),
    ghichu                         VARCHAR2(500 BYTE),
    tinh_id                        NUMBER,
    huyen_id                       NUMBER,
    xa_id                          NUMBER,
    tt_trienkhai                   NUMBER,
    thoigian_bd_sudung             DATE,
    dv_trienkhai                   NUMBER,
    loai_csyt                      NUMBER NOT NULL,
    tt_hoatdong                    NUMBER(1,0))
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Indexes for DM_CSYT

CREATE INDEX idx_dm_csyt_macsyt ON dm_csyt
  (
    ma_bv                           ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/



-- Constraints for DM_CSYT

ALTER TABLE dm_csyt
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/


-- End of DDL Script for Table MYHEALTH.DM_CSYT

-- Start of DDL Script for Table MYHEALTH.DM_CSYT_TRIENKHAI
-- Generated 02-Jul-2026 15:50:20 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_csyt_trienkhai
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_451893".nextval NOT NULL,
    ma_tinh                        VARCHAR2(50 BYTE) NOT NULL,
    ma_huyen                       VARCHAR2(50 BYTE),
    ma_xa                          VARCHAR2(50 BYTE),
    ma_csyt                        VARCHAR2(255 BYTE),
    ma_spdv                        VARCHAR2(50 BYTE),
    bd_trienkhai                   DATE NOT NULL,
    kt_trienkhai                   DATE,
    tt_trienkhai_id                NUMBER NOT NULL,
    dv_trienkhai_id                NUMBER,
    mota                           VARCHAR2(4000 BYTE),
    ghichu                         VARCHAR2(4000 BYTE),
    transaction_id                 VARCHAR2(100 BYTE),
    is_last                        NUMBER(1,0) DEFAULT 1,
    nguoidung_id                   VARCHAR2(255 BYTE),
    transaction_date               TIMESTAMP (6) DEFAULT SYSTIMESTAMP,
    ma_loaithuebao                 VARCHAR2(500 BYTE),
    dv_banhang_id                  NUMBER,
    dv_banhang_new_id              NUMBER,
    hinhthuc_cc                    VARCHAR2(50 BYTE),
    trangthai_ht                   NUMBER(2,0),
    kt_dungthu                     DATE,
    trangthai_goi                  NUMBER,
    ngay_ketthuc_goi               DATE,
    ngay_apdung_goi                DATE)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Indexes for DM_CSYT_TRIENKHAI

CREATE INDEX dm_csyt_trienkhai_index ON dm_csyt_trienkhai
  (
    nguoidung_id                    ASC,
    ma_spdv                         ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/



-- Constraints for DM_CSYT_TRIENKHAI


ALTER TABLE dm_csyt_trienkhai
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/




-- End of DDL Script for Table MYHEALTH.DM_CSYT_TRIENKHAI

-- Start of DDL Script for Table MYHEALTH.DM_CSYT_TRIENKHAI_IMP
-- Generated 02-Jul-2026 15:50:29 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_csyt_trienkhai_imp
    (ma_tinh                        VARCHAR2(50 BYTE) NOT NULL,
    ma_huyen                       VARCHAR2(50 BYTE),
    ma_xa                          VARCHAR2(50 BYTE),
    ma_csyt                        VARCHAR2(50 BYTE),
    ma_spdv                        VARCHAR2(50 BYTE),
    bd_trienkhai                   VARCHAR2(50 BYTE) NOT NULL,
    kt_trienkhai                   VARCHAR2(50 BYTE),
    tt_trienkhai_id                NUMBER NOT NULL,
    dv_trienkhai_id                NUMBER,
    mota                           VARCHAR2(4000 BYTE),
    ghichu                         VARCHAR2(4000 BYTE),
    transaction_id                 VARCHAR2(100 BYTE),
    is_last                        NUMBER(1,0),
    nguoidung_id                   VARCHAR2(255 BYTE),
    transaction_date               TIMESTAMP (6))
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- End of DDL Script for Table MYHEALTH.DM_CSYT_TRIENKHAI_IMP

-- Start of DDL Script for Table MYHEALTH.DM_CUM_DULIEU
-- Generated 02-Jul-2026 15:50:33 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_cum_dulieu
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_452373".nextval NOT NULL,
    ten                            VARCHAR2(255 BYTE) NOT NULL,
    mota                           VARCHAR2(255 BYTE),
    ip                             VARCHAR2(255 BYTE),
    cpu                            NUMBER(10,0),
    ram                            NUMBER(10,0),
    hdd                            NUMBER(10,0),
    trang_thai                     NUMBER(1,0) DEFAULT 1,
    ngay_khoitao                   DATE NOT NULL,
    ngay_ketthuc                   DATE,
    moitruong_id                   NUMBER,
    loai_hatang_id                 NUMBER NOT NULL,
    phanloai_hatang_id             NUMBER,
    thanhtien                      NUMBER(20,2) NOT NULL,
    nguoidung_id                   VARCHAR2(255 BYTE),
    create_date                    TIMESTAMP (6) DEFAULT SYSTIMESTAMP,
    nhom_cum_dl_id                 NUMBER)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for DM_CUM_DULIEU


ALTER TABLE dm_cum_dulieu
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/





-- End of DDL Script for Table MYHEALTH.DM_CUM_DULIEU

-- Start of DDL Script for Table MYHEALTH.DM_CUM_DULIEU_SPDV
-- Generated 02-Jul-2026 15:50:42 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_cum_dulieu_spdv
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_452376".nextval NOT NULL,
    ma_spdv                        VARCHAR2(50 BYTE) NOT NULL,
    ten_spdv                       VARCHAR2(1024 BYTE) NOT NULL,
    tyle                           NUMBER(5,2) NOT NULL,
    thanhtien                      NUMBER(20,2) NOT NULL,
    cum_dulieu_id                  NUMBER NOT NULL)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for DM_CUM_DULIEU_SPDV

ALTER TABLE dm_cum_dulieu_spdv
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/



-- End of DDL Script for Table MYHEALTH.DM_CUM_DULIEU_SPDV

-- Start of DDL Script for Table MYHEALTH.DM_DIA_PHUONG
-- Generated 02-Jul-2026 15:50:49 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_dia_phuong
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_448777".nextval NOT NULL,
    ma_dia_phuong                  VARCHAR2(10 BYTE),
    ten_dia_phuong                 VARCHAR2(255 BYTE),
    ten_dia_phuong_ns              VARCHAR2(255 BYTE),
    ten_viet_tat                   VARCHAR2(20 BYTE),
    cap_dia_phuong                 NUMBER(2,0),
    ten_viet_tat_day_du            VARCHAR2(50 BYTE),
    ten_dia_phuong_day_du          VARCHAR2(255 BYTE),
    ten_dia_phuong_day_du_ns       VARCHAR2(255 BYTE),
    dia_phuong_id                  NUMBER,
    ghi_chu                        VARCHAR2(1024 BYTE),
    nguoi_tao                      VARCHAR2(50 BYTE),
    ngay_tao                       TIMESTAMP (6) DEFAULT SYSTIMESTAMP,
    nguoi_sua                      VARCHAR2(50 BYTE),
    ngay_sua                       TIMESTAMP (6),
    trang_thai                     NUMBER(2,0) DEFAULT 1)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Indexes for DM_DIA_PHUONG

CREATE INDEX idx_dia_phuong_madiaphuong ON dm_dia_phuong
  (
    ma_dia_phuong                   ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/



-- Constraints for DM_DIA_PHUONG

ALTER TABLE dm_dia_phuong
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/


-- End of DDL Script for Table MYHEALTH.DM_DIA_PHUONG

-- Start of DDL Script for Table MYHEALTH.DM_DV_BANHANG
-- Generated 02-Jul-2026 15:50:51 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_dv_banhang
    (id                             NUMBER NOT NULL,
    stt                            NUMBER,
    ma_donvi                       VARCHAR2(20 BYTE),
    ten_donvi                      VARCHAR2(100 BYTE),
    ma_khuvuc                      NUMBER,
    is_last                        NUMBER(1,0),
    nguoidung_id                   VARCHAR2(255 BYTE),
    is_tinh_moi                    NUMBER(1,0),
    id_new                         NUMBER)
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for DM_DV_BANHANG

ALTER TABLE dm_dv_banhang
ADD CONSTRAINT pk_dm_dv_banhang PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
/


-- End of DDL Script for Table MYHEALTH.DM_DV_BANHANG

-- Start of DDL Script for Table MYHEALTH.DM_EHC
-- Generated 02-Jul-2026 15:50:54 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_ehc
    (ma_tinh                        VARCHAR2(50 BYTE) NOT NULL,
    ma_huyen                       VARCHAR2(50 BYTE),
    ma_xa                          VARCHAR2(50 BYTE),
    ma_csyt                        VARCHAR2(50 BYTE),
    ma_spdv                        VARCHAR2(50 BYTE),
    bd_trienkhai                   VARCHAR2(50 BYTE) NOT NULL,
    kt_trienkhai                   VARCHAR2(50 BYTE),
    tt_trienkhai_id                NUMBER NOT NULL,
    dv_trienkhai_id                NUMBER,
    mota                           VARCHAR2(4000 BYTE),
    ghichu                         VARCHAR2(4000 BYTE),
    transaction_id                 VARCHAR2(100 BYTE),
    is_last                        NUMBER(1,0),
    nguoidung_id                   VARCHAR2(255 BYTE),
    transaction_date               VARCHAR2(50 BYTE))
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- End of DDL Script for Table MYHEALTH.DM_EHC

-- Start of DDL Script for Table MYHEALTH.DM_GOI_DICHVU
-- Generated 02-Jul-2026 15:50:57 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_goi_dichvu
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_448955".nextval NOT NULL,
    ma                             VARCHAR2(255 BYTE),
    ten                            VARCHAR2(255 BYTE),
    mota                           VARCHAR2(255 BYTE),
    loai                           NUMBER(2,0),
    spdv_id                        NUMBER,
    trang_thai                     NUMBER(1,0) DEFAULT 1)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for DM_GOI_DICHVU

ALTER TABLE dm_goi_dichvu
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/



-- End of DDL Script for Table MYHEALTH.DM_GOI_DICHVU

-- Start of DDL Script for Table MYHEALTH.DM_HIS_COGB
-- Generated 02-Jul-2026 15:51:01 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_his_cogb
    (ma_bv                          VARCHAR2(10 BYTE),
    ngay_capnhat                   DATE,
    is_last                        NUMBER(1,0),
    tinh_giuong_tu_his_gui         NUMBER(1,0))
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- End of DDL Script for Table MYHEALTH.DM_HIS_COGB

-- Start of DDL Script for Table MYHEALTH.DM_QUYDOI
-- Generated 02-Jul-2026 15:51:02 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_quydoi
    (dv_id                          NUMBER ,
    he_so                          NUMBER,
    chia                           NUMBER)
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for DM_QUYDOI

ALTER TABLE dm_quydoi
ADD CONSTRAINT pk_dm_quydoi PRIMARY KEY (dv_id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
/


-- End of DDL Script for Table MYHEALTH.DM_QUYDOI

-- Start of DDL Script for Table MYHEALTH.DM_SPDV
-- Generated 02-Jul-2026 15:51:05 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_spdv
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_448899".nextval NOT NULL,
    ma                             VARCHAR2(50 BYTE) NOT NULL,
    ten                            VARCHAR2(1024 BYTE) NOT NULL,
    mota                           VARCHAR2(4000 BYTE),
    line_id                        NUMBER,
    line_ten                       VARCHAR2(2555 BYTE))
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for DM_SPDV


ALTER TABLE dm_spdv
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/

ALTER TABLE dm_spdv
ADD CONSTRAINT spdv_unique_ma UNIQUE (ma)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/


-- End of DDL Script for Table MYHEALTH.DM_SPDV

-- Start of DDL Script for Table MYHEALTH.DM_SPDV_CCTK
-- Generated 02-Jul-2026 15:51:10 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_spdv_cctk
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_457778".nextval NOT NULL,
    nam                            NUMBER(4,0) NOT NULL,
    ma_spdv                        VARCHAR2(50 BYTE) NOT NULL,
    cctk                           NUMBER(20,2) NOT NULL,
    ghichu                         VARCHAR2(4000 BYTE),
    nguoidung_id                   VARCHAR2(255 BYTE),
    create_date                    TIMESTAMP (6) DEFAULT SYSTIMESTAMP)
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for DM_SPDV_CCTK

ALTER TABLE dm_spdv_cctk
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
/


-- End of DDL Script for Table MYHEALTH.DM_SPDV_CCTK

-- Start of DDL Script for Table MYHEALTH.DM_SPDV_DONGIA
-- Generated 02-Jul-2026 15:51:15 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE dm_spdv_dongia
    (id                             NUMBER DEFAULT "MYHEALTH"."ISEQ$$_457781".nextval NOT NULL,
    ma                             VARCHAR2(50 BYTE) NOT NULL,
    ten                            VARCHAR2(1024 BYTE) NOT NULL,
    ma_spdv                        VARCHAR2(50 BYTE) NOT NULL,
    muc_duoi                       NUMBER(10,0) NOT NULL,
    muc_tren                       NUMBER(10,0) NOT NULL,
    dongia                         NUMBER(20,2) NOT NULL,
    activated                      NUMBER(1,0) DEFAULT 0,
    mota                           VARCHAR2(4000 BYTE),
    nguoidung_id                   VARCHAR2(255 BYTE),
    create_date                    TIMESTAMP (6) DEFAULT SYSTIMESTAMP,
    dongia_khac                    NUMBER(20,0))
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for DM_SPDV_DONGIA

ALTER TABLE dm_spdv_dongia
ADD PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
/


-- End of DDL Script for Table MYHEALTH.DM_SPDV_DONGIA

-- Start of DDL Script for Table MYHEALTH.PASSWORD_RESET_TOKENS
-- Generated 02-Jul-2026 15:51:20 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE password_reset_tokens
    (id                             VARCHAR2(36 CHAR) NOT NULL,
    user_id                        VARCHAR2(36 CHAR) NOT NULL,
    token_hash                     VARCHAR2(128 CHAR) NOT NULL,
    expires_at                     TIMESTAMP (6) NOT NULL,
    used_at                        TIMESTAMP (6),
    revoked_at                     TIMESTAMP (6),
    created_ip                     VARCHAR2(50 CHAR),
    user_agent                     VARCHAR2(500 CHAR),
    created_at                     TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL)
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Indexes for PASSWORD_RESET_TOKENS

CREATE INDEX idx_pw_reset_tokens_expires_at ON password_reset_tokens
  (
    expires_at                      ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
NOPARALLEL
LOGGING
/

CREATE INDEX idx_pw_reset_tokens_user_id ON password_reset_tokens
  (
    user_id                         ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
NOPARALLEL
LOGGING
/



-- Constraints for PASSWORD_RESET_TOKENS

ALTER TABLE password_reset_tokens
ADD CONSTRAINT password_reset_tokens_pk PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
/

ALTER TABLE password_reset_tokens
ADD CONSTRAINT password_reset_tokens_hash_uq UNIQUE (token_hash)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
/



-- End of DDL Script for Table MYHEALTH.PASSWORD_RESET_TOKENS

-- Start of DDL Script for Table MYHEALTH.PERMISSIONS
-- Generated 02-Jul-2026 15:51:28 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE permissions
    (id                             VARCHAR2(36 CHAR) NOT NULL,
    feature_group                  VARCHAR2(50 CHAR) NOT NULL,
    feature_code                   VARCHAR2(100 CHAR) NOT NULL,
    action                         VARCHAR2(50 CHAR) NOT NULL,
    description                    VARCHAR2(255 CHAR),
    created_at                     TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for PERMISSIONS

ALTER TABLE permissions
ADD CONSTRAINT permissions_pk PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/

ALTER TABLE permissions
ADD CONSTRAINT permissions_feature_code_uq UNIQUE (feature_code)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/


-- End of DDL Script for Table MYHEALTH.PERMISSIONS

-- Start of DDL Script for Table MYHEALTH.ROLE_PERMISSION
-- Generated 02-Jul-2026 15:51:33 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE role_permission
    (role_id                        VARCHAR2(36 CHAR) NOT NULL,
    permission_id                  VARCHAR2(36 CHAR) NOT NULL,
    granted                        NUMBER(1,0) DEFAULT 1 NOT NULL,
    data_scope                     VARCHAR2(20 CHAR) DEFAULT 'OWN' NOT NULL,
    updated_at                     TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL,
    updated_by                     VARCHAR2(36 CHAR))
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Indexes for ROLE_PERMISSION

CREATE INDEX idx_role_permission_role ON role_permission
  (
    role_id                         ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/



-- Constraints for ROLE_PERMISSION

ALTER TABLE role_permission
ADD CONSTRAINT role_permission_granted_ck CHECK (granted IN (0, 1))
/

ALTER TABLE role_permission
ADD CONSTRAINT role_permission_pk PRIMARY KEY (role_id, permission_id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/





-- End of DDL Script for Table MYHEALTH.ROLE_PERMISSION

-- Start of DDL Script for Table MYHEALTH.ROLES
-- Generated 02-Jul-2026 15:51:41 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE roles
    (id                             VARCHAR2(36 CHAR) NOT NULL,
    name                           VARCHAR2(50 CHAR) NOT NULL,
    display_name                   VARCHAR2(100 CHAR) NOT NULL,
    description                    VARCHAR2(255 CHAR),
    code                           VARCHAR2(50 CHAR),
    is_system                      NUMBER(1,0) DEFAULT 1 NOT NULL,
    created_at                     TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL,
    updated_at                     TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Constraints for ROLES

ALTER TABLE roles
ADD CONSTRAINT roles_pk PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/

ALTER TABLE roles
ADD CONSTRAINT roles_name_uq UNIQUE (name)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/


-- End of DDL Script for Table MYHEALTH.ROLES

-- Start of DDL Script for Table MYHEALTH.USER_PERMISSION_OVERRIDE
-- Generated 02-Jul-2026 15:51:46 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE user_permission_override
    (id                             VARCHAR2(36 CHAR) NOT NULL,
    user_id                        VARCHAR2(36 CHAR) NOT NULL,
    permission_id                  VARCHAR2(36 CHAR) NOT NULL,
    granted                        NUMBER(1,0) NOT NULL,
    reason                         VARCHAR2(500 CHAR),
    created_by                     VARCHAR2(36 CHAR) NOT NULL,
    created_at                     TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL,
    data_scope                     VARCHAR2(20 CHAR) NOT NULL)
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Indexes for USER_PERMISSION_OVERRIDE

CREATE INDEX idx_user_perm_override_user ON user_permission_override
  (
    user_id                         ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
NOPARALLEL
LOGGING
/



-- Constraints for USER_PERMISSION_OVERRIDE

ALTER TABLE user_permission_override
ADD CONSTRAINT upo_chk_scope CHECK (data_scope IN ('OWN','TEAM','ALL'))
/

ALTER TABLE user_permission_override
ADD CONSTRAINT upo_pk PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
/

ALTER TABLE user_permission_override
ADD CONSTRAINT upo_uq_usr_perm UNIQUE (user_id, permission_id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
/




-- End of DDL Script for Table MYHEALTH.USER_PERMISSION_OVERRIDE

-- Start of DDL Script for Table MYHEALTH.USER_PROVINCE_SCOPE
-- Generated 02-Jul-2026 15:51:58 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE user_province_scope
    (id                             VARCHAR2(36 CHAR) NOT NULL,
    user_id                        VARCHAR2(36 CHAR) NOT NULL,
    dia_phuong_id                  NUMBER(19,0) NOT NULL,
    version                        NUMBER(10,0) DEFAULT 1 NOT NULL,
    created_at                     TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL,
    created_by                     VARCHAR2(36 CHAR) NOT NULL,
    updated_at                     TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL,
    updated_by                     VARCHAR2(36 CHAR) NOT NULL)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Indexes for USER_PROVINCE_SCOPE

CREATE INDEX sec_usr_prov_scope_dp_ix ON user_province_scope
  (
    dia_phuong_id                   ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/

CREATE INDEX sec_usr_prov_scope_user_ix ON user_province_scope
  (
    user_id                         ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/



-- Constraints for USER_PROVINCE_SCOPE

ALTER TABLE user_province_scope
ADD CONSTRAINT sec_user_province_scope_pk PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/

ALTER TABLE user_province_scope
ADD CONSTRAINT sec_user_province_scope_uq UNIQUE (user_id, dia_phuong_id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/




-- End of DDL Script for Table MYHEALTH.USER_PROVINCE_SCOPE

-- Start of DDL Script for Table MYHEALTH.USER_ROLE
-- Generated 02-Jul-2026 15:52:07 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE user_role
    (user_id                        VARCHAR2(36 CHAR) NOT NULL,
    role_id                        VARCHAR2(36 CHAR) NOT NULL,
    assigned_at                    TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL,
    assigned_by                    VARCHAR2(36 CHAR))
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Indexes for USER_ROLE

CREATE INDEX idx_user_role_role_id ON user_role
  (
    role_id                         ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/



-- Constraints for USER_ROLE

ALTER TABLE user_role
ADD CONSTRAINT user_role_pk PRIMARY KEY (user_id, role_id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/




-- End of DDL Script for Table MYHEALTH.USER_ROLE

-- Start of DDL Script for Table MYHEALTH.USER_SESSION
-- Generated 02-Jul-2026 15:52:11 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE user_session
    (id                             VARCHAR2(36 CHAR) NOT NULL,
    user_id                        VARCHAR2(36 CHAR) NOT NULL,
    jwt_id                         VARCHAR2(64 CHAR) NOT NULL,
    refresh_token_hash             VARCHAR2(128 CHAR),
    issued_at                      TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL,
    expires_at                     TIMESTAMP (6) NOT NULL,
    revoked_at                     TIMESTAMP (6),
    revoke_reason                  VARCHAR2(255 CHAR),
    ip_address                     VARCHAR2(50 CHAR),
    user_agent                     VARCHAR2(500 CHAR))
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  NOPARALLEL
  LOGGING
/





-- Indexes for USER_SESSION

CREATE INDEX idx_user_session_expires ON user_session
  (
    expires_at                      ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/

CREATE INDEX idx_user_session_user_id ON user_session
  (
    user_id                         ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/



-- Constraints for USER_SESSION

ALTER TABLE user_session
ADD CONSTRAINT user_session_pk PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/

ALTER TABLE user_session
ADD CONSTRAINT user_session_jti_uq UNIQUE (jwt_id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/



-- End of DDL Script for Table MYHEALTH.USER_SESSION

-- Start of DDL Script for Table MYHEALTH.USERS
-- Generated 02-Jul-2026 15:52:17 from MYHEALTH@(DESCRIPTION =(ADDRESS_LIST =(ADDRESS = (PROTOCOL = TCP)(HOST = 10.159.12.144)(PORT = 1521)))(CONNECT_DATA =(SERVICE_NAME = PDB_MASTERDEV)))

CREATE TABLE users
    (id                             VARCHAR2(36 CHAR) NOT NULL,
    full_name                      VARCHAR2(100 CHAR) NOT NULL,
    username                       VARCHAR2(50 CHAR) NOT NULL,
    email                          VARCHAR2(200 CHAR) NOT NULL,
    phone                          VARCHAR2(20 CHAR) NOT NULL,
    password_hash                  CLOB NOT NULL,
    status                         VARCHAR2(20 CHAR) DEFAULT 'ACTIVE' NOT NULL,
    must_change_password           NUMBER(1,0) DEFAULT 1 NOT NULL,
    created_at                     TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL,
    updated_at                     TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL,
    created_by                     VARCHAR2(36 CHAR),
    password_changed_at            TIMESTAMP (6) DEFAULT SYSTIMESTAMP NOT NULL)
  SEGMENT CREATION IMMEDIATE
  PCTFREE     10
  INITRANS    1
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
  NOCACHE
  MONITORING
  LOB ("PASSWORD_HASH") STORE AS SECUREFILE SYS_LOB0000457675C00006$$
  (
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     106496
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
   NOCACHE LOGGING
   CHUNK 8192
  )
  NOPARALLEL
  LOGGING
/





-- Indexes for USERS

CREATE INDEX idx_users_password_changed ON users
  (
    password_changed_at             ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/

CREATE INDEX idx_users_status ON users
  (
    status                          ASC
  )
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
NOPARALLEL
LOGGING
/



-- Constraints for USERS

ALTER TABLE users
ADD CONSTRAINT chk_user_status CHECK (
            status IN (
                'ACTIVE',
                'LOCKED',
                'PENDING_PASSWORD_CHANGE'
            )
        )
/

ALTER TABLE users
ADD CONSTRAINT users_pk PRIMARY KEY (id)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/

ALTER TABLE users
ADD CONSTRAINT uq_users_email UNIQUE (email)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/

ALTER TABLE users
ADD CONSTRAINT uq_users_username UNIQUE (username)
USING INDEX
  PCTFREE     10
  INITRANS    2
  MAXTRANS    255
  TABLESPACE  myhealth
  STORAGE   (
    INITIAL     65536
    NEXT        1048576
    MINEXTENTS  1
    MAXEXTENTS  2147483645
  )
/


-- End of DDL Script for Table MYHEALTH.USERS

-- Foreign Key
ALTER TABLE dm_csyt_trienkhai
ADD CONSTRAINT fk_dv_banhang_id FOREIGN KEY (dv_banhang_id)
REFERENCES dm_dv_banhang (id)
/
ALTER TABLE dm_csyt_trienkhai
ADD CONSTRAINT fk_tt_trienkhai_id FOREIGN KEY (tt_trienkhai_id)
REFERENCES dm_chung (id)
/
ALTER TABLE dm_csyt_trienkhai
ADD CONSTRAINT fk_dv_trienkhai_id FOREIGN KEY (dv_trienkhai_id)
REFERENCES dm_chung (id)
/
-- Foreign Key
ALTER TABLE dm_cum_dulieu
ADD CONSTRAINT fk_cumdl_nhom FOREIGN KEY (nhom_cum_dl_id)
REFERENCES dm_chung (id)
/
ALTER TABLE dm_cum_dulieu
ADD CONSTRAINT fk_cumdl_moitruong FOREIGN KEY (moitruong_id)
REFERENCES dm_chung (id)
/
ALTER TABLE dm_cum_dulieu
ADD CONSTRAINT fk_cumdl_loai_hatang FOREIGN KEY (loai_hatang_id)
REFERENCES dm_chung (id)
/
ALTER TABLE dm_cum_dulieu
ADD CONSTRAINT fk_cumdl_phanloai_hatang FOREIGN KEY (phanloai_hatang_id)
REFERENCES dm_chung (id)
/
-- Foreign Key
ALTER TABLE dm_cum_dulieu_spdv
ADD CONSTRAINT fk_cumdl_spdv FOREIGN KEY (cum_dulieu_id)
REFERENCES dm_cum_dulieu (id)
/
-- Foreign Key
ALTER TABLE dm_goi_dichvu
ADD CONSTRAINT fk_goi_dichvu_spdv FOREIGN KEY (spdv_id)
REFERENCES dm_spdv (id)
/
-- Foreign Key
ALTER TABLE dm_spdv
ADD CONSTRAINT fk_line_id FOREIGN KEY (line_id)
REFERENCES dm_chung (id)
/
-- Foreign Key
ALTER TABLE password_reset_tokens
ADD CONSTRAINT password_reset_tokens_user_fk FOREIGN KEY (user_id)
REFERENCES users (id)
/
-- Foreign Key
ALTER TABLE role_permission
ADD CONSTRAINT role_permission_role_fk FOREIGN KEY (role_id)
REFERENCES roles (id)
/
ALTER TABLE role_permission
ADD CONSTRAINT role_permission_permission_fk FOREIGN KEY (permission_id)
REFERENCES permissions (id)
/
ALTER TABLE role_permission
ADD CONSTRAINT role_permission_scope_fk FOREIGN KEY (data_scope)
REFERENCES data_scope (code)
/
-- Foreign Key
ALTER TABLE user_permission_override
ADD CONSTRAINT upo_usr_fk FOREIGN KEY (user_id)
REFERENCES users (id)
/
ALTER TABLE user_permission_override
ADD CONSTRAINT upo_perm_fk FOREIGN KEY (permission_id)
REFERENCES permissions (id)
/
-- Foreign Key
ALTER TABLE user_province_scope
ADD CONSTRAINT sec_usr_prov_scope_user_fk FOREIGN KEY (user_id)
REFERENCES users (id)
/
ALTER TABLE user_province_scope
ADD CONSTRAINT sec_usr_prov_scope_dp_fk FOREIGN KEY (dia_phuong_id)
REFERENCES dm_dia_phuong (id)
/
-- Foreign Key
ALTER TABLE user_role
ADD CONSTRAINT user_role_user_fk FOREIGN KEY (user_id)
REFERENCES users (id)
/
ALTER TABLE user_role
ADD CONSTRAINT user_role_role_fk FOREIGN KEY (role_id)
REFERENCES roles (id)
/
-- Foreign Key
ALTER TABLE user_session
ADD CONSTRAINT user_session_user_fk FOREIGN KEY (user_id)
REFERENCES users (id)
/
-- End of DDL script for Foreign Key(s)
