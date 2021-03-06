/*
 * Copyright (C) 2019 The "MysteriumNetwork/payment-bindings" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package bindings

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewClient creates new remote client connection to ethereum node
func NewClient(rpcUrl string) (*ethclient.Client, error) {
	return ethclient.Dial(rpcUrl)
}
