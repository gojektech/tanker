package apps

import (
	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
)

// Service for apps
type Service struct {
	cnf       *config.Config
	datastore Datastore
}

// New - initialise a new apps service
func New(cnf *config.Config, db *sqlx.DB) *Service {
	datastore := NewDatastore(cnf, db)
	return &Service{cnf, datastore}
}

// Add a new app
func (s *Service) Add(appGroup, name, bundleID, platform string) (string, error) {
	return s.datastore.Add(appGroup, name, bundleID, platform)
}

// Delete an app
func (s *Service) Delete(id string) error {
	return s.datastore.Delete(id)
}

// View an app
func (s *Service) View(id string) (App, error) {
	return s.datastore.View(id)
}

// ViewAll apps
func (s *Service) ViewAll() ([]App, error) {
	return s.datastore.ViewAll()
}
