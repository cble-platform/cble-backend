// Code generated by ent, DO NOT EDIT.

package virtualizationprovider

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cble-platform/backend/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldLTE(FieldID, id))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEQ(FieldDisplayName, v))
}

// ProviderGitURL applies equality check predicate on the "provider_git_url" field. It's identical to ProviderGitURLEQ.
func ProviderGitURL(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEQ(FieldProviderGitURL, v))
}

// ProviderVersion applies equality check predicate on the "provider_version" field. It's identical to ProviderVersionEQ.
func ProviderVersion(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEQ(FieldProviderVersion, v))
}

// ConfigPath applies equality check predicate on the "config_path" field. It's identical to ConfigPathEQ.
func ConfigPath(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEQ(FieldConfigPath, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldContainsFold(FieldDisplayName, v))
}

// ProviderGitURLEQ applies the EQ predicate on the "provider_git_url" field.
func ProviderGitURLEQ(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEQ(FieldProviderGitURL, v))
}

// ProviderGitURLNEQ applies the NEQ predicate on the "provider_git_url" field.
func ProviderGitURLNEQ(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldNEQ(FieldProviderGitURL, v))
}

// ProviderGitURLIn applies the In predicate on the "provider_git_url" field.
func ProviderGitURLIn(vs ...string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldIn(FieldProviderGitURL, vs...))
}

// ProviderGitURLNotIn applies the NotIn predicate on the "provider_git_url" field.
func ProviderGitURLNotIn(vs ...string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldNotIn(FieldProviderGitURL, vs...))
}

// ProviderGitURLGT applies the GT predicate on the "provider_git_url" field.
func ProviderGitURLGT(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldGT(FieldProviderGitURL, v))
}

// ProviderGitURLGTE applies the GTE predicate on the "provider_git_url" field.
func ProviderGitURLGTE(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldGTE(FieldProviderGitURL, v))
}

// ProviderGitURLLT applies the LT predicate on the "provider_git_url" field.
func ProviderGitURLLT(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldLT(FieldProviderGitURL, v))
}

// ProviderGitURLLTE applies the LTE predicate on the "provider_git_url" field.
func ProviderGitURLLTE(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldLTE(FieldProviderGitURL, v))
}

// ProviderGitURLContains applies the Contains predicate on the "provider_git_url" field.
func ProviderGitURLContains(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldContains(FieldProviderGitURL, v))
}

// ProviderGitURLHasPrefix applies the HasPrefix predicate on the "provider_git_url" field.
func ProviderGitURLHasPrefix(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldHasPrefix(FieldProviderGitURL, v))
}

// ProviderGitURLHasSuffix applies the HasSuffix predicate on the "provider_git_url" field.
func ProviderGitURLHasSuffix(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldHasSuffix(FieldProviderGitURL, v))
}

// ProviderGitURLEqualFold applies the EqualFold predicate on the "provider_git_url" field.
func ProviderGitURLEqualFold(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEqualFold(FieldProviderGitURL, v))
}

// ProviderGitURLContainsFold applies the ContainsFold predicate on the "provider_git_url" field.
func ProviderGitURLContainsFold(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldContainsFold(FieldProviderGitURL, v))
}

// ProviderVersionEQ applies the EQ predicate on the "provider_version" field.
func ProviderVersionEQ(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEQ(FieldProviderVersion, v))
}

// ProviderVersionNEQ applies the NEQ predicate on the "provider_version" field.
func ProviderVersionNEQ(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldNEQ(FieldProviderVersion, v))
}

// ProviderVersionIn applies the In predicate on the "provider_version" field.
func ProviderVersionIn(vs ...string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldIn(FieldProviderVersion, vs...))
}

// ProviderVersionNotIn applies the NotIn predicate on the "provider_version" field.
func ProviderVersionNotIn(vs ...string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldNotIn(FieldProviderVersion, vs...))
}

// ProviderVersionGT applies the GT predicate on the "provider_version" field.
func ProviderVersionGT(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldGT(FieldProviderVersion, v))
}

// ProviderVersionGTE applies the GTE predicate on the "provider_version" field.
func ProviderVersionGTE(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldGTE(FieldProviderVersion, v))
}

// ProviderVersionLT applies the LT predicate on the "provider_version" field.
func ProviderVersionLT(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldLT(FieldProviderVersion, v))
}

// ProviderVersionLTE applies the LTE predicate on the "provider_version" field.
func ProviderVersionLTE(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldLTE(FieldProviderVersion, v))
}

// ProviderVersionContains applies the Contains predicate on the "provider_version" field.
func ProviderVersionContains(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldContains(FieldProviderVersion, v))
}

// ProviderVersionHasPrefix applies the HasPrefix predicate on the "provider_version" field.
func ProviderVersionHasPrefix(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldHasPrefix(FieldProviderVersion, v))
}

// ProviderVersionHasSuffix applies the HasSuffix predicate on the "provider_version" field.
func ProviderVersionHasSuffix(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldHasSuffix(FieldProviderVersion, v))
}

// ProviderVersionEqualFold applies the EqualFold predicate on the "provider_version" field.
func ProviderVersionEqualFold(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEqualFold(FieldProviderVersion, v))
}

// ProviderVersionContainsFold applies the ContainsFold predicate on the "provider_version" field.
func ProviderVersionContainsFold(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldContainsFold(FieldProviderVersion, v))
}

// ConfigPathEQ applies the EQ predicate on the "config_path" field.
func ConfigPathEQ(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEQ(FieldConfigPath, v))
}

// ConfigPathNEQ applies the NEQ predicate on the "config_path" field.
func ConfigPathNEQ(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldNEQ(FieldConfigPath, v))
}

// ConfigPathIn applies the In predicate on the "config_path" field.
func ConfigPathIn(vs ...string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldIn(FieldConfigPath, vs...))
}

// ConfigPathNotIn applies the NotIn predicate on the "config_path" field.
func ConfigPathNotIn(vs ...string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldNotIn(FieldConfigPath, vs...))
}

// ConfigPathGT applies the GT predicate on the "config_path" field.
func ConfigPathGT(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldGT(FieldConfigPath, v))
}

// ConfigPathGTE applies the GTE predicate on the "config_path" field.
func ConfigPathGTE(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldGTE(FieldConfigPath, v))
}

// ConfigPathLT applies the LT predicate on the "config_path" field.
func ConfigPathLT(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldLT(FieldConfigPath, v))
}

// ConfigPathLTE applies the LTE predicate on the "config_path" field.
func ConfigPathLTE(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldLTE(FieldConfigPath, v))
}

// ConfigPathContains applies the Contains predicate on the "config_path" field.
func ConfigPathContains(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldContains(FieldConfigPath, v))
}

// ConfigPathHasPrefix applies the HasPrefix predicate on the "config_path" field.
func ConfigPathHasPrefix(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldHasPrefix(FieldConfigPath, v))
}

// ConfigPathHasSuffix applies the HasSuffix predicate on the "config_path" field.
func ConfigPathHasSuffix(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldHasSuffix(FieldConfigPath, v))
}

// ConfigPathEqualFold applies the EqualFold predicate on the "config_path" field.
func ConfigPathEqualFold(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldEqualFold(FieldConfigPath, v))
}

// ConfigPathContainsFold applies the ContainsFold predicate on the "config_path" field.
func ConfigPathContainsFold(v string) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.FieldContainsFold(FieldConfigPath, v))
}

// HasBlueprints applies the HasEdge predicate on the "blueprints" edge.
func HasBlueprints() predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BlueprintsTable, BlueprintsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlueprintsWith applies the HasEdge predicate on the "blueprints" edge with a given conditions (other predicates).
func HasBlueprintsWith(preds ...predicate.Blueprint) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(func(s *sql.Selector) {
		step := newBlueprintsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VirtualizationProvider) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VirtualizationProvider) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VirtualizationProvider) predicate.VirtualizationProvider {
	return predicate.VirtualizationProvider(sql.NotPredicates(p))
}