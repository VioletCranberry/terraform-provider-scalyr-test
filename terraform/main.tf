locals {
  my_alert_addr = "my_email_addr@something"
}

resource "scalyr_alert_rule" "test_alert_rule_1" {
  description              = "Test Alert Rule #1"
  grace_period_minutes     = 0
  renotify_period_minutes  = 60
  resolution_delay_minutes = 0
  trigger                  = "count:5 minutes(test) > 0"
}

resource "scalyr_alert_rule_json" "test_alert_rule_2" {
  description = "Test Alert Rule #2"
  json = jsonencode({
    trigger                = "count:60 minutes(test) > 0"
    gracePeriodMinutes     = 0
    renotifyPeriodMinutes  = 60
    resolutionDelayMinutes = 30
  })
}

resource "scalyr_alert_group" "test_alert_group" {
  description = "Test Alert Group #2"
  alert_address = [
    local.my_alert_addr
  ]

  alert_rule {
    description              = "Test Nested Alert Rule #1"
    grace_period_minutes     = 0
    renotify_period_minutes  = 30
    resolution_delay_minutes = 0
    trigger                  = "count:5 minutes(test) > 0"
  }

  alert_rule {
    description              = "Test Nested Alert Rule #2"
    grace_period_minutes     = 0
    renotify_period_minutes  = 60
    resolution_delay_minutes = 0
    trigger                  = "count:2 minutes(test) > 0"
  }
}

resource "scalyr_config_file" "alerts" {
  file_path = "/scalyr/alerts"
  content   = <<EOF
{
  "alertAddress": "",
  "alerts": [
    ${scalyr_alert_rule.test_alert_rule_1.json},
    ${scalyr_alert_group.test_alert_group.json},
    ${scalyr_alert_rule_json.test_alert_rule_2.json}
  ]
}
EOF
}
