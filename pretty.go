package main

import (
	"regexp"
)

func isPrettyKey(fpr string) bool {
	patterns := []*regexp.Regexp{
		// Repetitions
		regexp.MustCompile(`000000|111111|222222|333333|444444|555555|666666|777777|888888|999999|AAAAAA|BBBBBB|CCCCCC|DDDDDD|EEEEEE|FFFFFF`),
		regexp.MustCompile(`123456|234567|345678|456789|56789A|6789AB|789ABC|89ABCD|9ABCDE|ABCDEF`),
		// Math & physics constants
		regexp.MustCompile(`314159|271828|60221|1380[67]|8314[45]|299792|6626[01]|66738|9109[34]|40490FDB|27315`),
		// Magic debug values
		regexp.MustCompile(`A55E55|ACCE55|BA0BAB|CA5CADE|DECADE|DEC0DE|D0D0E5|FACADE|0B5E55|5EABED|5EAF00D`),
		regexp.MustCompile(`BAAAAAAD|BAD22222|BADBADBADBAD|BAADF00D|BADDCAFE|BEEFCACE|C00010FF|CAFEBABE|CAFED00D|CAFEFEED`),
		// Famous words
		regexp.MustCompile(`DEADBEEF|C0FFEE`),
	}

	for _, v := range patterns {
		if v.MatchString(fpr) {
			return true
		}
	}
	return false
}
