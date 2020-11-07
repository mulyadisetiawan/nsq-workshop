package benefit

import "github.com/sharring_session/nsq/repository/benefit"

func parsePublishBenefitParam(userIDStr string) benefit.PublishParam {
	return benefit.PublishParam{
		UserIDStr: userIDStr,
	}
}
