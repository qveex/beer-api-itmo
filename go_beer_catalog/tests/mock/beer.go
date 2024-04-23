package mock

import "main/internal/repository"

var Beers = []*repository.Beer{
	{
		BeerId: 1,
		Name:   "name 1",
		Brand:  "brand 1",
		Price:  3,
		Type:   3,
		Degree: 3,
		Sweet:  true,
	},
	{
		BeerId: 2,
		Name:   "name 2",
		Brand:  "brand 2",
		Price:  1,
		Type:   1,
		Degree: 1,
		Sweet:  true,
	},
	{
		BeerId: 3,
		Name:   "name 3",
		Brand:  "brand 3",
		Price:  3,
		Type:   3,
		Degree: 3,
		Sweet:  true,
	},
}
