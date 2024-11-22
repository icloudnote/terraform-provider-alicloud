---
subcategory: "Application Real-Time Monitoring Service (ARMS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_arms_alert_robots"
sidebar_current: "docs-alicloud-datasource-arms-alert-robots"
description: |-
  Provides a list of Arms Alert Robots to the user.
---

# alicloud_arms_alert_robots

This data source provides the Arms Alert Robots of the current Alibaba Cloud user.

-> **NOTE:** Available in v1.135.0+.

## Example Usage

Basic Usage

```terraform
data "alicloud_arms_alert_robots" "nameRegex" {
  name_regex = "^my-AlertRobot"
}
output "arms_alert_robot_id" {
  value = data.alicloud_arms_alert_robots.nameRegex.robots.0.id
}
```

## Argument Reference

The following arguments are supported:

* `ids` - (Optional, ForceNew, Computed)  A list of Alert Robot IDs.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Alert Robot name.
* `alert_robot_name` - (Optional, ForceNew) The robot name.
* `robot_type` - (Optional, ForceNew) The robot type.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

## Argument Reference

The following attributes are exported in addition to the arguments listed above:

* `names` - A list of Alert Robot names.
* `robots` - A list of Arms Alert Robots. Each element contains the following attributes:
	* `robot_id` - The id of the robot.
	* `robot_name` - The name of the robot.
	* `robot_type` - The type of the robot.
	* `robot_addr` - The webhook url of the robot.
	* `daily_noc` - Specifies whether the alert robot receives daily notifications.
	* `daily_noc_time` - The time of the daily notification.
	* `create_time` - The creation time of the resource.
	* `id` - The ID of the Alert Robot.
