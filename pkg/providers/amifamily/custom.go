/*
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

package amifamily

import (
	v1 "k8s.io/api/core/v1"

	corev1beta1 "github.com/aws/karpenter-core/pkg/apis/v1beta1"
	"github.com/aws/karpenter-core/pkg/cloudprovider"
	"github.com/aws/karpenter/pkg/apis/v1beta1"

	"github.com/aws/karpenter/pkg/providers/amifamily/bootstrap"
)

type Custom struct {
	DefaultFamily
	*Options
}

// UserData returns the default userdata script for the AMI Family
func (c Custom) UserData(_ *corev1beta1.Kubelet, _ []v1.Taint, _ map[string]string, _ *string, _ []*cloudprovider.InstanceType, customUserData *string) bootstrap.Bootstrapper {
	return bootstrap.Custom{
		Options: bootstrap.Options{
			CustomUserData: customUserData,
		},
	}
}

func (c Custom) DefaultAMIs(_ string, _ bool) []DefaultAMIOutput {
	return nil
}

func (c Custom) DefaultBlockDeviceMappings() []*v1beta1.BlockDeviceMapping {
	// By returning nil, we ensure that EC2 will automatically choose the volumes defined by the AMI
	// and we don't need to describe the AMI ourselves.
	return nil
}

// EphemeralBlockDevice is the block device that the pods on the node will use. For an AMI of a custom family, this is unknown
// to us.
func (c Custom) EphemeralBlockDevice() *string {
	return nil
}
