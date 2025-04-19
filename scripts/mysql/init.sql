create table contact_apply
(
    id            bigint auto_increment comment '自增id'
        primary key,
    uuid          char(20)     not null comment '申请id',
    user_id       char(20)     not null comment '申请人id',
    contact_id    char(20)     not null comment '被申请id',
    contact_type  tinyint      not null comment '被申请类型，0.用户，1.群聊',
    status        tinyint      not null comment '申请状态，0.申请中，1.通过，2.拒绝，3.拉黑',
    message       varchar(100) null comment '申请信息',
    last_apply_at datetime     not null comment '最后申请时间',
    deleted_at    datetime     null comment '删除时间',
    constraint idx_uuid
        unique (uuid)
)
    comment '联系人申请表';

create index idx_contact_id
    on contact_apply (contact_id);

create index idx_deleted_at
    on contact_apply (deleted_at);

create index idx_status_apply_time
    on contact_apply (status, last_apply_at)
    comment '状态与申请时间联合索引';

create index idx_user_contact
    on contact_apply (user_id, contact_id, contact_type)
    comment '申请人与被申请人联合索引';

create index idx_user_id
    on contact_apply (user_id);

create table group_info
(
    id         bigint auto_increment comment '自增id'
        primary key,
    uuid       char(20)                                                                                not null comment '群组唯一id',
    name       varchar(20)                                                                             not null comment '群名称',
    notice     varchar(500)                                                                            null comment '群公告',
    members    json                                                                                    null comment '群组成员',
    member_cnt int       default 1                                                                     not null comment '群人数',
    owner_id   char(20)                                                                                not null comment '群主uuid',
    add_mode   tinyint   default 0                                                                     not null comment '加群方式，0.直接，1.审核',
    avatar     char(255) default 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png' not null comment '头像',
    status     tinyint   default 0                                                                     not null comment '状态，0.正常，1.禁用，2.解散',
    created_at datetime                                                                                not null comment '创建时间',
    deleted_at datetime                                                                                null comment '删除时间',
    constraint idx_uuid
        unique (uuid)
)
    comment '群组信息表';

create index idx_created_at
    on group_info (created_at);

create index idx_deleted_at
    on group_info (deleted_at);

create table message
(
    id          bigint auto_increment comment '自增id'
        primary key,
    uuid        char(20)     not null comment '消息uuid',
    session_id  char(20)     not null comment '会话uuid',
    type        tinyint      not null comment '消息类型，0.文本，1.文件，2.通话',
    content     text         null comment '消息内容',
    url         char(255)    null comment '消息url',
    send_id     char(20)     not null comment '发送者uuid',
    send_name   varchar(20)  not null comment '发送者昵称',
    send_avatar varchar(255) not null comment '发送者头像',
    receive_id  char(20)     not null comment '接受者uuid',
    file_type   char(10)     null comment '文件类型',
    file_name   varchar(50)  null comment '文件名',
    file_size   char(20)     null comment '文件大小',
    status      tinyint      not null comment '状态，0.未发送，1.已发送',
    created_at  datetime     not null comment '创建时间',
    av_data     text         null comment '通话传递数据',
    constraint idx_uuid
        unique (uuid)
)
    comment '消息表';

create index idx_created_at
    on message (created_at);

create index idx_receive_id
    on message (receive_id);

create index idx_send_id
    on message (send_id);

create index idx_session_created
    on message (session_id, created_at)
    comment '会话消息时间联合索引';

create index idx_session_id
    on message (session_id);

create table session
(
    id           bigint auto_increment comment '自增id'
        primary key,
    uuid         char(20)                               not null comment '会话uuid',
    send_id      char(20)                               not null comment '创建会话人id',
    receive_id   char(20)                               not null comment '接受会话人id',
    receive_name varchar(20)                            not null comment '名称',
    avatar       char(255) default 'default_avatar.png' not null comment '头像',
    created_at   datetime                               not null comment '创建时间',
    deleted_at   datetime                               null comment '删除时间',
    constraint idx_uuid
        unique (uuid)
)
    comment '用户会话表';

create index idx_created_at
    on session (created_at);

create index idx_deleted_at
    on session (deleted_at);

create index idx_receive_id
    on session (receive_id);

create index idx_send_id
    on session (send_id);

create index idx_session_pair
    on session (send_id, receive_id)
    comment '会话双方联合索引';

create table user_contact
(
    id           bigint auto_increment comment '自增id'
        primary key,
    user_id      char(20) not null comment '用户唯一id',
    contact_id   char(20) not null comment '对应联系id',
    contact_type tinyint  not null comment '联系类型，0.用户，1.群聊',
    status       tinyint  not null comment '联系状态，0.正常，1.拉黑，2.被拉黑，3.删除好友，4.被删除好友，5.被禁言，6.退出群聊，7.被踢出群聊',
    created_at   datetime not null comment '创建时间',
    deleted_at   datetime null comment '删除时间'
)
    comment '用户联系人表';

create index idx_contact_id
    on user_contact (contact_id);

create index idx_deleted_at
    on user_contact (deleted_at);

create index idx_user_contact
    on user_contact (user_id, contact_id, contact_type)
    comment '用户与联系人联合索引';

create index idx_user_id
    on user_contact (user_id);

create table user_info
(
    id         bigint auto_increment comment '自增id'
        primary key,
    uuid       char(20)                                                                                not null comment '用户唯一id',
    nickname   varchar(20)                                                                             not null comment '昵称',
    telephone  char(11)                                                                                not null comment '电话',
    email      char(30)                                                                                null comment '邮箱',
    avatar     char(255) default 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png' not null comment '头像',
    gender     tinyint                                                                                 null comment '性别，0.男，1.女',
    signature  varchar(100)                                                                            null comment '个性签名',
    password   char(18)                                                                                not null comment '密码',
    birthday   char(8)                                                                                 null comment '生日',
    created_at datetime                                                                                not null comment '创建时间',
    deleted_at datetime                                                                                null comment '删除时间',
    is_admin   tinyint                                                                                 not null comment '是否是管理员，0.不是，1.是',
    status     tinyint                                                                                 not null comment '状态，0.正常，1.禁用',
    constraint idx_uuid
        unique (uuid)
)
    comment '用户信息表';

create index idx_created_at
    on user_info (created_at);

create index idx_telephone
    on user_info (telephone);

