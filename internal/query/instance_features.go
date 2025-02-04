package query

import (
	"context"

	"github.com/zitadel/zitadel/internal/domain"
)

type InstanceFeatures struct {
	Details                         *domain.ObjectDetails
	LoginDefaultOrg                 FeatureSource[bool]
	TriggerIntrospectionProjections FeatureSource[bool]
	LegacyIntrospection             FeatureSource[bool]
	UserSchema                      FeatureSource[bool]
}

func (q *Queries) GetInstanceFeatures(ctx context.Context, cascade bool) (_ *InstanceFeatures, err error) {
	var system *SystemFeatures
	if cascade {
		system, err = q.GetSystemFeatures(ctx)
		if err != nil {
			return nil, err
		}
	}
	m := NewInstanceFeaturesReadModel(ctx, system)
	if err = q.eventstore.FilterToQueryReducer(ctx, m); err != nil {
		return nil, err
	}
	m.instance.Details = readModelToObjectDetails(m.ReadModel)
	return m.instance, nil
}
