package dtos

type EvaluateRequestDTO struct {
	FlagKey    string `json:"flag_key" binding:"required"`
	UserID     int    `json:"user_id" binding:"required"`
	Department string `json:"department,omitempty"`
}

type EvaluateResponseDTO struct {
	FlagKey string `json:"flag_key"`
	Enabled bool   `json:"enabled"`
	Reason  string `json:"reason"`
}
