package sentinelm

import (
	iconfig "ffly-plus/internal/config"

	sentinelAPI "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/colinrs/pkgx/logger"
)

// InitSentinelByCustom ...
func InitSentinelByCustom() error {
	// We should initialize Sentinel first.
	conf := config.NewDefaultConfig()
	// for testing, logging output to console
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger("Sentinel")
	err := sentinelAPI.InitWithConfig(conf)
	if err != nil {
		logger.Error(err)
		return err
	}
	var loadRules []*flow.Rule
	for _, ruls := range iconfig.Conf.SentinelRules {
		flowRule := &flow.Rule{
			Resource: ruls.Resource,
			Count:    float64(ruls.Count),
		}
		flowRule.MetricType = flow.QPS
		flowRule.ControlBehavior = flow.Reject
		if ruls.MetricType == "QPS" {
			flowRule.MetricType = flow.QPS
		}
		if ruls.ControlBehavior == "Reject" {
			flowRule.ControlBehavior = flow.Reject
		}
		loadRules = append(loadRules, flowRule)

	}
	_, err = flow.LoadRules(loadRules)
	if err != nil {
		logger.Error("Unexpected error: %+v", err)
		return err
	}
	logger.Info(flow.GetRules())
	return nil
}
