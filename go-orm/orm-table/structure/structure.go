package structure

/*
生成基本的表，user,group,group_relation,等
如果要自己添加字段，在相应的模型下面定义字段名，类型，特征等
example：在user下添加create_time字段，类型为时间
create_time   time.Time
运行程序以后会自动构建相应的表，请注意，如果添加了不需要的字段，删除之后再次构建字段还是会存在的
这时候需要运行删除字段的命令
*/
type User struct {
	Uid      int     `gorm:"size:64;not null;unique;primary_key"` //uid值，长度 255, 不为空，唯一值, 主键
	Uname    string  `gorm:"type:varchar(40);not null;unique"`
	Password string  `gorm:"type:varchar(128)"`
	Status   int     `gorm:"size:4"`
	Group    []Group `gorm:"many2many:user_groups"` //多对多，链接表user_groups
}

type Group struct {
	Gid             int              `gorm:"size:64;unique;primary_key"`
	Name            string           `gorm:"type:varchar(64);unique"`
	Status          int              `gorm:"size:4"`
	Group_relations []Group_relation `gorm:"ForeignKey:Parent_gid;AssociationForeignKey:Gid"` //一对多，外键parent_gid,关联Gid
}

type User_group struct {
	Status int `gorm:"size:4"`
}

type Group_relation struct {
	Gids       int   `gorm:"primary_key"`
	Groups     Group `gorm:"ForeignKey:Gids;AssociationForeignKey:Gid"`
	Parent_gid int
	Status     int
}

type Group_relation_log struct {
	Op_id    string `gorm:"type:varchar(64)"`
	Op_type  string `gorm:"type:varchar(32)"`
	Op_value string `gorm:"type:varchar(255)"`
}

type Group_log struct {
	Op_id    string `gorm:"type:varchar(64)"`
	Op_type  string `gorm:"type:varchar(32)"`
	Op_value string `gorm:"type:varchar(255)"`
}

type User_group_log struct {
	Op_id    string `gorm:"type:varchar(64)"`
	Op_type  string `gorm:"type:varchar(32)"`
	Op_value string `gorm:"type:varchar(255)"`
}

type User_log struct {
	Op_id    string `gorm:"type:varchar(64)"`
	Op_type  string `gorm:"type:varchar(32)"`
	Op_value string `gorm:"type:varchar(255)"`
}
