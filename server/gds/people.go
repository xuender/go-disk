package gds

// People 人
type People struct {
	Name  string `JSON:"name"`  // 姓名
	Faces []Face `JSON:"faces"` // 脸
}
