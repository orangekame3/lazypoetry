package poetry

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var haikuLines5 = []string{
	"Morning dew falls soft",
	"Silent moonlight gleams",
	"Autumn leaves descend",
	"Winter wind whispers",
	"Cherry blossoms bloom",
	"Ocean waves crash down",
	"Mountain peaks stand tall",
	"Fireflies dance bright",
}

var haikuLines7 = []string{
	"Through the quiet forest paths",
	"On petals of spring flowers",
	"Across the endless blue sky",
	"Between shadows and sunlight",
	"Above the misty valleys",
	"Through windows of sleepless nights",
	"Against the weathered shoreline",
	"Beneath the starlit canopy",
}

var limerickTemplates = [][]string{
	{
		"There once was a %s quite %s",
		"Who %s in a way that was %s",
		"They %s all day",
		"In their own special way",
		"And %s with a smile most %s",
	},
	{
		"A %s from the town quite %s",
		"Had a habit of being %s",
		"They would %s and %s",
		"Without any care",
		"And everyone thought it was %s",
	},
}

var nouns = []string{"cat", "poet", "writer", "dreamer", "artist", "wanderer", "scholar", "musician"}
var adjectives = []string{"lazy", "clever", "witty", "peaceful", "creative", "curious", "gentle", "wise"}
var verbs = []string{"write", "dream", "sing", "dance", "read", "think", "wander", "create"}
var places = []string{"Maine", "Spain", "Lane", "Plain", "Cane", "Rain"}

func GenerateHaiku() string {
	line1 := haikuLines5[rand.Intn(len(haikuLines5))]
	line2 := haikuLines7[rand.Intn(len(haikuLines7))]
	line3 := haikuLines5[rand.Intn(len(haikuLines5))]
	
	return strings.Join([]string{line1, line2, line3}, "\n")
}

func GenerateLimerick() string {
	template := limerickTemplates[rand.Intn(len(limerickTemplates))]
	
	noun := nouns[rand.Intn(len(nouns))]
	adj1 := adjectives[rand.Intn(len(adjectives))]
	adj2 := adjectives[rand.Intn(len(adjectives))]
	verb1 := verbs[rand.Intn(len(verbs))]
	verb2 := verbs[rand.Intn(len(verbs))]
	verb3 := verbs[rand.Intn(len(verbs))]
	
	// Simple word substitution
	result := make([]string, len(template))
	for i, line := range template {
		switch i {
		case 0:
			result[i] = fmt.Sprintf(line, noun, adj1)
		case 1:
			if strings.Contains(line, "Who %s") {
				result[i] = fmt.Sprintf(line, verb1, adj2)
			} else {
				result[i] = fmt.Sprintf(line, adj2)
			}
		case 2:
			if strings.Contains(line, "%s and %s") {
				result[i] = fmt.Sprintf(line, verb2, verb3)
			} else {
				result[i] = fmt.Sprintf(line, verb1)
			}
		case 3:
			result[i] = line // No substitution needed
		case 4:
			if strings.Contains(line, "And %s with a smile") {
				result[i] = fmt.Sprintf(line, verb2, adj2)
			} else {
				result[i] = fmt.Sprintf(line, adj2)
			}
		}
	}
	
	return strings.Join(result, "\n")
}

func GeneratePoem(form string) string {
	switch strings.ToLower(form) {
	case "limerick":
		return GenerateLimerick()
	default:
		return GenerateHaiku()
	}
}