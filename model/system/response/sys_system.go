package response

import "achilles/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
