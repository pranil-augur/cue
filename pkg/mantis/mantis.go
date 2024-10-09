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
	"fmt"
	"math/big"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

// CidrSubnet calculates a subnet address within a given IP network address prefix using cty and cidr libraries.
func CidrSubnet(prefix string, newbits int, netnum *big.Int) (string, error) {
	// Convert prefix string to cty.Value
	prefixVal := cty.StringVal(prefix)

	// Convert newbits int to cty.Value
	newbitsVal, err := gocty.ToCtyValue(newbits, cty.Number)
	if err != nil {
		return "", fmt.Errorf("error converting newbits to cty.Value: %w", err)
	}

	// Convert netnum *big.Int to cty.Value
	netnumVal, err := gocty.ToCtyValue(netnum, cty.Number)
	if err != nil {
		return "", fmt.Errorf("error converting netnum to cty.Value: %w", err)
	}

	// Prepare arguments for the original CidrSubnet logic
	args := []cty.Value{prefixVal, newbitsVal, netnumVal}

	// Call the original CidrSubnet logic
	result, err := originalCidrSubnet(args)
	if err != nil {
		return "", err
	}

	// Extract the string result from cty.Value
	return result.AsString(), nil
}

// originalCidrSubnet contains the business logic using cty and cidr libraries
func originalCidrSubnet(args []cty.Value) (cty.Value, error) {
	var newbits int
	if err := gocty.FromCtyValue(args[1], &newbits); err != nil {
		return cty.UnknownVal(cty.String), err
	}

	var netnum *big.Int
	if err := gocty.FromCtyValue(args[2], &netnum); err != nil {
		return cty.UnknownVal(cty.String), err
	}

	_, network, err := net.ParseCIDR(args[0].AsString())
	if err != nil {
		return cty.UnknownVal(cty.String), fmt.Errorf("invalid CIDR expression: %w", err)
	}

	newNetwork, err := cidr.SubnetBig(network, newbits, netnum)
	if err != nil {
		return cty.UnknownVal(cty.String), err
	}

	return cty.StringVal(newNetwork.String()), nil
}

