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
	"fmt"

	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type RoleListOptions struct {
		Offset     string
		Limit      int
		PathPrefix string
	}
	shellutils.R(&RoleListOptions{}, "cloud-role-list", "List roles", func(cli *aws.SRegion, args *RoleListOptions) error {
		roles, err := cli.GetClient().ListRoles(args.Offset, args.Limit, args.PathPrefix)
		if err != nil {
			return err
		}
		printList(roles.Roles, 0, 0, 0, []string{})
		return nil
	})

	type RoleNameOptions struct {
		ROLE string
	}

	shellutils.R(&RoleNameOptions{}, "cloud-role-show", "Show role", func(cli *aws.SRegion, args *RoleNameOptions) error {
		role, err := cli.GetClient().GetRole(args.ROLE)
		if err != nil {
			return err
		}
		printObject(role)
		document := role.GetDocument()
		if document != nil {
			printObject(document)
		}
		return nil
	})

	type RoleAttachPolicyListOptions struct {
		ROLE       string
		Marker     string
		MaxItems   int
		PathPrefix string
	}

	shellutils.R(&RoleAttachPolicyListOptions{}, "cloud-role-attach-policy-list", "List Role attach policy", func(cli *aws.SRegion, args *RoleAttachPolicyListOptions) error {
		policy, err := cli.GetClient().ListAttachedRolePolicies(args.ROLE, args.Marker, args.MaxItems, args.PathPrefix)
		if err != nil {
			return errors.Wrapf(err, "ListAttachedRolePolicies")
		}
		printList(policy.AttachedPolicies, 0, 0, 0, nil)
		if len(policy.Marker) > 0 {
			fmt.Println("marker: ", policy.Marker)
		}
		return nil
	})

	shellutils.R(&RoleNameOptions{}, "cloud-role-delete", "Delete role", func(cli *aws.SRegion, args *RoleNameOptions) error {
		return cli.GetClient().DeleteRole(args.ROLE)
	})

	shellutils.R(&cloudprovider.SRoleCreateOptions{}, "cloud-role-create", "Create role", func(cli *aws.SRegion, args *cloudprovider.SRoleCreateOptions) error {
		role, err := cli.GetClient().CreateRole(args)
		if err != nil {
			return err
		}
		printObject(role)
		return nil
	})

}
