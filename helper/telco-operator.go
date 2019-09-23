package helper

import "sort"

func TselPrefix (handphone string) bool {
	if len(hp) < 5 {
		return false
	}
	
	tselPrefix	:= []string{
		"62811",
		"62812",
		"62813",
		"62821",
		"62822",
		"62823",
		"62852",
		"62853",
		"62851",
	}
	pref := handphone[0:5]
	sort.Strings(tselPrefix)
	i := sort.SearchStrings(tselPrefix, pref)
	if i < len(tselPrefix) && tselPrefix[i] == pref {
		return true
	}

	return false
}

func IsatPrefix (handphone string) bool {
	if len(hp) < 5 {
		return false
	}

	isatPrefix	:= []string{
		"62814", // nomor untuk Indosat M2 Broadband
		"62815", // nomor Matrix dan Mentari
		"62816", // Matrix dan Mentari
		"62855", // adalah nomor Matrix
		"62856", // IM3
		"62857", // IM3
		"62858", // Mentari
	}
	pref := handphone[0:5]
	sort.Strings(isatPrefix)
	i := sort.SearchStrings(isatPrefix, pref)
	if i < len(isatPrefix) && isatPrefix[i] == pref {
		return true
	}

	return false
}

func XlPrefix (handphone string) bool {
	if len(hp) < 5 {
		return false
	}

	isatPrefix	:= []string{
		"62817",
		"62818",
		"62819",
		"62859", // Prabayar dan pascabayar
		"62877", // Prabayar
		"62878", // Prabayar
		"62879", // Prabayar
	}
	pref := handphone[0:5]
	sort.Strings(isatPrefix)
	i := sort.SearchStrings(isatPrefix, pref)
	if i < len(isatPrefix) && isatPrefix[i] == pref {
		return true
	}

	return false
}

func ThreePrefix (hp string) bool {
	if len(hp) < 5 {
		return false
	}

	triPrefix := []string {
		"62895",
		"62896",
		"62897",
		"62898",
		"62899",
	}
	pref := hp[0:5]
	sort.Strings(triPrefix)
	i := sort.SearchStrings(triPrefix, pref)
	if i < len(triPrefix) && triPrefix[i] == pref {
		return true
	}

	return false
}