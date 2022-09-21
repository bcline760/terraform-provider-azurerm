package hostpool

type HostPoolOperationPredicate struct {
	Etag      *string
	Id        *string
	Kind      *string
	Location  *string
	ManagedBy *string
	Name      *string
	Type      *string
}

func (p HostPoolOperationPredicate) Matches(input HostPool) bool {

	if p.Etag != nil && (input.Etag == nil && *p.Etag != *input.Etag) {
		return false
	}

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil && *p.Kind != *input.Kind) {
		return false
	}

	if p.Location != nil && (input.Location == nil && *p.Location != *input.Location) {
		return false
	}

	if p.ManagedBy != nil && (input.ManagedBy == nil && *p.ManagedBy != *input.ManagedBy) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}