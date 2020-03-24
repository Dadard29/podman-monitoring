package api

import "time"

type ContainersInfos struct {
	TimeSerie time.Time `gorm:"type:timestamp;column:TimeSerie;"`
	Name      string    `gorm:"type:varchar(50);column:Name;"`
	PodName   string    `gorm:"type:varchar(50);column:PodName;"`
	Created   time.Time `gorm:"type:date;column:Created;"`
	Status    string    `gorm:"type:varchar(30);column:Status;"`
	ImageName string    `gorm:"type:varchar(100);column:ImageName;"`
	Command   string    `gorm:"type:varchar(200);column:Command;"`
}

func (ContainersInfos) TableName() string {
	return "podman_containers_infos"
}

type PodsInfos struct {
	TimeSerie time.Time `gorm:"type:timestamp;column:TimeSerie;"`
	Name      string    `gorm:"type:varchar(50);column:Name;"`
	IPAddress string    `gorm:"type:varchar(30);column:IPAddress;"`
}

func (PodsInfos) TableName() string {
	return "podman_pods_infos"
}

type PodmanInfos struct {
	TimeSerie         time.Time `gorm:"type:timestamp;column:TimeSerie"`
	BuildahVersion    string    `gorm:"type:varchar(10);column:BuildahVersion"`
	CgroupVersion     string    `gorm:"type:varchar(5);column:CgroupVersion"`
	ConmonVersion     string    `gorm:"type:varchar(100);column:ConmonVersion"`
	Distribution      string    `gorm:"type:varchar(20);column:Distribution"`
	MemFree           int64     `gorm:"type:int;column:MemFree"`
	MemTotal          int64     `gorm:"type:int;column:MemTotal"`
	OCIRuntimeName    string    `gorm:"type:varchar(10);column:OCIRuntimeName"`
	OCIRuntimeVersion string    `gorm:"type:varchar(150);column:OCIRuntimeVersion"`
	SwapFree          int64     `gorm:"type:bigint;column:SwapFree"`
	SwapTotal         int64     `gorm:"type:bigint;column:SwapTotal"`
	Arch              string    `gorm:"type:varchar(10);column:Arch"`
	Cpus              int       `gorm:"type:int;column:Cpus"`
	Hostname          string    `gorm:"type:varchar(30);column:Hostname"`
	Kernel            string    `gorm:"type:varchar(30);column:Kernel"`
	Os                string    `gorm:"type:varchar(10);column:Os"`
	Rootless          bool      `gorm:"type:bool;column:Rootless"`
	Uptime            string    `gorm:"type:varchar(50);column:Uptime"`
}

func (PodmanInfos) TableName() string {
	return "podman_infos"
}
