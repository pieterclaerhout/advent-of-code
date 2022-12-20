package day19

type Resources struct {
	Ores     int
	Clay     int
	Obsidian int
	Geodes   int
}

func (r Resources) Clone() Resources {
	return Resources{
		Ores:     r.Ores,
		Clay:     r.Clay,
		Obsidian: r.Obsidian,
		Geodes:   r.Geodes,
	}
}

func mineResources(resources, robots Resources) Resources {
	return Resources{
		Ores:     resources.Ores + robots.Ores,
		Clay:     resources.Clay + robots.Clay,
		Obsidian: resources.Obsidian + robots.Obsidian,
		Geodes:   resources.Geodes + robots.Geodes,
	}
}
