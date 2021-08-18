// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entsample2/ent/laptop"
	"entsample2/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	laptopFields := schema.Laptop{}.Fields()
	_ = laptopFields
	// laptopDescBrand is the schema descriptor for brand field.
	laptopDescBrand := laptopFields[0].Descriptor()
	// laptop.BrandValidator is a validator for the "brand" field. It is called by the builders before save.
	laptop.BrandValidator = laptopDescBrand.Validators[0].(func(string) error)
	// laptopDescModel is the schema descriptor for model field.
	laptopDescModel := laptopFields[1].Descriptor()
	// laptop.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	laptop.ModelValidator = laptopDescModel.Validators[0].(func(string) error)
	// laptopDescCPU is the schema descriptor for cpu field.
	laptopDescCPU := laptopFields[2].Descriptor()
	// laptop.CPUValidator is a validator for the "cpu" field. It is called by the builders before save.
	laptop.CPUValidator = laptopDescCPU.Validators[0].(func(string) error)
	// laptopDescRAM is the schema descriptor for ram field.
	laptopDescRAM := laptopFields[3].Descriptor()
	// laptop.RAMValidator is a validator for the "ram" field. It is called by the builders before save.
	laptop.RAMValidator = laptopDescRAM.Validators[0].(func(string) error)
}