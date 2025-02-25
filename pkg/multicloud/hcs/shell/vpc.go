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
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/hcs"
)

func init() {
	type VpcListOptions struct {
	}
	shellutils.R(&VpcListOptions{}, "vpc-list", "List vpcs", func(cli *hcs.SRegion, args *VpcListOptions) error {
		vpcs, err := cli.GetVpcs()
		if err != nil {
			return nil
		}
		printList(vpcs, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&VpcListOptions{}, "external-vpc-list", "List vpcs", func(cli *hcs.SRegion, args *VpcListOptions) error {
		vpcs, err := cli.GetExternalVpcs()
		if err != nil {
			return nil
		}
		printList(vpcs, 0, 0, 0, nil)
		return nil
	})

	type VpcCreateOptions struct {
		NAME string
		CIDR string
		Desc string
	}

	shellutils.R(&VpcCreateOptions{}, "vpc-create", "Create vpc", func(cli *hcs.SRegion, args *VpcCreateOptions) error {
		vpc, err := cli.CreateVpc(args.NAME, args.CIDR, args.Desc)
		if err != nil {
			return err
		}
		printObject(vpc)
		return nil
	})

	type VpcIdOption struct {
		ID string
	}

	shellutils.R(&VpcIdOption{}, "vpc-delete", "Delete vpc", func(cli *hcs.SRegion, args *VpcIdOption) error {
		return cli.DeleteVpc(args.ID)
	})

}
