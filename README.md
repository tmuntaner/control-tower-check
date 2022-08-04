# Control Tower Check

This project checks AWS Organizational Units to verify that their accounts still have the necessary AWS Control Tower Cloudformation Stacks.

In order to upgrade an account to the latest version of Control Tower's configuration, the `StackSet-AWSControlTowerBP-BASELINE-CONFIG` and `StackSet-AWSControlTowerBP-BASELINE-CLOUDWATCH` Cloudformation stacks must be present and not in a deletion state.


## Build the Binary

**Note:**
* To build the binary, you need go version 1.18+.

```bash
make build
```
