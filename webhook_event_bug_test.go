package tapd

import (
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestWebhook_Event_BugCreateEvent(t *testing.T) {
	rawBody := `{
  "event": "bug::create",
  "event_from": "web",
  "referer": "https://www.tapd.cn/23402991/bugtrace/bugs/add",
  "workspace_id": "23402991",
  "current_user": "T8黄玲",
  "id": "1123402991001037239",
  "title": "交易已完成，但未收到款項",
  "issue_id": "",
  "is_new_status": "0",
  "is_replicate": "0",
  "create_link": "0",
  "is_jenkins": "0",
  "template_id": "1123402991001000320",
  "description": "<div><span style=\"background-color: #ffffff;\">【问题描述】：&nbsp;</span><br></div><div><!--StartFragment --><div><a href=\"https://admin.8591.com.tw/admin.php?module=wareDeal&amp;action=serverRecord&amp;id=2240000712\" target=\"_blank\" rel=\"noopener\">https://admin.8591.com.tw/admin.php?module=wareDeal&amp;action=serverRecord&amp;id=2240000712</a></div><div>賣家反饋交易完成後未收到款項，客服查詢帳戶明細屬實，麻煩轉交工程師確認處理，謝謝<br><br></div></div><div><span style=\"background-color: #ffffff;\"><br></span></div><div><p style=\"word-break: break-word; margin: 0px 0px 1em; padding: 0px; list-style: none; line-height: inherit; cursor: text; color: #182b50; font-family: 'Microsoft YaHei', 'Helvetica Neue', 'PingFang SC', sans-serif; font-size: 14px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: left; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: #ffffff; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\">【复现步驟】：</p><p><span style=\"color: #c7c7cc;\"><br></span></p><p><span style=\"color: #c7c7cc;\">示例：</span><span style=\"color: #c7c7cc;\">1、打开首页-&gt;查找手机游戏 &nbsp; &nbsp;</span><span style=\"color: #c7c7cc;\">2、点击某一类手游进入详情页</span></p><p><span style=\"color: #c7c7cc;\"><br></span></p><p style=\"word-break: break-word; margin: 0px 0px 1em; padding: 0px; list-style: none; line-height: inherit; cursor: text; color: #182b50; font-family: 'Microsoft YaHei', 'Helvetica Neue', 'PingFang SC', sans-serif; font-size: 14px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: left; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: #ffffff; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\"><span style=\"word-break: break-word; line-height: inherit; color: #182b50; font-family: 'Microsoft YaHei', 'Helvetica Neue', 'PingFang SC', sans-serif; font-size: 14px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: left; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: #ffffff; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; float: none; display: inline !important;\">【相关资料】：</span></p><p class=\"tox-clear-float\"><img src=\"/tfl/captures/2024-08/tapd_23402991_base64_1724741451_156.png\" width=\"80%\"></p><div style=\"word-break: break-word; margin: 0px 0px 1em; padding: 0px; line-height: inherit; cursor: text;\">【实际结果】：</div><div style=\"word-break: break-word; margin: 0px 0px 1em; padding: 0px; line-height: inherit; cursor: text;\"><br></div><div style=\"word-break: break-word; margin: 0px 0px 1em; padding: 0px; line-height: inherit; cursor: text;\"><span style=\"color: #c7c7cc;\">异常的运行结果库存修改异常</span></div><div style=\"word-break: break-word; margin: 0px 0px 1em; padding: 0px; line-height: inherit; cursor: text;\"><br></div><div style=\"word-break: break-word; margin: 0px 0px 1em; padding: 0px; line-height: inherit; cursor: text;\">【期望结果】：</div><div style=\"word-break: break-word; margin: 0px 0px 1em; padding: 0px; line-height: inherit; cursor: text;\"><br></div><div style=\"word-break: break-word; margin: 0px 0px 1em; padding: 0px; line-height: inherit; cursor: text;\"><span style=\"word-break: break-word; line-height: inherit; color: #c7c7cc; font-family: 'Microsoft YaHei', 'Helvetica Neue', 'PingFang SC', sans-serif; font-size: 14px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: left; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: #ffffff; float: none; display: inline !important;\">示例：客户能够正常修改商品的库存。 请尽快处理（如果是很要紧的， 则可以写 期望一个工作日内解决这种）。</span></div></div>",
  "iteration_id": "0",
  "custom_field_three": "线上缺陷",
  "severity": "serious",
  "priority": "",
  "custom_field_four": "",
  "current_owner": "T8曾令涛;T8龙文娇;T8列仲宇;",
  "cc": "",
  "de": "",
  "te": "",
  "custom_field_6": "",
  "platform": "PC",
  "bugtype": "",
  "originphase": "",
  "source": "",
  "custom_field_one": "",
  "description_type": "1",
  "project_id": "23402991",
  "is_draft": "0",
  "begin": "",
  "due": "",
  "status": "new",
  "reporter": "T8黄玲",
  "flows": "new",
  "resolution": "",
  "resolved": "",
  "closed": "",
  "in_progress_time": "",
  "verify_time": "",
  "reject_time": "",
  "audit_time": "",
  "suspend_time": "",
  "secret": "asdfasdfsadfasdf",
  "rio_token": "",
  "devproxy_host": "http://websocket-proxy",
  "queue_id": "281687330",
  "event_id": "168547800",
  "created": "2024-08-27 14:54:15"
}`
	var event BugCreateEvent
	err := json.Unmarshal([]byte(rawBody), &event)
	assert.NoError(t, err)
	spew.Dump(event)
}
