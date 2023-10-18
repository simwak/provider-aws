/*
Copyright 2023 The Crossplane Authors.

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

package batch

import (
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane-contrib/provider-aws/pkg/controller/batch/computeenvironment"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/batch/job"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/batch/jobdefinition"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/batch/jobqueue"
	"github.com/crossplane-contrib/provider-aws/pkg/utils/setup"
)

// Setup batch controllers.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	return setup.SetupControllers(
		mgr, o,
		computeenvironment.SetupComputeEnvironment,
		job.SetupJob,
		jobdefinition.SetupJobDefinition,
		jobqueue.SetupJobQueue,
	)
}