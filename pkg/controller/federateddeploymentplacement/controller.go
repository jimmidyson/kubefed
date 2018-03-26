
/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/


package federateddeploymentplacement

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/marun/fnord/pkg/apis/federation/v1alpha1"
	"github.com/marun/fnord/pkg/controller/sharedinformers"
	listers "github.com/marun/fnord/pkg/client/listers_generated/federation/v1alpha1"
)

// +controller:group=federation,version=v1alpha1,kind=FederatedDeploymentPlacement,resource=federateddeploymentplacements
type FederatedDeploymentPlacementControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about FederatedDeploymentPlacement
	lister listers.FederatedDeploymentPlacementLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *FederatedDeploymentPlacementControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing federateddeploymentplacements labels
	c.lister = arguments.GetSharedInformers().Factory.Federation().V1alpha1().FederatedDeploymentPlacements().Lister()
}

// Reconcile handles enqueued messages
func (c *FederatedDeploymentPlacementControllerImpl) Reconcile(u *v1alpha1.FederatedDeploymentPlacement) error {
	// Implement controller logic here
	log.Printf("Running reconcile FederatedDeploymentPlacement for %s\n", u.Name)
	return nil
}

func (c *FederatedDeploymentPlacementControllerImpl) Get(namespace, name string) (*v1alpha1.FederatedDeploymentPlacement, error) {
	return c.lister.FederatedDeploymentPlacements(namespace).Get(name)
}
