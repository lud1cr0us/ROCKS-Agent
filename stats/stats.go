package stats

import (
	"fmt"
	"rocks_agent/response"
	"github.com/shirou/gopsutil/disk"
)

func GetDiskSpace() response.ResponseStruct{
	parts, _ := disk.Partitions(true)
	for _, p := range parts {
			device := p.Mountpoint
			s, _ := disk.Usage(device)

			if s.Total == 0 {
					continue
			}
			percent := fmt.Sprintf("%0.3f%%", s.UsedPercent)
			return response.ResponseStruct{Code: response.SCS600, Message: percent}
	}
	return response.ResponseStruct{Code: response.SCS600, Message: "Failed to get Free Disk Space"}
}
