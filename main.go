package main

import "openapi"

s = openapi.UserApiService{}
ctx = context.Background()
user, _ = s.GetUsersUserId(ctx, 1).Execute()
