// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shell

import (
	"yunion.io/x/cloudmux/pkg/multicloud/bingocloud"
	"yunion.io/x/onecloud/pkg/util/shellutils"
)

func init() {
	type ImageListOptions struct {
		Id        string
		NextToken string
	}
	shellutils.R(&ImageListOptions{}, "image-list", "list images", func(cli *bingocloud.SRegion, args *ImageListOptions) error {
		images, _, err := cli.GetImages(args.Id, args.NextToken)
		if err != nil {
			return err
		}
		printList(images, 0, 0, 0, []string{})
		return nil
	})
}
