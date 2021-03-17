package query

// IdsQuery filters documents that only have the provided ids.
// Note, this query uses the _uid field.
// For more details, see
// https://www.elastic.co/guide/en/elasticsearch/reference/7.0/query-dsl-ids-query.html
type IdsQuery struct {
	values []string
}

// NewIdsQuery creates and initializes a new ids query.
func NewIdsQuery() *IdsQuery {
	return &IdsQuery{
		values: make([]string, 0),
	}
}

// Ids adds ids to the filter.
func (q *IdsQuery) Ids(ids ...string) *IdsQuery {
	q.values = append(q.values, ids...)
	return q
}

// Source returns JSON for the function score query.
func (q *IdsQuery) Source() (interface{}, error) {

	//"query": {
	//	"ids" : {
	//		"values" : ["1", "4", "100"]
	//	}
	//}

	source := make(map[string]interface{})
	query := make(map[string]interface{})
	source["ids"] = query

	// values
	query["values"] = q.values

	return source, nil
}
