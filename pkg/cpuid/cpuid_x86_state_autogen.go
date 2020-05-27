// automatically generated by stateify.

// +build 386 amd64

package cpuid

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *Cache) beforeSave() {}
func (x *Cache) save(m state.Map) {
	x.beforeSave()
	m.Save("Level", &x.Level)
	m.Save("Type", &x.Type)
	m.Save("FullyAssociative", &x.FullyAssociative)
	m.Save("Partitions", &x.Partitions)
	m.Save("Ways", &x.Ways)
	m.Save("Sets", &x.Sets)
	m.Save("InvalidateHierarchical", &x.InvalidateHierarchical)
	m.Save("Inclusive", &x.Inclusive)
	m.Save("DirectMapped", &x.DirectMapped)
}

func (x *Cache) afterLoad() {}
func (x *Cache) load(m state.Map) {
	m.Load("Level", &x.Level)
	m.Load("Type", &x.Type)
	m.Load("FullyAssociative", &x.FullyAssociative)
	m.Load("Partitions", &x.Partitions)
	m.Load("Ways", &x.Ways)
	m.Load("Sets", &x.Sets)
	m.Load("InvalidateHierarchical", &x.InvalidateHierarchical)
	m.Load("Inclusive", &x.Inclusive)
	m.Load("DirectMapped", &x.DirectMapped)
}

func (x *FeatureSet) beforeSave() {}
func (x *FeatureSet) save(m state.Map) {
	x.beforeSave()
	m.Save("Set", &x.Set)
	m.Save("VendorID", &x.VendorID)
	m.Save("ExtendedFamily", &x.ExtendedFamily)
	m.Save("ExtendedModel", &x.ExtendedModel)
	m.Save("ProcessorType", &x.ProcessorType)
	m.Save("Family", &x.Family)
	m.Save("Model", &x.Model)
	m.Save("SteppingID", &x.SteppingID)
	m.Save("Caches", &x.Caches)
	m.Save("CacheLine", &x.CacheLine)
}

func (x *FeatureSet) afterLoad() {}
func (x *FeatureSet) load(m state.Map) {
	m.Load("Set", &x.Set)
	m.Load("VendorID", &x.VendorID)
	m.Load("ExtendedFamily", &x.ExtendedFamily)
	m.Load("ExtendedModel", &x.ExtendedModel)
	m.Load("ProcessorType", &x.ProcessorType)
	m.Load("Family", &x.Family)
	m.Load("Model", &x.Model)
	m.Load("SteppingID", &x.SteppingID)
	m.Load("Caches", &x.Caches)
	m.Load("CacheLine", &x.CacheLine)
}

func init() {
	state.Register("pkg/cpuid.Cache", (*Cache)(nil), state.Fns{Save: (*Cache).save, Load: (*Cache).load})
	state.Register("pkg/cpuid.FeatureSet", (*FeatureSet)(nil), state.Fns{Save: (*FeatureSet).save, Load: (*FeatureSet).load})
}
