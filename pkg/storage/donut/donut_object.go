/*
 * Minimalist Object Storage, (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package donut

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// object internal struct
type object struct {
	name                string
	objectPath          string
	objectMetadata      map[string]string
	donutObjectMetadata map[string]string
}

// NewObject - instantiate a new object
func NewObject(objectName, p string) (Object, error) {
	if objectName == "" {
		return nil, errors.New("invalid argument")
	}
	o := object{}
	o.name = objectName
	o.objectPath = filepath.Join(p, objectName)
	return o, nil
}

func (o object) GetObjectMetadata() (map[string]string, error) {
	objectMetadata := make(map[string]string)
	objectMetadataBytes, err := ioutil.ReadFile(filepath.Join(o.objectPath, objectMetadataConfig))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(objectMetadataBytes, &objectMetadata); err != nil {
		return nil, err
	}
	o.objectMetadata = objectMetadata
	return o.objectMetadata, nil
}

func (o object) GetDonutObjectMetadata() (map[string]string, error) {
	donutObjectMetadata := make(map[string]string)
	donutObjectMetadataBytes, err := ioutil.ReadFile(filepath.Join(o.objectPath, donutObjectMetadataConfig))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(donutObjectMetadataBytes, &donutObjectMetadata); err != nil {
		return nil, err
	}
	o.donutObjectMetadata = donutObjectMetadata
	return o.donutObjectMetadata, nil
}
