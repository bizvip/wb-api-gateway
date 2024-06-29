/**********************************************************************************************************************
 * Copyright (c) 2024. Archer++. All rights reserved.   Author ORCID: https://orcid.org/0009-0003-8150-367X           *
 **********************************************************************************************************************/

package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)
