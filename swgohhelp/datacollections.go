package swgohhelp

import (
	"bytes"
	"encoding/json"
	"log"
)

// DataPlayerTitle is the data library information about player titles.
type DataPlayerTitle struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"nameKey,omitempty"`
	Desc    string `json:"descKey,omitempty"`
	Details string `json:"shortDescKey,omitempty"`
}

// DataPlayerTitles retrieves the data collection for player titles.
func (c *Client) DataPlayerTitles() (result map[string]DataPlayerTitle, err error) {
	cacheKey := "data.playerTitles"
	if ok := c.gameData.Get(cacheKey, &result); ok {
		return result, err
	}
	// Prepare data collection call
	payload, err := json.Marshal(map[string]interface{}{
		"collection": "playerTitleList",
		"language":   "eng_us",
		"match": map[string]interface{}{
			"hidden":     false,
			"obtainable": true,
		},
	})
	if err != nil {
		return nil, err
	}
	// Parse result
	resp, err := c.call("POST", "/swgoh/data", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	values := make([]DataPlayerTitle, 0)
	err = json.NewDecoder(resp.Body).Decode(&values)
	if err != nil {
		return
	}
	// Prepare response as map for easier usage
	result = make(map[string]DataPlayerTitle)
	for i := range values {
		result[values[i].ID] = values[i]
	}
	// Cache prepared response map
	c.gameData.Put(cacheKey, &result)
	log.Printf("swgohhelp: saving cache for updated titles")
	return
}

// DataUnitAbility is the ability display name and icon
type DataUnitAbility struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"nameKey,omitempty"`
	Icon string `json:"icon,omitempty"`
}

// DataUnitAbilities returns a map of ability IDs to their descriptions.
func (c *Client) DataUnitAbilities() (result map[string]DataUnitAbility, err error) {
	cacheKey := "data.unitAbilities"
	if ok := c.gameData.Get(cacheKey, &result); ok {
		return result, err
	}
	// Prepare data collection call
	payload, err := json.Marshal(map[string]interface{}{
		"collection": "abilityList",
		"language":   "eng_us",
		"project": map[string]int{
			"id":                 1,
			"nameKey":            1,
			"coolDown":           1,
			"icon":               1,
			"descriptiveTagList": 1,
			"aiParams":           1,
			"abilityReference":   1,
			"skillType":          1,
			"isZeta":             1,
		},
	})
	if err != nil {
		return nil, err
	}
	// Parse result
	resp, err := c.call("POST", "/swgoh/data", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	values := make([]DataUnitAbility, 0)
	err = json.NewDecoder(resp.Body).Decode(&values)
	if err != nil {
		return
	}
	// Prepare response as map for easier usage
	result = make(map[string]DataUnitAbility)
	for i := range values {
		result[values[i].ID] = values[i]
	}
	log.Printf("swgohhelp: saving cache for updated abilities")
	c.gameData.Put(cacheKey, &result)
	return
}

// DataUnitSkill is the map for units and their abilities
type DataUnitSkill struct {
	ID        string `json:"id,omitempty"`
	AbilityID string `json:"abilityReference,omitempty"`
	Type      int    `json:"skillType,omitempty"`
	IsZeta    bool   `json:"isZeta,omitempty"`
}

// DataUnitSkills returns a map of skill IDs to their ability IDs.
func (c *Client) DataUnitSkills() (result map[string]DataUnitSkill, err error) {
	cacheKey := "data.unitSkills"
	if ok := c.gameData.Get(cacheKey, &result); ok {
		return result, nil
	}
	// Prepare data collection call
	payload, err := json.Marshal(map[string]interface{}{
		"collection": "skillList",
		"language":   "eng_us",
		"project": map[string]int{
			"id":               1,
			"abilityReference": 1,
			"isZeta":           1,
		},
	})
	if err != nil {
		return nil, err
	}
	// Parse result
	resp, err := c.call("POST", "/swgoh/data", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	values := make([]DataUnitSkill, 0)
	err = json.NewDecoder(resp.Body).Decode(&values)
	if err != nil {
		return
	}
	// Prepare response as map for easier usage
	result = make(map[string]DataUnitSkill)
	for i := range values {
		result[values[i].ID] = values[i]
	}
	log.Printf("swgohhelp: saving cache for updated skills")
	c.gameData.Put(cacheKey, &result)
	return
}

// DataUnitCategory is the category "tags" labels for characters and ships.
type DataUnitCategory struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"descKey,omitempty"`
	Visible bool   `json:"visible"`
}

// DataUnitCategories returns a map of category IDs to their descriptions.
func (c *Client) DataUnitCategories() (result map[string]DataUnitCategory, err error) {
	cacheKey := "data.unitCategories"
	if ok := c.gameData.Get(cacheKey, &result); ok {
		return result, nil
	}
	// Prepare data collection call
	payload, err := json.Marshal(map[string]interface{}{
		"collection": "categoryList",
		"language":   "eng_us",
		"match": map[string]interface{}{
			"visible": true,
		},
		"project": map[string]int{
			"id":      1,
			"descKey": 1,
			"visible": 1,
		},
	})
	if err != nil {
		return nil, err
	}
	// Parse result
	resp, err := c.call("POST", "/swgoh/data", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	values := make([]DataUnitCategory, 0)
	err = json.NewDecoder(resp.Body).Decode(&values)
	if err != nil {
		return
	}
	// Prepare response as map for easier usage
	result = make(map[string]DataUnitCategory)
	for i := range values {
		result[values[i].ID] = values[i]
	}
	log.Printf("swgohhelp: saving cache for updated categoryList")
	c.gameData.Put(cacheKey, &result)
	return
}

// DataUnit is the unit basic data info.
type DataUnit struct {
	ID             string `json:"baseId,omitempty"`
	Name           string `json:"nameKey,omitempty"`
	MaxRarity      int    `json:"maxRarity,omitempty"`
	ForceAlignment int    `json:"forceAlignment,omitempty"`
	CombatType     int    `json:"combatType,omitempty"`
	CombatTypeName string `json:"combatTypeName,omitempty"`

	CategoryRefs []string `json:"categoryIdList,omitempty"`
	Categories   []string `json:"categoryList,omitempty"`

	SkillRefs []DataUnitSkillList `json:"skillReferenceList,omitempty"`
	Skills    []UnitSkill         `json:"skillList,omitempty"`
}

// DataUnitSkillList is an unit skill identifier and requirements.
type DataUnitSkillList struct {
	ID             string `json:"skillId,omitempty"`
	RequiredTier   int    `json:"requiredTier,omitempty"`
	RequiredRarity int    `json:"requiredRarity,omitempty"`
}

// DataUnits returns a map of unit IDs to their details in game.
func (c *Client) DataUnits() (result map[string]DataUnit, err error) {
	cacheKey := "data.units"
	if ok := c.gameData.Get(cacheKey, &result); ok {
		return result, err
	}
	// Prepare data collection call
	payload, err := json.Marshal(map[string]interface{}{
		"collection": "unitsList",
		"language":   "eng_us",
		"match": map[string]interface{}{
			"rarity":     7,
			"obtainable": true,
		},
		"project": map[string]int{
			"id":                 1,
			"thumbnailName":      1,
			"nameKey":            1,
			"maxRarity":          1,
			"categoryIdList":     1,
			"forceAlignment":     1,
			"skillReferenceList": 1,
			"baseId":             1,
			"combatType":         1,
		},
	})
	if err != nil {
		return nil, err
	}
	// Parse result
	resp, err := c.call("POST", "/swgoh/data", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	values := make([]DataUnit, 0)
	err = json.NewDecoder(resp.Body).Decode(&values)
	if err != nil {
		return
	}
	// Prepare response as map for easier usage
	categories, err := c.DataUnitCategories()
	if err != nil {
		return
	}
	skills, err := c.DataUnitSkills()
	if err != nil {
		return
	}
	abilities, err := c.DataUnitAbilities()
	if err != nil {
		return
	}
	result = make(map[string]DataUnit)
	for i := range values {
		switch values[i].CombatType {
		case 1:
			values[i].CombatTypeName = "Character"
		case 2:
			values[i].CombatTypeName = "Ship"
		}
		switch values[i].ForceAlignment {
		case 2:
			values[i].Categories = append(values[i].Categories, "Light Side")
		case 3:
			values[i].Categories = append(values[i].Categories, "Dark Side")
		}
		for j := range values[i].CategoryRefs {
			ref := values[i].CategoryRefs[j]
			if category, ok := categories[ref]; ok {
				if category.Name == "Placeholder" {
					continue
				}
				values[i].Categories = append(values[i].Categories, category.Name)
			}
		}
		for j := range values[i].SkillRefs {
			ref := values[i].SkillRefs[j]
			if skill, ok := skills[ref.ID]; ok {
				if ability, ok := abilities[skill.AbilityID]; ok {
					unitSkill := UnitSkill{
						ID:     skill.ID,
						Name:   ability.Name,
						IsZeta: skill.IsZeta,
					}
					values[i].Skills = append(values[i].Skills, unitSkill)
				}
			}
		}
		result[values[i].ID] = values[i]
	}
	log.Printf("swgohhelp: saving cache for updated unitList")
	c.gameData.Put(cacheKey, &result)
	return
}
