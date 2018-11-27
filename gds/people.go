package gds

import face "github.com/Kagami/go-face"

// People 人
type People struct {
	Name  string `JSON:"name"`  // 姓名
	Faces []Face `JSON:"faces"` // 脸
}

// Peoples 人脸
type Peoples []People

// Cats 标签
func (ps Peoples) Cats() []int32 {
	i := 0
	cats := []int32{}
	for _, p := range ps {
		for f := 0; f < len(p.Faces); f++ {
			cats = append(cats, int32(i))
			i++
		}
	}
	return cats
}

// Samples 例子
func (ps Peoples) Samples() []face.Descriptor {
	samples := []face.Descriptor{}
	for _, p := range ps {
		for _, f := range p.Faces {
			samples = append(samples, f.Descriptor)
		}
	}
	return samples
}

// PeopleID 获取人脸标识
func (ps *Peoples) PeopleID(face face.Face) []byte {
	// TODO _rec.SetSamples()
	return nil
}
