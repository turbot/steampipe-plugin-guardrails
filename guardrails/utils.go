package turbot

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/machinebox/graphql"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-guardrails/apiClient"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const (
	filterTimeFormat = "2006-01-02T15:04:05.000Z"
)

func connect(ctx context.Context, d *plugin.QueryData) (*apiClient.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "guardrails"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*apiClient.Client), nil
	}

	// Start with an empty Turbot config
	config := apiClient.ClientConfig{Credentials: apiClient.ClientCredentials{}}

	// Prefer config options given in Steampipe
	guardrailsConfig := GetConfig(d.Connection)
	if guardrailsConfig.Profile != nil {
		config.Profile = *guardrailsConfig.Profile
	}
	if guardrailsConfig.Workspace != nil {
		config.Credentials.Workspace = *guardrailsConfig.Workspace
	}
	if guardrailsConfig.AccessKey != nil {
		config.Credentials.AccessKey = *guardrailsConfig.AccessKey
	}
	if guardrailsConfig.SecretKey != nil {
		config.Credentials.SecretKey = *guardrailsConfig.SecretKey
	}

	clientOptions, err := getClientOptions(guardrailsConfig)
	if err != nil {
		return nil, fmt.Errorf("Error creating HTTP client options: %w", err)
	}

	// Create the client
	client, err := apiClient.CreateClient(config, clientOptions...)
	if err != nil {
		return nil, fmt.Errorf("Error creating Turbot Guardrails client: %s", err.Error())
	}
	if err = client.Validate(); err != nil {
		return nil, fmt.Errorf("Error validating Turbot Guardrails client: %s", err.Error())
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	// Done
	return client, nil
}

// getClientOptions returns the appropriate client options based on the guardrails configuration
func getClientOptions(guardrailsConfig guardrailsConfig) ([]graphql.ClientOption, error) {
	if guardrailsConfig.InsecureSkipVerify != nil && *guardrailsConfig.InsecureSkipVerify {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		clientWithOption := &http.Client{
			Transport: transport,
		}
		return []graphql.ClientOption{graphql.WithHTTPClient(clientWithOption)}, nil
	}
	return nil, nil
}

func getMapValue(_ context.Context, d *transform.TransformData) (interface{}, error) {
	param := d.Param.(string)
	inputMap := d.Value.(map[string]interface{})
	if inputMap[param] != nil {
		return inputMap[param], nil
	}
	return "", nil
}

func emptyMapIfNil(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	v := d.Value.(map[string]interface{})
	return v, nil
}

func emptyListIfNil(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	v := d.Value.([]string)
	return v, nil
}

func intToBool(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	v := d.Value.(int)
	return v > 0, nil
}

func convToString(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	var v interface{} = fmt.Sprint(d.Value)
	return v, nil
}

func attachedResourceIDs(_ context.Context, d *transform.TransformData) (interface{}, error) {
	objs := d.Value.([]GuardrailsIDObject)
	ids := []int64{}
	for _, o := range objs {
		id, err := strconv.ParseInt(o.Turbot.ID, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func pathToArray(_ context.Context, d *transform.TransformData) (interface{}, error) {
	pathStr := types.SafeString(d.Value)
	pathStrs := strings.Split(pathStr, ".")
	pathInts := []int64{}
	for _, s := range pathStrs {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		pathInts = append(pathInts, i)
	}
	return pathInts, nil
}

func escapeQualString(_ context.Context, quals map[string]*proto.QualValue, qualName string) string {
	s := quals[qualName].GetStringValue()
	s = strings.Replace(s, "\\", "\\\\", -1)
	s = strings.Replace(s, "'", "\\'", -1)
	return s
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize
// since getTurbotGuardrailsWorkspace is a call, caching should be per connection
var getTurbotGuardrailsWorkspaceMemoized = plugin.HydrateFunc(getTurbotGuardrailsWorkspaceUncached).Memoize(memoize.WithCacheKeyFunction(getTurbotGuardrailsWorkspaceCacheKey))

// Build a cache key for the call to getTurbotGuardrailsWorkspace.
func getTurbotGuardrailsWorkspaceCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getTurbotGuardrailsWorkspaceUrl"
	return key, nil
}

func getTurbotGuardrailsWorkspace(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	workspaceUrl, err := getTurbotGuardrailsWorkspaceMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}

	return workspaceUrl, nil
}

func getTurbotGuardrailsWorkspaceUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// Start with an empty Turbot config
	config := apiClient.ClientConfig{Credentials: apiClient.ClientCredentials{}}

	// Prefer config options given in Steampipe
	guardrailsConfig := GetConfig(d.Connection)
	if guardrailsConfig.Profile != nil {
		config.Profile = *guardrailsConfig.Profile
	}
	if guardrailsConfig.Workspace != nil {
		config.Credentials.Workspace = *guardrailsConfig.Workspace
	}
	if guardrailsConfig.AccessKey != nil {
		config.Credentials.AccessKey = *guardrailsConfig.AccessKey
	}
	if guardrailsConfig.SecretKey != nil {
		config.Credentials.SecretKey = *guardrailsConfig.SecretKey
	}

	credentials, err := apiClient.GetCredentials(config)
	if err != nil {
		return nil, nil
	}
	endpoint := credentials.Workspace // https://pikachu-turbot.cloud.turbot-dev.com/api/latest/graphql
	if endpoint != "" {
		workspaceUrl := strings.Split(endpoint, "/api/")[0]
		return workspaceUrl, nil
	}

	return nil, nil
}

// Get QualValueList as an list of items
func getQualListValues(ctx context.Context, quals map[string]*proto.QualValue, qualName string, qualType string) string {
	switch qualType {
	case "string":
		if quals[qualName].GetStringValue() != "" {
			return fmt.Sprintf("'%s'", escapeQualString(ctx, quals, qualName))
		} else if quals[qualName].GetListValue() != nil {
			values := make([]string, 0)
			for _, value := range quals[qualName].GetListValue().Values {
				str := value.GetStringValue()
				str = strings.Replace(str, "\\", "\\\\", -1)
				str = strings.Replace(str, "'", "\\'", -1)
				values = append(values, fmt.Sprintf("'%s'", str))
			}
			return strings.Join(values, ",")
		}
	case "int64":
		if quals[qualName].GetInt64Value() != 0 {
			return strconv.FormatInt(quals[qualName].GetInt64Value(), 10)
		} else if quals[qualName].GetListValue() != nil {
			values := make([]string, 0)
			for _, value := range quals[qualName].GetListValue().Values {
				values = append(values, strconv.FormatInt(value.GetInt64Value(), 10))
			}
			return strings.Join(values, ",")
		}
	}
	return ""
}
