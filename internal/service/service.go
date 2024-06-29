/**********************************************************************************************************************
 * Copyright (c) 2024. Archer++. All rights reserved.   Author ORCID: https://orcid.org/0009-0003-8150-367X           *
 **********************************************************************************************************************/

package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)
