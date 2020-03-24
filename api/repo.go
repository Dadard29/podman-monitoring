package api

import (
	"github.com/Dadard29/podman-monitoring/scraper"
	"log"
	"time"
)

func insert(value interface{}) {
	apiObject.orm.Create(value)
}

func storePodInfosRepo(podInfos []scraper.PodInfos) error {
	log.Println("storing pods infos...")
	for _, p := range podInfos {
		podInfosDb := PodsInfos{
			TimeSerie: time.Now(),
			Name:      p.Name,
			IPAddress: p.IPAddress,
		}
		insert(&podInfosDb)

		for _, c := range p.Containers {
			containerInfoDb := ContainersInfos{
				TimeSerie: time.Now(),
				Name:      c.Name,
				PodName:   p.Name,
				Created:   c.Created,
				Status:    c.Status,
				ImageName: c.Image.Name,
				Command:   c.Command,
			}

			insert(&containerInfoDb)
		}

	}

	return nil
}

func storePodmanInfosRepo(pi scraper.PodmanInfos) error {
	log.Println("storing podman infos...")
	podmanInfosDb := PodmanInfos{
		TimeSerie:         time.Now(),
		BuildahVersion:    pi.Host.BuildahVersion,
		CgroupVersion:     pi.Host.CgroupVersion,
		ConmonVersion:     pi.Host.Conmon.Version,
		Distribution:      pi.Host.Distribution.Distribution,
		MemFree:           pi.Host.MemFree,
		MemTotal:          pi.Host.MemTotal,
		OCIRuntimeName:    pi.Host.OCIRuntime.Name,
		OCIRuntimeVersion: pi.Host.OCIRuntime.Version,
		SwapFree:          pi.Host.SwapFree,
		SwapTotal:         pi.Host.SwapTotal,
		Arch:              pi.Host.Arch,
		Cpus:              pi.Host.Cpus,
		Hostname:          pi.Host.Hostname,
		Kernel:            pi.Host.Kernel,
		Os:                pi.Host.Os,
		Rootless:          pi.Host.Rootless,
		Uptime:            pi.Host.Uptime,
	}

	insert(podmanInfosDb)

	return nil
}

func storePodmanProxyInfosRepo() {
	// todo
}
