package api

import "SamWaf/service/waf_service"

type APIGroup struct {
	WafHostAPi
	WafStatApi
	WafLogAPi
	WafRuleAPi
	WafEngineApi
	WafWhiteIpApi
	WafWhiteUrlApi
	WafLdpUrlApi
}

var APIGroupAPP = new(APIGroup)
var (
	wafHostService     = waf_service.WafHostServiceApp
	wafLogService      = waf_service.WafLogServiceApp
	wafStatService     = waf_service.WafStatServiceApp
	wafRuleService     = waf_service.WafRuleServiceApp
	wafIpWhiteService  = waf_service.WafWhiteIpServiceApp
	wafUrlWhiteService = waf_service.WafWhiteUrlServiceApp
	wafLdpUrlService   = waf_service.WafLdpUrlServiceApp
)
