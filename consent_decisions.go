package picnic

type ConsentDecisions struct {
	MiscCommercialAds          bool `json:"MISC_COMMERCIAL_ADS"`
	MiscCommercialEmails       bool `json:"MISC_COMMERCIAL_EMAILS"`
	MiscCommercialMessages     bool `json:"MISC_COMMERCIAL_MESSAGES"`
	MiscReadAdvertisingID      bool `json:"MISC_READ_ADVERTISING_ID"`
	MiscWhatsAppCommunication  bool `json:"MISC_WHATS_APP_COMMUNICATION"`
	NespressoDataSharing       bool `json:"NESPRESSO_DATA_SHARING"`
	PersonalizedRankingConsent bool `json:"PERSONALIZED_RANKING_CONSENT"`
	PurchasesCategoryConsent   bool `json:"PURCHASES_CATEGORY_CONSENT"`
	WeeklyCommercialEmails     bool `json:"WEEKLY_COMMERCIAL_EMAILS"`
}
