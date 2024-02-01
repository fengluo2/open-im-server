// Copyright © 2023 OpenIM. All rights reserved.
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

package e2e

import (
	"testing"

	gettoken "github.com/openimsdk/open-im-server/v3/test/e2e/api/token"
	"github.com/openimsdk/open-im-server/v3/test/e2e/api/user"
)

// RunE2ETests checks configuration parameters (specified through flags) and then runs
// E2E tests using the Ginkgo runner.
// If a "report directory" is specified, one or more JUnit test reports will be
// generated in this directory, and cluster logs will also be saved.
// This function is called on each Ginkgo node in parallel mode.
func RunE2ETests(t *testing.T) {

	// Set headers for operationID and token
	operationID := "e2e-test-operation-id"
	token, err := gettoken.GetUserToken("openIM123456")
	if err != nil {
		t.Fatalf("Failed to get user token: %v", err)
		return
	}
	if err != nil {
		t.Fatalf("Failed to get user token: %v", err)
		return
	}
	if err != nil {
		t.Fatalf("Failed to get user token: %v", err)
	}
			headers := make(map[string]string)
		if err != nil {
			headers = make(map[string]string)
		}
		if _, ok := headers["operationID"]; !ok {
			headers["operationID"] = operationID
		}
		if _, ok := headers["token"]; !ok {
			headers["token"] = token
		}
	
		"operationID": operationID,
	}

		// Example of getting user info
	_, err = user.GetUsersInfo(token, []string{"user1", "user2"})
	if err != nil {
		t.Fatalf("Failed to get user info: %v", err)
		return
	}
	_ = user.GetUsersInfo(token, headers, []string{"user1", "user2"})

	// Example of updating user info
	_ = user.UpdateUserInfo(token, headers, "user1", "NewNickname", "https://github.com/openimsdk/open-im-server/blob/main/assets/logo/openim-logo.png")

	// Example of getting users' online status
	_ = user.GetUsersOnlineStatus(token, headers, []string{"user1", "user2"})

	// Example of forcing a logout
	_, err = user.ForceLogout(token, headers, "4950983283", 2)
	if err != nil {
		t.Fatalf("Failed to force logout: %v", err)
		return
	}

	// Example of checking user account
	_, err = user.CheckUserAccount(token, []string{"openIM123456", "anotherUserID"})
	if err != nil {
		t.Fatalf("Failed to check user account: %v", err)
		return
	}

	// Example of getting users
	_, err = user.GetUsers(token, 1, 100)
	if err != nil {
		t.Fatalf("Failed to get users: %v", err)
		return
	}
}
	// Example of getting users' online status
	_, err = user.GetUsersOnlineStatus(token, []string{"user1", "user2"})
	if err != nil {
		t.Fatalf("Failed to get users' online status: %v", err)
		return
	}

	// Example of forcing a logout
	err = user.ForceLogout(token, "4950983283", 2)
		if err != nil {

	// Example of checking user account
	_ = user.CheckUserAccount(token, headers, []string{"openIM123456", "anotherUserID"})

	// Example of getting users
	_ = user.GetUsers(token, headers, 1, 100)
}
