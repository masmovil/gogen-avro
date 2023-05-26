package resolver

import (
	"fmt"

	"github.com/actgardner/gogen-avro/v10/parser"
	avro "github.com/actgardner/gogen-avro/v10/schema"
)

// ResolveDefinition resolves the References in a Definition one level deep
func ResolveDefinition(def avro.Definition, defs map[avro.QualifiedName]avro.Definition) error {
	for _, child := range def.Children() {
		err := resolveReferences(child, defs)
		if err != nil {
			return err
		}
	}
	return nil
}

func resolveReferences(t avro.AvroType, defs map[avro.QualifiedName]avro.Definition) error {
	// If we reach a Reference to a Definition, resolve it and stop recursing
	if ref, ok := t.(*avro.Reference); ok {
		ref.Def, ok = defs[ref.TypeName]
		if !ok {
			var redefinedName avro.QualifiedName
			if redefinedName, ok = parser.RedefinedNames[ref.TypeName]; ok {
				ref.Def, ok = defs[redefinedName]
				if !ok {
					return fmt.Errorf("Unable to resolve type reference %v", ref.TypeName)
				}
			} else {
				return fmt.Errorf("Unable to resolve type reference %v", ref.TypeName)
			}
		}
		return nil
	}

	// Otherwise recursively search for further References
	for _, child := range t.Children() {
		err := resolveReferences(child, defs)
		if err != nil {
			return err
		}
	}

	return nil
}
