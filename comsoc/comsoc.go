package comsoc

import "errors"

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) int {
	var ret int
	for i, p := range prefs {
		if p == alt {
			ret = i
		}
	}
	return ret
}

// renvoie vrai ssi alt1 est préférée à alt2
func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	if rank(alt1, prefs) > rank(alt2, prefs) {
		return true
	} else {
		return false
	}
}

// renvoie les meilleures alternatives pour un décomtpe donné
func maxCount(count Count) (bestAlts []Alternative) {
	// n := len(count)
	// bestscores := make([]int, n)
	// n = 0
	// for i, v := range count {
	// 	if bestscores[n] < v {
	// 		bestscores[n] = v
	// 	} else if bestscores[n] == v {
	// 		n++
	// 		bestscores[n] = v
	// 	}
	// }
	// return bests[:n+1]
	var max int = 0
	var bestalts []Alternative

	for _, v := range count {
		if v > max {
			max = v
		}
	}

	for i, v := range count {
		if v == max {
			bestalts = append(bestalts, i)
		}
	}

	return bestalts
}

// vérifie les préférences d'un agent, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois
func checkProfile(prefs []Alternative, alts []Alternative) error {
	if len(prefs) != len(alts) {
		return errors.New("Non Complet ou Plusier fois apparition")
	}

	for _, alt := range alts {
		is_in := false
		for _, pre_agent := range prefs {
			if alt == pre_agent {
				is_in = true
			}
		}
		if !is_in {
			return errors.New("Non Complet")
		}
	}
	return nil
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative de alts apparaît exactement une fois par préférences
func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	for _, pref := range prefs {
		if checkProfile(pref, alts) != nil {
			return errors.New("Non Complet ou Plusier fois apparition")
		}
	}
	return nil
}

func MajoritySWF(p Profile) (count Count, err error) {
	c := make(Count)
	for _, pref := range p {
		c[pref[0]]++
	}
	return c, nil
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	count, _ := MajoritySWF(p)
	return maxCount(count), nil
}

func BordaSWF(p Profile) (count Count, err error) {
	nb := len(p[0])
	c := make(Count)
	for _, pref := range p {
		for rank, alt := range pref {
			c[alt] += (nb - rank - 1)
		}
	}
	return c, nil
}
func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	count, _ := BordaSWF(p)
	return maxCount(count), nil
}

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	c := make(Count)
	for n, pref := range p {
		for i := 0; i < thresholds[n]; i++ {
			c[pref[i]]++
		}
	}
	return c, nil
}
func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	count, _ := ApprovalSWF(p, thresholds)
	return maxCount(count), nil
}

func TieBreakFactory(orderedAlts []Alternative) func([]Alternative) (Alternative, error) {
	return func(a []Alternative) (Alternative, error) {
		var rk_max int
		for _, alt := range a {
			rk := rank(alt, orderedAlts)
			if rk > rk_max {
				rk_max = rk
			}
		}
		return orderedAlts[rk_max], nil
	}
}

// //func SWFFactory(func swf(p Profile) (Count, error), func ([]Alternative) (Alternative, error)) (func(Profile) ([]Alternative, error))

// //func SCFFactory(func scf(p Profile) ([]Alternative, error), func ([]Alternative) (Alternative, error)) (func(Profile) (Alternative, error))

// func CondorcetWinner(p Profile) (bestAlts []Alternative, err error)
