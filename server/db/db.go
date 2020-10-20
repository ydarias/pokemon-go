package db

import (
	"gorm.io/gorm"
)

type Connection struct {
	Dialector gorm.Dialector
	Config    gorm.Config
}

func (c Connection) Connect() *gorm.DB {
	db, err := gorm.Open(c.Dialector, &c.Config)
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&PokemonEntity{},
		&PokemonEntityType{},
		&PokemonEntityAttackType{},
		&PokemonEntityEvolutionRequirements{},
		&PokemonEntityAttack{})
	if err != nil {
		panic("failed to apply schema changes")
	}

	return db
}

type Populator struct {
	DbConnection *gorm.DB
}

func (p Populator) IsPopulated() bool {
	var count int64
	p.DbConnection.Model(&PokemonEntity{}).Count(&count)

	return count > 0
}

func (p Populator) Populate() {
	if p.IsPopulated() {
		return
	}

	pokemons := ParsePokemons()
	for _, pokemon := range pokemons {
		pokemonInstance := ToPokemonEntity(pokemon)
		pokemonInstance.Types = p.populateTypes(pokemon.Types)
		pokemonInstance.Resistant = p.populateAttackTypes(pokemon.Resistant)
		pokemonInstance.Weaknesses = p.populateAttackTypes(pokemon.Weaknesses)
		pokemonInstance.EvolutionRequirements = PokemonEntityEvolutionRequirements{
			Name:   pokemon.EvolutionRequirements.Name,
			Amount: pokemon.EvolutionRequirements.Amount,
		}

		p.DbConnection.Create(&pokemonInstance)
	}
	for _, pokemon := range pokemons {
		p.updateEvolutionsAndAttacks(pokemon)
	}
}

func (p Populator) populateTypes(pokemonTypes []string) []PokemonEntityType {
	var pokemonEntityTypes []PokemonEntityType
	for _, pokemonType := range pokemonTypes {
		pokemonEntityTypes = append(pokemonEntityTypes, p.findOrCreateType(pokemonType))
	}

	return pokemonEntityTypes
}

func (p Populator) findOrCreateType(pokemonType string) PokemonEntityType {
	var pokemonEntityType PokemonEntityType
	p.DbConnection.Where(PokemonEntityType{Name: pokemonType}).FirstOrCreate(&pokemonEntityType)

	return pokemonEntityType
}

func (p Populator) populateAttackTypes(attackTypes []string) []PokemonEntityAttackType {
	var pokemonEntityAttackTypes []PokemonEntityAttackType
	for _, attackType := range attackTypes {
		pokemonEntityAttackTypes = append(pokemonEntityAttackTypes, p.findOrCreateAttackType(attackType))
	}

	return pokemonEntityAttackTypes
}

func (p Populator) findOrCreateAttackType(pokemonAttack string) PokemonEntityAttackType {
	var pokemonEntityAttack PokemonEntityAttackType
	p.DbConnection.Where(PokemonEntityAttackType{Name: pokemonAttack}).FirstOrCreate(&pokemonEntityAttack)

	return pokemonEntityAttack
}

func (p Populator) updateEvolutionsAndAttacks(pokemon PokemonDefinition) {
	var pokemonEntity PokemonEntity
	p.DbConnection.Where("Identifier = ?", pokemon.Id).First(&pokemonEntity)
	evolutions := p.findEvolutions(pokemon.Evolutions)
	previousEvolutions := p.findEvolutions(pokemon.PreviousEvolutions)
	fastAttacks := p.createAttacks(pokemon.Attacks.Fast)
	specialAttacks := p.createAttacks(pokemon.Attacks.Special)

	p.DbConnection.Model(&pokemonEntity).Association("Evolutions").Append(evolutions)
	p.DbConnection.Model(&pokemonEntity).Association("PreviousEvolutions").Append(previousEvolutions)
	p.DbConnection.Model(&pokemonEntity).Association("FastAttacks").Append(fastAttacks)
	p.DbConnection.Model(&pokemonEntity).Association("SpecialAttacks").Append(specialAttacks)
}

func (p Populator) findEvolutions(evolutions []PokemonDefinitionEvolution) []PokemonEntity {
	var names []string
	for _, evolution := range evolutions {
		names = append(names, evolution.Name)
	}

	var pokemons []PokemonEntity
	p.DbConnection.Where("Name in ?", names).Find(&pokemons)

	return pokemons
}

func (p Populator) createAttacks(attackDefinitions []AttackDefinition) []PokemonEntityAttack {
	var attacks []PokemonEntityAttack
	var attackType PokemonEntityAttackType
	for _, attackDefinition := range attackDefinitions {
		p.DbConnection.Where("Name = ?", attackDefinition.Type).First(&attackType)
		attacks = append(attacks, PokemonEntityAttack{
			Name:   attackDefinition.Name,
			Damage: attackDefinition.Damage,
			Type:   attackType,
		})
	}

	return attacks
}
