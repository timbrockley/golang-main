/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package file

import (
	"os"

	"gopkg.in/yaml.v3"
)

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

var yamlTagResolvers = map[string]func(*yaml.Node) (*yaml.Node, error){}

//------------------------------------------------------------

func AddYamlResolvers(tag string, fn func(*yaml.Node) (*yaml.Node, error)) {
	yamlTagResolvers[tag] = fn
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

type CustomYamlTagProcessor struct {
	target interface{}
}

func (i *CustomYamlTagProcessor) UnmarshalYAML(value *yaml.Node) error {
	resolved, err := resolveYamlTags(value)
	if err != nil {
		return err
	}
	return resolved.Decode(i.target)
}

//------------------------------------------------------------

func resolveYamlTags(node *yaml.Node) (*yaml.Node, error) {
	for tag, fn := range yamlTagResolvers {
		if node.Tag == tag {
			return fn(node)
		}
	}
	if node.Kind == yaml.SequenceNode || node.Kind == yaml.MappingNode {
		var err error
		for i := range node.Content {
			node.Content[i], err = resolveYamlTags(node.Content[i])
			if err != nil {
				return nil, err
			}
		}
	}
	return node, nil
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func ReadYAMLFile(filePath string) (map[string]any, error) {

	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var yamlBytes []byte
	//----------
	yamlBytes, err = os.ReadFile(filePath)
	//----------
	yamlData := map[string]any{}
	//----------
	if err == nil {

		err = yaml.Unmarshal([]byte(yamlBytes), &CustomYamlTagProcessor{&yamlData})
	}
	//------------------------------------------------------------
	return yamlData, err
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
