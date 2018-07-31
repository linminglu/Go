create table t_currency_record
(
  tradeID       varchar(64)  not null
    primary key,
  playerID      bigint       not null,
  channel       int          null,
  currencyType  int          null
  comment '货币类型：1.金币，2.元宝，3，房卡',
  amount        int          null
  comment '变化金额',
  beforeBalance int          null
  comment '变化前余额',
  afterBalance  int          null
  comment '变化后余额',
  tradeTime     datetime     null
  comment '交易时间',
  status        int          null
  comment '1.成功，2失败',
  remark        varchar(256) null,
  constraint t_currency_record_tradeID_uindex
  unique (tradeID)
)
  comment '虚拟货币流水表';



create table t_game_detail
(
  detailID   varchar(64) not null
    primary key,
  playerID   bigint      not null,
  deskID     int         null,
  gameID     int         null,
  amount     int         null,
  isWinner   tinyint(1)  null,
  createTime datetime    null,
  createBy   varchar(64) null,
  updateTime datetime    null,
  updateBy   varchar(64) null
)
  comment '游戏记录明细表';




create table t_game_sumary
(
  sumaryID   bigint       not null
    primary key,
  deskID     bigint       null,
  gameID     int          null,
  levelID    int          null
  comment '场次ID',
  playerIDs  varchar(256) null
  comment '桌子内玩家，多个玩家用|分割',
  winnerIDs  varchar(256) null
  comment '赢家ID，多个赢家用|分割',
  createTime datetime     null,
  createBy   varchar(64)  null,
  updateTime datetime     null,
  updateBy   varchar(64)  null
)
  comment '游戏记录汇总表';



create table t_login_record
(
  recordID       bigint          not null
    primary key,
  playerID       bigint          not null,
  onlineDuration int default '0' null
  comment '在线时长',
  gamingDuration int default '0' null
  comment '游戏时长',
  area           varchar(64)     null,
  loginChannel   int             null
  comment '上一次登录游戏的渠道号：省ID + 渠道ID',
  loginType      int             null
  comment '玩家上一次登陆游戏时，所选方式。',
  loginTime      datetime        null,
  logoutTime     datetime        null,
  ip             varchar(16)     null,
  loginDevice    varchar(32)     null,
  deviceCode     varchar(128)    null,
  createTime     datetime        null,
  createBy       varchar(64)     null,
  updateTime     datetime        null,
  updateBy       varchar(64)     null,
  constraint t_login_record_recordID_uindex
  unique (recordID)
)
  comment '玩家登录记录表';