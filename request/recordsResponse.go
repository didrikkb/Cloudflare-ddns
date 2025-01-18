package Request

type Meta struct {
	AutoAdded           bool `json:"auto_added"`
	ManagedByApps       bool `json:"managed_by_apps"`
	ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
}

type Settings struct {
	FlattenCname bool `json:"flatten_cname"`
}

type ResultItem struct {
	ID         string   `json:"id"`
	ZoneID     string   `json:"zone_id"`
	ZoneName   string   `json:"zone_name"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Content    string   `json:"content"`
	Proxiable  bool     `json:"proxiable"`
	Proxied    bool     `json:"proxied"`
	TTL        int      `json:"ttl"`
	Settings   Settings `json:"settings"`
	Meta       Meta     `json:"meta"`
	Comment    *string  `json:"comment"` // Nullable, so use pointer
	Tags       []string `json:"tags"`
	CreatedOn  string   `json:"created_on"`
	ModifiedOn string   `json:"modified_on"`
}

type ResultInfo struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Count      int `json:"count"`
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
}

type RecordsResponse struct {
	Result     []ResultItem `json:"result"`
	Success    bool         `json:"success"`
	Errors     []string     `json:"errors"`
	Messages   []string     `json:"messages"`
	ResultInfo ResultInfo   `json:"result_info"`
}

type UpdateResponse struct {
	Result     ResultItem `json:"result"`
	Success    bool       `json:"success"`
	Errors     []string   `json:"errors"`
	Messages   []string   `json:"messages"`
	ResultInfo ResultInfo `json:"result_info"`
}
