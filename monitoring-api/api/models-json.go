package api

import "time"

type PodInfosJson struct {
	Name       string           `json:"name"`
	IPAddress  string           `json:"ip_address"`
	Containers []ContainerInfosJson `json:"containers"`
}

type ContainerInfosJson struct {
	Name    string     `json:"name"`
	Created time.Time  `json:"created"`
	Status  string     `json:"status"`
	Image   ImageInfosJson `json:"image"`
	Command string     `json:"command"`
	//Ports []string `json:"ports"`
}

type ImageInfosJson struct {
	Name string `json:"name"`
}

type PodmanInfosJson struct {
	Host struct {
		BuildahVersion string `json:"BuildahVersion"`
		CgroupVersion  string `json:"CgroupVersion"`
		Conmon         struct {
			Version string `json:"version"`
		} `json:"Conmon"`
		Distribution struct {
			Distribution string `json:"distribution"`
		} `json:"Distribution"`
		MemFree    int64 `json:"MemFree"`
		MemTotal   int64 `json:"MemTotal"`
		OCIRuntime struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"OCIRuntime"`
		SwapFree  int64  `json:"SwapFree"`
		SwapTotal int64  `json:"SwapTotal"`
		Arch      string `json:"arch"`
		Cpus      int    `json:"cpus"`
		Hostname  string `json:"hostname"`
		Kernel    string `json:"kernel"`
		Os        string `json:"os"`
		Rootless  bool   `json:"rootless"`
		Uptime    string `json:"uptime"`
	} `json:"host"`
}
