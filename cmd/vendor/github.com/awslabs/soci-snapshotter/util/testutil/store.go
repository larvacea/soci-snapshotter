/*
   Copyright The Soci Snapshotter Authors.

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

/*
   Copyright The containerd Authors.

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

package testutil

import (
	"path/filepath"

	"github.com/awslabs/soci-snapshotter/soci/store"
)

// GetContentStoreBlobPath returns the bottom level directory for the content store, e.g. "/blobs/sha256".
func GetContentStoreBlobPath(contentStoreType store.ContentStoreType) (string, error) {
	contentStorePath, err := store.GetContentStorePath(contentStoreType)
	if err != nil {
		return "", err
	}
	return filepath.Join(contentStorePath, "blobs", "sha256"), nil
}
