package service

import (
	"github.com/gin-gonic/gin"
	"trojan-panel/dao"
	"trojan-panel/model/vo"
	"trojan-panel/util"
)

// CronTrafficRank 流量排行榜 一小时更新一次
func CronTrafficRank() {
	// Rankings are now queried from indexed durable ledgers. Kept as a no-op
	// for compatibility with existing cron configuration.
}

func TrafficRank(period string) ([]vo.AccountTrafficRankVo, error) {
	trafficRank, err := dao.TrafficRank(period)
	if err != nil {
		return nil, err
	}
	for index := range trafficRank {
		trafficRank[index].Username = maskUsername(trafficRank[index].Username)
	}
	return trafficRank, nil
}

func maskUsername(username string) string {
	runes := []rune(username)
	if len(runes) <= 2 {
		return string(runes[:1]) + "****"
	}
	if len(runes) <= 4 {
		return string(runes[:1]) + "****" + string(runes[len(runes)-1:])
	}
	return string(runes[:2]) + "****" + string(runes[len(runes)-2:])
}

func ServerTrafficUsage(period string, nodeServerID *uint, pageNum, pageSize uint) (*vo.ServerTrafficUsagePageVo, error) {
	rows, total, err := dao.SelectServerTrafficUsage(period, nodeServerID, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return &vo.ServerTrafficUsagePageVo{BaseVoPage: vo.BaseVoPage{PageNum: pageNum, PageSize: pageSize, Total: total}, Rows: rows}, nil
}

func PanelGroup(c *gin.Context) (*vo.PanelGroupVo, error) {
	accountInfo, err := GetAccountInfo(c)
	if err != nil {
		return nil, err
	}
	account, err := SelectAccountById(&accountInfo.Id)
	if err != nil {
		return nil, err
	}
	nodeCount, err := CountNode()
	if err != nil {
		return nil, err
	}
	panelGroupVo := vo.PanelGroupVo{
		Quota:        *account.Quota,
		ResidualFlow: *account.Quota - *account.Upload - *account.Download,
		NodeCount:    nodeCount,
		ExpireTime:   *account.ExpireTime,
	}
	if util.IsAdmin(accountInfo.Roles) {
		var err error
		accountCount, err := CountAccountByUsername(nil)
		cpuUsed, err := util.GetCpuPercent()
		memUsed, err := util.GetMemPercent()
		diskUsed, err := util.GetDiskPercent()
		if err != nil {
			return nil, err
		}
		panelGroupVo.AccountCount = accountCount
		panelGroupVo.CpuUsed = cpuUsed
		panelGroupVo.MemUsed = memUsed
		panelGroupVo.DiskUsed = diskUsed
	}
	return &panelGroupVo, nil
}
