// Code generated by ent, DO NOT EDIT.

package certifyvex

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the certifyvex type in the database.
	Label = "certify_vex"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPackageID holds the string denoting the package_id field in the database.
	FieldPackageID = "package_id"
	// FieldArtifactID holds the string denoting the artifact_id field in the database.
	FieldArtifactID = "artifact_id"
	// FieldVulnerabilityID holds the string denoting the vulnerability_id field in the database.
	FieldVulnerabilityID = "vulnerability_id"
	// FieldKnownSince holds the string denoting the known_since field in the database.
	FieldKnownSince = "known_since"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldStatement holds the string denoting the statement field in the database.
	FieldStatement = "statement"
	// FieldStatusNotes holds the string denoting the status_notes field in the database.
	FieldStatusNotes = "status_notes"
	// FieldJustification holds the string denoting the justification field in the database.
	FieldJustification = "justification"
	// FieldOrigin holds the string denoting the origin field in the database.
	FieldOrigin = "origin"
	// FieldCollector holds the string denoting the collector field in the database.
	FieldCollector = "collector"
	// EdgePackage holds the string denoting the package edge name in mutations.
	EdgePackage = "package"
	// EdgeArtifact holds the string denoting the artifact edge name in mutations.
	EdgeArtifact = "artifact"
	// EdgeVulnerability holds the string denoting the vulnerability edge name in mutations.
	EdgeVulnerability = "vulnerability"
	// Table holds the table name of the certifyvex in the database.
	Table = "certify_vexes"
	// PackageTable is the table that holds the package relation/edge.
	PackageTable = "certify_vexes"
	// PackageInverseTable is the table name for the PackageVersion entity.
	// It exists in this package in order to avoid circular dependency with the "packageversion" package.
	PackageInverseTable = "package_versions"
	// PackageColumn is the table column denoting the package relation/edge.
	PackageColumn = "package_id"
	// ArtifactTable is the table that holds the artifact relation/edge.
	ArtifactTable = "certify_vexes"
	// ArtifactInverseTable is the table name for the Artifact entity.
	// It exists in this package in order to avoid circular dependency with the "artifact" package.
	ArtifactInverseTable = "artifacts"
	// ArtifactColumn is the table column denoting the artifact relation/edge.
	ArtifactColumn = "artifact_id"
	// VulnerabilityTable is the table that holds the vulnerability relation/edge.
	VulnerabilityTable = "certify_vexes"
	// VulnerabilityInverseTable is the table name for the VulnerabilityType entity.
	// It exists in this package in order to avoid circular dependency with the "vulnerabilitytype" package.
	VulnerabilityInverseTable = "vulnerability_types"
	// VulnerabilityColumn is the table column denoting the vulnerability relation/edge.
	VulnerabilityColumn = "vulnerability_id"
)

// Columns holds all SQL columns for certifyvex fields.
var Columns = []string{
	FieldID,
	FieldPackageID,
	FieldArtifactID,
	FieldVulnerabilityID,
	FieldKnownSince,
	FieldStatus,
	FieldStatement,
	FieldStatusNotes,
	FieldJustification,
	FieldOrigin,
	FieldCollector,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the CertifyVex queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPackageID orders the results by the package_id field.
func ByPackageID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPackageID, opts...).ToFunc()
}

// ByArtifactID orders the results by the artifact_id field.
func ByArtifactID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArtifactID, opts...).ToFunc()
}

// ByVulnerabilityID orders the results by the vulnerability_id field.
func ByVulnerabilityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVulnerabilityID, opts...).ToFunc()
}

// ByKnownSince orders the results by the known_since field.
func ByKnownSince(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKnownSince, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByStatement orders the results by the statement field.
func ByStatement(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatement, opts...).ToFunc()
}

// ByStatusNotes orders the results by the status_notes field.
func ByStatusNotes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatusNotes, opts...).ToFunc()
}

// ByJustification orders the results by the justification field.
func ByJustification(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJustification, opts...).ToFunc()
}

// ByOrigin orders the results by the origin field.
func ByOrigin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrigin, opts...).ToFunc()
}

// ByCollector orders the results by the collector field.
func ByCollector(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCollector, opts...).ToFunc()
}

// ByPackageField orders the results by package field.
func ByPackageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPackageStep(), sql.OrderByField(field, opts...))
	}
}

// ByArtifactField orders the results by artifact field.
func ByArtifactField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArtifactStep(), sql.OrderByField(field, opts...))
	}
}

// ByVulnerabilityField orders the results by vulnerability field.
func ByVulnerabilityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVulnerabilityStep(), sql.OrderByField(field, opts...))
	}
}
func newPackageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PackageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PackageTable, PackageColumn),
	)
}
func newArtifactStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArtifactInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ArtifactTable, ArtifactColumn),
	)
}
func newVulnerabilityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VulnerabilityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, VulnerabilityTable, VulnerabilityColumn),
	)
}
