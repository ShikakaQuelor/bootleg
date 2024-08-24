package taxes

import (
	"slices"
)

type ImageModule struct {
	Extensions interface{} `json:"extensions,omitempty" form:"extensions"`
	ProductID  interface{} `json:"productId,omitempty" form:"productId"`
	Sizes      interface{} `json:"sizes,omitempty" form:"sizes"`
	Thumbnail  interface{} `json:"thumbnail,omitempty" form:"thumbnail"`
}

type Image struct {
	FileType interface{} `json:"fileType,omitempty" form:"fileType"`
	ImageURL string      `json:"imageUrl,omitempty" form:"imageUrl"`
	Size     interface{} `json:"size,omitempty" form:"size"`
}

type TasteClock struct {
	Key   string `json:"key,omitempty" form:"key"`
	Value int    `json:"value,omitempty" form:"value"`
}

type Product struct {
	AlcoholPercentage               float64       `json:"alcoholPercentage,omitempty" form:"alcoholPercentage"`
	AlcoholTax                      float64       `json:"alcoholTax,omitempty" form:"alcoholTax"`
	Assortment                      string        `json:"assortment,omitempty" form:"assortment"`
	AssortmentText                  string        `json:"assortmentText,omitempty" form:"assortmentText"`
	BottleText                      string        `json:"bottleText,omitempty" form:"bottleText"`
	Category                        interface{}   `json:"category,omitempty" form:"category"`
	CategoryLevel1                  string        `json:"categoryLevel1,omitempty" form:"categoryLevel1"`
	CategoryLevel2                  string        `json:"categoryLevel2,omitempty" form:"categoryLevel2"`
	CategoryLevel3                  string        `json:"categoryLevel3,omitempty" form:"categoryLevel3"`
	CategoryLevel4                  interface{}   `json:"categoryLevel4,omitempty" form:"categoryLevel4"`
	Color                           string        `json:"color,omitempty" form:"color"`
	Country                         string        `json:"country,omitempty" form:"country"`
	CustomCategoryTitle             string        `json:"customCategoryTitle,omitempty" form:"customCategoryTitle"`
	DishPoints                      interface{}   `json:"dishPoints,omitempty" form:"dishPoints"`
	EthicalLabel                    interface{}   `json:"ethicalLabel,omitempty" form:"ethicalLabel"`
	Grapes                          []interface{} `json:"grapes,omitempty" form:"grapes"`
	ImageModules                    ImageModule   `json:"imageModules,omitempty" form:"imageModules"`
	Images                          []Image       `json:"images,omitempty" form:"images"`
	IsAlcoholFree                   bool          `json:"isAlcoholFree,omitempty" form:"isAlcoholFree"`
	IsBsAssortment                  bool          `json:"isBsAssortment,omitempty" form:"isBsAssortment"`
	IsClimateSmartPackaging         bool          `json:"isClimateSmartPackaging,omitempty" form:"isClimateSmartPackaging"`
	IsCompletelyOutOfStock          bool          `json:"isCompletelyOutOfStock,omitempty" form:"isCompletelyOutOfStock"`
	IsDiscontinued                  bool          `json:"isDiscontinued,omitempty" form:"isDiscontinued"`
	IsEthical                       bool          `json:"isEthical,omitempty" form:"isEthical"`
	IsFsAssortment                  bool          `json:"isFsAssortment,omitempty" form:"isFsAssortment"`
	IsFsTsAssortment                bool          `json:"isFsTsAssortment,omitempty" form:"isFsTsAssortment"`
	IsKosher                        bool          `json:"isKosher,omitempty" form:"isKosher"`
	IsManufacturingCountry          bool          `json:"isManufacturingCountry,omitempty" form:"isManufacturingCountry"`
	IsNews                          bool          `json:"isNews,omitempty" form:"isNews"`
	IsOrganic                       bool          `json:"isOrganic,omitempty" form:"isOrganic"`
	IsPaAssortment                  bool          `json:"isPaAssortment,omitempty" form:"isPaAssortment"`
	IsRecommendedByTasteProfile     bool          `json:"isRecommendedByTasteProfile,omitempty" form:"isRecommendedByTasteProfile"`
	IsRegionalRestricted            bool          `json:"isRegionalRestricted,omitempty" form:"isRegionalRestricted"`
	IsSupplierTemporaryNotAvailable bool          `json:"isSupplierTemporaryNotAvailable,omitempty" form:"isSupplierTemporaryNotAvailable"`
	IsSustainableChoice             bool          `json:"isSustainableChoice,omitempty" form:"isSustainableChoice"`
	IsTemporaryOutOfStock           bool          `json:"isTemporaryOutOfStock,omitempty" form:"isTemporaryOutOfStock"`
	IsTsAssortment                  bool          `json:"isTsAssortment,omitempty" form:"isTsAssortment"`
	IsTsLsAssortment                bool          `json:"isTsLsAssortment,omitempty" form:"isTsLsAssortment"`
	IsTseAssortment                 bool          `json:"isTseAssortment,omitempty" form:"isTseAssortment"`
	IsTssAssortment                 bool          `json:"isTssAssortment,omitempty" form:"isTssAssortment"`
	IsTstAssortment                 bool          `json:"isTstAssortment,omitempty" form:"isTstAssortment"`
	IsTsvAssortment                 bool          `json:"isTsvAssortment,omitempty" form:"isTsvAssortment"`
	IsWebLaunch                     bool          `json:"isWebLaunch,omitempty" form:"isWebLaunch"`
	OriginLevel1                    interface{}   `json:"originLevel1,omitempty" form:"originLevel1"`
	OriginLevel2                    interface{}   `json:"originLevel2,omitempty" form:"originLevel2"`
	OriginalPrice                   float32       `json:"originalPrice,omitempty" form:"originalPrice"`
	OtherSelections                 interface{}   `json:"otherSelections,omitempty" form:"otherSelections"`
	PackagingLevel1                 string        `json:"packagingLevel1,omitempty" form:"packagingLevel1"`
	Price                           float32       `json:"price,omitempty" form:"price"`
	ProducerName                    string        `json:"producerName,omitempty" form:"producerName"`
	ProductID                       string        `json:"productId,omitempty" form:"productId"`
	ProductLaunchDate               string        `json:"productLaunchDate,omitempty" form:"productLaunchDate"`
	ProductNameBold                 string        `json:"productNameBold,omitempty" form:"productNameBold"`
	ProductNameThin                 interface{}   `json:"productNameThin,omitempty" form:"productNameThin"`
	ProductNumber                   string        `json:"productNumber,omitempty" form:"productNumber"`
	ProductNumberShort              string        `json:"productNumberShort,omitempty" form:"productNumberShort"`
	RecycleFee                      int           `json:"recycleFee,omitempty" form:"recycleFee"`
	RestrictedParcelQuantity        int           `json:"restrictedParcelQuantity,omitempty" form:"restrictedParcelQuantity"`
	Seal                            interface{}   `json:"seal,omitempty" form:"seal"`
	SellStartTime                   string        `json:"sellStartTime,omitempty" form:"sellStartTime"`
	SkatteverketType                string        `json:"skatteverketType,omitempty" form:"skatteverketType"`
	SugarContent                    int           `json:"sugarContent,omitempty" form:"sugarContent"`
	SugarContentGramPer100Ml        int           `json:"sugarContentGramPer100ml,omitempty" form:"sugarContentGramPer100ml"`
	SupplierName                    string        `json:"supplierName,omitempty" form:"supplierName"`
	SystembolagetType               string        `json:"systembolagetType,omitempty" form:"systembolagetType"`
	SystembolagetUnitCut            float32       `json:"systembolagetUnitCut,omitempty" form:"systembolagetUnitCut"`
	SystembolagetPercentageCut      float32       `json:"systembolagetPercentageCut,omitempty" form:"systembolagetPercentageCut"`
	Taste                           string        `json:"taste,omitempty" form:"taste"`
	TasteClockBitter                int           `json:"tasteClockBitter,omitempty" form:"tasteClockBitter"`
	TasteClockBody                  int           `json:"tasteClockBody,omitempty" form:"tasteClockBody"`
	TasteClockCasque                int           `json:"tasteClockCasque,omitempty" form:"tasteClockCasque"`
	TasteClockFruitacid             int           `json:"tasteClockFruitacid,omitempty" form:"tasteClockFruitacid"`
	TasteClockGroupBitter           interface{}   `json:"tasteClockGroupBitter,omitempty" form:"tasteClockGroupBitter"`
	TasteClockGroupSmokiness        interface{}   `json:"tasteClockGroupSmokiness,omitempty" form:"tasteClockGroupSmokiness"`
	TasteClockRoughness             int           `json:"tasteClockRoughness,omitempty" form:"tasteClockRoughness"`
	TasteClockSmokiness             int           `json:"tasteClockSmokiness,omitempty" form:"tasteClockSmokiness"`
	TasteClockSweetness             int           `json:"tasteClockSweetness,omitempty" form:"tasteClockSweetness"`
	TasteClocks                     TasteClock    `json:"tasteClocks,omitempty" form:"tasteClocks"`
	TasteSymbols                    []string      `json:"tasteSymbols,omitempty" form:"tasteSymbols"`
	Usage                           string        `json:"usage,omitempty" form:"usage"`
	Vat                             float32       `json:"moms,omitempty" form:"moms"`
	Vintage                         interface{}   `json:"vintage,omitempty" form:"vintage"`
	Volume                          int           `json:"volume,omitempty" form:"volume"`
	VolumeText                      string        `json:"volumeText,omitempty" form:"volumeText"`
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
	product.AlcoholTax = (float64(product.Volume) / 1000) * (product.AlcoholPercentage / 100) * taxInfo.TaxPerPercentAndLiter
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
