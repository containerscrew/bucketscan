<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [BUCKETSCAN](#bucketscan)
- [THIS APP AND README IS STILL IN PROGRESS](#this-app-and-readme-is-still-in-progress)
- [Local usage](#local-usage)
- [⚠️ DISCLAIMER](#-disclaimer)
  - [Educational Use Only](#educational-use-only)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# BUCKETSCAN

# THIS APP AND README IS STILL IN PROGRESS

Happy bucket fuzzer. For the moment only supports `AWS S3 bucket`.

Pending to do:

* gcp buckets
* azure buckets
* other endpoints that can be useful

# Local usage

```shell
git clone https://github.com/containerscrew/bucketscan.git
cd bucketscan/
go run main.go -k containerscrew -d assets/fuzz.txt
```

# ⚠️ DISCLAIMER
## Educational Use Only

The use of this tool or software is intended solely for educational and learning purposes.

It is not intended for any illegal, unethical, or malicious activities. By using this tool, you agree that:

You will use it exclusively for educational and non-malicious purposes. You will not engage in any illegal or harmful activities using this tool.You understand that the tool is provided as-is, without any warranties or guarantees of any kind.You accept full responsibility for any actions you undertake with this tool and any consequences that may arise.You release the creators, maintainers, and distributors of this tool from any legal liability or responsibility for its use.You will comply with all applicable laws and regulations while using this tool. Please be aware that misuse of this tool may have serious legal and ethical consequences. Always ensure that you have the appropriate permissions and consent before using it in any context.

By using this tool, you acknowledge and agree to the terms and conditions outlined above.
