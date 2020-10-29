package viewmodels

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wagoodman/dive/dive/filetree"
	"github.com/wagoodman/dive/dive/image"
)

const (
	CompareSingleLayer LayerCompareMode = iota
	CompareAllLayers
)

type LayerCompareMode int

type LayersViewModel struct {
	mode LayerCompareMode
	layers []*image.Layer
	index int
}

func NewLayersViewModel(layers []*image.Layer) *LayersViewModel {
	return &LayersViewModel{
		mode: CompareSingleLayer,
		layers: layers,
	}
}

func (lm *LayersViewModel) GetMode() LayerCompareMode {
	return lm.mode
}

func (lm *LayersViewModel) SwitchMode() {
	lm.mode = (lm.mode + 1)%2  //this just cycles the mode
}

func (lm *LayersViewModel) GetCompareIndicies() filetree.TreeIndexKey {
	intMax := func (i,j int) int {
		if i > j {
			return i
		}
		return j
	}

	bottomStart := 0
	bottomStop := 0
	if lm.mode == CompareSingleLayer {
		bottomStop = intMax(lm.index-1, 0)
	}
	return filetree.NewTreeIndexKey(bottomStart,bottomStop,lm.index,lm.index)
}

func (lm *LayersViewModel) SetLayerIndex(index int) bool {
	if 0 <= index && index < len(lm.layers) {
		logrus.Debug("setting index, old: %d, new: %d", lm.index, index)
		lm.index = index
		return true
	}
	return false
}

func (lm *LayersViewModel) GetPrintableLayers() []fmt.Stringer {
	var result []fmt.Stringer
	for _, layer := range lm.layers {
		result = append(result, fmt.Stringer(layer))
	}
	return result
}

func (lm *LayersViewModel) GetCurrentLayer() *image.Layer {
	return lm.layers[lm.index]
}
