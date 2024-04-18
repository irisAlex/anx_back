package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/ncr"

type SupplierResponse struct {
	Ncr ncr.Supplier `json:"supplier"`
}

type TypeResponse struct {
	Ncr ncr.TypeM `json:"typem"`
}

type ProjectResponse struct {
	Ncr ncr.Project `json:"project"`
}

type ManageResponse struct {
	Ncr ncr.Manage `json:"manage"`
}
type ComplaintResponse struct {
	Ncr ncr.Complaint `json:"compliant"`
}

type MessageResponse struct {
	Ncr ncr.Message `json:"message"`
}
