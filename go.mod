module mayadata.io/openebs-upgrade

go 1.13

require (
	github.com/ghodss/yaml v1.0.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/pkg/errors v0.8.1
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	openebs.io/metac v0.1.1-0.20200122022231-6cec7468b84a
)

replace openebs.io/metac => github.com/AmitKumarDas/metac v0.1.1-0.20200122022231-6cec7468b84a
