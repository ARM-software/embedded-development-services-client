package codegen

import (
	"encoding/json"
	"slices"
	"strings"

	"github.com/ARM-software/golang-utils/utils/commonerrors"
	"github.com/getkin/kin-openapi/openapi3"
)

type JobItem struct {
	JobItemSchema string
}

type JobItems = []JobItem

type JobItemsParams = struct {
	JobItems
}

const (
	JobItemName = "x-job" // See https://github.com/Arm-Debug/API-Uniform-Contract?tab=readme-ov-file#api-extensions
)

func AddJobItemsToParams(d *Data) (err error) {
	return AddValuesToParams(d, func(swagger *openapi3.T) (interface{}, error) { return GetJobItems(swagger) }, "jobs.go")
}

func GetJobItems(swagger *openapi3.T) (jobItems JobItemsParams, err error) {
	var jobs JobItems
	for schemaName, schema := range swagger.Components.Schemas {
		schemaVal := schema.Value
		if schemaVal == nil {
			err = commonerrors.Newf(commonerrors.ErrUndefined, "The schema response for '%s' exists but has no value", schemaName)
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

	slices.SortFunc(jobs, func(a, b JobItem) int {
		return strings.Compare(a.JobItemSchema, b.JobItemSchema)
	})

	jobItems = JobItemsParams{
		jobs,
	}

	return
}
