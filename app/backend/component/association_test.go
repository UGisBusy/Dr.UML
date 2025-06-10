// VIBE CODING

package component

import (
	"testing"

	"Dr.uml/backend/component/attribute"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

func Test_NewAssociation(t *testing.T) {
	gadget := newEmptyGadget(Class, utils.Point{X: 0, Y: 0})
	tests := []struct {
		name                 string
		parents              [2]*Gadget
		assType              AssociationType
		startPoint, endPoint utils.Point
		wantErr              bool
	}{
		{"valid association", [2]*Gadget{gadget, gadget}, Extension, utils.Point{X: 0, Y: 0}, utils.Point{X: 1, Y: 1}, false},
		{"same point", [2]*Gadget{gadget, gadget}, Extension, utils.Point{X: 0, Y: 0}, utils.Point{X: 0, Y: 0}, true},
		{"nil parent", [2]*Gadget{nil, gadget}, Extension, utils.Point{X: 0, Y: 0}, utils.Point{X: 1, Y: 1}, true},
		{"invalid assType", [2]*Gadget{gadget, gadget}, 0, utils.Point{X: 0, Y: 0}, utils.Point{X: 1, Y: 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAssociation(tt.parents, tt.assType, tt.startPoint, tt.endPoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("error mismatch: got %v, want error %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Association_Getters(t *testing.T) {
	gadget := newEmptyGadget(Class, utils.Point{X: 0, Y: 0})
	ass := &Association{
		assType: Extension,
		layer:   5,
		parents: [2]*Gadget{gadget, gadget},
	}

	t.Run("GetAssType", func(t *testing.T) {
		if ass.GetAssType() != Extension {
			t.Errorf("expected %v, got %v", Extension, ass.GetAssType())
		}
	})

	t.Run("GetLayer", func(t *testing.T) {
		if ass.GetLayer() != 5 {
			t.Errorf("expected %v, got %v", 5, ass.GetLayer())
		}
	})

	t.Run("GetParentStart", func(t *testing.T) {
		if ass.GetParentStart() != gadget {
			t.Errorf("expected %v, got %v", gadget, ass.GetParentStart())
		}
	})

	t.Run("GetParentEnd", func(t *testing.T) {
		if ass.GetParentEnd() != gadget {
			t.Errorf("expected %v, got %v", gadget, ass.GetParentEnd())
		}
	})
}

func Test_Association_Setters(t *testing.T) {
	gadget := newEmptyGadget(Class, utils.Point{X: 0, Y: 0})
	ass := &Association{
		assType:         Extension,
		parents:         [2]*Gadget{gadget, gadget},
		startPointRatio: [2]float64{0, 0},
		endPointRatio:   [2]float64{1, 1},
	}

	t.Run("SetAssType", func(t *testing.T) {
		err := ass.SetAssType(Extension)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ass.GetAssType() != Extension {
			t.Errorf("expected %v, got %v", Extension, ass.GetAssType())
		}
	})

	t.Run("SetLayer", func(t *testing.T) {
		err := ass.SetLayer(10)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ass.GetLayer() != 10 {
			t.Errorf("expected %v, got %v", 10, ass.GetLayer())
		}
	})

	t.Run("SetParentStart", func(t *testing.T) {
		newStPoint := utils.Point{X: 2, Y: 2}
		newSt := newEmptyGadget(Class, newStPoint)

		err := ass.SetParentStart(newSt, newStPoint)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ass.GetParentStart() != newSt {
			t.Errorf("expected %v, got %v", newSt, ass.GetParentStart())
		}
	})

	t.Run("SetParentEnd", func(t *testing.T) {
		newEnPoint := utils.Point{X: 4, Y: 4}
		newEn := newEmptyGadget(Class, newEnPoint)

		err := ass.SetParentEnd(newEn, newEnPoint)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ass.GetParentEnd() != newEn {
			t.Errorf("expected %v, got %v", newEn, ass.GetParentEnd())
		}
	})
}

func Test_Association_AddAttribute(t *testing.T) {
	stPoint := utils.Point{X: 0, Y: 0}
	enPoint := utils.Point{X: 200, Y: 200}
	stGadget, _ := NewGadget(Class, stPoint, 0, drawdata.DefaultGadgetColor, "")
	enGadget, _ := NewGadget(Class, enPoint, 0, drawdata.DefaultGadgetColor, "")
	ass, _ := NewAssociation([2]*Gadget{stGadget, enGadget}, Extension, stPoint, enPoint)

	ratio := 0.5
	content := "test attribute"
	var att *attribute.AssAttribute
	t.Run("Add Attribute", func(t *testing.T) {
		err := ass.AddAttribute(ratio, content)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(ass.attributes) != 1 {
			t.Errorf("ass.attributes does not change")
		}

		att = ass.attributes[0]
		if att.GetContent() != content {
			t.Errorf("unexpected content %v, got %v", content, att.GetContent())
		}
		if att.GetRatio() != ratio {
			t.Errorf("unexpected ratio %v, got %v", ratio, att.GetRatio())
		}
	})
}

func Test_Association_RemoveAttribute(t *testing.T) {
	att := &attribute.AssAttribute{}
	gadget := newEmptyGadget(Class, utils.Point{X: 0, Y: 0})
	ass := &Association{
		parents:         [2]*Gadget{gadget, gadget},
		attributes:      []*attribute.AssAttribute{att},
		startPointRatio: [2]float64{0, 0},
		endPointRatio:   [2]float64{1, 1},
	}

	t.Run("Remove valid attribute", func(t *testing.T) {
		err := ass.RemoveAttribute(0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(ass.attributes) != 0 {
			t.Errorf("expected 0 attributes, got %v", len(ass.attributes))
		}
	})

	t.Run("Remove invalid index", func(t *testing.T) {
		err := ass.RemoveAttribute(1)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func Test_Association_MoveAttribute(t *testing.T) {
	att := &attribute.AssAttribute{}
	gadget := newEmptyGadget(Class, utils.Point{X: 0, Y: 0})
	ass := &Association{
		parents:         [2]*Gadget{gadget, gadget},
		attributes:      []*attribute.AssAttribute{att},
		startPointRatio: [2]float64{0, 0},
		endPointRatio:   [2]float64{1, 1},
	}

	t.Run("Move valid attribute", func(t *testing.T) {
		err := ass.MoveAttribute(0, 0.5)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Move invalid index", func(t *testing.T) {
		err := ass.MoveAttribute(1, 0.5)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func Test_Association_Cover(t *testing.T) {
	gadget := newEmptyGadget(Class, utils.Point{X: 0, Y: 0})
	ass := &Association{
		parents: [2]*Gadget{gadget, gadget},
		drawdata: drawdata.Association{
			StartX: 0, StartY: 0,
			EndX: 10, EndY: 10,
		},
	}

	t.Run("Point inside threshold", func(t *testing.T) {
		covered, err := ass.Cover(utils.Point{X: 5, Y: 5})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !covered {
			t.Errorf("expected point to be covered")
		}
	})

	t.Run("Point outside threshold", func(t *testing.T) {
		covered, err := ass.Cover(utils.Point{X: 20, Y: 20})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if covered {
			t.Errorf("expected point to not be covered")
		}
	})
}

func Test_Association_UpdateDrawData(t *testing.T) {
	g1, _ := NewGadget(Class, utils.Point{X: 0, Y: 0}, 0, "#FF00FF", "sample header")
	g2, _ := NewGadget(Class, utils.Point{X: 10, Y: 10}, 0, "#FF00FF", "sample header")
	ass := &Association{
		parents:         [2]*Gadget{g1, g2},
		startPointRatio: [2]float64{0, 0},
		endPointRatio:   [2]float64{1, 1},
	}

	t.Run("Update with valid data", func(t *testing.T) {
		err := ass.updateDrawData()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func Test_Association_RegisterUpdateParentDraw(t *testing.T) {
	gadget := newEmptyGadget(Class, utils.Point{X: 0, Y: 0})
	ass := &Association{
		parents: [2]*Gadget{gadget, gadget},
	}

	t.Run("Register valid function", func(t *testing.T) {
		err := ass.RegisterUpdateParentDraw(func() duerror.DUError { return nil })
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Register nil function", func(t *testing.T) {
		err := ass.RegisterUpdateParentDraw(nil)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func Test_Association_Selected(t *testing.T) {
	stPoint := utils.Point{X: 0, Y: 0}
	enPoint := utils.Point{X: 200, Y: 200}
	stGadget, _ := NewGadget(Class, stPoint, 0, drawdata.DefaultGadgetColor, "")
	enGadget, _ := NewGadget(Class, enPoint, 0, drawdata.DefaultGadgetColor, "")
	ass, _ := NewAssociation([2]*Gadget{stGadget, enGadget}, Extension, stPoint, enPoint)
	ass.RegisterUpdateParentDraw(func() duerror.DUError { return nil })

	if ass.GetIsSelected() {
		t.Errorf("ass should be unselect when created")
	}

	err := ass.SetIsSelected(true)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !ass.GetIsSelected() {
		t.Errorf("ass SetIsSelected(true) doesnt work")
	}

	err = ass.SetIsSelected(false)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if ass.GetIsSelected() {
		t.Errorf("ass SetIsSelected(false) doesnt work")
	}
}

func Test_Association_Attribute(t *testing.T) {

	stPoint := utils.Point{X: 0, Y: 0}
	enPoint := utils.Point{X: 200, Y: 200}
	stGadget, _ := NewGadget(Class, stPoint, 0, drawdata.DefaultGadgetColor, "")
	enGadget, _ := NewGadget(Class, enPoint, 0, drawdata.DefaultGadgetColor, "")
	ass, _ := NewAssociation([2]*Gadget{stGadget, enGadget}, Extension, stPoint, enPoint)

	ratio := 0.5
	content := "old content"
	index := 0
	ass.AddAttribute(ratio, content)
	att := ass.attributes[index]
	t.Run("SetAttrContent", func(t *testing.T) {
		newContent := "new content"
		err := ass.SetAttrContent(index, newContent)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if att.GetContent() != newContent {
			t.Errorf("expected content: %v, got %v", newContent, att.GetContent())
		}
		content = newContent
	})
	t.Run("SetAttrSize", func(t *testing.T) {
		newSize := 69
		err := ass.SetAttrSize(index, newSize)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if att.GetSize() != newSize {
			t.Errorf("expected size: %v, got %v", newSize, att.GetSize())
		}
	})
	t.Run("SetAttrStyle", func(t *testing.T) {
		newStyle := 0x4
		err := ass.SetAttrStyle(index, newStyle)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if int(att.GetStyle()) != newStyle {
			t.Errorf("expected size: %v, got %v", newStyle, int(att.GetStyle()))
		}
	})
	t.Run("SetAttrFontFile", func(t *testing.T) {
		// dont care
	})
	t.Run("SetAttrRatio", func(t *testing.T) {
		newRatio := 0.69
		err := ass.SetAttrRatio(index, newRatio)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if att.GetRatio() != newRatio {
			t.Errorf("expected size: %v, got %v", newRatio, att.GetRatio())
		}
	})
}
func Test_Association_ToSavedAssociation(t *testing.T) {
	gadget := newEmptyGadget(Class, utils.Point{X: 1, Y: 2})
	ass := &Association{
		assType:         Extension,
		layer:           3,
		parents:         [2]*Gadget{gadget, gadget},
		startPointRatio: [2]float64{.2, .3},
		endPointRatio:   [2]float64{.2, .3},
		attributes: []*attribute.AssAttribute{
			{
				// Fill with minimal fields if needed
			},
		},
	}
	parents := [2]int{10, 20}
	saved := ass.ToSavedAssociation(parents)
	if saved.AssType != int(Extension) {
		t.Errorf("expected AssType %v, got %v", Extension, saved.AssType)
	}
	if saved.Layer != 3 {
		t.Errorf("expected Layer 3, got %v", saved.Layer)
	}
	if saved.StartPointRatio != [2]float64{.2, .3} || saved.EndPointRatio != [2]float64{.2, .3} {
		t.Errorf("expected StartPointRatio and EndPointRatio [.2 .3], got %v and %v", saved.StartPointRatio, saved.EndPointRatio)
	}
	if len(saved.Parents) != 2 || saved.Parents[0] != 10 || saved.Parents[1] != 20 {
		t.Errorf("expected Parents [10 20], got %v", saved.Parents)
	}
	if len(saved.Attributes) != 1 {
		t.Errorf("expected 1 attribute, got %v", len(saved.Attributes))
	}
}
