package mmdbwriter

import "github.com/maxmind/mmdbwriter/mmdbtype"

type dataMapKey string

// Please note, if you change the order of these fields, please check
// alignment as we end up storing quite a few in memory.
type dataMapValue struct {
	data mmdbtype.DataType
	key  dataMapKey

	// Alternatively, we could use a weak map for the data map, but I
	// don't see any very good options at the moment. We should revist
	// if something happens with https://github.com/golang/go/issues/43615
	refCount uint32
}

// dataMap is used to deduplicate data inserted into the tree to reduce
// memory usage using keys generated by keyWriter.
type dataMap struct {
	data      map[dataMapKey]*dataMapValue
	keyWriter KeyGenerator
}

func newDataMap(keyWriter KeyGenerator) *dataMap {
	return &dataMap{
		data:      map[dataMapKey]*dataMapValue{},
		keyWriter: keyWriter,
	}
}

// store stores the value in the dataMap and returns the dataMapValue for it.
// If the value is already in the dataMap, the reference count for it is
// incremented.
func (dm *dataMap) store(v mmdbtype.DataType) (*dataMapValue, error) {
	key, err := dm.keyWriter.Key(v)
	if err != nil {
		return nil, err
	}

	dmv, ok := dm.data[dataMapKey(key)]
	if !ok {
		dmKey := dataMapKey(key)
		dmv = &dataMapValue{
			key:  dmKey,
			data: v,
		}
		dm.data[dmKey] = dmv
	}

	dmv.refCount++

	return dmv, nil
}

// remove removes a reference to the value. If the reference count
// drops to zero, the value is removed from the dataMap.
func (dm *dataMap) remove(v *dataMapValue) {
	// This is here mostly so that we don't have to guard against it
	// elsewhere.
	if v == nil {
		return
	}
	v.refCount--

	if v.refCount == 0 {
		delete(dm.data, v.key)
	}
}