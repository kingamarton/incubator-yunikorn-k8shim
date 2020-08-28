/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package utils

import (
	"testing"

	"gotest.tools/assert"
	v1 "k8s.io/api/core/v1"
	apis "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/apache/incubator-yunikorn-k8shim/pkg/common"
)

func TestConvert2Pod(t *testing.T) {
	pod, err := Convert2Pod(&v1.Node{})
	assert.Assert(t, err != nil)
	assert.Assert(t, pod == nil)

	pod, err = Convert2Pod(&v1.Pod{})
	assert.Assert(t, err == nil)
	assert.Assert(t, pod != nil)
}

func TestIsAssignedPod(t *testing.T) {
	assigned := IsAssignedPod(&v1.Pod{
		Spec: v1.PodSpec{
			NodeName: "some-node",
		},
	})
	assert.Equal(t, assigned, true)

	assigned = IsAssignedPod(&v1.Pod{
		Spec: v1.PodSpec{},
	})
	assert.Equal(t, assigned, false)

	assigned = IsAssignedPod(&v1.Pod{})
	assert.Equal(t, assigned, false)
}

func TestGetApplicationIDFromPod(t *testing.T) {
	appIDInLabel := "labelAppID"
	appIDInAnnotation := "annotationAppID"
	appIDInSelector := "selectorAppID"
	sparkIDInAnnotation := "sparkAnnotationAppID"
	testCases := []struct {
		name          string
		pod           *v1.Pod
		expectedError bool
		expectedAppID string
	}{
		{"AppID defined in label", &v1.Pod{
			ObjectMeta: apis.ObjectMeta{
				Labels: map[string]string{common.LabelApplicationID: appIDInLabel},
			},
		}, false, appIDInLabel},
		{"AppID defined in annotation", &v1.Pod{
			ObjectMeta: apis.ObjectMeta{
				Annotations: map[string]string{common.LabelApplicationID: appIDInAnnotation},
			},
		}, false, appIDInAnnotation},
		{"AppID defined in label and annotation", &v1.Pod{
			ObjectMeta: apis.ObjectMeta{
				Annotations: map[string]string{common.LabelApplicationID: appIDInAnnotation},
				Labels:      map[string]string{common.LabelApplicationID: appIDInLabel},
			},
		}, false, appIDInAnnotation},
		{"Spark AppID defined in annotation", &v1.Pod{
			ObjectMeta: apis.ObjectMeta{
				Annotations: map[string]string{common.SparkAnnotationAppID: sparkIDInAnnotation},
			},
		}, false, sparkIDInAnnotation},
		{"Spark AppID defined in label and annotation", &v1.Pod{
			ObjectMeta: apis.ObjectMeta{
				Annotations: map[string]string{common.SparkAnnotationAppID: sparkIDInAnnotation},
				Labels:      map[string]string{common.LabelApplicationID: appIDInLabel},
			},
		}, false, sparkIDInAnnotation},
		{"Spark AppID defined in spark app selector", &v1.Pod{
			ObjectMeta: apis.ObjectMeta{
				Labels: map[string]string{common.SparkLabelAppID: appIDInSelector},
			},
		}, false, appIDInSelector},
		{"Spark AppID defined in spark app selector and annotation", &v1.Pod{
			ObjectMeta: apis.ObjectMeta{
				Labels:      map[string]string{common.SparkLabelAppID: appIDInSelector},
				Annotations: map[string]string{common.SparkAnnotationAppID: sparkIDInAnnotation},
			},
		}, false, sparkIDInAnnotation},
		{"Spark AppID defined in spark app selector, label and annotation", &v1.Pod{
			ObjectMeta: apis.ObjectMeta{
				Labels:      map[string]string{common.SparkLabelAppID: appIDInSelector, common.LabelApplicationID: appIDInLabel},
				Annotations: map[string]string{common.SparkAnnotationAppID: sparkIDInAnnotation},
			},
		}, false, sparkIDInAnnotation},
		{"No AppID defined", &v1.Pod{}, true, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			appID, err := GetApplicationIDFromPod(tc.pod)
			if tc.expectedError {
				assert.Assert(t, err != nil, "An error is expected")
			} else {
				assert.NilError(t, err, "No error is expected")
			}
			assert.DeepEqual(t, appID, tc.expectedAppID)
		})
	}
}
