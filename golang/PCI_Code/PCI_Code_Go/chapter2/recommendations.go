// A dictionary of movie critics and their Ratings of a small
// set of movies

package main

import (
	"log"
	"math"
)

type CriticType map[string](map[string]float64)

func getCriticList() CriticType {
	critics := CriticType{
		"Lisa Rose": {"Lady in the Water": 2.5, "Snakes on a Plane": 3.5,
			"Just My Luck": 3.0, "Superman Returns": 3.5, "You, Me and Dupree": 2.5,
			"The Night Listener": 3.0},
		"Gene Seymour": {"Lady in the Water": 3.0, "Snakes on a Plane": 3.5,
			"Just My Luck": 1.5, "Superman Returns": 5.0, "The Night Listener": 3.0,
			"You, Me and Dupree": 3.5},
		"Michael Phillips": {"Lady in the Water": 2.5, "Snakes on a Plane": 3.0,
			"Superman Returns": 3.5, "The Night Listener": 4.0},
		"Claudia Puig": {"Snakes on a Plane": 3.5, "Just My Luck": 3.0,
			"The Night Listener": 4.5, "Superman Returns": 4.0,
			"You, Me and Dupree": 2.5},
		"Mick LaSalle": {"Lady in the Water": 3.0, "Snakes on a Plane": 4.0,
			"Just My Luck": 2.0, "Superman Returns": 3.0, "The Night Listener": 3.0,
			"You, Me and Dupree": 2.0},
		"Jack Matthews": {"Lady in the Water": 3.0, "Snakes on a Plane": 4.0,
			"The Night Listener": 3.0, "Superman Returns": 5.0, "You, Me and Dupree": 3.5},
		"Toby": {"Snakes on a Plane": 4.5, "You, Me and Dupree": 1.0, "Superman Returns": 4.0}}

	return critics
}

func criticsTest(critics CriticType) {
	log.Println(critics["Lisa Rose"]["Lady in the Water"])
	log.Println(critics["Lisa Rose"])
	log.Println(critics["Toby"]["Snakes on a Plane"])
	critics["Toby"]["Snakes on a Plane"] = 3.5
	log.Println(critics["Toby"]["Snakes on a Plane"])
}

func euclideanTest() {
	log.Println(1 / (1 + math.Sqrt(math.Pow(4.5-4, 2)+math.Pow(1-2, 2))))
}

/**
 * Returns a distance-based similarity score for person1 and person2
 */
func sim_distance(prefs CriticType, person1, person2 string) float64 {
	var sumOfSquares float64 = 0
	for movie, _ := range prefs[person1] {
		_, ok := prefs[person2][movie]
		if ok {
			//  Add up the squares of all the differences
			sumOfSquares += math.Pow((prefs[person1][movie] - prefs[person2][movie]), 2)
			//log.Printf("User: %s,\tMovie: %s,\tRating: %g\n",person1, movie, rating)			
		}
	}

	return 1 / (1 + math.Sqrt(sumOfSquares))
}

/**
 * Returns the Pearson correlation coefficient for p1 and p2
 */

func sim_pearson(prefs CriticType, person1, person2 string) float64 {
	var n, sum1, sum2, sum1Sq, sum2Sq, pSum float64= 0, 0, 0, 0, 0, 0
	
	for movie, _ := range prefs[person1] {
		_, ok := prefs[person2][movie]
		if ok {
			// Add up all the preferences
			sum1 += prefs[person1][movie]
			sum2 += prefs[person2][movie]
			
			// Sum up the squares
			sum1Sq += math.Pow(prefs[person1][movie], 2)
			sum2Sq += math.Pow(prefs[person2][movie], 2)
			
			// Sum up the products
			pSum += prefs[person1][movie] * prefs[person2][movie]
			n += 1
		}
	}	
	if n == 0 {
		return 0
	}
	// Calculate Pearson scor
	num := pSum - (sum1*sum2/n)
	den := math.Sqrt((sum1Sq-math.Pow(sum1, 2)/n)*(sum2Sq-math.Pow(sum2, 2)/n))
	if den == 0 {
		return 0
	}
	return (num/den)
}

func main() {
	var critics = getCriticList()
	log.Println(critics)
	for critic, movies := range critics {
		log.Printf("Critic: %s\n", critic)
		for movie, rating := range movies {
			log.Printf("Movie: %s,\tRating: %g\n", movie, rating)
		}
	}
	criticsTest(critics)

	euclideanTest()
	
	
	var cri = sim_distance(critics, "Lisa Rose", "Gene Seymour")
	var pea = sim_pearson(critics, "Lisa Rose", "Gene Seymour")
	log.Printf("[Lisa Rose][Gene Seymour]\t %f\t%f\n", cri, pea)

	cri = sim_distance(critics, "Mick LaSalle", "Toby")
	pea = sim_pearson(critics, "Mick LaSalle", "Toby")
	log.Printf("[Mick LaSalle][Toby]\t %f\t%f\n", cri, pea)
	
	cri = sim_distance(critics, "Michael Phillips", "Claudia Puig")
	pea = sim_pearson(critics, "Michael Phillips", "Claudia Puig")
	log.Printf("[Michael Phillips][Claudia Puig]\t %f\t%f\n", cri, pea)
	

}
