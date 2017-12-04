package generation

import (
	"gomine/vectors"
	"gomine/interfaces"
)

type Generator struct {
	name  string
	spawn vectors.TripleVector
	Chunk interfaces.IChunk
	Level interfaces.ILevel
}

type IGenerator interface {
	GetName() string
	GetSpawn() vectors.TripleVector
	GenerateChunk(x, z int)
	PopulateChunk()
}

func NewGenerator(name string) *Generator {
	return &Generator{name, nil, nil, nil}
}

func (gen *Generator) GetName() string {
	return gen.name
}

func (gen *Generator) SetSpawn(v vectors.TripleVector) {
	gen.spawn = v
}

func (gen *Generator) GetSpawn() vectors.TripleVector {
	return gen.spawn
}

func (gen *Generator) GetLevel() interfaces.ILevel {
	return gen.Level
}

func (gen *Generator) SetLevel(level interfaces.ILevel) {
	gen.Level = level
}

func (gen *Generator) GenerateChunk(x, z int) {
}

func (gen *Generator) PopulateChunk() {
}