package v1alpha1

import (
	"fmt"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// todo: merge these 4 arguments into a struct `DBConnectionInfo`
// and the Specs can inherit this
func (in *Prepare) commandsForPrepareSysbench(host string, port string, user string, password string) []string {
	commands := []string{
		"sysbench",
		"--db-driver=mysql",
		fmt.Sprint("--oltp-table-size=", in.Params["tablesize"]),
		fmt.Sprint("--oltp-tables-count=", in.Params["tablescount"]),
		fmt.Sprint("--threads=", in.Params["threads"]),
		fmt.Sprint("--mysql-host=", host),
		fmt.Sprint("--mysql-port=", port),
		fmt.Sprint("--mysql-user=", user),
		fmt.Sprint("--mysql-password=", password),
		"/usr/share/sysbench/tests/include/oltp_legacy/parallel_prepare.lua",
		"run",
	}
	return commands
}
func (in *Prepare) createSysBenchmarkJob(host string, port string, user string, password string) (*batchv1.Job, error) {
	result := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			// todo: make these part of the api
			Name:      "db-benchmark-prepare-sysbench",
			Namespace: "default",
			// todo: add labels and annotation support
			Labels:      nil,
			Annotations: nil,
		},
		Spec: batchv1.JobSpec{
			// todo: add selector support
			Selector: &metav1.LabelSelector{
				MatchLabels:      nil,
				MatchExpressions: nil,
			},
			ManualSelector: nil,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:    "sysbench",
						Image:   in.Image,
						Command: in.commandsForPrepareSysbench(host, port, user, password),
					}},
					RestartPolicy: "Never",
				},
			},
		},
	}
	return &result, nil
}
func (in *Prepare) CreateJob(host string, port string, user string, password string) (*batchv1.Job, error) {
	switch in.Type {
	case "sysbench":
		return in.createSysBenchmarkJob(host, port, user, password)
	default:
		panic("unimplemented")
	}
}
