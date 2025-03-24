package codegen

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/ARM-software/golang-utils/utils/commonerrors"
	"github.com/getkin/kin-openapi/openapi3"
)

type JobItem struct {
	JobItemSchema string
}

type JobItems []JobItem

func (m JobItems) Len() int {
	return len(m)
}

func (m JobItems) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m JobItems) Less(i, j int) bool {
	return m[i].JobItemSchema < m[j].JobItemSchema
}

type JobItemsParams = struct {
	JobItems
}

const (
	JobItemName = "x-job"
)

func AddJobItemsToParams(d *Data) (err error) {
	return AddValuesToParams(d, func(swagger *openapi3.T) (interface{}, error) { return generateJobItems(swagger) }, "deprecations")
}

func generateJobItems(swagger *openapi3.T) (params JobItemsParams, err error) {
	JobItems, err := getJobItemsFromPaths(swagger)
	if err != nil {
		return
	}

	sort.Sort(JobItems)

	return JobItems, nil
}

func getJobItemsFromPaths(swagger *openapi3.T) (jobItems JobItemsParams, err error) {
	var jobs JobItems
	for schemaName, schema := range swagger.Components.Schemas {
		schemaVal := schema.Value
		if schemaVal == nil {
			err = fmt.Errorf("%w: The schema response for '%s' exists but has no value", commonerrors.ErrUndefined, schemaName)
			return
		}

		var isRedacted bool
		if c, ok := schemaVal.ExtensionProps.Extensions[redactFlag].(json.RawMessage); ok {
			err = json.Unmarshal(c, &isRedacted)
			if err != nil {
				return
			}
		}
		if isRedacted {
			continue
		}

		var isExtendedJobItem bool
		if c, ok := schemaVal.ExtensionProps.Extensions[JobItemName].(json.RawMessage); ok {
			err = json.Unmarshal(c, &isExtendedJobItem)
			if err != nil {
				return
			}
		}
		if isExtendedJobItem {
			jobs = append(jobs, JobItem{schemaName})
		}
	}

	jobItems = JobItemsParams{
		jobs,
	}

	return
}
