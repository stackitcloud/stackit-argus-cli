package yamltojson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReceivers(t *testing.T) {
	value := "    name: \"receiver1\"\n    email_configs:\n      - to: \"to\"\n        from: \"from\"\n        " +
		"smarthost: \"smarthost\"\n        auth_username: \"username\"\n        auth_password: \"pass\"\n        " +
		"auth_identity: \"identity\"\n      - to: \"to\"\n        from: \"from\"\n        smarthost: \"smarthost\"\n" +
		"        auth_username: \"username\"\n        auth_password: \"pass\"\n        auth_identity: \"identity\"\n" +
		"    opsgenie_configs:\n      - api_key: \"key\"\n        api_url: \"url\"\n        tags: \"tag\"\n" +
		"      - api_key: \"key\"\n        api_url: \"url\"\n        tags: \"tag\"\n    webhook_configs:\n" +
		"      - url: \"url\"\n        ms_teams: true\n      - url: \"url\"\n        ms_teams: false"

	expectedResult := "{\"name\":\"receiver1\",\"emailConfigs\":[{\"to\":\"to\",\"from\":\"from\",\"smarthost\":" +
		"\"smarthost\",\"authUsername\":\"username\",\"authPassword\":\"pass\",\"authIdentity\":\"identity\"}," +
		"{\"to\":\"to\",\"from\":\"from\",\"smarthost\":\"smarthost\",\"authUsername\":\"username\"," +
		"\"authPassword\":\"pass\",\"authIdentity\":\"identity\"}],\"opsgenieConfigs\":[{\"apiKey\":\"key\"," +
		"\"apiUrl\":\"url\",\"tags\":\"tag\"},{\"apiKey\":\"key\",\"apiUrl\":\"url\",\"tags\":\"tag\"}]," +
		"\"webhookConfigs\":[{\"url\":\"url\",\"msTeams\":true},{\"url\":\"url\",\"msTeams\":false}]}"

	res, err := Receivers([]byte(value))
	assert.NoError(t, err, "test working case")
	assert.Equal(t, []byte(expectedResult), res)

	// test case with wrong format

	value = "name \"receiver1\""

	_, err = Receivers([]byte(value))
	assert.Error(t, err, "test case with wrong format")
}

func TestRoutes(t *testing.T) {
	value := "receiver: \"receiver\"\ngroup_by:\n  - \"group\"\n  - \"by\"\ngroup_wait: \"group wait\"\n" +
		"group_interval: \"group_interval\"\nrepeat_interval: \"repeat_interval\"\nmatch:\n  \"key1\": \"value1\"" +
		"\n  \"key2\": \"value2\"\nmatch_re:\n  \"key1\": \"value1\"\n  \"key2\": \"value2\"\nmatchers:\n" +
		"  - \"matcher1\"\n  - \"matcher2\"\nroutes:\n  - receiver: \"receiver\"\n    group_by:\n      - \"group\"" +
		"\n      - \"by\"\n    group_wait: \"group wait\"\n    group_interval: \"group_interval\"\n    " +
		"repeat_interval: \"repeat_interval\"\n    match:\n      \"key1\": \"value1\"\n      \"key2\": \"value2\"\n" +
		"    match_re:\n      \"key1\": \"value1\"\n      \"key2\": \"value2\"\n    matchers:\n      - \"matcher1\"\n" +
		"      - \"matcher2\"\n    routes:\n      - receiver: \"receiver\"\n        group_by:\n" +
		"          - \"group\"\n          - \"by\"\n        group_wait: \"group wait\"\n        " +
		"group_interval: \"group_interval\"\n        repeat_interval: \"repeat_interval\"\n        " +
		"match:\n          \"key1\": \"value1\"\n          \"key2\": \"value2\"\n        match_re:\n          " +
		"\"key1\": \"value1\"\n          \"key2\": \"value2\"\n        matchers:\n          - \"matcher1\"\n" +
		"          - \"matcher2\"\n  - receiver: \"receiver\"\n    group_by:\n      - \"group\"\n      - \"by\"\n" +
		"    group_wait: \"group wait\"\n    group_interval: \"group_interval\"\n    " +
		"repeat_interval: \"repeat_interval\"\n    match:\n      \"key1\": \"value1\"\n      \"key2\": \"value2\"\n" +
		"    match_re:\n      \"key1\": \"value1\"\n      \"key2\": \"value2\"\n    matchers:\n      - \"matcher1\"\n" +
		"      - \"matcher2\"\n    routes:\n      - receiver: \"receiver\"\n        group_by:\n          " +
		"- \"group\"\n          - \"by\"\n        group_wait: \"group wait\"\n        " +
		"group_interval: \"group_interval\"\n        repeat_interval: \"repeat_interval\"\n        " +
		"match:\n          \"key1\": \"value1\"\n          \"key2\": \"value2\"\n        match_re:\n          " +
		"\"key1\": \"value1\"\n          \"key2\": \"value2\"\n        matchers:\n          - \"matcher1\"\n" +
		"          - \"matcher2\""

	expectedResult := "{\"receiver\":\"receiver\",\"groupBy\":[\"group\",\"by\"],\"groupWait\":\"group wait\"," +
		"\"groupInterval\":\"group_interval\",\"repeatInterval\":\"repeat_interval\",\"match\":{\"key1\":\"value1\"," +
		"\"key2\":\"value2\"},\"matchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"matchers\":[\"matcher1\"," +
		"\"matcher2\"],\"routes\":[{\"receiver\":\"receiver\",\"groupBy\":[\"group\",\"by\"]," +
		"\"groupWait\":\"group wait\",\"groupInterval\":\"group_interval\",\"repeatInterval\":\"repeat_interval\"," +
		"\"match\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"matchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"}," +
		"\"matchers\":[\"matcher1\",\"matcher2\"],\"routes\":[{\"receiver\":\"receiver\"," +
		"\"groupBy\":[\"group\",\"by\"],\"groupWait\":\"group wait\",\"groupInterval\":\"group_interval\"," +
		"\"repeatInterval\":\"repeat_interval\",\"match\":{\"key1\":\"value1\",\"key2\":\"value2\"}," +
		"\"matchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"matchers\":[\"matcher1\",\"matcher2\"]," +
		"\"routes\":null}]},{\"receiver\":\"receiver\",\"groupBy\":[\"group\",\"by\"],\"groupWait\":\"group wait\"," +
		"\"groupInterval\":\"group_interval\",\"repeatInterval\":\"repeat_interval\",\"match\":{\"key1\":\"value1\"," +
		"\"key2\":\"value2\"},\"matchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"matchers\":[\"matcher1\"," +
		"\"matcher2\"],\"routes\":[{\"receiver\":\"receiver\",\"groupBy\":[\"group\",\"by\"]," +
		"\"groupWait\":\"group wait\",\"groupInterval\":\"group_interval\",\"repeatInterval\":\"repeat_interval\"," +
		"\"match\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"matchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"}," +
		"\"matchers\":[\"matcher1\",\"matcher2\"],\"routes\":null}]}]}"

	res, err := Routes([]byte(value))
	assert.NoError(t, err, "test working case")
	assert.Equal(t, []byte(expectedResult), res)

	// test case with wrong format

	value = "receiver \"receiver\""

	_, err = Routes([]byte(value))
	assert.Error(t, err, "test case with wrong format")
}

func TestAlertConfig(t *testing.T) {
	value := "global:\n  resolve_timeout: \"5m\"\n  smtp_from: \"form smtp\"\n  smtp_smarthost: \"very smart\"\n" +
		"  smtp_auth_username: \"name\"\n  smtp_auth_password: \"password\"\n  " +
		"smtp_auth_identity: \"smart identity\"\n  opsgenie_api_key: \"api key\"\n  opsgenie_api_url: \"api url\"\n" +
		"route:\n  receiver: \"receiver\"\n  group_by:\n    - \"group\"\n    - \"by\"\n  group_wait: \"group wait\"\n" +
		"  group_interval: \"group_interval\"\n  repeat_interval: \"repeat_interval\"\n  match:\n    " +
		"\"key1\": \"value1\"\n    \"key2\": \"value2\"\n  match_re:\n    \"key1\": \"value1\"\n    " +
		"\"key2\": \"value2\"\n  matchers:\n    - \"matcher1\"\n    - \"matcher2\"\n  routes:\n    " +
		"- receiver: \"receiver\"\n      group_by:\n        - \"group\"\n        - \"by\"\n      " +
		"group_wait: \"group wait\"\n      group_interval: \"group_interval\"\n      " +
		"repeat_interval: \"repeat_interval\"\n      match:\n        \"key1\": \"value1\"\n        " +
		"\"key2\": \"value2\"\n      match_re:\n        \"key1\": \"value1\"\n        " +
		"\"key2\": \"value2\"\n      matchers:\n        - \"matcher1\"\n        - \"matcher2\"\n      " +
		"routes:\n        - receiver: \"receiver\"\n          group_by:\n            - \"group\"\n            " +
		"- \"by\"\n          group_wait: \"group wait\"\n          group_interval: \"group_interval\"\n          " +
		"repeat_interval: \"repeat_interval\"\n          match:\n            \"key1\": \"value1\"\n            " +
		"\"key2\": \"value2\"\n          match_re:\n            \"key1\": \"value1\"\n            " +
		"\"key2\": \"value2\"\n          matchers:\n            - \"matcher1\"\n            - \"matcher2\"\n    " +
		"- receiver: \"receiver\"\n      group_by:\n        - \"group\"\n        - \"by\"\n      " +
		"group_wait: \"group wait\"\n      group_interval: \"group_interval\"\n      " +
		"repeat_interval: \"repeat_interval\"\n      match:\n        \"key1\": \"value1\"\n        " +
		"\"key2\": \"value2\"\n      match_re:\n        \"key1\": \"value1\"\n        \"key2\": \"value2\"\n      " +
		"matchers:\n        - \"matcher1\"\n        - \"matcher2\"\n      routes:\n        " +
		"- receiver: \"receiver\"\n          group_by:\n            - \"group\"\n            " +
		"- \"by\"\n          group_wait: \"group wait\"\n          group_interval: \"group_interval\"\n          " +
		"repeat_interval: \"repeat_interval\"\n          match:\n            \"key1\": \"value1\"\n            " +
		"\"key2\": \"value2\"\n          match_re:\n            \"key1\": \"value1\"\n            " +
		"\"key2\": \"value2\"\n          matchers:\n            - \"matcher1\"\n            " +
		"- \"matcher2\"\nreceivers:\n  - name: \"receiver1\"\n    email_configs:\n      - to: \"to\"\n        " +
		"from: \"from\"\n        smarthost: \"smarthost\"\n        auth_username: \"username\"\n        " +
		"auth_password: \"pass\"\n        auth_identity: \"identity\"\n      - to: \"to\"\n        from: \"from\"\n" +
		"        smarthost: \"smarthost\"\n        auth_username: \"username\"\n        auth_password: \"pass\"\n" +
		"        auth_identity: \"identity\"\n    opsgenie_configs:\n      - api_key: \"key\"\n        " +
		"api_url: \"url\"\n        tags: \"tag\"\n      - api_key: \"key\"\n        api_url: \"url\"\n        " +
		"tags: \"tag\"\n    webhook_configs:\n      - url: \"url\"\n        ms_teams: true\n      - url: \"url\"\n" +
		"        ms_teams: false\n  - name: \"receiver2\"\n    email_configs:\n      - to: \"to\"\n        " +
		"from: \"from\"\n        smarthost: \"smarthost\"\n        auth_username: \"username\"\n        " +
		"auth_password: \"pass\"\n        auth_identity: \"identity\"\n      - to: \"to\"\n        from: \"from\"\n" +
		"        smarthost: \"smarthost\"\n        auth_username: \"username\"\n        auth_password: \"pass\"\n" +
		"        auth_identity: \"identity\"\n    opsgenie_configs:\n      - api_key: \"key\"\n        " +
		"api_url: \"url\"\n        tags: \"tag\"\n      - api_key: \"key\"\n        api_url: \"url\"\n        " +
		"tags: \"tag\"\n    webhook_configs:\n      - url: \"url\"\n        ms_teams: true\n      - url: \"url\"\n" +
		"        ms_teams: false\ninhibit_rules:\n  - source_match:\n      \"key1\": \"value1\"\n      " +
		"\"key2\": \"value2\"\n    source_match_re:\n      \"key1\": \"value1\"\n      \"key2\": \"value2\"\n    " +
		"target_match:\n      \"key1\": \"value1\"\n      \"key2\": \"value2\"\n    target_match_re:\n      " +
		"\"key1\": \"value1\"\n      \"key2\": \"value2\"\n    equal:\n      - \"1\"\n      - \"2\"\n  " +
		"- source_match:\n      \"key1\": \"value1\"\n      \"key2\": \"value2\"\n    source_match_re:\n      " +
		"\"key1\": \"value1\"\n      \"key2\": \"value2\"\n    target_match:\n      \"key1\": \"value1\"\n      " +
		"\"key2\": \"value2\"\n    target_match_re:\n      \"key1\": \"value1\"\n      \"key2\": \"value2\"\n    " +
		"equal:\n      - \"1\"\n      - \"2\""

	expectedResult := "{\"global\":{\"resolveTimeout\":\"5m\",\"smtpFrom\":\"form smtp\",\"" +
		"smtpSmarthost\":\"very smart\",\"smtpAuthUsername\":\"name\",\"smtpAuthPassword\":\"password\"," +
		"\"smtpAuthIdentity\":\"smart identity\",\"opsgenieApiKey\":\"api key\",\"opsgenieApiUrl\":\"api url\"}," +
		"\"route\":{\"receiver\":\"receiver\",\"groupBy\":[\"group\",\"by\"],\"groupWait\":\"group wait\"," +
		"\"groupInterval\":\"group_interval\",\"repeatInterval\":\"repeat_interval\",\"match\":{\"key1\":\"value1\"," +
		"\"key2\":\"value2\"},\"matchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"matchers\":[\"matcher1\"," +
		"\"matcher2\"],\"routes\":[{\"receiver\":\"receiver\",\"groupBy\":[\"group\",\"by\"],\"" +
		"groupWait\":\"group wait\",\"groupInterval\":\"group_interval\",\"repeatInterval\":\"repeat_interval\"," +
		"\"match\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"matchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"" +
		"},\"matchers\":[\"matcher1\",\"matcher2\"],\"routes\":[{\"receiver\":\"receiver\"," +
		"\"groupBy\":[\"group\",\"by\"],\"groupWait\":\"group wait\",\"groupInterval\":\"group_interval\"," +
		"\"repeatInterval\":\"repeat_interval\",\"match\":{\"key1\":\"value1\",\"key2\":\"value2\"}," +
		"\"matchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"matchers\":[\"matcher1\",\"matcher2\"]," +
		"\"routes\":null}]},{\"receiver\":\"receiver\",\"groupBy\":[\"group\",\"by\"],\"groupWait\":\"group wait\"," +
		"\"groupInterval\":\"group_interval\",\"repeatInterval\":\"repeat_interval\",\"match\":{\"key1\":\"value1\"," +
		"\"key2\":\"value2\"},\"matchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"matchers\":[\"matcher1\"," +
		"\"matcher2\"],\"routes\":[{\"receiver\":\"receiver\",\"groupBy\":[\"group\",\"by\"],\"" +
		"groupWait\":\"group wait\",\"groupInterval\":\"group_interval\",\"repeatInterval\":\"repeat_interval\"," +
		"\"match\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"matchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"}," +
		"\"matchers\":[\"matcher1\",\"matcher2\"],\"routes\":null}]}]},\"receivers\":[{\"name\":\"receiver1\"," +
		"\"emailConfigs\":[{\"to\":\"to\",\"from\":\"from\",\"smarthost\":\"smarthost\",\"authUsername\":\"username\"," +
		"\"authPassword\":\"pass\",\"authIdentity\":\"identity\"},{\"to\":\"to\",\"from\":\"from\"," +
		"\"smarthost\":\"smarthost\",\"authUsername\":\"username\",\"authPassword\":\"pass\"," +
		"\"authIdentity\":\"identity\"}],\"opsgenieConfigs\":[{\"apiKey\":\"key\",\"apiUrl\":\"url\"," +
		"\"tags\":\"tag\"},{\"apiKey\":\"key\",\"apiUrl\":\"url\",\"tags\":\"tag\"}]," +
		"\"webhookConfigs\":[{\"url\":\"url\",\"msTeams\":true},{\"url\":\"url\",\"msTeams\":false}]}," +
		"{\"name\":\"receiver2\",\"emailConfigs\":[{\"to\":\"to\",\"from\":\"from\",\"smarthost\":\"smarthost\"," +
		"\"authUsername\":\"username\",\"authPassword\":\"pass\",\"authIdentity\":\"identity\"},{\"to\":\"to\"," +
		"\"from\":\"from\",\"smarthost\":\"smarthost\",\"authUsername\":\"username\",\"authPassword\":\"pass\"," +
		"\"authIdentity\":\"identity\"}],\"opsgenieConfigs\":[{\"apiKey\":\"key\",\"apiUrl\":\"url\"," +
		"\"tags\":\"tag\"},{\"apiKey\":\"key\",\"apiUrl\":\"url\",\"tags\":\"tag\"}]," +
		"\"webhookConfigs\":[{\"url\":\"url\",\"msTeams\":true},{\"url\":\"url\",\"msTeams\":false}]}]," +
		"\"inhibitRules\":[{\"sourceMatch\":{\"key1\":\"value1\",\"key2\":\"value2\"}," +
		"\"sourceMatchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"targetMatch\":{\"key1\":\"value1\"," +
		"\"key2\":\"value2\"},\"targetMatchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"equal\":[\"1\"," +
		"\"2\"]},{\"sourceMatch\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"sourceMatchRe\":{\"key1\":\"value1\"," +
		"\"key2\":\"value2\"},\"targetMatch\":{\"key1\":\"value1\",\"key2\":\"value2\"}," +
		"\"targetMatchRe\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"equal\":[\"1\",\"2\"]}]}"

	res, err := AlertConfig([]byte(value))
	assert.NoError(t, err, "test working case")
	assert.Equal(t, []byte(expectedResult), res)

	// test case with wrong format

	value = "global\n  resolve_timeout \"5m\""

	_, err = AlertConfig([]byte(value))
	assert.Error(t, err, "test case with wrong format")
}

func TestScrapeConfig(t *testing.T) {
	value := "static_configs:\n  - targets:\n      - \"target1\"\n      - \"target2\"\n    labels:\n      \"key1\":\n" +
		"        - \"value1\"\n        - \"value2\"\n      \"key2\":\n        - \"value1\"\n        - \"value2\"\n" +
		"  - targets:\n      - \"target1\"\n      - \"target2\"\n    labels:\n      \"key1\":\n        - \"value1\"\n" +
		"        - \"value2\"\n      \"key2\":\n        - \"value1\"\n        - \"value2\"\njob_name: \"job\"\n" +
		"scheme: \"https\"\nscrape_interval: \"5m\"\nscrape_timeout: \"10s\"\nmetrics_path: \"/metrics\"\n" +
		"sample_limit: 42\nbasic_auth:\n  username: \"username\"\n  password: \"password\"\noauth2:\n  " +
		"client_id: \"id\"\n  client_secret: \"secret\"\n  token_url: \"url\"\n  scopes:\n    - \"1\"\n    " +
		"- \"2\"\n  tls_config:\n    insecure_skip_verify: true\ntls_config:\n  insecure_skip_verify: true\n" +
		"bearer_token: \"token\"\nmetric_relabel_configs:\n  - source_labels:\n      - \"label1\"\n      " +
		"- \"label2\"\n    separator: \";\"\n    target_label: \"label\"\n    regex: \".*\"\n    modulus: 42\n" +
		"    replacement: \"$1\"\n    action: \"keep\"\n  - source_labels:\n      - \"label1\"\n      " +
		"- \"label2\"\n    separator: \";\"\n    target_label: \"label\"\n    regex: \".*\"\n    modulus: 42\n" +
		"    replacement: \"$1\"\n    action: \"keep\"\nparams:\n  \"key1\":\n    - \"value1\"\n    - \"value2\"\n" +
		"  \"key2\":\n    - \"value1\"\n    - \"value2\"\nhttp_sd_configs:\n  - url: \"url\"\n    " +
		"refresh_interval: \"5m\"\n    basic_auth:\n      username: \"username\"\n      password: \"password\"\n" +
		"    oauth2:\n      client_id: \"id\"\n      client_secret: \"secret\"\n      token_url: \"url\"\n      " +
		"scopes:\n        - \"1\"\n        - \"2\"\n      tls_config:\n        insecure_skip_verify: true\n    " +
		"tls_config:\n      insecure_skip_verify: true\n  - url: \"url\"\n    refresh_interval: \"5m\"\n    " +
		"basic_auth:\n      username: \"username\"\n      password: \"password\"\n    oauth2:\n      " +
		"client_id: \"id\"\n      client_secret: \"secret\"\n      token_url: \"url\"\n      scopes:\n        " +
		"- \"1\"\n        - \"2\"\n      tls_config:\n        insecure_skip_verify: true\n    tls_config:\n      " +
		"insecure_skip_verify: true\nhonor_labels: true\nhonor_timestamps: true"

	expectedResult := "{\"staticConfigs\":[{\"targets\":[\"target1\",\"target2\"],\"labels\":{\"key1\":[\"value1\"," +
		"\"value2\"],\"key2\":[\"value1\",\"value2\"]}},{\"targets\":[\"target1\",\"target2\"],\"labels\":" +
		"{\"key1\":[\"value1\",\"value2\"],\"key2\":[\"value1\",\"value2\"]}}],\"jobName\":\"job\"," +
		"\"scheme\":\"https\",\"scrapeInterval\":\"5m\",\"scrapeTimeout\":\"10s\",\"metricsPath\":\"/metrics\"," +
		"\"sampleLimit\":42,\"basicAuth\":{\"username\":\"username\",\"password\":\"password\"},\"oauth2\":" +
		"{\"clientId\":\"id\",\"clientSecret\":\"secret\",\"tokenUrl\":\"url\",\"scopes\":[\"1\",\"2\"],\"tlsConfig\"" +
		":{\"insecureSkipVerify\":true}},\"tlsConfig\":{\"insecureSkipVerify\":true},\"bearerToken\":\"token\"," +
		"\"metricRelabelConfigs\":[{\"sourceLabels\":[\"label1\",\"label2\"],\"separator\":\";\",\"targetLabel\":" +
		"\"label\",\"regex\":\".*\",\"modulus\":42,\"replacement\":\"$1\",\"action\":\"keep\"},{\"sourceLabels\":" +
		"[\"label1\",\"label2\"],\"separator\":\";\",\"targetLabel\":\"label\",\"regex\":\".*\",\"modulus\":42," +
		"\"replacement\":\"$1\",\"action\":\"keep\"}],\"httpSdConfigs\":[{\"url\":\"url\",\"refreshInterval\":\"5m\"," +
		"\"basicAuth\":{\"username\":\"username\",\"password\":\"password\"},\"oauth2\":{\"clientId\":\"id\"," +
		"\"clientSecret\":\"secret\",\"tokenUrl\":\"url\",\"scopes\":[\"1\",\"2\"],\"tlsConfig\":{\"" +
		"insecureSkipVerify\":true}},\"tlsConfig\":{\"insecureSkipVerify\":true}},{\"url\":\"url\",\"" +
		"refreshInterval\":\"5m\",\"basicAuth\":{\"username\":\"username\",\"password\":\"password\"},\"" +
		"oauth2\":{\"clientId\":\"id\",\"clientSecret\":\"secret\",\"tokenUrl\":\"url\",\"scopes\":[\"1\",\"2\"]," +
		"\"tlsConfig\":{\"insecureSkipVerify\":true}},\"tlsConfig\":{\"insecureSkipVerify\":true}}],\"params\":" +
		"{\"key1\":[\"value1\",\"value2\"],\"key2\":[\"value1\",\"value2\"]},\"honorLabels\":true," +
		"\"honorTimestamps\":true}"

	res, err := ScrapeConfig([]byte(value))
	assert.NoError(t, err, "test working case")
	assert.Equal(t, []byte(expectedResult), res)

	// test case with wrong format

	value = "staticConfigs \"receiver\""

	_, err = ScrapeConfig([]byte(value))
	assert.Error(t, err, "test case with wrong format")
}

func TestScrapeConfigs(t *testing.T) {
	value := "- static_configs:\n    - targets:\n        - \"target1\"\n      labels:\n        \"key1\":\n          " +
		"- \"value1\"\n  job_name: \"job\"\n  scheme: \"https\"\n  scrape_interval: \"5m\"\n  scrape_timeout: \"" +
		"10s\"\n  metrics_path: \"/metrics\"\n  sample_limit: 42\n  basic_auth:\n    username: \"username\"\n    " +
		"password: \"password\"\n  oauth2:\n    client_id: \"id\"\n    client_secret: \"secret\"\n    token_url: \"" +
		"url\"\n    scopes:\n      - \"1\"\n    tls_config:\n      insecure_skip_verify: true\n  tls_config:\n    " +
		"insecure_skip_verify: true\n  bearer_token: \"token\"\n  metric_relabel_configs:\n    - source_labels:\n" +
		"        - \"label1\"\n      separator: \";\"\n      target_label: \"label\"\n      regex: \".*\"\n      " +
		"modulus: 42\n      replacement: \"$1\"\n      action: \"keep\"\n    - source_labels:\n        " +
		"- \"label1\"\n      separator: \";\"\n      target_label: \"label\"\n      regex: \".*\"\n      " +
		"modulus: 42\n      replacement: \"$1\"\n      action: \"keep\"\n  params:\n    \"key1\":\n      " +
		"- \"value1\"\n  http_sd_configs:\n    - url: \"url\"\n      refresh_interval: \"5m\"\n      basic_auth:\n" +
		"        username: \"username\"\n        password: \"password\"\n      oauth2:\n        client_id: \"id\"\n" +
		"        client_secret: \"secret\"\n        token_url: \"url\"\n        scopes:\n          - \"1\"\n" +
		"        tls_config:\n          insecure_skip_verify: true\n      tls_config:\n        " +
		"insecure_skip_verify: true\n    - url: \"url\"\n      refresh_interval: \"5m\"\n      basic_auth:\n" +
		"        username: \"username\"\n        password: \"password\"\n      oauth2:\n        client_id: \"id\"\n" +
		"        client_secret: \"secret\"\n        token_url: \"url\"\n        scopes:\n          - \"1\"\n" +
		"        tls_config:\n          insecure_skip_verify: true\n      tls_config:\n        insecure_skip_verify:" +
		" true\n  honor_labels: true\n  honor_timestamps: true\n- static_configs:\n    - targets:\n        " +
		"- \"target1\"\n      labels:\n        \"key1\":\n          - \"value1\"\n  job_name: \"job\"\n  scheme: " +
		"\"https\"\n  scrape_interval: \"5m\"\n  scrape_timeout: \"10s\"\n  metrics_path: \"/metrics\"\n  " +
		"sample_limit: 42\n  basic_auth:\n    username: \"username\"\n    password: \"password\"\n  oauth2:\n    " +
		"client_id: \"id\"\n    client_secret: \"secret\"\n    token_url: \"url\"\n    scopes:\n      - \"1\"\n    " +
		"tls_config:\n      insecure_skip_verify: true\n  tls_config:\n    insecure_skip_verify: true\n  bearer_token: " +
		"\"token\"\n  metric_relabel_configs:\n    - source_labels:\n        - \"label1\"\n      separator: \";\"\n" +
		"      target_label: \"label\"\n      regex: \".*\"\n      modulus: 42\n      replacement: \"$1\"\n      " +
		"action: \"keep\"\n    - source_labels:\n        - \"label1\"\n      separator: \";\"\n      target_label: " +
		"\"label\"\n      regex: \".*\"\n      modulus: 42\n      replacement: \"$1\"\n      action: \"keep\"\n  " +
		"params:\n    \"key1\":\n      - \"value1\"\n  http_sd_configs:\n    - url: \"url\"\n      refresh_interval: " +
		"\"5m\"\n      basic_auth:\n        username: \"username\"\n        password: \"password\"\n      oauth2:\n" +
		"        client_id: \"id\"\n        client_secret: \"secret\"\n        token_url: \"url\"\n        scopes:\n" +
		"          - \"1\"\n        tls_config:\n          insecure_skip_verify: true\n      tls_config:\n        " +
		"insecure_skip_verify: true\n    - url: \"url\"\n      refresh_interval: \"5m\"\n      basic_auth:\n        " +
		"username: \"username\"\n        password: \"password\"\n      oauth2:\n        client_id: \"id\"\n        " +
		"client_secret: \"secret\"\n        token_url: \"url\"\n        scopes:\n          - \"1\"\n        " +
		"tls_config:\n          insecure_skip_verify: true\n      tls_config:\n        insecure_skip_verify: true\n  " +
		"honor_labels: true\n  honor_timestamps: true"

	expectedResult := "[{\"staticConfigs\":[{\"targets\":[\"target1\"],\"labels\":{\"key1\":[\"value1\"]}}],\"" +
		"jobName\":\"job\",\"scheme\":\"https\",\"scrapeInterval\":\"5m\",\"scrapeTimeout\":\"10s\",\"metricsPath\"" +
		":\"/metrics\",\"sampleLimit\":42,\"basicAuth\":{\"username\":\"username\",\"password\":\"password\"}," +
		"\"oauth2\":{\"clientId\":\"id\",\"clientSecret\":\"secret\",\"tokenUrl\":\"url\",\"scopes\":[\"1\"]," +
		"\"tlsConfig\":{\"insecureSkipVerify\":true}},\"tlsConfig\":{\"insecureSkipVerify\":true},\"bearerToken\":" +
		"\"token\",\"metricRelabelConfigs\":[{\"sourceLabels\":[\"label1\"],\"separator\":\";\",\"targetLabel\":" +
		"\"label\",\"regex\":\".*\",\"modulus\":42,\"replacement\":\"$1\",\"action\":\"keep\"},{\"sourceLabels\"" +
		":[\"label1\"],\"separator\":\";\",\"targetLabel\":\"label\",\"regex\":\".*\",\"modulus\":42,\"replacement\"" +
		":\"$1\",\"action\":\"keep\"}],\"httpSdConfigs\":[{\"url\":\"url\",\"refreshInterval\":\"5m\",\"basicAuth\"" +
		":{\"username\":\"username\",\"password\":\"password\"},\"oauth2\":{\"clientId\":\"id\",\"clientSecret\":" +
		"\"secret\",\"tokenUrl\":\"url\",\"scopes\":[\"1\"],\"tlsConfig\":{\"insecureSkipVerify\":true}},\"tlsConfig\"" +
		":{\"insecureSkipVerify\":true}},{\"url\":\"url\",\"refreshInterval\":\"5m\",\"basicAuth\":{\"username\":\"" +
		"username\",\"password\":\"password\"},\"oauth2\":{\"clientId\":\"id\",\"clientSecret\":\"secret\",\"" +
		"tokenUrl\":\"url\",\"scopes\":[\"1\"],\"tlsConfig\":{\"insecureSkipVerify\":true}},\"tlsConfig\":" +
		"{\"insecureSkipVerify\":true}}],\"params\":{\"key1\":[\"value1\"]},\"honorLabels\":true,\"honorTimestamps\"" +
		":true},{\"staticConfigs\":[{\"targets\":[\"target1\"],\"labels\":{\"key1\":[\"value1\"]}}],\"jobName\":" +
		"\"job\",\"scheme\":\"https\",\"scrapeInterval\":\"5m\",\"scrapeTimeout\":\"10s\",\"metricsPath\":" +
		"\"/metrics\",\"sampleLimit\":42,\"basicAuth\":{\"username\":\"username\",\"password\":\"password\"}," +
		"\"oauth2\":{\"clientId\":\"id\",\"clientSecret\":\"secret\",\"tokenUrl\":\"url\",\"scopes\":[\"1\"]," +
		"\"tlsConfig\":{\"insecureSkipVerify\":true}},\"tlsConfig\":{\"insecureSkipVerify\":true},\"bearerToken\":" +
		"\"token\",\"metricRelabelConfigs\":[{\"sourceLabels\":[\"label1\"],\"separator\":\";\",\"targetLabel\":" +
		"\"label\",\"regex\":\".*\",\"modulus\":42,\"replacement\":\"$1\",\"action\":\"keep\"},{\"sourceLabels\":[\"" +
		"label1\"],\"separator\":\";\",\"targetLabel\":\"label\",\"regex\":\".*\",\"modulus\":42,\"replacement\":" +
		"\"$1\",\"action\":\"keep\"}],\"httpSdConfigs\":[{\"url\":\"url\",\"refreshInterval\":\"5m\",\"basicAuth\":" +
		"{\"username\":\"username\",\"password\":\"password\"},\"oauth2\":{\"clientId\":\"id\",\"clientSecret\":" +
		"\"secret\",\"tokenUrl\":\"url\",\"scopes\":[\"1\"],\"tlsConfig\":{\"insecureSkipVerify\":true}},\"tlsConfig\"" +
		":{\"insecureSkipVerify\":true}},{\"url\":\"url\",\"refreshInterval\":\"5m\",\"basicAuth\":{\"username\":" +
		"\"username\",\"password\":\"password\"},\"oauth2\":{\"clientId\":\"id\",\"clientSecret\":\"secret\"," +
		"\"tokenUrl\":\"url\",\"scopes\":[\"1\"],\"tlsConfig\":{\"insecureSkipVerify\":true}},\"tlsConfig\":{" +
		"\"insecureSkipVerify\":true}}],\"params\":{\"key1\":[\"value1\"]},\"honorLabels\":true,\"honorTimestamps\"" +
		":true}]"

	res, err := ScrapeConfigs([]byte(value))
	assert.NoError(t, err, "test working case")
	assert.Equal(t, []byte(expectedResult), res)

	// test case with wrong format

	value = "staticConfigs \"receiver\""

	_, err = ScrapeConfigs([]byte(value))
	assert.Error(t, err, "test case with wrong format")
}
