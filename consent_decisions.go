package picnic

type ConsentDecisions struct {
	MiscCommercialAds          bool `json:"MISC_COMMERCIAL_ADS"`
	PurchasesCategoryConsent   bool `json:"PURCHASES_CATEGORY_CONSENT"`
	MiscCommercialEmails       bool `json:"MISC_COMMERCIAL_EMAILS"`
	MiscReadAdvertisingID      bool `json:"MISC_READ_ADVERTISING_ID"`
	PersonalizedRankingConsent bool `json:"PERSONALIZED_RANKING_CONSENT"`
	MiscCommercialMessages     bool `json:"MISC_COMMERCIAL_MESSAGES"`
	WeeklyCommercialEmails     bool `json:"WEEKLY_COMMERCIAL_EMAILS"`
}
