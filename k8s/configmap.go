/*
Copyright 2020 The MayaData Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8s

import (
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigMap is a wrapper over k8s ConfigMap
type ConfigMap struct {
	Object *corev1.ConfigMap `json:"object"`
}

// createOrUpdate checks if the resource provided is present or not, if
// not present then it creates the resource otherwise updates it.
func (cm *ConfigMap) createOrUpdate() error {
	existingCm, err := Clientset.CoreV1().ConfigMaps(cm.Object.Namespace).Get(cm.Object.Name, metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			_, err = Clientset.CoreV1().ConfigMaps(cm.Object.Namespace).Create(cm.Object)
			if err != nil {
				return errors.Errorf("Error creating configmap: %s: %+v", cm.Object.Name, err)
			}
		} else {
			return errors.Errorf("Error getting configmap: %s: %+v", cm.Object.Name, err)
		}

	}
	// Set the resource version of the object to be updated
	cm.Object.SetResourceVersion(existingCm.GetResourceVersion())
	_, err = Clientset.CoreV1().ConfigMaps(cm.Object.Namespace).Update(cm.Object)
	if err != nil {
		return errors.Errorf("Error updating configmap: %s: %+v", cm.Object.Name, err)
	}
	return nil
}

// DeployConfigMap creates/updates a given configmap based on
// the given YAML.
func DeployConfigMap(YAML string) error {
	cm := &corev1.ConfigMap{}
	err := yaml.Unmarshal([]byte(YAML), cm)
	if err != nil {
		return errors.Errorf(
			"Error unmarshalling configMap YAML: %+v", err)
	}
	configMap := &ConfigMap{
		Object: cm,
	}
	err = configMap.createOrUpdate()
	if err != nil {
		return err
	}
	return nil
}