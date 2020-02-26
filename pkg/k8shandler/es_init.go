package k8shandler

import (
	"github.com/openshift/elasticsearch-operator/pkg/utils"
	batch "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const DefaultBootstrapImage = "quay.io/openshift/origin-elasticsearch-bootstrap:latest"
const DefaultOnRestartAction = "init.sh"
const DefaultOnRollingRestartAction = "init.sh"

func (elasticsearchRequest *ElasticsearchRequest) CreateOrUpdateESInitJob(err error) {
	bsImage := elasticsearchRequest.cluster.Spec.Bootstrap.Image
	if bsImage == "" {
		bsImage = DefaultBootstrapImage
	}

	onRestart := elasticsearchRequest.cluster.Spec.Bootstrap.OnFullRestart
	if onRestart == "" {
		onRestart = DefaultOnRestartAction
	}

	onRollingRestart := elasticsearchRequest.cluster.Spec.Bootstrap.OnRollingRestart
	if onRollingRestart == "" {
		onRollingRestart = DefaultOnRollingRestartAction
	}

	// Detect if it's a rolling restart...
	initJob := NewInitJob(elasticsearchRequest.cluster.Namespace, bsImage, []string{"bash", "-c"}, []string{onRestart})
}

func NewInitJob(namespace, imageName string, command, args []string) batch.Job {
	return batch.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "es-init-job",
			Namespace: namespace,
		},
		Spec: batch.JobSpec{
			Parallelism: utils.GetInt32(1),
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:    "es-init",
							Image:   imageName,
							Command: command,
							Args:    args,
						},
					},
				},
			},
		},
	}
}
