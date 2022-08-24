package phonenumbers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//Splits a range of numbers and returns two ints
//example "02080668800 - 02080668850"
func SplitRange(input string, deliminator string) (int, int) {

	numbers := strings.Split(input, deliminator)

	return ConverttoE164int(numbers[0]), ConverttoE164int(numbers[1])

}

//returns slice of each number between two number ranges
//can be formatted 02080668800 - 02080668850 or "02080668800 - 850"
func ListRange(firstnum int, secondnum int) []string {

	var divideby int
	var output []string

	if secondnum < 9999 {

		if DigitCount(secondnum) == 2 {
			divideby = 100
		} else if DigitCount(secondnum) == 3 {
			divideby = 1000
		} else if DigitCount(secondnum) == 4 {
			divideby = 10000
		}

		base := firstnum / divideby
		startingdigits := RetrieveTrailingDigits(firstnum, DigitCount(secondnum))

		for count := ConverttoE164int(startingdigits); count <= secondnum; count++ {
			output = append(output, fmt.Sprintf("%v%v", base, count))
		}
		return output

	} else {

		for count := firstnum; count <= secondnum; count++ {
			output = append(output, fmt.Sprint(count))
		}
		return output

	}

}

//Retrieves the trailing digits of a number in string format
func RetrieveTrailingDigits(number int, decimalplace int) string {

	numberstring := fmt.Sprintf("%v", number)

	trailingdigits := numberstring[len(numberstring)-decimalplace:]

	return trailingdigits

}

//Returns count for decimal places of a number
func DigitCount(number int) int {

	count := 0

	for number != 0 {
		number /= 10
		count += 1
	}

	return count

}

//Convert telephone number in string format to int in e164
//expected format "02080668800" or "200" for when number ranges are written like "02080668800 - 8850"
func ConverttoE164int(number string) int {

	number = strings.TrimSpace(number)

	if len(number) < 5 {

		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		return n
	}

	number = ConverttoE164(number)

	n, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}

	return n

}

//Takes a number in string format, converts it to e164 format.
//02080668866 to 442080668866
func ConverttoE164(number string) string {
	number = number[1:]
	number = "44" + number

	return number
}

//Helper function returns country codes for telephone numbers.
//Can pass string location ie Canada to retrieve []int ISO prefix
//Pass an prefix as an int i.e 44 to return the country name in []string (can be multiple countries i.e canada or america are 1)
//or pass nil for a full map map[string]int
func GeoData(option interface{}) (interface{}, error) {

	countrymap := make(map[string]int)

	countrymap["Canada"] = 1
	countrymap["United States"] = 1
	countrymap["Bahamas"] = 1242
	countrymap["Barbados"] = 1246
	countrymap["Anguilla"] = 1264
	countrymap["Antigua and Barbuda"] = 1268
	countrymap["Barbuda"] = 1268
	countrymap["British Virgin Islands"] = 1284
	countrymap["US Virgin Islands"] = 1340
	countrymap["Cayman Islands"] = 1345
	countrymap["Bermuda"] = 1441
	countrymap["Grenada"] = 1473
	countrymap["Turks and Caicos Islands"] = 1649
	countrymap["Montserrat"] = 1664
	countrymap["Northern Mariana Islands"] = 1670
	countrymap["Northern Cyprus"] = 90392
	countrymap["Guam"] = 1671
	countrymap["American Samoa"] = 1684
	countrymap["Sint Maarten (Netherlands)"] = 1721
	countrymap["Saint Lucia"] = 1758
	countrymap["Dominica"] = 1767
	countrymap["Saint Vincent and the Grenadines"] = 1784
	countrymap["Puerto Rico"] = 1787
	countrymap["Puerto Rico 1"] = 1939
	countrymap["Midway Island USA"] = 1808
	countrymap["Wake Island USA"] = 1808
	countrymap["Dominican Republic"] = 1809
	countrymap["Dominican Republic 1"] = 1829
	countrymap["Dominican Republic 2"] = 1849
	countrymap["Trinidad and Tobago"] = 1868
	countrymap["Nevis"] = 1869
	countrymap["Saint Kitts and Nevis"] = 1869
	countrymap["Jamaica"] = 1658
	countrymap["Jamaica"] = 1876
	countrymap["Egypt"] = 20
	countrymap["South Sudan"] = 211
	countrymap["Morocco"] = 212
	countrymap["Algeria"] = 213
	countrymap["Tunisia"] = 216
	countrymap["Libya"] = 218
	countrymap["Gambia"] = 220
	countrymap["Senegal"] = 221
	countrymap["Mauritania"] = 222
	countrymap["Mali"] = 223
	countrymap["Guinea"] = 224
	countrymap["Ivory Coast (Côte d'Ivoire)"] = 225
	countrymap["Burkina Faso"] = 226
	countrymap["Niger"] = 227
	countrymap["Togo"] = 228
	countrymap["Benin"] = 229
	countrymap["Mauritius"] = 230
	countrymap["Liberia"] = 231
	countrymap["Sierra Leone"] = 232
	countrymap["Ghana"] = 233
	countrymap["Nigeria"] = 234
	countrymap["Chad"] = 235
	countrymap["Central African Republic"] = 236
	countrymap["Cameroon"] = 237
	countrymap["Cape Verde"] = 238
	countrymap["São Tomé and Príncipe"] = 239
	countrymap["Equatorial Guinea"] = 240
	countrymap["Gabon"] = 241
	countrymap["Congo"] = 242
	countrymap["Congo Democratic Republic of the"] = 243
	countrymap["Angola"] = 244
	countrymap["Guinea-Bissau"] = 245
	countrymap["British Indian Ocean Territory"] = 246
	countrymap["Diego Garcia"] = 246
	countrymap["Ascension"] = 247
	countrymap["Seychelles"] = 248
	countrymap["Sudan"] = 249
	countrymap["Rwanda"] = 250
	countrymap["Ethiopia"] = 251
	countrymap["Somalia"] = 252
	countrymap["Djibouti"] = 253
	countrymap["Kenya"] = 254
	countrymap["Tanzania"] = 255
	countrymap["Zanzibar"] = 25524
	countrymap["Uganda"] = 256
	countrymap["Burundi"] = 257
	countrymap["Mozambique"] = 258
	countrymap["Zambia"] = 260
	countrymap["Madagascar"] = 261
	countrymap["Réunion"] = 262
	countrymap["Mayotte"] = 262269
	countrymap["Mayotte 1"] = 262639
	countrymap["Zimbabwe"] = 263
	countrymap["Namibia"] = 264
	countrymap["Malawi"] = 265
	countrymap["Lesotho"] = 266
	countrymap["Botswana"] = 267
	countrymap["Eswatini"] = 268
	countrymap["Comoros"] = 269
	countrymap["South Africa"] = 27
	countrymap["Saint Helena"] = 290
	countrymap["Tristan da Cunha"] = 2908
	countrymap["Eritrea"] = 291
	countrymap["Aruba"] = 297
	countrymap["Faroe Islands"] = 298
	countrymap["Greenland"] = 299
	countrymap["Greece"] = 30
	countrymap["Netherlands"] = 31
	countrymap["Belgium"] = 32
	countrymap["France"] = 33
	countrymap["Spain"] = 34
	countrymap["Gibraltar"] = 350
	countrymap["Portugal"] = 351
	countrymap["Luxembourg"] = 352
	countrymap["Ireland"] = 353
	countrymap["Iceland"] = 354
	countrymap["Albania"] = 355
	countrymap["Malta"] = 356
	countrymap["Cyprus"] = 357
	countrymap["Finland"] = 358
	countrymap["Åland"] = 35818
	countrymap["Bulgaria"] = 359
	countrymap["Hungary"] = 36
	countrymap["Lithuania"] = 370
	countrymap["Latvia"] = 371
	countrymap["Estonia"] = 372
	countrymap["Moldova"] = 373
	countrymap["Transnistria"] = 3732
	countrymap["Transnistria 1"] = 3735
	countrymap["Armenia"] = 374
	countrymap["Artsakh"] = 37447
	countrymap["Artsakh2"] = 37497
	countrymap["Belarus"] = 375
	countrymap["Andorra"] = 376
	countrymap["Monaco"] = 377
	countrymap["San Marino"] = 378
	countrymap["Ukraine"] = 380
	countrymap["Serbia"] = 381
	countrymap["Montenegro"] = 382
	countrymap["Kosovo"] = 383
	countrymap["Croatia"] = 385
	countrymap["Slovenia"] = 386
	countrymap["Bosnia and Herzegovina"] = 387
	countrymap["North Macedonia"] = 389
	countrymap["Italy"] = 39
	countrymap["Romania"] = 40
	countrymap["Switzerland"] = 41
	countrymap["Czech Republic"] = 420
	countrymap["Slovakia"] = 421
	countrymap["Liechtenstein"] = 423
	countrymap["Austria"] = 43
	countrymap["United Kingdom"] = 44
	countrymap["Guernsey"] = 441481
	countrymap["Guernsey 1"] = 447781
	countrymap["Guernsey 2"] = 447839
	countrymap["Guernsey 3"] = 447911
	countrymap["Jersey"] = 441534
	countrymap["Isle of Man"] = 441624
	countrymap["Isle of Man 1"] = 447524
	countrymap["Isle of Man 2"] = 447624
	countrymap["Isle of Man 3"] = 447924
	countrymap["Northern Ireland"] = 4428
	countrymap["Denmark"] = 45
	countrymap["Sweden"] = 46
	countrymap["Norway"] = 47
	countrymap["Jan Mayen"] = 4779
	countrymap["Svalbard"] = 4779
	countrymap["Poland"] = 48
	countrymap["Germany"] = 49
	countrymap["Falkland Islands"] = 500
	countrymap["South Georgia and the South Sandwich Islands"] = 500
	countrymap["Belize"] = 501
	countrymap["Guatemala"] = 502
	countrymap["El Salvador"] = 503
	countrymap["Honduras"] = 504
	countrymap["Nicaragua"] = 505
	countrymap["Costa Rica"] = 506
	countrymap["Panama"] = 507
	countrymap["Saint Pierre and Miquelon"] = 508
	countrymap["Haiti"] = 509
	countrymap["Peru"] = 51
	countrymap["Mexico"] = 52
	countrymap["Cuba"] = 53
	countrymap["Argentina"] = 54
	countrymap["Brazil"] = 55
	countrymap["Chile"] = 56
	countrymap["Easter Island"] = 56
	countrymap["Colombia"] = 57
	countrymap["Venezuela"] = 58
	countrymap["Guadeloupe"] = 590
	countrymap["Saint Barthélemy"] = 590
	countrymap["Saint Martin (France)"] = 590
	countrymap["Bolivia"] = 591
	countrymap["Guyana"] = 592
	countrymap["Ecuador"] = 593
	countrymap["French Guiana"] = 594
	countrymap["Paraguay"] = 595
	countrymap["French Antilles"] = 596
	countrymap["Martinique"] = 596
	countrymap["Suriname"] = 597
	countrymap["Uruguay"] = 598
	countrymap["Sint Eustatius"] = 5993
	countrymap["Caribbean Netherlands"] = 5993
	countrymap["Caribbean Netherlands 1"] = 5994
	countrymap["Caribbean Netherlands 2"] = 5997
	countrymap["Saba"] = 5994
	countrymap["Bonaire"] = 5997
	countrymap["Curaçao"] = 5999
	countrymap["Malaysia"] = 60
	countrymap["Australia"] = 61
	countrymap["Cocos (Keeling) Islands"] = 6189162
	countrymap["Christmas Island"] = 6189164
	countrymap["Indonesia"] = 62
	countrymap["Philippines"] = 63
	countrymap["Chatham Island New Zealand"] = 64
	countrymap["New Zealand"] = 64
	countrymap["Pitcairn Islands"] = 64
	countrymap["Singapore"] = 65
	countrymap["Thailand"] = 66
	countrymap["East Timor (Timor-Leste)"] = 670
	countrymap["Australian External Territories"] = 672
	countrymap["Australian Antarctic Territory"] = 6721
	countrymap["Norfolk Island"] = 6723
	countrymap["Brunei Darussalam"] = 673
	countrymap["Nauru"] = 674
	countrymap["Papua New Guinea"] = 675
	countrymap["Tonga"] = 676
	countrymap["Solomon Islands"] = 677
	countrymap["Vanuatu"] = 678
	countrymap["Fiji"] = 679
	countrymap["Palau"] = 680
	countrymap["Wallis and Futuna"] = 681
	countrymap["Cook Islands"] = 682
	countrymap["Niue"] = 683
	countrymap["Samoa"] = 685
	countrymap["Kiribati"] = 686
	countrymap["New Caledonia"] = 687
	countrymap["Tuvalu"] = 688
	countrymap["French Polynesia"] = 689
	countrymap["Tokelau"] = 690
	countrymap["Micronesia Federated States of"] = 691
	countrymap["Marshall Islands"] = 692
	countrymap["Kazakhstan"] = 76
	countrymap["Kazakhstan 2"] = 77
	countrymap["Japan"] = 81
	countrymap["Korea South"] = 82
	countrymap["Vietnam"] = 84
	countrymap["Korea North"] = 850
	countrymap["Hong Kong"] = 852
	countrymap["Macau"] = 853
	countrymap["Cambodia"] = 855
	countrymap["Laos"] = 856
	countrymap["China"] = 86
	countrymap["Bangladesh"] = 880
	countrymap["Taiwan"] = 886
	countrymap["Turkey"] = 90
	countrymap["India"] = 91
	countrymap["Pakistan"] = 92
	countrymap["Afghanistan"] = 93
	countrymap["Sri Lanka"] = 94
	countrymap["Myanmar"] = 95
	countrymap["Maldives"] = 960
	countrymap["Lebanon"] = 961
	countrymap["Jordan"] = 962
	countrymap["Syria"] = 963
	countrymap["Iraq"] = 964
	countrymap["Kuwait"] = 965
	countrymap["Saudi Arabia"] = 966
	countrymap["Yemen"] = 967
	countrymap["Oman"] = 968
	countrymap["Palestine State of"] = 970
	countrymap["United Arab Emirates"] = 971
	countrymap["Israel"] = 972
	countrymap["Bahrain"] = 973
	countrymap["Qatar"] = 974
	countrymap["Bhutan"] = 975
	countrymap["Mongolia"] = 976
	countrymap["Nepal"] = 977
	countrymap["Iran"] = 98
	countrymap["Tajikistan"] = 992
	countrymap["Turkmenistan"] = 993
	countrymap["Azerbaijan"] = 994
	countrymap["Georgia"] = 995
	countrymap["South Ossetia"] = 99534
	countrymap["Kyrgyzstan"] = 996
	countrymap["Uzbekistan"] = 998

	switch option.(type) {
	case string:

		var countrycodearray []int

		//There's a couple of countries that have multiple prefixes so rather than query the map directly we're looking for a contains
		for countryname, countrycode := range countrymap {
			if strings.Contains(countryname, option.(string)) {

				countrycodearray = append(countrycodearray, countrycode)

			}
		}

		if len(countrycodearray) == 0 {
			return "", errors.New("Country Not found")
		}

		return countrycodearray, nil
	case int:

		var countryarray []string

		for countryname, countrycode := range countrymap {
			if countrycode == option.(int) {
				countryarray = append(countryarray, countryname)
			}
		}

		if len(countryarray) == 0 {
			return "", errors.New("Prefix Not Found")
		}
		return countryarray, nil
	}

	return countrymap, nil
}
