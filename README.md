### Dataset/Scalyr Terraform Provider (proof-of-concept)

Inspired by https://github.com/ansoni/terraform-provider-scalyr

The provider is used to interact with the resources supported by [Dataset](https://www.dataset.com) and its [API](https://app.scalyr.com/help/api). The provider needs to be configured with the proper access key with write permissions before it can be used. It requires terraform 0.12 or later. See `./terraform/providers.tf` for configuration example.

TODO next:
1. Add provider resources' tests 
2. Support configuration of Scalyr [dashboards](https://app.scalyr.com/help/dashboards#editingjson) 

Currently supported methods and resources (see `./terraform/main.tf`)
1. [DataSet's Alerts feature](https://app.scalyr.com/help/alerts)
    
  - Alert as a single resource.
```
    resource "scalyr_alert_rule" "test_alert_rule_1" {
      description              = "Test Alert Rule #1"
      alert_address            = [
          "my_email_addr_1@something.com",
          "my_email_addr_2@something.com"
      ]
      grace_period_minutes     = 0
      renotify_period_minutes  = 60
      resolution_delay_minutes = 0
      trigger                  = "count:5 minutes(test) > 0"
    }
```
  - Alert group sharing the same e-mail address.
   
```
    resource "scalyr_alert_group" "test_alert_group" {
      description   = "Test Alert Group #1"
      alert_address = ["my_email_addr_1@something.com"]

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
```
  - Alert with JSON definition (useful to set silence rules that are currently not supported along with host templating)
```
    resource "scalyr_alert_rule_json" "test_alert_rule_2" {
      description = "Test Alert Rule #2"
      json = jsonencode({
        trigger                = "count:60 minutes(test) > 0"
        gracePeriodMinutes     = 0
        renotifyPeriodMinutes  = 60
        resolutionDelayMinutes = 30
      })
    }
```
2. Configuration files set by [/api/getFile](https://app.scalyr.com/help/api#getFile) and [api/putFile](https://app.scalyr.com/help/api#putFile). This makes alert templating very easy.

```
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

```

Notes:

Need to check the issues with building the provider library on M1 machines.
For now the easiest way to build/test it:
```
GOOS=linux GOARCH=amd64 go build -o terraform-provider-scalyr
docker run -it --rm -v "$PWD":/terraform --entrypoint /bin/sh hashicorp/terraform:0.14.7
cd terraform/terraform
mkdir -p ~/.terraform.d/plugins/terraform.local/local/scalyr/1.0.0/linux_amd64/
cp ../terraform-provider-scalyr ~/.terraform.d/plugins/terraform.local/local/scalyr/1.0.0/linux_amd64/terraform-provider-scalyr
rm -rf ./terraform.tfstate .terraform .terraform.
lock.hcl

terraform init
terraform plan
terraform apply 
```
