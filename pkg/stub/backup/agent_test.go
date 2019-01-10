package backup

import (
	"testing"

	"github.com/Percona-Lab/percona-server-mongodb-operator/pkg/apis/psmdb/v1alpha1"

	"github.com/stretchr/testify/assert"
)

func TestHasS3Backups(t *testing.T) {
	c := &Controller{
		psmdb: &v1alpha1.PerconaServerMongoDB{
			Spec: v1alpha1.PerconaServerMongoDBSpec{
				Backup: &v1alpha1.BackupSpec{
					Tasks: []*v1alpha1.BackupTaskSpec{},
				},
			},
		},
	}
	assert.False(t, c.hasS3Backups())

	c.psmdb.Spec.Backup.Tasks = append(c.psmdb.Spec.Backup.Tasks, &v1alpha1.BackupTaskSpec{
		Name:            t.Name(),
		DestinationType: v1alpha1.BackupDestinationS3,
	})
	assert.True(t, c.hasS3Backups())
}

func TestNewAgentContainer(t *testing.T) {
	c := &Controller{
		psmdb: &v1alpha1.PerconaServerMongoDB{
			Spec: v1alpha1.PerconaServerMongoDBSpec{
				Backup: &v1alpha1.BackupSpec{},
				Mongod: &v1alpha1.MongodSpec{
					Net: &v1alpha1.MongodSpecNet{
						Port: int32(0),
					},
				},
				Secrets: &v1alpha1.SecretsSpec{
					Users: "users-secret",
				},
			},
		},
	}
	replset := &v1alpha1.ReplsetSpec{
		Name: t.Name() + "-rs",
	}
	container := c.NewAgentContainer(replset)
	assert.NotNil(t, container)
	assert.NotNil(t, container.SecurityContext.RunAsUser)
	assert.Equal(t, backupImagePrefix+":backup-agent", container.Image)

	// test with version set
	c.psmdb.Spec.Backup.Version = "0.0.0"
	container = c.NewAgentContainer(replset)
	assert.NotNil(t, container)
	assert.Equal(t, backupImagePrefix+":0.0.0-backup-agent", container.Image)
}
