/*
 * Copyright (C) 2020 The "MysteriumNetwork/node" Authors.
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

package mbtime

import (
	"errors"
)

func nanotime() (uint64, error) {
	// TODO: This is actually incorrect. UnixNano doesn't return monotonic ticks. For Windows we can
	// TODO: to get nanotime from runtime by using go:linkname nanotime runtime.nanotime but first need
	// TODO: to fix window compile.
	return 0, errors.New("not implemented")
}
