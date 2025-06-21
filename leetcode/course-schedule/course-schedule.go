package courseschedule

import "fmt"

type Course struct {
	id        int
	prereqs   []*Course
	flag      bool
	canFinish *bool
}

func NewCourse(id int) *Course {
	return &Course{
		id:        id,
		prereqs:   make([]*Course, 0),
		canFinish: nil,
	}
}

func (c *Course) AddPrereq(p *Course) {
	c.prereqs = append(c.prereqs, p)
}

func (c *Course) IsCanFinish() bool {
	if c.canFinish != nil {
		return *c.canFinish
	}
	if len(c.prereqs) == 0 {
		trueVal := true
		c.canFinish = &trueVal
		return true
	}

	if c.flag {
		return false
	}

	c.flag = true
	isCanFinish := true
	for _, p := range c.prereqs {
		isCanFinish = isCanFinish && p.IsCanFinish()
	}
	c.canFinish = &isCanFinish
	c.flag = false

	return isCanFinish
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	courses := make([]*Course, numCourses)
	for i := range courses {
		courses[i] = NewCourse(i)
	}

	for _, prereq := range prerequisites {
		c := courses[prereq[0]]
		p := courses[prereq[1]]
		c.AddPrereq(p)
	}

	for i, c := range courses {
		fmt.Println("=> Validate course", i)
		if !c.IsCanFinish() {
			return false
		}
	}

	return true
}
