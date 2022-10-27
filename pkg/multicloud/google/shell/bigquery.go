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
	"yunion.io/x/cloudmux/pkg/multicloud/google"
	"yunion.io/x/onecloud/pkg/util/shellutils"
)

func init() {
	type BigQueryOptions struct {
		SQL string `help:"sql to query"`
	}
	shellutils.R(&BigQueryOptions{}, "gcp-bigquery", "Exec bigquery", func(cli *google.SRegion, args *BigQueryOptions) error {
		results, err := cli.BigQuery(args.SQL)
		if err != nil {
			return err
		}

		printList(results, 0, 0, 0, nil)
		return nil
	})
}
