/* 
 * Sysdig REST API
 *
 * The Sysdig REST API
 *
 * OpenAPI spec version: 1.0.0
 * Contact: info@sysdig.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

type Alert struct {

	ModifiedOn int64 `json:"modifiedOn,omitempty"`

	Name string `json:"name,omitempty"`

	SegmentCondition []Condition `json:"segmentCondition,omitempty"`

	Timespan int64 `json:"timespan,omitempty"`

	Severity int32 `json:"severity,omitempty"`

	CreatedOn int64 `json:"createdOn,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	NotificationCount int32 `json:"notificationCount,omitempty"`

	TeamId int64 `json:"teamId,omitempty"`

	Version int32 `json:"version,omitempty"`

	SegmentBy []string `json:"segmentBy,omitempty"`

	Type_ string `json:"type,omitempty"`

	Id int64 `json:"id,omitempty"`

	Condition Condition `json:"condition,omitempty"`
}
