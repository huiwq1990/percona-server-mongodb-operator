package backup

import (
	"context"

	"github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	api "github.com/percona/percona-server-mongodb-operator/pkg/apis/psmdb/v1"
)

// HasActiveJobs returns true if there are running backups or restores
// in given cluster and namestpace
func HasActiveJobs(cl client.Client, cluster, namespace string) (bool, error) {
	bcps := &api.PerconaServerMongoDBBackupList{}
	err := cl.List(context.TODO(),
		&client.ListOptions{
			Namespace: namespace,
		},
		bcps,
	)
	if err != nil {
		return false, errors.Wrap(err, "get backup list")
	}
	for _, b := range bcps.Items {
		if b.Spec.PSMDBCluster == cluster && b.Status.State != api.BackupStateReady && b.Status.State != api.BackupStateError {
			return true, nil
		}
	}

	rstrs := &api.PerconaServerMongoDBRestoreList{}
	err = cl.List(context.TODO(),
		&client.ListOptions{
			Namespace: namespace,
		},
		rstrs,
	)
	if err != nil {
		return false, errors.Wrap(err, "get backup list")
	}
	for _, r := range rstrs.Items {
		if r.Spec.ClusterName == cluster && r.Status.State != api.RestoreStateReady && r.Status.State != api.RestoreStateError {
			return true, nil
		}
	}

	return false, nil
}
