module github.com/tvastar/test

go 1.12

require (
	github.com/OpenPeeDeeP/depguard v0.0.0-20181229194401-1f388ab2d810 // indirect
	github.com/StackExchange/wmi v0.0.0-20181212234831-e0a55b97c705 // indirect
	github.com/coreos/etcd v3.3.12+incompatible // indirect
	github.com/fatih/color v1.7.0 // indirect
	github.com/go-critic/go-critic v0.3.4 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-toolsmith/astcast v1.0.0 // indirect
	github.com/go-toolsmith/astcopy v1.0.0 // indirect
	github.com/go-toolsmith/astfmt v1.0.0 // indirect
	github.com/go-toolsmith/astp v1.0.0 // indirect
	github.com/go-toolsmith/pkgload v1.0.0 // indirect
	github.com/go-toolsmith/typep v1.0.0 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/mock v1.2.0 // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/golangci/errcheck v0.0.0-20181223084120-ef45e06d44b6 // indirect
	github.com/golangci/go-tools v0.0.0-20190124090046-35a9f45a5db0 // indirect
	github.com/golangci/gocyclo v0.0.0-20180528144436-0a533e8fa43d // indirect
	github.com/golangci/gofmt v0.0.0-20181222123516-0b8337e80d98 // indirect
	github.com/golangci/golangci-lint v1.16.1-0.20190402065613-de1d1ad903cd // indirect
	github.com/golangci/gosec v0.0.0-20180901114220-8afd9cbb6cfb // indirect
	github.com/golangci/lint-1 v0.0.0-20181222135242-d2cdd8c08219 // indirect
	github.com/golangci/revgrep v0.0.0-20180812185044-276a5c0a1039 // indirect
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pty v1.1.4 // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/nbutton23/zxcvbn-go v0.0.0-20180912185939-ae427f1e4c1d // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pelletier/go-toml v1.3.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/rogpeppe/go-internal v1.3.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1
	github.com/sergi/go-diff v1.0.0
	github.com/shirou/gopsutil v2.18.12+incompatible // indirect
	github.com/shurcooL/go v0.0.0-20190330031554-6713ea532688 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.4.1 // indirect
	github.com/sourcegraph/go-diff v0.5.1 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.3 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.3.2 // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20190404164418-38d8ce5564a5 // indirect
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 // indirect
	golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6 // indirect
	golang.org/x/sys v0.0.0-20190405154228-4b34438f7a67 // indirect
	golang.org/x/tools v0.0.0-20190408220357-e5b8258f4918
	mvdan.cc/unparam v0.0.0-20190310220240-1b9ccfa71afe // indirect
	sourcegraph.com/sourcegraph/go-diff v0.5.1 // indirect
	sourcegraph.com/sqs/pbtypes v1.0.0 // indirect
)

// need this because go mod seems like a P.O.S
replace sourcegraph.com/sourcegraph/go-diff@v0.5.1 => github.com/sourcegraph/go-diff v0.5.0

replace sourcegraph.com/sourcegraph/go-diff v0.5.1 => github.com/sourcegraph/go-diff v0.5.1
