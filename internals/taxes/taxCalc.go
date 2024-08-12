package taxes

import (
	"slices"
)

type Product struct {
	AlcoholPercentage   float64       `json:"alcoholPercentage"`
	AlcoholTax          float64       `json:"alcoholTax"`
	Assortment          string        `json:"assortment"`
	AssortmentText      string        `json:"assortmentText"`
	BottleText          string        `json:"bottleText"`
	Category            interface{}   `json:"category"`
	CategoryLevel1      string        `json:"categoryLevel1"`
	CategoryLevel2      string        `json:"categoryLevel2"`
	CategoryLevel3      string        `json:"categoryLevel3"`
	CategoryLevel4      interface{}   `json:"categoryLevel4"`
	Color               string        `json:"color"`
	Country             string        `json:"country"`
	CustomCategoryTitle string        `json:"customCategoryTitle"`
	DishPoints          interface{}   `json:"dishPoints"`
	EthicalLabel        interface{}   `json:"ethicalLabel"`
	Grapes              []interface{} `json:"grapes"`
	ImageModules        struct {
		Extensions interface{} `json:"extensions"`
		ProductID  interface{} `json:"productId"`
		Sizes      interface{} `json:"sizes"`
		Thumbnail  interface{} `json:"thumbnail"`
	} `json:"imageModules"`
	Images []struct {
		FileType interface{} `json:"fileType"`
		ImageURL string      `json:"imageUrl"`
		Size     interface{} `json:"size"`
	} `json:"images"`
	IsAlcoholFree                   bool        `json:"isAlcoholFree"`
	IsBsAssortment                  bool        `json:"isBsAssortment"`
	IsClimateSmartPackaging         bool        `json:"isClimateSmartPackaging"`
	IsCompletelyOutOfStock          bool        `json:"isCompletelyOutOfStock"`
	IsDiscontinued                  bool        `json:"isDiscontinued"`
	IsEthical                       bool        `json:"isEthical"`
	IsFsAssortment                  bool        `json:"isFsAssortment"`
	IsFsTsAssortment                bool        `json:"isFsTsAssortment"`
	IsKosher                        bool        `json:"isKosher"`
	IsManufacturingCountry          bool        `json:"isManufacturingCountry"`
	IsNews                          bool        `json:"isNews"`
	IsOrganic                       bool        `json:"isOrganic"`
	IsPaAssortment                  bool        `json:"isPaAssortment"`
	IsRecommendedByTasteProfile     bool        `json:"isRecommendedByTasteProfile"`
	IsRegionalRestricted            bool        `json:"isRegionalRestricted"`
	IsSupplierTemporaryNotAvailable bool        `json:"isSupplierTemporaryNotAvailable"`
	IsSustainableChoice             bool        `json:"isSustainableChoice"`
	IsTemporaryOutOfStock           bool        `json:"isTemporaryOutOfStock"`
	IsTsAssortment                  bool        `json:"isTsAssortment"`
	IsTsLsAssortment                bool        `json:"isTsLsAssortment"`
	IsTseAssortment                 bool        `json:"isTseAssortment"`
	IsTssAssortment                 bool        `json:"isTssAssortment"`
	IsTstAssortment                 bool        `json:"isTstAssortment"`
	IsTsvAssortment                 bool        `json:"isTsvAssortment"`
	IsWebLaunch                     bool        `json:"isWebLaunch"`
	OriginLevel1                    interface{} `json:"originLevel1"`
	OriginLevel2                    interface{} `json:"originLevel2"`
	OriginalPrice                   float32     `json:"originalPrice"`
	OtherSelections                 interface{} `json:"otherSelections"`
	PackagingLevel1                 string      `json:"packagingLevel1"`
	Price                           float32     `json:"price"`
	ProducerName                    string      `json:"producerName"`
	ProductID                       string      `json:"productId"`
	ProductLaunchDate               string      `json:"productLaunchDate"`
	ProductNameBold                 string      `json:"productNameBold"`
	ProductNameThin                 interface{} `json:"productNameThin"`
	ProductNumber                   string      `json:"productNumber"`
	ProductNumberShort              string      `json:"productNumberShort"`
	RecycleFee                      int         `json:"recycleFee"`
	RestrictedParcelQuantity        int         `json:"restrictedParcelQuantity"`
	Seal                            interface{} `json:"seal"`
	SellStartTime                   string      `json:"sellStartTime"`
	SkatteverketType                string      `json:"skatteverketType"`
	SugarContent                    int         `json:"sugarContent"`
	SugarContentGramPer100Ml        int         `json:"sugarContentGramPer100ml"`
	SupplierName                    string      `json:"supplierName"`
	SystembolagetType               string      `json:"systembolagetType"`
	SystembolagetUnitCut            float32     `json:"systembolagetUnitCut"`
	SystembolagetPercentageCut      float32     `json:"systembolagetPercentageCut"`
	Taste                           string      `json:"taste"`
	TasteClockBitter                int         `json:"tasteClockBitter"`
	TasteClockBody                  int         `json:"tasteClockBody"`
	TasteClockCasque                int         `json:"tasteClockCasque"`
	TasteClockFruitacid             int         `json:"tasteClockFruitacid"`
	TasteClockGroupBitter           interface{} `json:"tasteClockGroupBitter"`
	TasteClockGroupSmokiness        interface{} `json:"tasteClockGroupSmokiness"`
	TasteClockRoughness             int         `json:"tasteClockRoughness"`
	TasteClockSmokiness             int         `json:"tasteClockSmokiness"`
	TasteClockSweetness             int         `json:"tasteClockSweetness"`
	TasteClocks                     []struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	} `json:"tasteClocks"`
	TasteSymbols []string    `json:"tasteSymbols"`
	Usage        string      `json:"usage"`
	Vat          float32     `json:"moms"`
	Vintage      interface{} `json:"vintage"`
	Volume       int         `json:"volume"`
	VolumeText   string      `json:"volumeText"`
}

type VariableResponse struct {
	Systembolaget Systembolaget `json:"systembolaget"`
	Skatteverket  Skatteverket  `json:"skatteverket"`
}

type Systembolaget struct {
	PercentageCut float32  `json:"percentageCut"`
	UnitCuts      UnitCuts `json:"unitcuts"`
}

type UnitCuts struct {
	Spirits           float32 `json:"spirits"`
	Wine              float32 `json:"wine"`
	Cider             float32 `json:"cider"`
	Beer              float32 `json:"beer"`
	WineNonAlcoholic  float32 `json:"wineNonAlcoholic"`
	CiderNonAlcoholic float32 `json:"ciderNonAlcoholic"`
}

type Skatteverket struct {
	Moms          float32              `json:"moms"`
	TaxCategories map[string][]TaxInfo `json:"taxCategories"`
	Beer          []TaxInfo            `json:"beer"`
	Wine          []TaxInfo            `json:"wine"`
	MiddleClass   []TaxInfo            `json:"middleClass"`
	Spirits       []TaxInfo            `json:"spirits"`
}

type TaxInfo struct {
	AlcFrom               float64 `json:"alcFrom"`
	AlcTo                 float64 `json:"alc"`
	TaxPerPercentAndLiter float64 `json:"taxPerPercentAndLiter"`
}

var CategoriesMap = map[string]func(product *Product, taxInfo TaxInfo){
	"Öl":                   calcPerPercentageAndLiter,
	"Vin":                  calcPerLiter,
	"Sprit":                calcPerPureLiter,
	"Cider & blanddrycker": calcPerLiter,
	"Mellanklass":          calcPerLiter,
}

var Middleclass = []string{"Madeira", "Portvin", "Sherry", "Vermouth", "Övrigt starkvin"}

var Variables = VariableResponse{
	Systembolaget: Systembolaget{
		PercentageCut: 0.147,
		UnitCuts: UnitCuts{
			Spirits:           6.37,
			Wine:              5.40,
			Cider:             1.52,
			Beer:              0.95,
			WineNonAlcoholic:  5.92,
			CiderNonAlcoholic: 2.46,
		},
	},
	Skatteverket: Skatteverket{
		Moms: 0.25,
		TaxCategories: map[string][]TaxInfo{
			"Öl": {
				{AlcFrom: 0, AlcTo: 2.8, TaxPerPercentAndLiter: 0},
				{AlcFrom: 2.8, AlcTo: 100, TaxPerPercentAndLiter: 2.28},
			},
			"Vin": {
				{AlcFrom: 0, AlcTo: 2.25, TaxPerPercentAndLiter: 0},
				{AlcFrom: 2.25, AlcTo: 4.5, TaxPerPercentAndLiter: 10.38},
				{AlcFrom: 4.5, AlcTo: 7, TaxPerPercentAndLiter: 15.34},
				{AlcFrom: 7, AlcTo: 8.5, TaxPerPercentAndLiter: 21.12},
				{AlcFrom: 8.5, AlcTo: 15, TaxPerPercentAndLiter: 29.58},
			},
			"Sprit": {
				{AlcFrom: 0, AlcTo: 1.2, TaxPerPercentAndLiter: 0},
				{AlcFrom: 1.2, AlcTo: 100, TaxPerPercentAndLiter: 526.97},
			},
			"Mellanklass": {
				{AlcFrom: 0, AlcTo: 15, TaxPerPercentAndLiter: 37.34},
				{AlcFrom: 15, AlcTo: 22, TaxPerPercentAndLiter: 61.90},
				{AlcFrom: 22, AlcTo: 100, TaxPerPercentAndLiter: 526.79},
			},
		},
		Beer: []TaxInfo{
			{AlcFrom: 0, AlcTo: 2.8, TaxPerPercentAndLiter: 0},
			{AlcFrom: 2.8, AlcTo: 100, TaxPerPercentAndLiter: 2.28},
		},
		Wine: []TaxInfo{
			{AlcFrom: 0, AlcTo: 2.25, TaxPerPercentAndLiter: 0},
			{AlcFrom: 2.25, AlcTo: 4.5, TaxPerPercentAndLiter: 10.38},
			{AlcFrom: 4.5, AlcTo: 7, TaxPerPercentAndLiter: 15.34},
			{AlcFrom: 7, AlcTo: 8.5, TaxPerPercentAndLiter: 21.12},
			{AlcFrom: 8.5, AlcTo: 15, TaxPerPercentAndLiter: 29.58},
			{AlcFrom: 15, AlcTo: 18, TaxPerPercentAndLiter: 61.90},
		},
		MiddleClass: []TaxInfo{
			{AlcFrom: 0, AlcTo: 15, TaxPerPercentAndLiter: 37.34},
			{AlcFrom: 15, AlcTo: 22, TaxPerPercentAndLiter: 61.90},
		},
		Spirits: []TaxInfo{
			{AlcFrom: 0, AlcTo: 1.2, TaxPerPercentAndLiter: 0},
			{AlcFrom: 1.2, AlcTo: 100, TaxPerPercentAndLiter: 526.97},
		},
	},
}

func calcVat(product *Product) {
	product.Vat = product.Price * Variables.Skatteverket.Moms
}

func calcPerLiter(product *Product, taxInfo TaxInfo) {
	product.AlcoholTax = float64(product.Volume/1000) * taxInfo.TaxPerPercentAndLiter
}

func calcPerPercentageAndLiter(product *Product, taxInfo TaxInfo) {
	product.AlcoholTax = float64(product.Volume) / 1000 * product.AlcoholPercentage * taxInfo.TaxPerPercentAndLiter
}

func calcPerPureLiter(product *Product, taxInfo TaxInfo) {
	product.AlcoholTax = float64(product.Volume/1000) * (product.AlcoholPercentage / 100) * taxInfo.TaxPerPercentAndLiter
}

func findTaxLevel(product *Product, taxInfo []TaxInfo) TaxInfo {
	idx := slices.IndexFunc(taxInfo, func(taxInfo TaxInfo) bool {
		return taxInfo.AlcFrom < product.AlcoholPercentage && taxInfo.AlcTo >= product.AlcoholPercentage
	})

	if idx < 0 {
		return taxInfo[len(taxInfo)-1]
	}

	return taxInfo[idx]
}

func calcTax(product *Product) {
	if product.IsAlcoholFree {
		return
	}
	k := product.SkatteverketType
	CategoriesMap[k](product, findTaxLevel(product, Variables.Skatteverket.TaxCategories[k]))
}

func findSkatteverketType(product *Product) string {
	var categorySlice = []string{product.CategoryLevel1, product.CategoryLevel2, product.CategoryLevel3}

	for _, i := range categorySlice {
		for _, j := range Middleclass {
			if i == j {
				return "Mellanklass"
			}
		}
	}

	if product.SystembolagetType == "Cider & blanddrycker" {
		return "Vin"
	}

	return product.SystembolagetType
}

func findSystembolagetType(product *Product) {
	var categorySlice = []string{product.CategoryLevel1, product.CategoryLevel2, product.CategoryLevel3}
	product.IsAlcoholFree = slices.Contains(categorySlice, "Alkoholfritt")

	for k := range CategoriesMap {
		if slices.Contains(categorySlice, k) {
			product.SystembolagetType = k
			break
		}
	}
}

func calcCut(product *Product) {
	product.SystembolagetUnitCut = unitCut(product.SystembolagetType, product.IsAlcoholFree)
	product.SystembolagetPercentageCut = percentageCut(product)
}

func percentageCut(product *Product) float32 {
	percentageCutPrice := product.Price - product.Vat - product.SystembolagetUnitCut - float32(product.AlcoholTax)
	beforePercentageCutprice := percentageCutPrice / (1 + Variables.Systembolaget.PercentageCut)
	return beforePercentageCutprice * Variables.Systembolaget.PercentageCut
}

func unitCut(sType string, aFree bool) float32 {
	if aFree {
		if sType == "Vin" {
			return Variables.Systembolaget.UnitCuts.WineNonAlcoholic
		}
		if sType == "Cider" || sType == "Beer" {
			return Variables.Systembolaget.UnitCuts.CiderNonAlcoholic
		}
		return 0
	}
	if sType == "Vin" {
		return Variables.Systembolaget.UnitCuts.Wine
	}
	if sType == "Öl" {
		return Variables.Systembolaget.UnitCuts.Beer
	}
	if sType == "Sprit" {
		return Variables.Systembolaget.UnitCuts.Spirits
	}
	if sType == "Cider" {
		return Variables.Systembolaget.UnitCuts.Cider
	}
	return 0
}

func calculateTaxes(product *Product) error {
	calcVat(product)
	calcTax(product)

	return nil
}

func CalculateTaxesAndCut(product *Product) error {
	findSystembolagetType(product)
	product.SkatteverketType = findSkatteverketType(product)

	calculateTaxes(product)
	calcCut(product)

	product.OriginalPrice = product.Price - product.Vat - product.SystembolagetPercentageCut - product.SystembolagetUnitCut - float32(product.AlcoholTax)

	return nil
}
