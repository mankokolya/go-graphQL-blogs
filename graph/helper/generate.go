package helper

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

// mutateHook adds a tag to the fields
// inside the model
func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	// iterate through model
	for _, model := range b.Models {
		// iterate through fields inside the model
		for _, field := range model.Fields {
			name := field.Name
			if name == "id" {
				name = "_id,omitempty"
			}
			// add the "bson" tag to the field
			field.Tag += ` bson:"` + name + `"`
		}
	}
	// return the model with the additional tag
	return b
}

func main() {
	// load a default configuration
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	// assign "mutateHook()" to MutateHook field
	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	// generate a model from the schema
	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(&p),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
