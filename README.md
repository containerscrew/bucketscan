<p align="center" >
  <img src="assets/logo.png" alt="logo" width="250"/>
<h3 align="center">bucketscan</h3>
<p align="center">Bucket scanner</p>
<p align="center">Build with ❤ in Golang</p>
</p>

<p align="center" >
  <img alt="Go report card" src="https://goreportcard.com/badge/github.com/containerscrew/bucketscan">
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/containerscrew/bucketscan">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/containerscrew/bucketscan">
</p>

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [⚠️ DISCLAIMER](#-disclaimer)
  - [Educational Use Only](#educational-use-only)
- [BUCKETSCAN](#bucketscan)
- [THIS APP AND README IS STILL IN PROGRESS](#this-app-and-readme-is-still-in-progress)
- [Local usage](#local-usage)
- [Contribution](#contribution)
- [LICENSE](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# ⚠️ DISCLAIMER
## Educational Use Only

The use of this tool or software is intended solely for educational and learning purposes.

It is not intended for any illegal, unethical, or malicious activities. By using this tool, you agree that:

You will use it exclusively for educational and non-malicious purposes. You will not engage in any illegal or harmful activities using this tool.You understand that the tool is provided as-is, without any warranties or guarantees of any kind.You accept full responsibility for any actions you undertake with this tool and any consequences that may arise.You release the creators, maintainers, and distributors of this tool from any legal liability or responsibility for its use.You will comply with all applicable laws and regulations while using this tool. Please be aware that misuse of this tool may have serious legal and ethical consequences. Always ensure that you have the appropriate permissions and consent before using it in any context.

By using this tool, you acknowledge and agree to the terms and conditions outlined above.

# BUCKETSCAN

# THIS APP AND README IS STILL IN PROGRESS

Bucket fuzzer. For the moment only supports `AWS S3 bucket`.

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

# Contribution

Pull requests are welcome! Any code refactoring, improvement, implementation.

# LICENSE

[LICENSE](./LICENSE)
