package main

type AlbumInfoResponse struct {
	AlbumInfo struct {
		ID              string      `json:"id"`
		Title           string      `json:"title"`
		Description     string      `json:"description"`
		Datetime        int         `json:"datetime"`
		Cover           string      `json:"cover"`
		CoverEdited     interface{} `json:"cover_edited"`
		CoverWidth      int         `json:"cover_width"`
		CoverHeight     int         `json:"cover_height"`
		AccountURL      interface{} `json:"account_url"`
		AccountID       interface{} `json:"account_id"`
		Privacy         string      `json:"privacy"`
		Layout          string      `json:"layout"`
		Views           int         `json:"views"`
		Link            string      `json:"link"`
		Favorite        bool        `json:"favorite"`
		Nsfw            bool        `json:"nsfw"`
		Section         interface{} `json:"section"`
		ImagesCount     int         `json:"images_count"`
		InGallery       bool        `json:"in_gallery"`
		IsAd            bool        `json:"is_ad"`
		IncludeAlbumAds bool        `json:"include_album_ads"`
		IsAlbum         bool        `json:"is_album"`
		Images          []struct {
			ID          string        `json:"id"`
			Title       string        `json:"title"`
			Description string        `json:"description"`
			Datetime    int           `json:"datetime"`
			Type        string        `json:"type"`
			Animated    bool          `json:"animated"`
			Width       int           `json:"width"`
			Height      int           `json:"height"`
			Size        int           `json:"size"`
			Views       int           `json:"views"`
			Bandwidth   int           `json:"bandwidth"`
			Vote        interface{}   `json:"vote"`
			Favorite    bool          `json:"favorite"`
			Nsfw        interface{}   `json:"nsfw"`
			Section     interface{}   `json:"section"`
			AccountURL  interface{}   `json:"account_url"`
			AccountID   interface{}   `json:"account_id"`
			IsAd        bool          `json:"is_ad"`
			InMostViral bool          `json:"in_most_viral"`
			HasSound    bool          `json:"has_sound"`
			Tags        []interface{} `json:"tags"`
			AdType      int           `json:"ad_type"`
			AdURL       string        `json:"ad_url"`
			Edited      string        `json:"edited"`
			InGallery   bool          `json:"in_gallery"`
			Link        string        `json:"link"`
		} `json:"images"`
		AdConfig struct {
			SafeFlags       []string      `json:"safeFlags"`
			HighRiskFlags   []interface{} `json:"highRiskFlags"`
			UnsafeFlags     []string      `json:"unsafeFlags"`
			WallUnsafeFlags []interface{} `json:"wallUnsafeFlags"`
			ShowsAds        bool          `json:"showsAds"`
			ShowAdLevel     int           `json:"showAdLevel"`
		} `json:"ad_config"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}
