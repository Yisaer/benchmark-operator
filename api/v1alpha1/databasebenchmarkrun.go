package v1alpha1

import (
	"fmt"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (in *Run) commandsForRunSysbench(host string, port string, user string, password string) []string {
	commands := []string{
		"sysbench",
		"--db-driver=mysql",
		"--mysql-table-engine=innodb",
		fmt.Sprint("--report-interval=", in.Params["reportinterval"]),
		fmt.Sprint("--oltp-table-size=", in.Params["tablesize"]),
		fmt.Sprint("--oltp-tables-count=", in.Params["tablescount"]),
		fmt.Sprint("--threads=", in.Params["threads"]),
		fmt.Sprint("--time=", in.Params["time"]),
		fmt.Sprint("--mysql-host=", host),
		fmt.Sprint("--mysql-port=", port),
		fmt.Sprint("--mysql-user=", user),
		fmt.Sprint("--mysql-password=", password),
		"/usr/share/sysbench/tests/include/oltp_legacy/oltp.lua",
		"run",
	}
	return commands
}
func (in *Run) createSysBenchmarkJob(host string, port string, user string, password string) (*batchv1.Job, error) {
	result := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			// todo: make these part of the api
			Name:      "db-benchmark-run-sysbench",
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
						Command: in.commandsForRunSysbench(host, port, user, password),
					}},
					RestartPolicy: "Never",
				},
			},
		},
	}
	return &result, nil
}
func (in *Run) CreateJob(host string, port string, user string, password string) (*batchv1.Job, error) {
	switch in.Type {
	case "sysbench":
		return in.createSysBenchmarkJob(host, port, user, password)
	default:
		panic("unimplemented")
	}
}
