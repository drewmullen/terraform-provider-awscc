package codegen

import (
	"fmt"
	"io"
	"sort"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/naming"
)

// Features of the emitted code.
type Features int

const (
	HasUpdatableProperty Features = 1 << iota // At least one property can be updated.
)

type Emitter struct {
	CfResource *cfschema.Resource
	Writer     io.Writer
}

// EmitRootSchema generates the Terraform Plugin SDK code for a CloudFormation root schema
// and emits the generated code to the emitter's Writer. Code features are returned.
// The root schema is the map of root property names to Attributes.
func (e *Emitter) EmitRootPropertiesSchema() (Features, error) {
	return e.emitSchema([]string{}, e.CfResource.Properties)
}

// emitAttribute generates the Terraform Plugin SDK code for a CloudFormation property's Attributes
// and emits the generated code to the emitter's Writer. Code features are returned.
func (e *Emitter) emitAttribute(path []string, name string, property *cfschema.Property) (Features, error) {
	var features Features

	e.printf("{\n")
	e.printf("// Property: %s\n", name)
	if e.CfResource.PrimaryIdentifier.ContainsPath(path) {
		e.printf("// PrimaryIdentifier: %t\n", true)
	}
	e.printf("// CloudFormation resource type schema:\n")
	e.printf("/*\n")
	e.printf("%v\n", property)
	e.printf("*/\n")

	if description := property.Description; description != nil {
		e.printf("Description: `%s`,\n", *description)
	}

	switch propertyType := property.Type.String(); propertyType {
	//
	// Primitive types.
	//
	case cfschema.PropertyTypeBoolean:
		e.printf("Type: types.BoolType,\n")

	case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
		e.printf("Type: types.NumberType,\n")

	case cfschema.PropertyTypeString:
		e.printf("Type: types.StringType,\n")

	//
	// Complex types.
	//
	case cfschema.PropertyTypeArray:
		if property.UniqueItems != nil && *property.UniqueItems {
			//
			// Set.
			//
			switch itemType := property.Items.Type.String(); itemType {
			case cfschema.PropertyTypeObject:
				if len(property.Items.PatternProperties) > 0 {
					return 0, fmt.Errorf("%s is of unsupported type: set of map", name)
				}

				if len(property.Items.Properties) == 0 {
					return 0, fmt.Errorf("%s is of unsupported type: set of undefined schema", name)
				}

				e.printf("Attributes: schema.SetNestedAttributes(\n")

				f, err := e.emitSchema(path, property.Items.Properties)

				if err != nil {
					return 0, err
				}

				features |= f

				e.printf(",\n")
				e.printf("schema.SetNestedAttributesOptions{\n")
				if property.MinItems != nil {
					e.printf("MinItems: %d,\n", *property.MinItems)
				}
				if property.MaxItems != nil {
					e.printf("MaxItems: %d,\n", *property.MaxItems)
				}
				e.printf("},\n")
				e.printf("),\n")

			default:
				return 0, fmt.Errorf("%s is of unsupported type: set of %s", name, itemType)
			}
		} else {
			//
			// List.
			//
			switch itemType := property.Items.Type.String(); itemType {
			case cfschema.PropertyTypeBoolean:
				e.printf("Type: types.ListType{ElemType:types.BoolType},\n")

			case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
				e.printf("Type: types.ListType{ElemType:types.NumberType},\n")

			case cfschema.PropertyTypeString:
				e.printf("Type: types.ListType{ElemType:types.StringType},\n")

			case cfschema.PropertyTypeObject:
				if len(property.Items.PatternProperties) > 0 {
					return 0, fmt.Errorf("%s is of unsupported type: list of map", name)
				}

				if len(property.Items.Properties) == 0 {
					return 0, fmt.Errorf("%s is of unsupported type: list of undefined schema", name)
				}

				e.printf("Attributes: schema.ListNestedAttributes(\n")

				f, err := e.emitSchema(path, property.Items.Properties)

				if err != nil {
					return 0, err
				}

				features |= f

				e.printf(",\n")
				e.printf("schema.ListNestedAttributesOptions{\n")
				if property.MinItems != nil {
					e.printf("MinItems: %d,\n", *property.MinItems)
				}
				if property.MaxItems != nil {
					e.printf("MaxItems: %d,\n", *property.MaxItems)
				}
				e.printf("},\n")
				e.printf("),\n")

			default:
				return 0, fmt.Errorf("%s is of unsupported type: list of %s", name, itemType)
			}
		}

	case cfschema.PropertyTypeObject:
		if patternProperties := property.PatternProperties; len(patternProperties) > 0 {
			//
			// Map.
			//
			if len(property.Properties) > 0 {
				return 0, fmt.Errorf("%s has both Properties and PatternProperties", name)
			}

			n := 0
			for pattern, property := range patternProperties {
				e.printf("// Pattern: %q\n", pattern)
				if n == 0 {
					switch propertyType := property.Type.String(); propertyType {
					case cfschema.PropertyTypeBoolean:
						e.printf("Type: types.MapType{ElemType:types.BoolType},\n")

					case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
						e.printf("Type: types.MapType{ElemType:types.NumberType},\n")

					case cfschema.PropertyTypeString:
						e.printf("Type: types.MapType{ElemType:types.StringType},\n")

					default:
						return 0, fmt.Errorf("%s is of unsupported type: key-value map of %s", name, propertyType)
					}
				} else {
					e.printf("// Ignored.\n")
				}
				n++
			}

			break
		}

		//
		// Object.
		//
		if len(property.Properties) == 0 {
			// Schemaless object => key-value map of string.
			e.printf("Type: types.MapType{ElemType:types.StringType},\n")

			break
		}

		e.printf("Attributes: schema.SingleNestedAttributes(\n")
		f, err := e.emitSchema(path, property.Properties)

		if err != nil {
			return 0, err
		}

		features |= f

		e.printf(",\n")
		e.printf("),\n")

	default:
		return 0, fmt.Errorf("%s is of unsupported type: %s", name, propertyType)
	}

	createOnly := e.CfResource.CreateOnlyProperties.ContainsPath(path)
	readOnly := e.CfResource.ReadOnlyProperties.ContainsPath(path)
	required := e.CfResource.IsRequired(name)
	writeOnly := e.CfResource.WriteOnlyProperties.ContainsPath(path)

	if required {
		e.printf("Required: true,\n")
	} else if !readOnly {
		e.printf("Optional: true,\n")
	}

	if readOnly && !required {
		e.printf("Computed: true,\n")
	}

	if createOnly {
		e.printf("// %s is a force-new attribute.\n", name)
	}

	if writeOnly {
		e.printf("// %s is a write-only attribute.\n", name)
	}

	if !createOnly && !readOnly {
		features |= HasUpdatableProperty
	}

	e.printf("}")

	return features, nil
}

// emitSchema generates the Terraform Plugin SDK code for a CloudFormation property's schema.
// and emits the generated code to the emitter's Writer. Code features are returned.
// A schema is a map of property names to Attributes.
// Property names are sorted prior to code generation to reduce diffs.
func (e *Emitter) emitSchema(pathPrefix []string, properties map[string]*cfschema.Property) (Features, error) {
	names := make([]string, 0)
	for name := range properties {
		names = append(names, name)
	}
	sort.Strings(names)

	var features Features

	e.printf("map[string]schema.Attribute{\n")
	for _, name := range names {
		e.printf("%q: ", naming.CloudFormationPropertyToTerraformAttribute(name))

		f, err := e.emitAttribute(append(pathPrefix, name), name, properties[name])

		if err != nil {
			return 0, err
		}

		features |= f

		e.printf(",\n")
	}
	e.printf("}")

	return features, nil
}

// printf emits a formatted string to the underlying writer.
func (g *Emitter) printf(format string, a ...interface{}) (int, error) {
	return io.WriteString(g.Writer, fmt.Sprintf(format, a...))
}
