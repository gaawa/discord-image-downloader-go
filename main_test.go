package main

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func init() {
	RegexpUrlTwitter, _ = regexp.Compile(REGEXP_URL_TWITTER)
	RegexpUrlTistory, _ = regexp.Compile(REGEXP_URL_TISTORY)
	RegexpUrlTistoryWithCDN, _ = regexp.Compile(REGEXP_URL_TISTORY_WITH_CDN)
	RegexpUrlGfycat, _ = regexp.Compile(REGEXP_URL_GFYCAT)
	RegexpUrlInstagram, _ = regexp.Compile(REGEXP_URL_INSTAGRAM)
	RegexpUrlImgurSingle, _ = regexp.Compile(REGEXP_URL_IMGUR_SINGLE)
	RegexpUrlImgurAlbum, _ = regexp.Compile(REGEXP_URL_IMGUR_ALBUM)
	RegexpUrlGoogleDrive, _ = regexp.Compile(REGEXP_URL_GOOGLEDRIVE)
	RegexpUrlPossibleTistorySite, _ = regexp.Compile(REGEXP_URL_POSSIBLE_TISTORY_SITE)
}

type urlsTestpair struct {
	value  string
	result map[string]string
}

var getTwitterUrlsTests = []urlsTestpair{
	{
		"https://pbs.twimg.com/media/CulDBM6VYAA-YhY.jpg:orig",
		map[string]string{"https://pbs.twimg.com/media/CulDBM6VYAA-YhY.jpg:orig": "CulDBM6VYAA-YhY.jpg"},
	},
	{
		"https://pbs.twimg.com/media/CulDBM6VYAA-YhY.jpg",
		map[string]string{"https://pbs.twimg.com/media/CulDBM6VYAA-YhY.jpg:orig": "CulDBM6VYAA-YhY.jpg"},
	},
	{
		"http://pbs.twimg.com/media/CulDBM6VYAA-YhY.jpg",
		map[string]string{"https://pbs.twimg.com/media/CulDBM6VYAA-YhY.jpg:orig": "CulDBM6VYAA-YhY.jpg"},
	},
}

func TestGetTwitterUrls(t *testing.T) {
	for _, pair := range getTwitterUrlsTests {
		v, err := getTwitterUrls(pair.value)
		if err != nil {
			t.Errorf("For %v, expected %v, got %v", pair.value, nil, err)
		}
		if !reflect.DeepEqual(v, pair.result) {
			t.Errorf("For %s, expected %s, got %s", pair.value, pair.result, v)
		}
	}
}

var getTistoryUrlsTests = []urlsTestpair{
	{
		"http://cfile25.uf.tistory.com/original/235CA739582E86992EFC4E",
		map[string]string{"http://cfile25.uf.tistory.com/original/235CA739582E86992EFC4E": ""},
	},
	{
		"http://cfile25.uf.tistory.com/image/235CA739582E86992EFC4E",
		map[string]string{"http://cfile25.uf.tistory.com/original/235CA739582E86992EFC4E": ""},
	},
}

func TestGetTistoryUrls(t *testing.T) {
	for _, pair := range getTistoryUrlsTests {
		v, err := getTistoryUrls(pair.value)
		if err != nil {
			t.Errorf("For %v, expected %v, got %v", pair.value, nil, err)
		}
		if !reflect.DeepEqual(v, pair.result) {
			t.Errorf("For %s, expected %s, got %s", pair.value, pair.result, v)
		}
	}
}

var getGfycatUrlsTests = []urlsTestpair{
	{
		"https://gfycat.com/SandyChiefBoubou",
		map[string]string{"https://giant.gfycat.com/SandyChiefBoubou.gif": ""},
	},
}

func TestGetGfycatUrls(t *testing.T) {
	for _, pair := range getGfycatUrlsTests {
		v, err := getGfycatUrls(pair.value)
		if err != nil {
			t.Errorf("For %v, expected %v, got %v", pair.value, nil, err)
		}
		if !reflect.DeepEqual(v, pair.result) {
			t.Errorf("For %s, expected %s, got %s", pair.value, pair.result, v)
		}
	}
}

var getInstagramUrlsPictureTests = []urlsTestpair{
	{
		"https://www.instagram.com/p/BHhDAmhAz33/?taken-by=s_sohye",
		map[string]string{"https://www.instagram.com/p/BHhDAmhAz33/media/?size=l&taken-bys_sohye": ""},
	},
	{
		"https://www.instagram.com/p/BHhDAmhAz33/",
		map[string]string{"https://www.instagram.com/p/BHhDAmhAz33/media/?size=l": ""},
	},
}

var getInstagramUrlsVideoTests = []urlsTestpair{
	{
		"https://www.instagram.com/p/BL2_ZIHgYTp/?taken-by=s_sohye",
		map[string]string{"14811404_233311497085396_338650092456116224_n.mp4": ""},
	},
	{
		"https://www.instagram.com/p/BL2_ZIHgYTp/",
		map[string]string{"14811404_233311497085396_338650092456116224_n.mp4": ""},
	},
}

func TestGetInstagramUrls(t *testing.T) {
	for _, pair := range getInstagramUrlsPictureTests {
		v, err := getInstagramUrls(pair.value)
		if err != nil {
			t.Errorf("For %v, expected %v, got %v", pair.value, nil, err)
		}
		if !reflect.DeepEqual(v, pair.result) {
			t.Errorf("For %s, expected %s, got %s", pair.value, pair.result, v)
		}
	}
	for _, pair := range getInstagramUrlsVideoTests {
		v, err := getInstagramUrls(pair.value)
		if err != nil {
			t.Errorf("For %v, expected %v, got %v", pair.value, nil, err)
		}
		for keyResult, valueResult := range pair.result {
			for keyExpected, valueExpected := range v {
				if strings.Contains(keyResult, keyExpected) || valueResult != valueExpected { // CDN location can vary
					t.Errorf("For %s, expected %s, got %s", pair.value, pair.result, v)
				}
			}
		}
	}
}

var getImgurSingleUrlsTests = []urlsTestpair{
	{
		"http://imgur.com/viZictl",
		map[string]string{"http://imgur.com/download/viZictl": ""},
	},
	{
		"https://imgur.com/viZictl",
		map[string]string{"https://imgur.com/download/viZictl": ""},
	},
	{
		"https://i.imgur.com/viZictl.jpg",
		map[string]string{"https://i.imgur.com/download/viZictl.jpg": ""},
	},
	{
		"http://imgur.com/uYwt2VV",
		map[string]string{"http://imgur.com/download/uYwt2VV": ""},
	},
	{
		"http://i.imgur.com/uYwt2VV.gifv",
		map[string]string{"http://i.imgur.com/download/uYwt2VV": ""},
	},
}

func TestGetImgurSingleUrls(t *testing.T) {
	for _, pair := range getImgurSingleUrlsTests {
		v, err := getImgurSingleUrls(pair.value)
		if err != nil {
			t.Errorf("For %v, expected %v, got %v", pair.value, nil, err)
		}
		if !reflect.DeepEqual(v, pair.result) {
			t.Errorf("For %s, expected %s, got %s", pair.value, pair.result, v)
		}
	}
}

var getImgurAlbumUrlsTests = []urlsTestpair{
	{
		"http://imgur.com/a/ALTpi",
		map[string]string{
			"http://i.imgur.com/FKoguPh.jpg": "",
			"http://i.imgur.com/5FNL6Pe.jpg": "",
			"http://i.imgur.com/YA0V0g9.jpg": "",
			"http://i.imgur.com/Uc2iDhD.jpg": "",
			"http://i.imgur.com/J9JRSSJ.jpg": "",
			"http://i.imgur.com/Xrx0uyE.jpg": "",
			"http://i.imgur.com/3xDSq1O.jpg": "",
		},
	},
}

func TestGetImgurAlbumUrls(t *testing.T) {
	for _, pair := range getImgurAlbumUrlsTests {
		v, err := getImgurAlbumUrls(pair.value)
		if err != nil {
			t.Errorf("For %v, expected %v, got %v", pair.value, nil, err)
		}
		if !reflect.DeepEqual(v, pair.result) {
			t.Errorf("For %s, expected %s, got %s", pair.value, pair.result, v)
		}
	}
}

var getGoogleDriveUrlsTests = []urlsTestpair{
	{
		"https://drive.google.com/file/d/0B8TnwsJqlFllSUtvUEhoSU40WkE/view",
		map[string]string{"https://drive.google.com/uc?export=download&id=0B8TnwsJqlFllSUtvUEhoSU40WkE": ""},
	},
}

func TestGetGoogleDriveUrls(t *testing.T) {
	for _, pair := range getGoogleDriveUrlsTests {
		v, err := getGoogleDriveUrls(pair.value)
		if err != nil {
			t.Errorf("For %v, expected %v, got %v", pair.value, nil, err)
		}
		if !reflect.DeepEqual(v, pair.result) {
			t.Errorf("For %s, expected %s, got %s", pair.value, pair.result, v)
		}
	}
}

var getTistoryWithCDNUrlsTests = []urlsTestpair{
	{
		"http://img1.daumcdn.net/thumb/R720x0.q80/?scode=mtistory&fname=http%3A%2F%2Fcfile24.uf.tistory.com%2Fimage%2F2658554B580BDC4C0924CA",
		map[string]string{"http://cfile24.uf.tistory.com/original/2658554B580BDC4C0924CA": ""},
	},
}

func TestGetTistoryWithCDNUrls(t *testing.T) {
	for _, pair := range getTistoryWithCDNUrlsTests {
		v, err := getTistoryWithCDNUrls(pair.value)
		if err != nil {
			t.Errorf("For %v, expected %v, got %v", pair.value, nil, err)
		}
		if !reflect.DeepEqual(v, pair.result) {
			t.Errorf("For %s, expected %s, got %s", pair.value, pair.result, v)
		}
	}
}

var getPossibleTistorySiteUrlsTests = []urlsTestpair{
	{
		"http://shaq32.tistory.com/395",
		map[string]string{
			"http://cfile24.uf.tistory.com/original/2658554B580BDC4C0924CA": "A10Y4472.jpg",
			"http://cfile21.uf.tistory.com/original/26469D4B580BDC4D0D0387": "A27R9111.jpg",
			"http://cfile29.uf.tistory.com/original/2744A54B580BDC4F0EFAEE": "A10Y5545.jpg",
			"http://cfile28.uf.tistory.com/original/2758534B580BDC500969E4": "A10Y5613.jpg",
		},
	},
	{
		"http://shaq32.tistory.com/m/395",
		map[string]string{
			"http://cfile24.uf.tistory.com/original/2658554B580BDC4C0924CA": "",
			"http://cfile21.uf.tistory.com/original/26469D4B580BDC4D0D0387": "",
			"http://cfile29.uf.tistory.com/original/2744A54B580BDC4F0EFAEE": "",
			"http://cfile28.uf.tistory.com/original/2758534B580BDC500969E4": "",
		},
	},
	{
		"http://slmn.de/123",
		map[string]string{},
	},
}

func TestGetPossibleTistorySiteUrls(t *testing.T) {
	for _, pair := range getPossibleTistorySiteUrlsTests {
		v, err := getPossibleTistorySiteUrls(pair.value)
		if err != nil {
			t.Errorf("For %v, expected %v, got %v", pair.value, nil, err)
		}
		if !reflect.DeepEqual(v, pair.result) {
			t.Errorf("For %s, expected %s, got %s", pair.value, pair.result, v)
		}
	}
}