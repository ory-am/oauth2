// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package fosite

import (
	"context"
)

func (f *Fosite) NewDeviceResponse(ctx context.Context, r DeviceRequester, session Session) (DeviceResponder, error) {
	var resp = &DeviceResponse{}

	r.SetSession(session)
	for _, h := range f.Config.GetDeviceEndpointHandlers(ctx) {
		if err := h.HandleDeviceEndpointRequest(ctx, r, resp); err != nil {
			return nil, err
		}
	}

	return resp, nil
}