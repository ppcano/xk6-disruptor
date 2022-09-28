package disruptor

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/dop251/goja"
	"github.com/grafana/xk6-disruptor/pkg/kubernetes"
	"github.com/grafana/xk6-disruptor/pkg/testutils/kubernetes/builders"
	"github.com/sirupsen/logrus"

	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/js/modulestest"
	"go.k6.io/k6/lib"
	"go.k6.io/k6/lib/testutils"
	"go.k6.io/k6/metrics"

	"k8s.io/client-go/kubernetes/fake"
)

// testVU creates a test VU
func testVU() modules.VU {
	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper{})

	testLog := logrus.New()
	testLog.AddHook(&testutils.SimpleLogrusHook{
		HookedLevels: []logrus.Level{logrus.WarnLevel},
	})
	testLog.SetOutput(ioutil.Discard)

	state := &lib.State{
		Options: lib.Options{
			SystemTags: metrics.NewSystemTagSet(metrics.TagVU),
		},
		Logger: testLog,
		Tags:   lib.NewTagMap(nil),
	}

	return	&modulestest.VU{
			RuntimeField: rt,
			InitEnvField: &common.InitEnvironment{},
			CtxField:     context.Background(),
			StateField:   state,
		}
}

// instantiates a module with a fake kubernetes and a test VU
func setTestModule(k8s *kubernetes.FakeKubernetes, vu modules.VU) error {
	m := ModuleInstance{
		k8s: k8s,
		vu: vu,
	}
	err := vu.Runtime().Set("PodDisruptor", m.Exports().Named["PodDisruptor"])
	return err
}


const listTargetsScript =`
const selector = {
   namespace: "default",
   select: {
     labels: {
	app: "test"
     }
   }
} 

const disruptor = new PodDisruptor(selector)
const targets = disruptor.targets()
if (targets.length != 1) {
   throw new Error("expected list to have one target")
} 
`

func Test_PodDisruptor(t *testing.T) {
	pod := builders.NewPodBuilder("pod-with-app-label").
		WithLabels(map[string]string{
			"app": "test",
		}).
		Build()
	client := fake.NewSimpleClientset(pod)
	k8s, _ := kubernetes.NewFakeKubernetes(client)
	vu := testVU()	
	setTestModule(k8s, vu)

	_, err := vu.Runtime().RunString(listTargetsScript)
	if err != nil {
		t.Errorf("failed %v", err)
	}
}

