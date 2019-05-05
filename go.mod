module devops.lk/kubemeister

go 1.12

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/imdario/mergo v0.3.7 // indirect
	golang.org/x/crypto v0.0.0-20190426145343-a29dc8fdc734 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	k8s.io/api v0.0.0-20190425012535-181e1f9c52c1
	k8s.io/apimachinery v0.0.0-20190425132440-17f84483f500
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20190308190857-21c4ce38f2a7 // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190425172711-65184652c889
