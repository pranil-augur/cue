// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mantis 

import (
	"math/big"

	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/internal/pkg"
)

func init() {
	pkg.Register("mantis", p)
}

var _ = adt.TopKind // in case the adt package isn't used

var p = &pkg.Package{
	Native: []*pkg.Builtin{
		{
			Name: "CidrSubnet",
			Params: []pkg.Param{
				{Kind: adt.StringKind},
				{Kind: adt.IntKind},
				{Kind: adt.IntKind},
			},
			Result: adt.StringKind,
			Func: func(c *pkg.CallCtxt) {
				prefix := c.String(0)
				newbits := c.Int(1)
				netnum := big.NewInt(int64(c.Int(2)))
				if c.Do() {
					c.Ret, c.Err = CidrSubnet(prefix, newbits, netnum)
				}
			},
		},
	},
}

