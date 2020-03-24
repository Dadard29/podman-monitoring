package scraper

import (
	"errors"
	"fmt"
	"github.com/containers/libpod/libpod"
	"log"
)

func logg(prefix string, msg interface{}) {
	log.Println(prefix, msg)

}

func (s Scraper) getPodmanVersion() (PodmanInfos, error) {
	rawInfos, err := s.libpodRuntime.Info()
	var infos PodmanInfos

	for _, t := range rawInfos {
		if t.Type == "host" {
			infos.Host.BuildahVersion = t.Data["BuildahVersion"].(string)
			infos.Host.CgroupVersion = t.Data["CgroupVersion"].(string)

			conmon := t.Data["Conmon"].(map[string]interface{})
			infos.Host.Conmon.Version = conmon["version"].(string)

			distrib := t.Data["Distribution"].(map[string]interface{})
			infos.Host.Distribution.Distribution = distrib["distribution"].(string)

			infos.Host.MemFree = t.Data["MemFree"].(int64)
			infos.Host.MemTotal = t.Data["MemTotal"].(int64)

			oci := t.Data["OCIRuntime"].(map[string]interface{})
			infos.Host.OCIRuntime.Name = oci["name"].(string)
			infos.Host.OCIRuntime.Version = oci["version"].(string)

			infos.Host.SwapFree = t.Data["SwapFree"].(int64)
			infos.Host.SwapTotal = t.Data["SwapTotal"].(int64)

			infos.Host.Arch = t.Data["arch"].(string)
			infos.Host.Cpus = t.Data["cpus"].(int)
			infos.Host.Hostname = t.Data["hostname"].(string)
			infos.Host.Kernel = t.Data["kernel"].(string)
			infos.Host.Os = t.Data["os"].(string)
			infos.Host.Rootless = t.Data["rootless"].(bool)
			infos.Host.Uptime = t.Data["uptime"].(string)
		}
	}

	return infos, err
}

func (s Scraper) getInfraContainerIp(pod *libpod.Pod) (string, error) {
	infraId, err := pod.InfraContainerID()
	if err != nil {
		return "", err
	}
	infraContainer, err := s.libpodRuntime.GetContainer(infraId)
	if err != nil {
		return "", err
	}

	ips, err := infraContainer.IPs()
	if err != nil {
		return "", err
	}

	if len(ips) < 0 {
		return "", errors.New("no ip found for this pod - running in rootless ?")
	}

	return ips[0].IP.String(), nil
}

func (s Scraper) getContainerInfos(containerId string) (*ContainerInfos, error) {
	container, err := s.libpodRuntime.GetContainer(containerId)
	if err != nil {
		return nil, err
	}

	_, imageName := container.Image()

	return &ContainerInfos{
		Name:    container.Name(),
		Created: container.CreatedTime(),
		Status:  "",
		Image: ImageInfos{
			Name: imageName,
		},
		Command: fmt.Sprintf("%v", container.Command()),
	}, nil
}

func (s Scraper) listPodsInfos() ([]PodInfos, error) {
	var podListInfos = make([]PodInfos, 0)

	pods, err := s.libpodRuntime.GetAllPods()
	if err != nil {
		return nil, err
	}

	for _, p := range pods {
		// pod name
		name := p.Name()

		// infra ip
		ip, err := s.getInfraContainerIp(p)
		if err != nil {
			logg("ERROR", err)
			continue
		}

		// get inspect infos
		ps, err := p.Inspect()
		if err != nil {
			logg("ERROR", err)
			continue
		}

		infraId, err := p.InfraContainerID()
		if err != nil {
			logg("ERROR", err)
			continue
		}

		// list the containers
		var containerListInfos = make([]ContainerInfos, 0)
		for _, c := range ps.Containers {
			if c.ID == infraId {
				// not interested in infra
				continue
			}

			infos, err := s.getContainerInfos(c.ID)
			if err != nil {
				logg("ERROR", err)
				continue
			}
			infos.Status = c.State
			containerListInfos = append(containerListInfos, *infos)
		}

		podInfos := PodInfos{
			Name:       name,
			IPAddress:  ip,
			Containers: containerListInfos,
		}

		podListInfos = append(podListInfos, podInfos)
	}

	return podListInfos, nil
}
